package services

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/core"
	"log"
	"strings"
)

type SlackService interface {
	NotifySale(invoice *models.Invoice)
}

type slackService struct {
	client *slack.Client
	config *core.SlackConfig
}

func (n *slackService) NotifySale(invoice *models.Invoice) {
	name := invoice.Shipping.Name
	ref := invoice.Ref
	text := fmt.Sprintf("%s ha realizado una compra! (ref: %s)", name, ref)
	link := strings.ReplaceAll(n.config.SaleUrl, "{ref}", ref)
	block := fmt.Sprintf("*%s* ha realizado una compra! (ref: %s), <%s|mira en detalle aquí>", name, ref, link)
	_, _, err := n.client.PostMessage(
		n.config.Channel,
		slack.MsgOptionText(text, false),
		slack.MsgOptionBlocks(slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type: slack.MarkdownType,
				Text: block,
			},
		}),
	)
	if err != nil {
		log.Printf("failed sending slack notification for invoice (ref: %s): %s", ref, err.Error())
	}
}

func NewSlackService(client *slack.Client, config *core.SlackConfig) SlackService {
	return &slackService{client, config}
}
