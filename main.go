package main

import (
  "encoding/json"
  "log"
  "fmt"
  "net/http"
  "os"

)

type InfoApi struct {
    Uptime float64 `json:"uptime,omitempty"`
    Info string `json:"info,omitempty"`
    Version string `json:"version,omitempty"`
}



func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func initApi(w http.ResponseWriter, r *http.Request) {
  	http.Header.Add(w.Header(),"content-type","application/json")
  	infoApi:=InfoApi {
    		 Uptime: 100.0,
    		 Info: "Service for IGC tracks.",
    		 Version: "v1",
  	}
	
	json.NewEncoder(w).Encode(infoApi)
}

func getApi() {
	addr, err := determineListenAddress()
	if err != nil {
    		log.Fatal(err)
  	}
	resp, err := http.Get(addr+"/api")
   	if err != nil {
      		log.Fatal(err)
   	}

   	var infoApi Infoapi
   	err := json.NewDecoder(resp.Body).Decode(&infoApi)
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

  
}
