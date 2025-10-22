package product

import (
	productEntity "bisnis-be/internal/entity/product"
	"bisnis-be/pkg/errors"
	"context"
)

func (d Data) GetProductByProdID(ctx context.Context, productid string) (productEntity.Product, string, error) {
	var (
		result  string
		product productEntity.Product
		err     error
	)
	rows, err := (*d.stmt)[getProductByProdID].QueryxContext(ctx, productid)
	if err != nil {
		return product, result, errors.Wrap(err, "[DATA] [GetProductParameterByProdID][QueryxContext]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&product); err != nil {
			return product, result, errors.Wrap(err, "[DATA] [GetGoldUserByEmailLogin]")
		}
	}
	if product == (productEntity.Product{}) {
		return product, "Empty product params data", errors.Wrap(err, "[DATA] [GetGoldUserByEmailLogin]")
	}
	return product, result, err
}
