package jasmine

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Address string
	Routes  Routes
	Muxer   *http.ServeMux

	Logger *zerolog.Logger

	ctx   context.Context
	close context.CancelCauseFunc
}

func (s *Server) Start() {
	// Make sure logger is set
	if s.Logger == nil {
		s.Logger = &log.Logger
	}
	// Contexts for graceful shutdown
	s.ctx, s.close = context.WithCancelCause(context.Background())

	// Set the Muxer
	if s.Muxer == nil {
		s.Muxer = http.NewServeMux()
	}

	// Create the HTTP server
	server := &http.Server{
		Addr:    s.Address,
		Handler: s.Muxer,
	}

	// Register routes to the muxer
	s.Routes.RegisterRoutes(s.Muxer, s.Logger)

	s.Logger.Info().Msgf("Starting Server on %s", server.Addr)

	// Serve
	if err := server.ListenAndServe(); err != nil {
		s.Logger.Fatal().Err(err).Msgf("Failed to start server on %s", server.Addr)
	}
}

func (s *Server) Stop() {
	s.close(fmt.Errorf("Server told to stop"))
	s.Logger.Info().Msg("Server stopped")
}