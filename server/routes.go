package server

import (
	"github.com/rs/cors"
)

const (
	MustAuth     = true
	OptionalAuth = false
)

func (s *Server) routes() {
	s.router.Use(cors.AllowAll().Handler)
	// s.router.Use(Logger(os.Stdout))
	apiRouter := s.router.PathPrefix("/api/v1").Subrouter()

	// optionalAuth := apiRouter.PathPrefix("").Subrouter()
	// optionalAuth.Use(s.authenticate(OptionalAuth))
	// {
	// optionalAuth.Handle("/profiles/{username}", s.getProfile()).Methods("GET")
	// optionalAuth.Handle("/tags", s.listTags()).Methods("GET")
	// }

	noAuth := apiRouter.PathPrefix("").Subrouter()
	{
		noAuth.Handle("/health", healthCheck())
		noAuth.Handle("/games", s.getGames()).Methods("GET")
	}

	// authApiRoutes := apiRouter.PathPrefix("").Subrouter()
	// authApiRoutes.Use(s.authenticate(MustAuth))
	// {
	// authApiRoutes.Handle("/user", s.getCurrentUser()).Methods("GET")
	// authApiRoutes.Handle("/user", s.updateUser()).Methods("PUT", "PATCH")
	// }
}
