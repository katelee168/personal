package main

import (
       "fmt"
       "net/http"
       "os"
       "log"
)

func main() {
     fs := http.FileServer(http.Dir("static"))
     http.Handle("/", fs)

     fmt.Println("listening...")
     err := http.ListenAndServe(":"+os.Getenv("PORT"),nil)
     if err != nil {
     	panic(err)
     }
}

