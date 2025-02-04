package auth

import (
	"encoding/json"
	"fmt"
	"github.com/maximegorov13/go-api/configs"
	"github.com/maximegorov13/go-api/pkg/res"
	"net/http"
	"regexp"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), http.StatusPaymentRequired)
			return
		}
		if payload.Email == "" {
			res.Json(w, "Email required", http.StatusPaymentRequired)
			return
		}
		match, _ := regexp.MatchString(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`, payload.Email)
		if !match {
			res.Json(w, "Wrong email", http.StatusPaymentRequired)
			return
		}
		if payload.Password == "" {
			res.Json(w, "Password required", http.StatusPaymentRequired)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
