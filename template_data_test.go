package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Alqo",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada Lagi",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	
	body, _ := io.ReadAll(recorder.Result().Body)
	stringBody := string(body)

	fmt.Println(stringBody)
}

type Address struct {
	Street string
}

type Page struct {
	Title string
	Name string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Alqo",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	
	body, _ := io.ReadAll(recorder.Result().Body)
	stringBody := string(body)

	fmt.Println(stringBody)
}