package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"time"

	"github.com/gosimple/slug"
)

type ItemService interface {
	Create(ctx context.Context, item *entity.Item) error
	FindOneByID(ctx context.Context, id entity.ID) (*entity.Item, error)
	FindOneActiveByID(ctx context.Context, id entity.ID) (*entity.Item, error)
	FindOneBySlug(ctx context.Context, slug string) (*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) error
	Find(ctx context.Context) ([]*entity.Item, error)
	FindActive(ctx context.Context) ([]*entity.Item, error)
	Decrement(ctx context.Context, id entity.ID, size string, quantity int) error
	Increment(ctx context.Context, id entity.ID, size string, quantity int) error
}

type itemService struct {
	items entity.ItemRepository
}

func (p *itemService) Create(ctx context.Context, item *entity.Item) error {
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

func (p *itemService) FindOneByID(ctx context.Context, id entity.ID) (*entity.Item, error) {
	return p.items.FindOneByID(ctx, id)
}

func (p *itemService) FindOneActiveByID(ctx context.Context, id entity.ID) (*entity.Item, error) {
	return p.items.FindOneActiveByID(ctx, id)
}

func (p *itemService) FindOneBySlug(ctx context.Context, slug string) (*entity.Item, error) {
	return p.items.FindOneBySlug(ctx, slug)
}

func (p *itemService) Find(ctx context.Context) ([]*entity.Item, error) {
	return p.items.Find(ctx)
}

func (p *itemService) FindActive(ctx context.Context) ([]*entity.Item, error) {
	return p.items.FindActive(ctx)
}

func (p *itemService) Update(ctx context.Context, item *entity.Item) error {
	item.Slug = slug.Make(item.Name)
	item.UpdatedAt = time.Now()

	count, err := p.items.CountByNameIgnore(ctx, item.ID, item.Name)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("ItemExists: there is already a item with name %s", item.Name)
	}

	return p.items.Update(ctx, item)
}

func (p *itemService) Decrement(ctx context.Context, id entity.ID, size string, quantity int) error {
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

func (p *itemService) Increment(ctx context.Context, id entity.ID, size string, quantity int) error {
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

func NewItemService(items entity.ItemRepository) ItemService {
	return &itemService{items}
}
