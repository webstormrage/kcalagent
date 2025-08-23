package service

import (
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResp struct {
	Token string `json:"token"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := appContext.Get()
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.Login == "" || req.Password == "" {
		http.Error(w, "login and password required", http.StatusBadRequest)
		return
	}

	usr, err := kcaldb.GetUserByLogin(req.Login)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := IssueJWT(usr.Login)
	if err != nil {
		http.Error(w, "failed to issue token", http.StatusInternalServerError)
		ctx.Logger.Printf("[error]: failed to issue token %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(loginResp{Token: token})
}
