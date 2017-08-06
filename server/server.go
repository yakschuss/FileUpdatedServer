package main

import (
  "fmt"
  "log"
  "net/http"
  "monitor_server/fileChecker"
)

type databaseCheckedHandler struct{}

func main() {
  fmt.Println(fileChecker.IsModified("database_updated.txt"))

  mux := http.NewServeMux()

  mux.Handle("/database_updated", databaseCheckedHandler{})

  log.Fatal(http.ListenAndServe(":8080", mux))

}

func (h databaseCheckedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}
