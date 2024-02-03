package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dhekerlaadhibi/LearnGo/forms/message"
	g "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func render(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal("error encountred while parsing the template :", err)

	}
	t.Execute(w, data)

}
func contacHandler(w http.ResponseWriter, r *http.Request) {
	//if get then display contact .html
	if r.Method == "GET" {
		render(w, "./templates/contact.html", nil)
	}
	// file message in package message
	if r.Method == "POST" {
		msg := &message.Message{
			Email:   r.PostFormValue("email"),
			Content: r.PostFormValue("content"),
		}
		if !msg.Validate() {
render(w,"./templates/contact.html",msg)
return
		}
		http.Redirect(w,r,"/confirm",http.StatusSeeOther)
	}

}
func confirmHandler(w http.ResponseWriter , r *http.Request){
	render(w , "./templates/confirm.html",nil)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/contact", contacHandler).Methods("GET", "POST")
	router.HandleFunc("/confirm", confirmHandler)
	//LoggingHandler is used to add logging functionality to a router in a Go (Golang
	// ywarik fl terminal w9tli taaml request chnoa ysir
	loggedRouter := g.LoggingHandler(os.Stdout, router)
	/*The LoggingHandler is being used to wrap the existing router
	   with logging functionality. This means that every request and response handled by this router will be logged to the standard output (os.Stdout). The logs generated can include information such as the HTTP method,
	  request URL, status code,
	   and other relevant details.*/

	http.ListenAndServe(PORT, loggedRouter)
}
