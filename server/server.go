package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "monitor_server/fileChecker"
  "github.com/spf13/afero"
)

type databaseCheckedHandler struct{
  fileName string
  handler fileChecker.FileChecker
}

func main() {
  mux := http.NewServeMux()

  mux.Handle("/database_updated", databaseCheckedHandler{
    fileName: "database_updated.txt",
    handler: initChecker(),
  })

  log.Fatal(http.ListenAndServe(":8080", mux))

}

func (h databaseCheckedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  modifiedData := h.handler.IsModified(h.fileName)

  fmt.Println(modifiedData)

  var jsonData, err = json.Marshal(modifiedData)

  if err != nil {
    panic(err)
  }

  fmt.Println(jsonData)

  w.Write(jsonData)
}

func initChecker() fileChecker.FileChecker {
  fileOS := afero.NewOsFs()
  theChecker := fileChecker.FileChecker{
    Checker: fileOS,
  }

  return theChecker
}
