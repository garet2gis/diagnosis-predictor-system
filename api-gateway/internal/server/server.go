package server

import (
	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/pkg/logger"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	router chi.Router
	cfg    *config.Config
	logger *logrus.Logger
}

func NewApp(cfg *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

// Start конфигурирует роутер приложения, внедряя в него хэндлеры.
// После чего запускает сервер.
func (s *Server) Start() {
	s.InitRouter()

	s.startServerGracefully(uint(s.cfg.GracefulShutdownTimeoutSeconds), s.router)
}

func (s *Server) InitRouter() {
	router := chi.NewRouter()

	router.Use(middleware.InitRequestID)
	router.Use(logger.WithLogger(s.logger))

	router.Route("/api/v1", func(router chi.Router) {
		router.Post("/some_ai_model", middleware.ErrorMiddleware(s.TestModel))
	})

	s.router = router
}

// startServerGracefully запускает сервер с параметрами конфига.
// timeout - время для завершения graceful shutdown
func (s *Server) startServerGracefully(timeout uint, router chi.Router) {
	addr := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Duration(s.cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(s.cfg.WriteTimeoutSeconds) * time.Second,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, time.Duration(timeout)*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				s.logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			s.logger.Fatal(err)
		}
		s.logger.Info("server shutdown gracefully")
		serverStopCtx()
	}()

	s.logger.Infof("start server on %s", addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatal(err)
	}

	<-serverCtx.Done()
}
