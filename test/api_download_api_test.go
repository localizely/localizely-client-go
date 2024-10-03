/*
Localizely API

Testing DownloadAPIAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package localizely

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/localizely/localizely-client-go"
)

func Test_localizely_DownloadAPIAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DownloadAPIAPIService GetLocalizationFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		httpRes, err := apiClient.DownloadAPIAPI.GetLocalizationFile(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
