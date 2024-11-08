package mygowebfw

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// starts the server on the specified port
func Run(port int) {
	fmt.Println("\033[32mServer starting on http://localhost" + ":" +
		strconv.Itoa(port) + "\033[0m")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// adds a route to the server that puts the component in a root layout
func AddPage(path, pageName string, opts map[string]string) {
	DefPage(pageName, opts)
	http.HandleFunc("GET "+path, func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("\033[32mGET: '" + path + "'\033[0m")
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, Render("_page_"+pageName, nil))
	})
}

// adds a route to the server
func AddRoute(path, componentName string) {
	http.HandleFunc("GET "+path, func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("\033[32mGET: '" + path + "'\033[0m")
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, Render(componentName, nil))
	})
}
