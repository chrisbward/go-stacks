package stacksblockchainapi

import (
	"context"
	"fmt"
	"stacksblockchainapi/errors"
	"stacksblockchainapi/models"

	"github.com/apimatic/go-core-runtime/utilities"
)

// SearchController represents a controller struct.
type SearchController struct {
	baseController
}

// NewSearchController creates a new instance of SearchController.
// It takes a baseController as a parameter and returns a pointer to the SearchController.
func NewSearchController(baseController baseController) *SearchController {
	searchController := SearchController{baseController: baseController}
	return &searchController
}

// SearchById takes context, id, includeMetadata as parameters and
// returns an models.ApiResponse with models.SearchResult data and
// an error if there was an issue with the request or response.
// Search blocks, transactions, contracts, or accounts by hash/ID
func (s *SearchController) SearchById(
	ctx context.Context,
	id string,
	includeMetadata *bool) (
	models.ApiResponse[models.SearchResult],
	error) {
	req := s.prepareRequest(ctx, "GET", fmt.Sprintf("/extended/v1/search/%v", id))
	req.Authenticate(true)
	if includeMetadata != nil {
		req.QueryParam("include_metadata", *includeMetadata)
	}

	var result models.SearchResult
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SearchResult](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not found")
	}
	return models.NewApiResponse(result, resp), err
}
