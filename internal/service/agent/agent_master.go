package agent

import (
	"bisnis-be/internal/config"
	agentEntity "bisnis-be/internal/entity/agent"
	"bisnis-be/internal/entity/auth"
	"bisnis-be/pkg/errors"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtApplicationName = "BISNIS-BE"
	jwtSigningMethod   = jwt.SigningMethodHS256
	// jwtSecret          = []byte("a7fecfed-14c8-4f54-84a7-e43fe9cf1823")
	// jwtSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

func (s Service) LoginAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (agentEntity.ResponseLogin, string, error) {
	var (
		cfg           *config.Config // Configuration object
		responseLogin agentEntity.ResponseLogin
		agentResult   agentEntity.Agent
		result        string
		token         auth.Token
		err           error
	)

	agentResult, result, err = s.agent.CheckAgent(ctx, agentLogin)
	if err != nil {
		return responseLogin, result, errors.Wrap(err, "[Service][LoginAgent]")
	}
	if result == "Success" {
		now := time.Now()
		expiration := now.Add(12 * time.Hour)
		claims := jwt.MapClaims{
			"iss":  jwtApplicationName, // issuer
			"sub":  agentLogin.AgentID, // subject
			"user": agentLogin.AgentID, // username
			"nbf":  now.Unix(),         // not before
			"iat":  now.Unix(),         // issued at
			"exp":  expiration.Unix(),  // expiration
		}
		cfg, _ = config.Get()
		jwtSecret := []byte(cfg.JWT.Secret)
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessToken, err := jwtToken.SignedString(jwtSecret)
		if err != nil {
			return responseLogin, result, errors.Wrap(err, "[SERVICE][LoginAgent][SignedString]")
		}
		// 8️⃣ Construct token response
		token = auth.Token{
			AccessToken:         accessToken,
			ExpiresIn:           expiration.Unix() - now.Unix(),
			ExpiresAt:           expiration.Unix(),
			TokenType:           "Bearer",
			ForceChangePassword: 0,
		}
		responseLogin = agentEntity.ResponseLogin{
			AgentID:    agentLogin.AgentID,
			AgentName:  agentResult.AgentName,
			AgentToken: token.TokenType + " " + accessToken,
		}
	}
	// return responseLogin, err
	return responseLogin, result, err
}

func (s Service) CheckAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (string, error) {
	var (
		result string
		err    error
	)
	_, result, err = s.agent.CheckAgent(ctx, agentLogin)
	if err != nil {
		return result, errors.Wrap(err, "[Service][LoginAgent]")
	}
	return result, err
}

func (s Service) AddAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (agentEntity.AgentResp, string, error) {
	var (
		result        agentEntity.AgentResp
		agentResponse string
		err           error
	)
	agentResponse, err = s.agent.AddAgent(ctx, agentLogin)
	if err != nil {
		return result, agentResponse, errors.Wrap(err, "[Service][LoginAgent]")
	}
	if agentResponse != "Success" {
		agentLogin = agentEntity.AgentRequest{}
	}
	if agentResponse == "Success" {
		result = agentEntity.AgentResp{
			AgentID:   agentLogin.AgentID,
			AgentName: agentLogin.AgentName,
		}
	}
	return result, agentResponse, err
}

func (s Service) DeleteAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (agentEntity.AgentResp, string, error) {
	var (
		result        agentEntity.AgentResp
		agentData     agentEntity.Agent
		agentResponse string
		err           error
	)
	fmt.Println("MASOKDELETE")
	agentData, agentResponse, err = s.agent.DeleteAgent(ctx, agentLogin)
	if err != nil {
		return result, agentResponse, errors.Wrap(err, "[Service][LoginAgent]")
	}
	if agentResponse != "Success" {
		agentLogin = agentEntity.LoginAgent{}
	}
	if agentResponse == "Success" {
		result = agentEntity.AgentResp{
			AgentID:   agentLogin.AgentID,
			AgentName: agentData.AgentName,
		}
	}
	return result, agentResponse, err
}

func (s Service) UpdateAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (agentEntity.AgentResp, string, error) {
	var (
		result        agentEntity.AgentResp
		agentResponse string
		err           error
	)
	_, agentResponse, err = s.agent.UpdateAgent(ctx, agentLogin)
	if err != nil {
		return result, agentResponse, errors.Wrap(err, "[Service][LoginAgent]")
	}
	if agentResponse != "Success" {
		agentLogin = agentEntity.AgentRequest{}
	}
	if agentResponse == "Success" {
		result = agentEntity.AgentResp{
			AgentID:   agentLogin.AgentID,
			AgentName: agentLogin.AgentName,
		}
	}
	fmt.Println("test", agentResponse)
	return result, agentResponse, err
}
