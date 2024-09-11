package main

import "net/http"

type Account struct {
	Username string
	Password string
}

func handleCreateAccount(w http.ResponseWriter, r http.Request) {

}
