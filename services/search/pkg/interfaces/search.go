package interfaces

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/dto"
)

type SearchService interface {
	Search(query string) dto.SearchResult
}
