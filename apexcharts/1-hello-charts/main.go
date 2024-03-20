package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/seven-to-learn/apexcharts/1-hello-charts/templates"
)

func main() {
	component := templates.Hello()
	http.Handle("/", templ.Handler(component))
	http.ListenAndServe(":3000", nil)
}
