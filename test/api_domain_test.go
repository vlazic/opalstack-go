/*
Opalstack API

Testing DomainAPIService

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

func Test_OpalStack_DomainAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainAPIService DomainCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DomainAPI.DomainCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainAPIService DomainDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DomainAPI.DomainDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainAPIService DomainList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DomainAPI.DomainList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainAPIService DomainRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.DomainAPI.DomainRead(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
