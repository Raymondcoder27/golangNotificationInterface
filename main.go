package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Account struct {
	Username string
	Password string
}

func handleCreateAccount(w http.ResponseWriter, r http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("Failed to decode the response body", "err", err)
		return
	}

	if err := notifyAccountCreated(account); err != nil {
		slog.Error("Failed to notify account created", "err", err)
		return
	}
}
