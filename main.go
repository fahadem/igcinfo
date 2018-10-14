package main

import (
  "encoding/json"
  "log"
  "fmt"
  "net/http"
  "os"
  //"strings"
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
}

func getApi(w http.ResponseWriter, r *http.Request) {
	

   	var infoApi InfoApi
   	err := json.NewDecoder(r.Body).Decode(&infoApi)
   	if err != nil {
      		log.Fatal(err)
   	}

   	fmt.Println(infoApi)
}
/*var(
	addr = "https://glacial-wave-53134.herokuapp.com"
	infoapi interface{}
)*/
func main() {
  	addr, err := determineListenAddress()
  	if err != nil {
    		log.Fatal(err)
  	}


  	http.HandleFunc("/api", initApi)

  	log.Fatal(http.ListenAndServe(addr,nil))
  	http.HandleFunc("/api", getApi)
	/*router:=httprouter.New()
	router.GET("/api",show)	
	err:=http.ListenAndServe(*addr,router)
	if err != nil {
      		log.Fatal(err)
   	}*/
	/*resp, err := http.Get(addr+"/api")
   	if err != nil {
      		log.Fatal(err)
   	}

   	var infoApi interface{}
   	err = json.NewDecoder(resp.Body).Decode(&infoApi)
   	if err != nil {
      		log.Fatal(err)
   	}

   	fmt.Println(infoApi)

	resp,err = http.Post(addr+"/api/igc","application/json",strings.NewReader("{\"url\": http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc}"))
	if err != nil {
      		log.Fatal(err)
   	}*/
}

/*func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w,"Read info api: %s",infoapi)
}*/
