package repository_prt

import (
	"auditLog/internal/core/domain"
	"auditLog/internal/outboundadpt/audit/mongo"
	"context"
)

type AuditLogRepository interface {
	InsertOne(context.Context, domain.Audit) (string, error)
	FindAllByUserID(ctx context.Context, userID domain.ExternalID, page *domain.Page) ([]mongo.AuditLog, uint32, error)
	FindAllByEntityID(ctx context.Context, Entity string, EntityID domain.ExternalID, page *domain.Page) ([]mongo.AuditLog, uint32, error)
}
