package simpleserver

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleServer(t *testing.T) {
	t.Run("test get APIServer", func(t *testing.T) {
		jsonBody := []byte(`{"client_message": "hello, server!"}`)
		bodyReader := bytes.NewReader(jsonBody)

		request, _ := http.NewRequest(http.MethodPost,
			"/checkmarx",
			bodyReader)
		response := httptest.NewRecorder()

		handler(response, request)

		res := response.Result()
		defer res.Body.Close()

		data, _ := io.ReadAll(res.Body)
		fmt.Println("response: ", string(data))
	})

}
