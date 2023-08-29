/*
Opalstack API

Testing QuarantinedmailAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package OpalStack

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/vlazic/opalstack-go"
)

func Test_OpalStack_QuarantinedmailAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QuarantinedmailAPIService QuarantinedmailDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.QuarantinedmailAPI.QuarantinedmailDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuarantinedmailAPIService QuarantinedmailList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.QuarantinedmailAPI.QuarantinedmailList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuarantinedmailAPIService QuarantinedmailRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.QuarantinedmailAPI.QuarantinedmailRead(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuarantinedmailAPIService QuarantinedmailSubmit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.QuarantinedmailAPI.QuarantinedmailSubmit(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
