package bisnis

import (
	jaegerLog "bisnis-be/pkg/log"

	"github.com/opentracing/opentracing-go"
)

type IbisnisSvc interface {
	// LoginUser(ctx context.Context, _user, _password string, _host string) (auth.Token, map[string]interface{}, error)
}

type Handler struct {
	bisnisSvc IbisnisSvc
	tracer    opentracing.Tracer
	logger    jaegerLog.Factory
}

// New for bridging product handler initialization
func New(is IbisnisSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		bisnisSvc: is,
		tracer:    tracer,
		logger:    logger,
	}
}
