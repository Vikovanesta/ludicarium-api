package server

import (
	"log"
	"net/http"

	"github.com/vikovanesta/ludicarium-api/igdb"
)

var igdbClient = igdb.NewClient(nil)

func (s *Server) getGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var games []*igdb.Game
		var err error

		games, err = igdbClient.Games.Index(igdb.SetLimit(10), igdb.SetFields("*"))

		if err != nil {
			writeJson(w, http.StatusInternalServerError, M{"error": err.Error()})
			serverError(w, err)
			return
		}

		log.Println(games)
		writeJson(w, http.StatusOK, M{"games": games})
	}
}
