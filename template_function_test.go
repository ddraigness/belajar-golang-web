package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Surya" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Alqo",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	stringBody := string(body)

	fmt.Println(stringBody)
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", map[string]interface{}{
		"Name": "Alqo",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	stringBody := string(body)

	fmt.Println(stringBody)
}

// menambahkan global function
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Alqo Mulky",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	stringBody := string(body)

	fmt.Println(stringBody)
}

func TemplateFunctionCreateGlobalPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Alqo",
	})

}

func TestTemplateFunctionCreateGlobalPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipelines(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	stringBody := string(body)

	fmt.Println(stringBody)
}