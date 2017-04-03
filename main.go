package main

import (
	"log"
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",rootHandler)
	http.HandleFunc("/init",initHandler)
	http.HandleFunc("/init2",init2Handler)
	http.ListenAndServe(":8080",nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"Hello world!")
}

func initHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Hey, this is the init function")
}

func init2Handler(w http.ResponseWriter, r *http.Request){
	log.Println("Hey, this is the init2 functionn")
}