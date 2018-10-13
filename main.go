package main

import (
  "encoding/json"
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

func initApi(w http.ResponseWriter, r *http.Request) {
  	
  	info:= map[string]interface{}{
    		"uptime": "<uptime>",
    		"info": "Service for IGC tracks.",
    		"version": "v1",
  	}
	fmt.Fprintln(w,info)
}

func getApi(w http.ResponseWriter, r *http.Request) {
	addr, err := determineListenAddress()
	if err != nil {
    		log.Fatal(err)
  	}
	resp, err := http.Get(addr+"/api")
   	if err != nil {
      		log.Fatal(err)
   	}

   	var infoApi map[string]interface{}
   	err = json.NewDecoder(resp.Body).Decode(&infoApi)
   	if err != nil {
      		log.Fatal(err)
   	}

   	fmt.Println(infoApi)
}
func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }


  http.HandleFunc("/api", initApi)
  log.Fatal(http.ListenAndServe(addr,nil))

  http.HandleFunc("/api", getApi)
  log.Fatal(http.ListenAndServe(addr,nil))
}
