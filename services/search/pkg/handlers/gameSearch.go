package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (g Game) GameSearch(rw http.ResponseWriter, r *http.Request)  {
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	vars := r.URL.Query()
	query := vars.Get("q")
	results := g.searchService.SearchGames(query)

	encoder := json.NewEncoder(rw)
	err := encoder.Encode(results)

	if err != nil {
		g.logger.Errorw("Error during Encoding", "error", err)
		http.Error(rw, fmt.Sprintf("{error: %v}", "a error occured during the search."), http.StatusInternalServerError)
		return
	}
}
