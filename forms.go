package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dhekerlaadhibi/LearnGo/forms/message"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
const(
	PORT=":8080"
)
func render(w http.ResponseWriter ,filename string , data interface{}){
t,err :=template.ParseFiles(filename)
if err != nil{
	log.Fatal("error encountred while parsing the template :",err)

}
t.Execute(w,data)



}


func main(){
	router:= mux.NewRouter()
	router.HandleFunc("/contact",contacHandler).Methods("GET","POST")
}






