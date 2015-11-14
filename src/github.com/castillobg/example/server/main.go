package main

import (
  "fmt"
  "net"
  "net/http"
)

func johnnyHandler(w http.ResponseWriter, r *http.Request) {
  html := `<html>
    </body>
      <img src="https://media.giphy.com/media/RoajqIorBfSE/giphy.gif"></img>
    </body>
    </html>`
  fmt.Fprintf(w, "%s from %s", html, getLocalIP())
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, it's me: %s from %s", r.URL.Path[1:], getLocalIP())
}

func getLocalIP() string {
  addrs, err := net.InterfaceAddrs()
  if err != nil {
    return ""
  }
  for _, address := range addrs {
    if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        return ipnet.IP.String()
      }
    }
  }
  return ""
}

func main() {
  http.HandleFunc("/johnny", johnnyHandler)
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
