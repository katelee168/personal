package main

import (
       "fmt"
       "net/http"
       "os"
       "html/template"
)

func main() {
     http.HandleFunc("/", indexHandler)
     http.Handle("/Users/Katherine/Desktop/Projects/go/src/personal/static/stylesheets/", http.StripPrefix("/static/stylesheets/", http.FileServer(http.Dir("/Users/Katherine/Desktop/Projects/go/src/personal/static/stylesheets"))))

     fmt.Println("listening...")
     err := http.ListenAndServe(":"+os.Getenv("PORT"),nil)
     if err != nil {
     	panic(err)
     }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
     index := template.Must(template.ParseFiles(
     	   "/Users/Katherine/Desktop/Projects/go/src/personal/static/example.html",))
     index.Execute(w,nil)
}