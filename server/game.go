package server

import (
	"log"
	"net/http"

	"github.com/vikovanesta/ludicarium-api/igdb"
)

func (s *Server) getGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var games []*igdb.Game
		var err error

		games, err = s.gameService.List()

		if err != nil {
			writeJson(w, http.StatusInternalServerError, M{"error": err.Error()})
			serverError(w, err)
			return
		}

		log.Println(games)
		writeJson(w, http.StatusOK, M{"games": games})
	}
}
