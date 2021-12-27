package mongodb

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceRepository struct {
	collection *mongo.Collection
}

func (m InvoiceRepository) Create(ctx context.Context, invoice *entity.Invoice) error {
	invoice.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, invoice)
	return err
}

func (m InvoiceRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*entity.Invoice, error) {
	var invoice entity.Invoice
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&invoice)
	if err == nil {
		return &invoice, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m InvoiceRepository) FindOneByRef(ctx context.Context, ref string) (*entity.Invoice, error) {
	var invoice entity.Invoice
	err := m.collection.FindOne(ctx, bson.M{"ref": ref}).Decode(&invoice)
	if err == nil {
		return &invoice, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m InvoiceRepository) Search(ctx context.Context, opts entity.SearchInvoiceOptions) ([]*entity.Invoice, error) {
	filter := bson.M{}
	if len(opts.Status) > 0 {
		filter["status"] = opts.Status
	}

	if opts.Query != "" {
		regex := primitive.Regex{Pattern: opts.Query, Options: "i"}
		filter["$or"] = bson.A{
			bson.D{{"ref", regex}},
			bson.D{{"shipping.email", regex}},
			bson.D{{"shipping.phone", regex}},
			bson.D{{"shipping.name", regex}},
		}
	}

	invoices := make([]*entity.Invoice, 0)
	cursor, err := m.collection.Find(ctx, filter, &options.FindOptions{
		Sort: bson.D{{"createdat", -1}},
	})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var invoice entity.Invoice
		err := cursor.Decode(&invoice)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (m *InvoiceRepository) Update(ctx context.Context, invoice *entity.Invoice) error {
	filter := bson.M{"_id": invoice.ID}
	update := bson.M{"$set": invoice}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

func NewMongoInvoiceRepository(db *mongo.Database) *InvoiceRepository {
	return &InvoiceRepository{collection: db.Collection("invoices")}
}
