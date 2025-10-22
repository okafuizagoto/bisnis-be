package bisnis

import (
	agentEntity "bisnis-be/internal/entity/agent"
	// bisnisEntity "bisnis-be/internal/entity/bisnis"
	"bisnis-be/pkg/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	auth "bisnis-be/internal/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// Getgoldgym godoc
// @Summary Get entries of all goldgyms
// @Description Get entries of all goldgyms
// @Tags bisnis
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 2"00"
// @Router /v1/profiles [get]
func (h *Handler) InsertGoldGym(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error

		loginAgent agentEntity.LoginAgent
		addAgent   agentEntity.AgentRequest
		// addTransaction bisnisEntity.AddTransaction
		respLogin string
		resp      response.Response
		types     string
	)
	defer resp.RenderJSON(w, r)

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("Getgoldgym", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Your code here
	types = r.FormValue("type")
	switch types {
	case "loginagent":
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &loginAgent)
		result, respLogin, err = h.agentSvc.LoginAgent(ctx, loginAgent)
		// 	if err != nil {
		// 		log.Println("err", err)
		// 	}
		if err != nil || respLogin != "Success" {
			// resp.SetError(err, http.StatusInternalServerError)
			// resp.StatusCode = 5"00"
			// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			// return
			resp.Msg = respLogin
			resp.Code = "01"
			// resp.Status = true
			resp.StatusCode = http.StatusNotImplemented // 5"01"
			return
		}
		if respLogin == "Success" {
			resp.Msg = respLogin
			resp.Code = "00"
		}
	case "addagent":
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp.Msg = "Missing Authorization header"
			resp.Code = "01"
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			resp.Status = false
			resp.Msg = "Invalid token format"
			resp.Code = "01"
			return
		}

		tokenStr := parts[1]

		// --- Validasi token ---
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			resp.Status = false
			resp.Msg = "Invalid token: " + err.Error()
			resp.Code = "01"
			return
		}

		userID := claims["user"].(string)
		fmt.Println("userID", userID)
		agentStruct := agentEntity.LoginAgent{
			AgentID:       userID,
			AgentPassword: "",
		}
		validation, err := h.agentSvc.CheckAgent(ctx, agentStruct)
		if err != nil || (validation != "Success" && validation != "Incorrect password") {
			// resp.SetError(err, http.StatusInternalServerError)
			// resp.StatusCode = 5"00"
			// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			// return
			resp.Msg = validation
			resp.Code = "01"
			// resp.Status = true
			resp.StatusCode = http.StatusNotImplemented // 5"01"
			return
		}
		if validation == "Success" || validation == "Incorrect password" {
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &addAgent)
			result, respLogin, err = h.agentSvc.AddAgent(ctx, addAgent)
			if err != nil || respLogin != "Success" {
				// resp.SetError(err, http.StatusInternalServerError)
				// resp.StatusCode = 5"00"
				// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
				// return
				resp.Msg = respLogin
				resp.Code = "01"
				// resp.Status = true
				resp.StatusCode = http.StatusNotImplemented // 5"01"
				return
			}
			if respLogin == "Success" {
				resp.Msg = respLogin
				resp.Code = "00"
			}
		}
	case "updateagent":
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp.Msg = "Missing Authorization header"
			resp.Code = "01"
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			resp.Status = false
			resp.Msg = "Invalid token format"
			resp.Code = "01"
			return
		}

		tokenStr := parts[1]

		// --- Validasi token ---
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			resp.Status = false
			resp.Msg = "Invalid token: " + err.Error()
			resp.Code = "01"
			return
		}

		userID := claims["user"].(string)
		fmt.Println("userID", userID)
		agentStruct := agentEntity.LoginAgent{
			AgentID:       userID,
			AgentPassword: "",
		}
		validation, err := h.agentSvc.CheckAgent(ctx, agentStruct)
		if err != nil || (validation != "Success" && validation != "Incorrect password") {
			// resp.SetError(err, http.StatusInternalServerError)
			// resp.StatusCode = 5"00"
			// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			// return
			resp.Msg = validation
			resp.Code = "01"
			// resp.Status = true
			resp.StatusCode = http.StatusNotImplemented // 5"01"
			return
		}
		if validation == "Success" || validation == "Incorrect password" {
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &addAgent)
			result, respLogin, err = h.agentSvc.UpdateAgent(ctx, addAgent)
			if err != nil || respLogin != "Success" {
				// resp.SetError(err, http.StatusInternalServerError)
				// resp.StatusCode = 5"00"
				// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
				// return
				resp.Msg = respLogin
				resp.Code = "01"
				// resp.Status = true
				resp.StatusCode = http.StatusNotImplemented // 5"01"
				return
			}
			if respLogin == "Success" {
				resp.Msg = respLogin
				resp.Code = "00"
			}
		}
		// case "addtransaction":
		// 	authHeader := r.Header.Get("Authorization")
		// 	if authHeader == "" {
		// 		resp.Msg = "Missing Authorization header"
		// 		resp.Code = "01"
		// 		return
		// 	}
		// 	parts := strings.Split(authHeader, " ")
		// 	if len(parts) != 2 || parts[0] != "Bearer" {
		// 		resp.Status = false
		// 		resp.Msg = "Invalid token format"
		// 		resp.Code = "01"
		// 		return
		// 	}

		// 	tokenStr := parts[1]

		// 	// --- Validasi token ---
		// 	claims, err := auth.ValidateJWT(tokenStr)
		// 	if err != nil {
		// 		resp.Status = false
		// 		resp.Msg = "Invalid token: " + err.Error()
		// 		resp.Code = "01"
		// 		return
		// 	}

		// 	userID := claims["user"].(string)
		// 	fmt.Println("userID", userID)
		// 	agentStruct := agentEntity.LoginAgent{
		// 		AgentID:       userID,
		// 		AgentPassword: "",
		// 	}
		// 	validation, err := h.agentSvc.CheckAgent(ctx, agentStruct)
		// 	if err != nil || (validation != "Success" && validation != "Incorrect password") {
		// 		// resp.SetError(err, http.StatusInternalServerError)
		// 		// resp.StatusCode = 5"00"
		// 		// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		// 		// return
		// 		resp.Msg = validation
		// 		resp.Code = "01"
		// 		// resp.Status = true
		// 		resp.StatusCode = http.StatusNotImplemented // 5"01"
		// 		return
		// 	}
		// 	if validation == "Success" || validation == "Incorrect password" {
		// 		body, _ := ioutil.ReadAll(r.Body)
		// 		json.Unmarshal(body, &addTransaction)
		// 		result, respLogin, err = h.bisnisSvc.AddTransaction(ctx, addTransaction)
		// 		if err != nil || respLogin != "Success" {
		// 			// resp.SetError(err, http.StatusInternalServerError)
		// 			// resp.StatusCode = 5"00"
		// 			// log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		// 			// return
		// 			resp.Msg = respLogin
		// 			resp.Code = "01"
		// 			// resp.Status = true
		// 			resp.StatusCode = http.StatusNotImplemented // 5"01"
		// 			return
		// 		}
		// 		if respLogin == "Success" {
		// 			resp.Msg = respLogin
		// 			resp.Code = "00"
		// 		}
		// 	}
	}

	if types != "loginagent" {
		if err != nil {
			resp.SetError(err, http.StatusInternalServerError)
			resp.StatusCode = 01
			log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			return
		}
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	return
}

// ValidateToken menerima header Authorization, mengembalikan claims jika valid
func ValidateToken(authHeader string) (jwt.MapClaims, error) {
	fmt.Println("MASOK-1")
	if authHeader == "" {
		return nil, fmt.Errorf("missing Authorization header")
	}
	fmt.Println("MASOK-2")

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid token format")
	}
	fmt.Println("MASOK-3")

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("MASOK-3-1")

			return nil, fmt.Errorf("invalid signing method")
		}
		fmt.Println("MASOK-3-2")
		// return []byte(os.Getenv("TOKEN_SECRET")), nil
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	fmt.Println("TOKENERR", token)
	if err != nil {
		fmt.Println("MASOK-3-3", err)
		return nil, err
	}
	fmt.Println("MASOK-4")

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	fmt.Println("MASOK-5")

	return claims, nil
}
