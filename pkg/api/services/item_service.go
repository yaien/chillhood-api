package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
	"github.com/yaien/clothes-store-api/pkg/api/models"
)

type ItemService interface {
	Create(ctx context.Context, item *models.Item) error
	FindOneByID(ctx context.Context, id models.ID) (*models.Item, error)
	FindOneActiveByID(ctx context.Context, id models.ID) (*models.Item, error)
	FindOneBySlug(ctx context.Context, slug string) (*models.Item, error)
	Update(ctx context.Context, item *models.Item) error
	Find(ctx context.Context) ([]*models.Item, error)
	FindActive(ctx context.Context) ([]*models.Item, error)
	Decrement(ctx context.Context, id models.ID, size string, quantity int) error
	Increment(ctx context.Context, id models.ID, size string, quantity int) error
}

type itemService struct {
	items models.ItemRepository
}

func (p *itemService) Create(ctx context.Context, item *models.Item) error {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	item.Slug = slug.Make(item.Name)
	count, err := p.items.CountByName(ctx, item.Name)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("ItemExists: there is already a item with name %s", item.Name)
	}
	return p.items.Create(ctx, item)
}

func (p *itemService) FindOneByID(ctx context.Context, id models.ID) (*models.Item, error) {
	return p.items.FindOneByID(ctx, id)
}

func (p *itemService) FindOneActiveByID(ctx context.Context, id models.ID) (*models.Item, error) {
	return p.items.FindOneActiveByID(ctx, id)
}

func (p *itemService) FindOneBySlug(ctx context.Context, slug string) (*models.Item, error) {
	return p.items.FindOneBySlug(ctx, slug)
}

func (p *itemService) Find(ctx context.Context) ([]*models.Item, error) {
	return p.items.Find(ctx)
}

func (p *itemService) FindActive(ctx context.Context) ([]*models.Item, error) {
	return p.items.FindActive(ctx)
}

func (p *itemService) Update(ctx context.Context, item *models.Item) error {
	item.Slug = slug.Make(item.Name)
	item.UpdatedAt = time.Now()

	count, err := p.items.CountByName(ctx, item.Name)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("ItemExists: there is already a item with name %s", item.Name)
	}

	return p.items.Update(ctx, item)
}

func (p *itemService) Decrement(ctx context.Context, id models.ID, size string, quantity int) error {
	item, err := p.FindOneByID(ctx, id)
	if err != nil {
		return err
	}
	sz, err := item.Size(size)
	if err != nil {
		return err
	}
	if sz.Existence < quantity {
		return errors.New("INVALID_QUANTITY")
	}

	sz.Existence -= quantity
	return p.Update(ctx, item)
}

func (p *itemService) Increment(ctx context.Context, id models.ID, size string, quantity int) error {
	item, err := p.FindOneByID(ctx, id)
	if err != nil {
		return err
	}
	sz, err := item.Size(size)
	if err != nil {
		return err
	}

	sz.Existence += quantity
	return p.Update(ctx, item)
}

func NewItemService(items models.ItemRepository) ItemService {
	return &itemService{items}
}
