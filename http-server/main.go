package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

//func (s api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	//w.Write([]byte("Hello from the server"))
//	switch r.Method {
//	case http.MethodGet:
//		switch r.URL.Path {
//		case "/":
//			w.Write([]byte("index page"))
//			return
//		case "/users":
//			w.Write([]byte("users page"))
//			return
//		default:
//			w.Write([]byte("404 page"))
//			return
//		}
//	default:
//		w.Write([]byte("Only GET method is supported."))
//		return
//	}
//}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list..."))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

func main() {
	api := &api{addr: ":8080"}

	// Initialize the ServeMux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	log.Fatal(srv.ListenAndServe())

	//if err := http.ListenAndServe(s.addr, s); err != nil {
	//	log.Fatal(err)
	//}

	//log.Fatal(http.ListenAndServe(s.addr, s))

	//srv := &http.Server{
	//	Addr:    api.addr,
	//	Handler: api,
	//}
	//log.Fatal(srv.ListenAndServe())

}
