package main

import (
       "fmt"
       "net/http"
       "os"
       "html/template"
)

func main() {
     http.HandleFunc("/", indexHandler)
     http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", http.FileServer(http.Dir("static/stylesheets"))))

     fmt.Println("listening...")
     err := http.ListenAndServe(":"+os.Getenv("PORT"),nil)
     if err != nil {
     	panic(err)
     }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
     index := template.Must(template.ParseFiles(
     	   "/static/example.html",))
     index.Execute(w,nil)
}