package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type RateLimiterMiddleware struct {
	client map[string]*clientState
	mu     sync.Mutex
	limit  int
	window time.Duration
}

type clientState struct {
	requests []time.Time
	mu       sync.Mutex
}

func NewRateLimiterMiddleware(limit int, window time.Duration) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		client: make(map[string]*clientState),
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiterMiddleware) GetClientState(ip string) *clientState {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if client, ok := rl.client[ip]; ok {
		return client
	}

	client := &clientState{
		requests: make([]time.Time, 0),
	}

	rl.client[ip] = client

	return client
}

func (rl *RateLimiterMiddleware) Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		state := rl.GetClientState(ip)

		state.mu.Lock()
		defer state.mu.Unlock()

		now := time.Now()
		for len(state.requests) > 0 && now.Sub(state.requests[0]) > rl.window {
			state.requests = state.requests[1:]
		}

		if len(state.requests) >= rl.limit {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}

		state.requests = append(state.requests, now)
		next.ServeHTTP(w, r)
	})
}

func RateLimiterMiddlewareMain() {

	rateLimiter := NewRateLimiterMiddleware(10, time.Minute)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Value string `json:"value"`
		}{
			Value: "Server response",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(data)

		//fmt.Fprintf(w, "Hello world")
	})

	serve := http.Server{
		Addr:    ":8080",
		Handler: rateLimiter.Middleware(mux),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Serve start in PORT 8080")
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error in open server %v", err)
		}
	}()

	<-stop
	fmt.Println("Notify signal stop")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := serve.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
