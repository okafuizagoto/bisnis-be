package bisnis

import (
	auth "bisnis-be/internal/auth"
	httpHelper "bisnis-be/internal/delivery/http"
	agentEntity "bisnis-be/internal/entity/agent"
	"bisnis-be/pkg/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
func (h *Handler) DeleteGoldGym(w http.ResponseWriter, r *http.Request) {
	var (
		result      interface{}
		metadata    interface{}
		deleteAgent agentEntity.LoginAgent
		respLogin   string
		err         error
		resp        response.Response
		types       string
	)
	defer resp.RenderJSON(w, r)

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("Getgoldgym", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	types = r.FormValue("type")
	switch types {
	case "deleteagent":
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
			json.Unmarshal(body, &deleteAgent)
			result, respLogin, err = h.agentSvc.DeleteAgent(ctx, deleteAgent)
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
	}

	if types != "deleteagent" {
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			//
			log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
			h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
			return
		}
	}
	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	return
}
