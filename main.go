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

func getApi(w http.ResponseWriter, r *http.Request) {
	

   	var m Message
   	err := json.NewDecoder(r.Body).Decode(&m)
   	if err != nil {
      		log.Fatal(err)
   	}

   	fmt.Println(m)
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
