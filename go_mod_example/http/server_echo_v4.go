package http

import "net/http"

func ExampleDefaultServer() {
	http.HandleFunc("/product", ProductHandle)
	http.ListenAndServe(":9095", nil)
}

func ProductHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world</h1>"))
}
