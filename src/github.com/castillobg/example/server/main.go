package main

import (
  "fmt"
  "net/http"
)

func johnnyHandler(w http.ResponseWriter, r *http.Request) {
  html := `<html>
    </body>
      <img src="https://media.giphy.com/media/RoajqIorBfSE/giphy.gif"></img>
    </body>
    </html>`
  fmt.Fprintf(w, "%s", html)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, it's me: %s", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/johnny", johnnyHandler)
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
