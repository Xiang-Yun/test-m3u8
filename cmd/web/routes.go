package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // 允許所有來源
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin"},
		AllowCredentials: false,
		MaxAge:           300, // 預檢請求的結果可以被瀏覽器緩存多久（以秒為單位）
	}
	mux.Use(cors.New(corsOptions).Handler)

	mux.Get("/virtual-terminal", app.VirtualTerminal)

	mux.Route("/Media", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				next.ServeHTTP(w, r)
			})
		})
		r.Handle("/*", http.StripPrefix("/Media", http.FileServer(http.Dir("./Media"))))
	})

	return mux
}
