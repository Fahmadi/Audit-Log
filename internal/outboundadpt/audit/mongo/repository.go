package mongo

import (
	"auditLog/db/storage/mongo_store"
	"auditLog/internal/core/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	AUDIT_LOGS_COLLECTION = "audit"
)

type AuditLogRepository struct {
	client *mongo.Client
	dbName mongo_store.DatabaseName
}

func New(o mongo_store.Options) AuditLogRepository {
	return AuditLogRepository{
		client: o.Client,
		dbName: o.DatabaseName,
	}
}

func (r AuditLogRepository) InsertOne(ctx context.Context, a domain.Audit) (string, error) {
	new, err := FromModel(a)
	if err != nil {
		return "", err
	}
	inserted, err := r.client.Database(string(r.dbName)).Collection(AUDIT_LOGS_COLLECTION).InsertOne(ctx, new)
	if err != nil {
		return "", err
	}

	id, ok := inserted.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("cannot convert return inserted id to object id")
	}
	return id.Hex(), nil
}

func (r AuditLogRepository) FindAllByUserID(ctx context.Context, userID domain.ExternalID, page *domain.Page) ([]AuditLog, uint32, error) {
	if page == nil {
		*page = domain.NewDefaultPage()
	}

	filter := bson.M{
		"userId": userID,
	}
	options.Find().SetSkip(int64(page.Page) * int64(page.Size))
	options.Find().SetLimit(int64(page.Size))
	cur, err := r.client.Database(string(r.dbName)).Collection(AUDIT_LOGS_COLLECTION).Find(ctx, filter)

	if err != nil {
		return nil, 0, err
	}
	auList := make([]AuditLog, 0)

	if err := cur.All(ctx, &auList); err != nil {
		return nil, 0, err
	}
	total, err := r.client.Database(string(r.dbName)).Collection(AUDIT_LOGS_COLLECTION).CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return auList, uint32(total), nil
}

func (r AuditLogRepository) FindAllByEntityID(ctx context.Context, entity string, entityID domain.ExternalID, page *domain.Page) ([]AuditLog, uint32, error) {
	if page == nil {
		*page = domain.NewDefaultPage()
	}
	filter := bson.M{
		"entity":   string(entity),
		"entityId": entityID,
	}
	options.Find().SetSkip(int64(page.Page) * int64(page.Size))
	options.Find().SetLimit(int64(page.Size))
	cur, err := r.client.Database(string(r.dbName)).Collection(AUDIT_LOGS_COLLECTION).Find(ctx, filter)

	if err != nil {
		return nil, 0, err
	}
	auList := make([]AuditLog, 0)

	if err := cur.All(ctx, &auList); err != nil {
		return nil, 0, err
	}
	total, err := r.client.Database(string(r.dbName)).Collection(AUDIT_LOGS_COLLECTION).CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return auList, uint32(total), nil
}
