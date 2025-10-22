package bisnis

import (
	agentEntity "bisnis-be/internal/entity/agent"
	bisnisEntity "bisnis-be/internal/entity/bisnis"
	productEntity "bisnis-be/internal/entity/product"
	productparameterEntity "bisnis-be/internal/entity/productparameter"
	"bisnis-be/pkg/errors"
	"context"
	"strconv"
)

// "strings"

// "bisnis-be/internal/entity/auth/v2"

// "github.com/dgrijalva/jwt-go"

// "go.opentelemetry.io/otel/attribute"
// "go.opentelemetry.io/otel/trace"

func (s Service) AddTransaction(ctx context.Context, addTransaction bisnisEntity.AddTransaction) (bisnisEntity.TransactionResp, string, error) {
	var (
		// agentData            agentEntity.Agent
		responseTransaction  bisnisEntity.TransactionResp
		productParameterData []productparameterEntity.ProductParameter
		productData          productEntity.Product
		result               string
		resultProdParams     string
		resultProd           string
		lowestParams         int
		HighestParams        int
		transID              int
		err                  error
	)
	agentLogin := agentEntity.LoginAgent{
		AgentID:       addTransaction.AgentID,
		AgentPassword: "",
	}
	_, result, err = s.agent.CheckAgent(ctx, agentLogin)
	if err != nil {
		return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][CheckAgent]")
	}

	if result == "Incorrect password" || result == "Success" {
		productParameterData, resultProdParams, err = s.productparameter.GetProductParameterByProdID(ctx, addTransaction.ProductID)
		if err != nil {
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][GetProductParameterByProdID]")
		}

		if resultProdParams == "Empty product params data" {
			result = resultProdParams
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction]")
		}
		lowestParams = 0
		HighestParams = 0
		for _, y := range productParameterData {
			parameterValues, _ := strconv.Atoi(y.ParameterValue)
			if parameterValues != 0 {
				if lowestParams == 0 {
					lowestParams = parameterValues
				}
				if lowestParams > 0 {
					if parameterValues < lowestParams {
						lowestParams = parameterValues
					}
				}
				if parameterValues > HighestParams {
					HighestParams = parameterValues
				}
			}
		}
		if addTransaction.Usia < lowestParams {
			return responseTransaction, "Underage", errors.Wrap(err, "[Service][AddTransaction][addTransaction.Usia < lowestParams]")
		}
		if addTransaction.Usia > HighestParams {
			return responseTransaction, "Age exceeds maximum limit", errors.Wrap(err, "[Service][AddTransaction][addTransaction.Usia > HighestParams]")
		}
		productData, resultProd, err = s.product.GetProductByProdID(ctx, addTransaction.ProductID)
		if err != nil {
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][GetProductParameterByProdID]")
		}
		if resultProd == "Empty product params data" {
			return responseTransaction, resultProd, errors.Wrap(err, "[Service][AddTransaction][GetProductParameterByProdID]")
		}
		if addTransaction.Premium < productData.Premium {
			return responseTransaction, "Premium input must be equal or greater than product premium nominal", err
		}

		if int(addTransaction.Premium*100)%int(productData.Premium*100) != 0 {
			return responseTransaction, "Premium input must be a multiple of product nominal", err
		}

		transID, err = s.bisnis.AddTransaction(ctx, addTransaction)
		if err != nil {
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][CheckAgent]")
		}
		responseTransaction.TransID = strconv.Itoa(transID)
	}
	if result == "Agent does not exist" {
		return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction]")
	}
	result = "Success"
	return responseTransaction, result, err
}

func (s Service) DeleteTransaction(ctx context.Context, deleteTransaction bisnisEntity.DeleteTransaction) (bisnisEntity.TransactionResp, string, error) {
	var (
		err       error
		transID   int
		transResp bisnisEntity.TransactionResp
		result    string
	)

	transID, result, err = s.bisnis.DeleteTransaction(ctx, deleteTransaction)
	if err != nil {
		return transResp, result, errors.Wrap(err, "[Service][DeleteTransaction]")
	}

	transResp.TransID = strconv.Itoa(transID)
	return transResp, result, errors.Wrap(err, "[Service][DeleteTransaction]")
}

func (s Service) UpdateTransaction(ctx context.Context, addTransaction bisnisEntity.UpdateTransaction) (bisnisEntity.TransactionResp, string, error) {
	var (
		// agentData            agentEntity.Agent
		responseTransaction  bisnisEntity.TransactionResp
		productParameterData []productparameterEntity.ProductParameter
		productData          productEntity.Product
		result               string
		resultProdParams     string
		resultProd           string
		lowestParams         int
		HighestParams        int
		transID              int
		responseUpdate       string
		err                  error
	)
	agentLogin := agentEntity.LoginAgent{
		AgentID:       addTransaction.AgentID,
		AgentPassword: "",
	}
	_, result, err = s.agent.CheckAgent(ctx, agentLogin)
	if err != nil {
		return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][CheckAgent]")
	}

	if result == "Incorrect password" || result == "Success" {
		productParameterData, resultProdParams, err = s.productparameter.GetProductParameterByProdID(ctx, addTransaction.ProductID)
		if err != nil {
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][GetProductParameterByProdID]")
		}

		if resultProdParams == "Empty product params data" {
			result = resultProdParams
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction]")
		}
		lowestParams = 0
		HighestParams = 0
		for _, y := range productParameterData {
			parameterValues, _ := strconv.Atoi(y.ParameterValue)
			if parameterValues != 0 {
				if lowestParams == 0 {
					lowestParams = parameterValues
				}
				if lowestParams > 0 {
					if parameterValues < lowestParams {
						lowestParams = parameterValues
					}
				}
				if parameterValues > HighestParams {
					HighestParams = parameterValues
				}
			}
		}
		if addTransaction.Usia < lowestParams {
			return responseTransaction, "Underage", errors.Wrap(err, "[Service][AddTransaction][addTransaction.Usia < lowestParams]")
		}
		if addTransaction.Usia > HighestParams {
			return responseTransaction, "Age exceeds maximum limit", errors.Wrap(err, "[Service][AddTransaction][addTransaction.Usia > HighestParams]")
		}
		productData, resultProd, err = s.product.GetProductByProdID(ctx, addTransaction.ProductID)
		if err != nil {
			return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction][GetProductParameterByProdID]")
		}
		if resultProd == "Empty product params data" {
			return responseTransaction, resultProd, errors.Wrap(err, "[Service][AddTransaction][GetProductParameterByProdID]")
		}
		if addTransaction.Premium < productData.Premium {
			return responseTransaction, "Premium input must be equal or greater than product premium nominal", err
		}

		if int(addTransaction.Premium*100)%int(productData.Premium*100) != 0 {
			return responseTransaction, "Premium input must be a multiple of product nominal", err
		}

		transID, responseUpdate, err = s.bisnis.UpdateTransaction(ctx, addTransaction)
		if err != nil {
			return responseTransaction, responseUpdate, errors.Wrap(err, "[Service][AddTransaction][CheckAgent]")
		}
		responseTransaction.TransID = strconv.Itoa(transID)
	}
	if result == "Agent does not exist" {
		return responseTransaction, result, errors.Wrap(err, "[Service][AddTransaction]")
	}
	result = responseUpdate
	return responseTransaction, responseUpdate, err
}
