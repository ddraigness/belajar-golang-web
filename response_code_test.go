package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(400) // BadRequest
		fmt.Fprint(writer, "name is empty")
	} else {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(stringBody)
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000?name=Alqo", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(stringBody)
}