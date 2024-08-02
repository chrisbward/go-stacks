package api

import (
	"context"
	"testing"

	config "github.com/chrisbward/go-stacks/stacks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_MempoolAPIService(t *testing.T) {

	configuration := config.NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test MempoolAPIService GetMempoolFeePriorities", func(t *testing.T) {
		resp, httpRes, _ := apiClient.MempoolAPI.GetMempoolFeePriorities(context.Background()).Execute()

		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
