package main

import (
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"ai-kcal-agent/pkg/service"
	"errors"
	"log"
	"net"
	"net/http"
	"time"
)

func logMiddleware(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Printf("[http]: %s %s %s %d\n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start).Milliseconds())
	})
}

func runServer() {
	ctx := appContext.Get()
	err := kcaldb.SetupDb()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/web/", service.WebHandler("./frontend/dist/", "/web/"))

	mux.HandleFunc("/auth/login", service.HandleLogin)
	mux.HandleFunc("/get-daily-summary", service.GetDailyReportHandler)
	mux.HandleFunc("/meals/add", service.AddMealHandler)
	server := &http.Server{
		Addr:         ":" + ctx.ServerPort,
		Handler:      logMiddleware(ctx.Logger, mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}
	// ЛОГ РЕАЛЬНО ПОСЛЕ УСПЕШНОГО БИНДИНГА
	ctx.Logger.Printf("HTTP server listening on %s", ln.Addr())

	// Блокирующий запуск
	if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
		ctx.Logger.Fatalf("server error: %v", err)
	}
}
