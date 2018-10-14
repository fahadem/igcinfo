package main

import (
  "encoding/json"
  "log"
  "fmt"
  "net/http"
  "os"
  "time"
)

type InfoApi struct {
    Uptime time.Time `json:"uptime,omitempty"`
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
    		 Uptime: time.Now(),
    		 Info: "Service for IGC tracks.",
    		 Version: "v1",
  	}
	
	json.NewEncoder(w).Encode(infoApi)

	/*var infoApi2 InfoApi
   	err := json.NewDecoder(r.Body).Decode(&infoApi2)
   	if err != nil {
      		log.Fatal(err)
   	}

   	fmt.Println(infoApi)*/
}

func getApi(w http.ResponseWriter, r *http.Request) {
	

   	var infoApi InfoApi
   	err := json.NewDecoder(r.Body).Decode(&infoApi)
   	if err != nil {
      		log.Fatal(err)
   	}

   	fmt.Println(infoApi)
}

func postIgc(w http.ResponseWriter, r *http.Request) {
  	http.Header.Add(w.Header(),"content-type","application/json")

  	url:="http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"
	json.NewEncoder(w).Encode(url)
}


func main() {
  	addr, err := determineListenAddress()
  	if err != nil {
    		log.Fatal(err)
  	}


  	http.HandleFunc("/api", initApi)
  	//http.HandleFunc("/api/igc", postIgc)
	
  	log.Fatal(http.ListenAndServe(addr,nil))
  	
}

