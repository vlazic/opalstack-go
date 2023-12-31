/*
Opalstack API

Testing DnsrecordAPIService

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

func Test_OpalStack_DnsrecordAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DnsrecordAPIService DnsrecordCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DnsrecordAPI.DnsrecordCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsrecordAPIService DnsrecordDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DnsrecordAPI.DnsrecordDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsrecordAPIService DnsrecordList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DnsrecordAPI.DnsrecordList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsrecordAPIService DnsrecordRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.DnsrecordAPI.DnsrecordRead(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsrecordAPIService DnsrecordUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DnsrecordAPI.DnsrecordUpdate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
