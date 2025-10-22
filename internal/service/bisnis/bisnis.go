package bisnis

import (
	"bisnis-be/internal/entity"
	agentEntity "bisnis-be/internal/entity/agent"
	bisnisEntity "bisnis-be/internal/entity/bisnis"
	productEntity "bisnis-be/internal/entity/product"
	productParamsEntity "bisnis-be/internal/entity/productparameter"
	jaegerLog "bisnis-be/pkg/log"
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	// "go.opentelemetry.io/otel/trace"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	AddTransaction(ctx context.Context, addTransaction bisnisEntity.AddTransaction) (int, error)
	DeleteTransaction(ctx context.Context, deleteTransactionz bisnisEntity.DeleteTransaction) (int, string, error)
	UpdateTransaction(ctx context.Context, addTransaction bisnisEntity.UpdateTransaction) (int, string, error)
}

type AgentData interface {
	CheckAgent(ctx context.Context, agentUser agentEntity.LoginAgent) (agentEntity.Agent, string, error)
}

type ProductParamaterData interface {
	GetProductParameterByProdID(ctx context.Context, productid string) ([]productParamsEntity.ProductParameter, string, error)
}

type ProductData interface {
	GetProductByProdID(ctx context.Context, productid string) (productEntity.Product, string, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	bisnis           Data
	agent            AgentData
	productparameter ProductParamaterData
	product          ProductData
	tracer           opentracing.Tracer
	// tracer trace.Tracer
	logger jaegerLog.Factory
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(goldgymData Data, agentData AgentData, productParamaterData ProductParamaterData, productData ProductData, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	// Assign variable dari parameter ke object
	return Service{
		bisnis:           goldgymData,
		agent:            agentData,
		productparameter: productParamaterData,
		product:          productData,
		tracer:           tracer,
		logger:           logger,
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
