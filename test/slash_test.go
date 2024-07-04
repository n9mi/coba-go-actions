package test

import (
	"coba-go-actions/model"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestSlash(t *testing.T) {
	url := "/"

	fmt.Println(url)

	request := newRequest(
		fiber.MethodGet,
		url,
		"",
	)

	response, err := app.Test(request)
	require.Nil(t, err)
	require.Equal(t, fiber.StatusOK, response.StatusCode)

	responseBody, err := io.ReadAll(response.Body)
	require.Nil(t, err)

	testResponse := new(model.DataResponse)
	err = json.Unmarshal(responseBody, testResponse)
	require.Nil(t, err)

	require.Equal(t, "OK", testResponse.Data)
}
