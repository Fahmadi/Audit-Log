package audit_log_srv

import (
	"auditLog/internal/core/domain"
	repository_prt "auditLog/internal/core/port/outboundprt/repositoryprt"
	"time"

	auditlog_prt "auditLog/internal/core/port/inboundprt/audit_log_prt"
	"context"
)

type Service struct {
	repo repository_prt.AuditLogRepository
}

var _ auditlog_prt.AuditLogService = (*Service)(nil)

func New(repo repository_prt.AuditLogRepository) Service {
	return Service{
		repo: repo,
	}
}

func (ls Service) LogAudit(ctx context.Context, audit domain.Audit) error {
	_, err := ls.repo.InsertOne(ctx, audit)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) ListByUserId(ctx context.Context, userID domain.ExternalID, page *domain.Page) ([]domain.Audit, uint32, error) {
	dbRes, total, err := s.repo.FindAllByUserID(ctx, userID, page)
	res := make([]domain.Audit, 0)

	for i, r := range dbRes {
		t, _ := time.Parse("2006-01-02", r.CreatedAt)

		res[i] = domain.Audit{
			Action:    r.Action,
			UserID:    domain.ExternalID(r.UserID),
			EntityID:  domain.ExternalID(r.EntityID),
			Change:    r.Change,
			Entity:    r.Entity,
			Timestamp: t,
			Message:   "Message",
		}
	}
	if err != nil {
		return []domain.Audit{}, 0, nil
	}

	return res, total, nil
}

func (s Service) ListByEntityId(ctx context.Context, entity string, entityID domain.ExternalID, page *domain.Page) ([]domain.Audit, uint32, error) {
	dbRes, total, err := s.repo.FindAllByEntityID(ctx, entity, entityID, page)
	res := make([]domain.Audit, 0)

	for i, r := range dbRes {
		t, _ := time.Parse("2006-01-02", r.CreatedAt)

		res[i] = domain.Audit{
			Action:    r.Action,
			UserID:    domain.ExternalID(r.UserID),
			EntityID:  domain.ExternalID(r.EntityID),
			Change:    r.Change,
			Entity:    r.Entity,
			Timestamp: t,
			Message:   "Message",
		}
	}
	if err != nil {
		return []domain.Audit{}, 0, nil
	}

	return res, total, nil
}
