package main

import (
  "encoding/json"
  "log"
  "fmt"
  "net/http"
  "os"
  "io/ioutil"

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
  	
  	infoApi:=InfoApi {
    		 Uptime: <uptime>,
    		 Info: "Service for IGC tracks.",
    		 Version: "v1",
  	}
	fmt.Fprintln(w,infoApi)
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
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

  
}
