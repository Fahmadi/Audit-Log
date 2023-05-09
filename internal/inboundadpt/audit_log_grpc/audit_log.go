package audit_log_grpc

import (
	"auditLog/internal/core/domain"
	auditlog_prt "auditLog/internal/core/port/inboundprt/audit_log_prt"
	auditv1 "auditLog/internal/proto/audit/v1"
	"auditLog/pkg/grpcserver"
	"context"
	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	auditLogService auditlog_prt.AuditLogService
	auditv1.UnimplementedAuditServiceServer
}

func New(auditLogService auditlog_prt.AuditLogService, server grpcserver.Server) Service {
	service := Service{
		auditLogService: auditLogService,
	}
	auditv1.RegisterAuditServiceServer(server.Instance(), service)
	return service
}

// func (s Service) InitGateway(ctx context.Context, gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {
// 	return auditv1.RegisterAuditServiceHandler(ctx, gwmux, conn)
// }

func (s Service) LogAudit(ctx context.Context, msg []byte) error {
	var res domain.Audit
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return err
	}
	err = s.auditLogService.LogAudit(ctx, res)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) ListAuditLogsByUserID(ctx context.Context, req *auditv1.ListAuditLogsByUserIDRequest) (*auditv1.ListAuditLogsByUserIDResponse, error) {
	userId := req.GetUserId()
	var page domain.Page
	if req.GetPage() != nil {
		page = domain.NewPage(req.GetPage().Page, req.GetPage().Size)
	} else {
		page = domain.NewDefaultPage()
	}
	audits, total, err := s.auditLogService.ListByUserId(ctx, domain.ExternalID(userId), &page)
	if err != nil {
		return nil, err
	}
	res := make([]*auditv1.Audit, len(audits))
	for i, a := range audits {
		h := auditv1.Audit{
			Action:    a.Action,
			UserId:    string(a.UserID),
			Entity:    a.Entity,
			Message:   a.Change,
			Timestamp: timestamppb.New(a.Timestamp),
		}
		res[i] = &h

	}
	return &auditv1.ListAuditLogsByUserIDResponse{
		Audits: res,
		Page: &auditv1.Page{
			Page:  page.Page,
			Size:  page.Size,
			Total: total,
		},
	}, nil
}

func (s Service) ListAuditLogsByEntityID(ctx context.Context, req *auditv1.ListAuditLogsByEntityIDRequest) (*auditv1.ListAuditLogsByEntityIDResponse, error) {
	entity := req.GetEntity()
	entityId := req.GetEntityId()

	var page domain.Page
	if req.GetPage() != nil {
		page = domain.NewPage(req.GetPage().Page, req.GetPage().Size)
	} else {
		page = domain.NewDefaultPage()
	}
	audits, total, err := s.auditLogService.ListByEntityId(ctx, entity, domain.ExternalID(entityId), &page)
	if err != nil {
		return nil, err
	}
	res := make([]*auditv1.Audit, len(audits))
	for i, a := range audits {
		h := auditv1.Audit{
			Action:    a.Action,
			UserId:    string(a.UserID),
			Entity:    a.Entity,
			Message:   a.Change,
			Timestamp: timestamppb.New(a.Timestamp),
		}
		res[i] = &h

	}
	return &auditv1.ListAuditLogsByEntityIDResponse{
		Audits: res,
		Page: &auditv1.Page{
			Page:  page.Page,
			Size:  page.Size,
			Total: total,
		},
	}, nil
}
