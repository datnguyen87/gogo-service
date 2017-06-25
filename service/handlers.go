package service

import (
	"github.com/unrolled/render"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/cloudnativego/gogo-engine"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		err := json.Unmarshal(payload, &newMatchRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse create match request")
			return
		}

		if !newMatchRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new match request")
			return
		}

		newMatch := gogo.NewMatch(newMatchRequest.GridSize, newMatchRequest.PlayerBlack, newMatchRequest.PlayerWhite)
		repo.addMatch(newMatch)
		w.Header().Add("location", "/matches/" + newMatch.ID)
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{
			ID: newMatch.ID,
			GridSize: newMatch.GridSize,
			PlayerBlack: newMatch.PlayerBlack,
			PlayerWhite: newMatch.PlayerWhite,
		})
	}
}