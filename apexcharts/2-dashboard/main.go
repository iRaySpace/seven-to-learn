package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/irayspace/seven-to-learn/apexcharts/2-dashboard/views"
)

func main() {
	component := views.Dashboard()
	http.Handle("/", templ.Handler(component))
	http.ListenAndServe(":3000", nil)
}
