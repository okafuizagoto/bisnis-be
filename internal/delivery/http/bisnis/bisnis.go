package bisnis

import (
	agentEntity "bisnis-be/internal/entity/agent"
	jaegerLog "bisnis-be/pkg/log"
	"context"

	"github.com/opentracing/opentracing-go"
)

type IbisnisSvc interface {
	// LoginAgent(ctx context.Context, agentUser agentEntity.LoginAgent) error
}

type IgoldgymSvcStock interface {
}

type IagentSvc interface {
	LoginAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (agentEntity.ResponseLogin, string, error)
	CheckAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (string, error)
	AddAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (agentEntity.AgentResp, string, error)
	DeleteAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (agentEntity.AgentResp, string, error)
	UpdateAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (agentEntity.AgentResp, string, error)
}

type (
	// Handler ...
	Handler struct {
		bisnisSvc       IbisnisSvc
		goldgymSvcStock IgoldgymSvcStock
		agentSvc        IagentSvc
		tracer          opentracing.Tracer
		logger          jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IbisnisSvc, isst IgoldgymSvcStock, isag IagentSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		bisnisSvc:       is,
		goldgymSvcStock: isst,
		agentSvc:        isag,
		tracer:          tracer,
		logger:          logger,
	}
}
