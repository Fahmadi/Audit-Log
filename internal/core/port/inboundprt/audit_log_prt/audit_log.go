package audit_log_prt

import (
	"auditLog/internal/core/domain"
	"context"
)

type AuditLogService interface {
	LogAudit(ctx context.Context, audit domain.Audit) error
	ListByUserId(ctx context.Context, userID domain.ExternalID, page *domain.Page) ([]domain.Audit, uint32, error)
	ListByEntityId(ctx context.Context, entity string, entityID domain.ExternalID, page *domain.Page) ([]domain.Audit, uint32, error)
}
