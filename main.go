package main

import (
	"backend/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "Total number of requests processed by the API",
	}, []string{"path", "status"})

	HttpRequestErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_error_total",
		Help: "Total number of errors returned by the API",
	}, []string{"path", "status"})
)

var customRegistry = prometheus.NewRegistry()

func init() {
	customRegistry.MustRegister(httpRequestTotal, HttpRequestErrorTotal)
}

func main() {
	r := chi.NewRouter()

	r.Use(RequestMetricsMiddleware)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("not able to load .env file")
		return
	}

	r.Get("/metrics", promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{}).ServeHTTP)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	    response:= map[string]interface{}{
			"status": 201,
			"PORT": cfg.PORT,
			"secret": cfg.SECRET,
		}

		json.NewEncoder(w).Encode(response)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Printf("server started at 8000")
	PORTstr := fmt.Sprintf(":%s", cfg.PORT)
	http.ListenAndServe(PORTstr, r)
}

func PrometheusHandler() http.HandlerFunc {
	h := promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})
	return h.ServeHTTP
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func RequestMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		status := wrapped.statusCode
		if status < 400 {
			httpRequestTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		} else {
			HttpRequestErrorTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		}
	})
}
