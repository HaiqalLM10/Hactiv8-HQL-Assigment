package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type contextData string

const (
	CTX_TOTAL contextData = "TOTAL"
	CTX_LOG   contextData = "LOG"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", trace(middleware(middleware2(HelloHandler))))

	log.Println("server running at port :4444")

	http.ListenAndServe(":4444", mux)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("from handler")
	logData := r.Context().Value(CTX_LOG).(map[string]interface{})

	if true {
		delete(logData, "message")
		logData["error"] = errors.New("bad request deh").Error()
		logByte, _ := json.Marshal(logData)
		log.Println(string(logByte))
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"payload": r.Context().Value(CTX_TOTAL),
	})
}

func trace(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logData = map[string]interface{}{
			"method":   r.Method,
			"url":      r.URL.Path,
			"message":  "incoming request",
			"trace_id": uuid.New(),
		}
		now := time.Now()

		logByte, _ := json.Marshal(logData)
		log.Println(string(logByte))

		ctx := r.Context()
		ctx = context.WithValue(ctx, CTX_LOG, logData)

		r = r.WithContext(ctx)

		next(w, r)
		end := time.Since(now).Seconds()
		logData["response_time"] = end
		logData["message"] = "finish request"

		logByte, _ = json.Marshal(logData)
		log.Println(string(logByte))
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before request")
		val := 10

		ctx := r.Context()
		ctx = context.WithValue(ctx, CTX_TOTAL, val)

		r = r.WithContext(ctx)

		next(w, r)
		log.Println("after request")
	}
}
func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before request middleware 2")
		ctx := r.Context()

		val := ctx.Value(CTX_TOTAL)
		log.Println("got value from middleware 1 :", val)

		valInt, err := strconv.Atoi(fmt.Sprintf("%v", val))
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"error":   err.Error(),
			})
			return
		}

		valInt *= 5

		ctx = context.WithValue(ctx, CTX_TOTAL, valInt)
		r = r.WithContext(ctx)

		next(w, r)
		log.Println("after request middleware 2")

		var key ConfigKey = "hello"
		Config(key)
	}
}

type ConfigKey string

func (c ConfigKey) toString() string {
	return string(c)
}

func Config(key ConfigKey) string {
	if key == "" {
		return "default"
	}

	return os.Getenv(key.toString())
}
