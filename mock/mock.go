/*
 * Copyright (C) 2026 Holger de Carne
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package mock provides a mock implementation of the Paperless NGX REST API for testing.
package mock

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/tdrn-org/go-paperless-ngx/api"
)

// APIKey defines the API Key used to authorized towards a mock server.
const APIKey string = "paperlessngx"

// Server represents a mock instance.
type Server struct {
	httpListener net.Listener
	apiURL       *url.URL
	logger       *slog.Logger
	stoppedWG    sync.WaitGroup
	httpServer   *http.Server
}

// Start starts and returns a new mock instance.
//
// Start panics in case of an error. The returned server
// is listening on localhost using a dynamic port.
func Start() *Server {
	httpListener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	address := httpListener.Addr().String()
	apiURL, err := url.Parse(fmt.Sprintf("http://%s", address))
	logger := slog.Default().With(slog.String("server", address))
	if err != nil {
		log.Fatal(err)
	}
	server := &Server{
		httpListener: httpListener,
		apiURL:       apiURL,
		logger:       logger,
	}
	server.setupHttpServer()
	server.stoppedWG.Go(server.listenAndServe)
	return server
}

// APIURL gets the API URL for this mock instance.
func (s *Server) APIURL() *url.URL {
	return s.apiURL
}

// Stop stops this mock instance.
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
	s.stoppedWG.Wait()
}

func (s *Server) setupHttpServer() {
	strictHandlerMiddlewares := []api.StrictMiddlewareFunc{
		s.logOperationMiddleware,
	}
	strictHandlerOptions := api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  s.errorHandler,
		ResponseErrorHandlerFunc: s.errorHandler,
	}
	strictHandler := api.NewStrictHandlerWithOptions(s, strictHandlerMiddlewares, strictHandlerOptions)
	handlerMiddlewares := []api.MiddlewareFunc{
		s.checkAuthorization,
	}
	handlerOptions := api.ChiServerOptions{
		BaseURL:          "",
		Middlewares:      handlerMiddlewares,
		ErrorHandlerFunc: s.errorHandler,
	}
	handler := api.HandlerWithOptions(strictHandler, handlerOptions)
	s.httpServer = &http.Server{
		Handler: handler,
	}
}

func (s *Server) logOperationMiddleware(f api.StrictHandlerFunc, operationID string) api.StrictHandlerFunc {
	s.logger.Info("mock call", slog.String("operation", operationID))
	return f
}

func (s *Server) checkAuthorization(handler http.Handler) http.Handler {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get(api.AuthorizationHeader)
		if authorization == "Token "+APIKey {
			handler.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
		}
	}
	return http.HandlerFunc(handlerFunc)
}

func (s *Server) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	s.logger.Error("mock handler failure", slog.String("url", r.URL.String()), slog.Any("err", err))
	w.WriteHeader(http.StatusInternalServerError)
}

func (s *Server) listenAndServe() {
	s.logger.Info("http server starting...")
	err := s.httpServer.Serve(s.httpListener)
	if !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("http server failure", slog.Any("err", err))
		return
	}
	s.logger.Info("http server stopped")
}
