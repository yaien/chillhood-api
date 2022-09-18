package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"io"
	"log"
	"net/http"
	"strings"
)

type WhatsappService interface {
	NotifySale(ctx context.Context, inv *entity.Invoice)
}

func NewWhatsappService(config *infrastructure.WhatsappConfig, users UserService) WhatsappService {
	return &wsService{config: config, users: users}
}

type wsService struct {
	config *infrastructure.WhatsappConfig
	users  UserService
}

func (ws *wsService) NotifySale(ctx context.Context, inv *entity.Invoice) {
	users, err := ws.users.FindReportable(ctx)
	if err != nil {
		log.Printf("failed finding reportable users: %s", err)
		return
	}

	for _, user := range users {
		err = ws.SendTemplate(ctx, user.Phone, inv)
		if err != nil {
			log.Printf("failed sending template to phone %s: %s", user.Phone, err)
			return
		}
	}

	return
}

func (ws *wsService) SendTemplate(ctx context.Context, phone string, inv *entity.Invoice) error {

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"to":                phone,
		"type":              "template",
		"template": map[string]any{
			"name": ws.config.Template,
			"language": map[string]any{
				"code": ws.config.TemplateLanguageCode,
			},
		},
		"components": map[string]any{
			"type": "body",
			"parameters": []map[string]any{
				{"type": "text", "text": inv.Shipping.Name},
				{"type": "text", "text": inv.Ref},
				{"type": "text", "text": strings.ReplaceAll(ws.config.SaleUrl, "{ref}", inv.Ref)},
			},
		},
	}

	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(payload)
	if err != nil {
		return fmt.Errorf("failed encoding payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ws.config.MessagesEndpoint, &buff)
	if err != nil {
		return fmt.Errorf("failed creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ws.config.Token)

	var h http.Client
	res, err := h.Do(req)
	if err != nil {
		return fmt.Errorf("failed doing request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("response failed with status: %d, content: %s", res.StatusCode, string(body))
	}

	return nil
}
