package productparameter

import (
	productParamsEntity "bisnis-be/internal/entity/productparameter"
	"bisnis-be/pkg/errors"
	"context"
)

func (d Data) GetProductParameterByProdID(ctx context.Context, productid string) ([]productParamsEntity.ProductParameter, string, error) {
	var (
		productParamArr []productParamsEntity.ProductParameter
		result          string
		productParam    productParamsEntity.ProductParameter
		err             error
	)
	rows, err := (*d.stmt)[getProductParameterByProdID].QueryxContext(ctx, productid)
	if err != nil {
		return productParamArr, result, errors.Wrap(err, "[DATA] [GetProductParameterByProdID][QueryxContext]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&productParam); err != nil {
			return productParamArr, result, errors.Wrap(err, "[DATA] [GetGoldUserByEmailLogin]")
		}
		productParamArr = append(productParamArr, productParam)
	}
	if len(productParamArr) == 0 {
		result = "Empty product params data"
	}
	if len(productParamArr) > 0 {
		result = "Data product exist"
	}
	return productParamArr, result, err
}
