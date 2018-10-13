package main

import (
  "log"
  "fmt"
  "net/http"
  "os"
  "io/ioutil"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func getApi(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "meta information about the API")
  content, _ := ioutil.ReadAll(r.Body)
  fmt.Println(string(content))
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }

  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }

  http.HandleFunc("/api", getApi)
  log.Fatal(hhtp.ListenAndServe(addr,nil)
}
