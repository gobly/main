package main

import (
	"net/http"
	"github.com/gobly/core"
	"os"
	"github.com/gobly/ui"
)

var layoutSingle = ui.LoadSingle("html/main.html")

func main() {
	router := core.NewRouter()
	router.AddHandler("/", rootHandler)

	// Initialize modules
	ui.CreateContext("/ui", router)

	core.ShowWelcome(os.Stdout, router)
	http.ListenAndServe(":8080", router)
}

func rootHandler(out http.ResponseWriter, _ *http.Request) {
	layoutSingle.Execute(out, nil)
}
