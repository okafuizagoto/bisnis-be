package bisnis

import (
	jaegerLog "bisnis-be/pkg/log"

	"github.com/opentracing/opentracing-go"
)

type IgoldgymSvc interface {
}

type IgoldgymSvcStock interface {
}

type (
	// Handler ...
	Handler struct {
		goldgymSvc      IgoldgymSvc
		goldgymSvcStock IgoldgymSvcStock
		tracer          opentracing.Tracer
		logger          jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IgoldgymSvc, isst IgoldgymSvcStock, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		goldgymSvc:      is,
		goldgymSvcStock: isst,
		tracer:          tracer,
		logger:          logger,
	}
}
