package httpserver

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lgustavopalmieri/comex-ease/adapters/web/ncm/handlers"
	"github.com/lgustavopalmieri/comex-ease/internal/domain/ncm/usecase"
)

type WebServer struct {
	NcmUsecase usecase.NcmUsecaseInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handlers.MakeNcmHandlers(r, w.NcmUsecase)

	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":5500",
		Handler:           r,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
