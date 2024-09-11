package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Account struct {
	Username string
	Password string
}

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type SimpleAccountNotifier struct {
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("Failed to decode the response body", "err", err)
		return
	}

	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		slog.Error("Failed to notify account created", "err", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func handleCreateAccount(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func notifyAccountCreated(account Account) error {
	time.Sleep(time.Millisecond * 500)
	slog.Info("New account created,", "username", account.Username, "email", account.email)
	return nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /account", handleCreateAccount)
	http.ListenAndServe(":3000", mux)
}
