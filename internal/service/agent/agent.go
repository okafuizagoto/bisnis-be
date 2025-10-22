package agent

import (
	"bisnis-be/internal/entity"
	agentEntity "bisnis-be/internal/entity/agent"
	jaegerLog "bisnis-be/pkg/log"
	"context"
	"errors"
	"time"

	"github.com/opentracing/opentracing-go"
	// "go.opentelemetry.io/otel/trace"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type agentzData interface {
	CheckAgent(ctx context.Context, agentUser agentEntity.LoginAgent) (agentEntity.Agent, string, error)
	AddAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (string, error)
	DeleteAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (agentEntity.Agent, string, error)
	UpdateAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (agentEntity.Agent, string, error)
}

type rediszData interface {
	AddToRedis(ctx context.Context, data interface{}, key string, ttl time.Duration) (err error)
	GetFromRedis(ctx context.Context, key string, dest interface{}) (err error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	agent  agentzData
	redis  rediszData
	tracer opentracing.Tracer
	// tracer trace.Tracer
	logger jaegerLog.Factory
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(agentData agentzData, redisData rediszData, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	// Assign variable dari parameter ke object
	return Service{
		agent:  agentData,
		redis:  redisData,
		tracer: tracer,
		logger: logger,
	}
}

func (s Service) checkPermission(ctx context.Context, _permissions ...string) error {
	claims := ctx.Value(entity.ContextKey("claims"))
	if claims != nil {
		actions := claims.(entity.ContextValue).Get("permissions").(map[string]interface{})
		for _, action := range actions {
			permissions := action.([]interface{})
			for _, permission := range permissions {
				for _, _permission := range _permissions {
					if permission.(string) == _permission {
						return nil
					}
				}
			}
		}
	}
	return errors.New("401 unauthorized")
}
