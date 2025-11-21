package http

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler) // maps the URL to the handler function (When a client (browser, Postman, etc.) sends a request to /, the handler is called.)
	http.ListenAndServe(":8080", nil) //starts TSP server on port 8080

	/* Every request triggers a request–response cycle:
	Client sends request (GET /)
	Go’s HTTP server receives it
	Calls the registered handler function
	Sends back the response 
*/
}