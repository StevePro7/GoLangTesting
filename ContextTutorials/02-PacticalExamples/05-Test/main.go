package main

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type key int

const (
	requestID key = iota
	jwt
)

func isDatabaseUp(ctx context.Context) (bool, error) {
	// retrieve the request ID value
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return false, fmt.Errorf("requestID in context does not have the expected type")
	}
	log.Printf("req %s - checking db status", reqID)
	return true, nil
}

func isMonitoringUp(ctx context.Context) (bool, error) {
	// retrieve the request ID value
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return false, fmt.Errorf("requestID in context does not have the expected type")
	}
	log.Printf("req %s - checking monitoring status", reqID)
	return true, nil
}

func status(w http.ResponseWriter, req *http.Request) {
	// Add request id to context
	ctx := context.WithValue(req.Context(), requestID, uuid.NewV4().String())
	// Add credentials to context
	ctx = context.WithValue(ctx, jwt, req.Header.Get("Authorization"))

	upDB, err := isDatabaseUp(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	upAuth, err := isMonitoringUp(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "DB up: %t | Monitor up %t\n", upDB, upAuth)
	if err != nil {
		return
	}
}
