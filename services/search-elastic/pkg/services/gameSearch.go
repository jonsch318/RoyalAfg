package services

import (
	"bytes"
	"context"
	"encoding/json"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/services/search-elastic/pkg/dto"
)

type GameSearch struct {
	logger *zap.SugaredLogger
	//Only for testing purposes
	//IndexedGames []string
	//Games map[string]dto.GameResult
	es *elasticsearch.Client
}

func NewGameSearch(logger *zap.SugaredLogger, es *elasticsearch.Client) *GameSearch {
	return &GameSearch{
		logger: logger,
		//IndexedGames: LoadExampleDbIndexes(),
		//Games: LoadExampleDb(),
		es: es,
	}
}

func (s *GameSearch) SearchGames(query string) []dto.GameResult {
	/*	for _, i := range s.IndexedGames {
		if strings.HasPrefix(i, query){
			results = append(results, s.Games[i])
		}
	}*/

	buf, err := s.buildQuery(query)

	if err != nil {
		s.logger.Errorw("Error during query initialization", "error", err)
		return nil
	}

	res, err := s.es.Search(
		s.es.Search.WithContext(context.Background()),
		s.es.Search.WithIndex("games"),
		s.es.Search.WithBody(buf),
		s.es.Search.WithTrackTotalHits(true),
		s.es.Search.WithPretty())

	if err != nil {
		s.logger.Errorw("Error during search initialization", "error", err)
		return nil
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			s.logger.Errorw("Error parsing the response body", "error", err)
		} else {
			// Print the response status and error information.
			s.logger.Errorw("Error during search",
				"status", res.Status(),
				"error_type", e["error"].(map[string]interface{})["type"],
				"error", e["error"].(map[string]interface{})["reason"],
			)
		}
		return nil
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		s.logger.Errorw("Error parsing the response body", "error", err)
		return nil
	}

	suggestions := r["suggest"].(map[string]interface{})["autocomplete"].([]interface{})[0].(map[string]interface{})["options"].([]interface{})

	results := make([]dto.GameResult, len(suggestions))

	// Print the ID and document source for each hit.
	for i, hit := range suggestions {
		suggestion := hit.(map[string]interface{})["_source"].(map[string]interface{})
		name := suggestion["name"].(string)
		url := suggestion["url"].(string)
		maxPlayers := int(suggestion["maxPlayers"].(float64))
		results[i] = dto.GameResult{Name: name, URL: url, MaxPlayers: maxPlayers}
	}

	return results
}

func (s *GameSearch) buildQuery(prefix string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"suggest": map[string]interface{}{
			"autocomplete": map[string]interface{}{
				"prefix": prefix,
				"completion": map[string]interface{}{
					"field": "suggest",
					"fuzzy": map[string]interface{}{
						"fuzziness": "auto",
					},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}
	return &buf, nil
}
