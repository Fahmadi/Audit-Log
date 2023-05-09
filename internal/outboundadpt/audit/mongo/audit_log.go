package mongo

import (
	"auditLog/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditLog struct {
	ID        primitive.ObjectID `bson:"_id"`
	EntityID  string             `bson:"entityId"`
	UserID    string             `bson:"userId"`
	Change    string             `bson:"change"`
	Action    string             `bson:"action"`
	Entity    string             `bson:"entity"`
	CreatedAt string             `bson:"createdAt"`
}

func FromModel(a domain.Audit) (AuditLog, error) {
	id := primitive.NewObjectID()
	return AuditLog{
		ID:        id,
		EntityID:  string(a.EntityID),
		UserID:    string(a.UserID),
		Change:    a.Change,
		Action:    a.Action,
		Entity:    a.Entity,
		CreatedAt: a.Timestamp.String(),
	}, nil
}
