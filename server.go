package main

import (
    "fmt"
    "log"
	"net/http"
	"time"
	"encoding/json"
)

type Time1 struct  {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request){
		t := time.Now().Format(time.RFC3339);
		
		json1, err := json.MarshalIndent(Time1{
			Time: t,
		},  "", " ")
		if err != nil {
			fmt.Fprintf(w, err.Error());
			return
		}
        fmt.Fprintf(w, string(json1))
    })

	fmt.Printf("Starting server at port 8795\n")
    if err := http.ListenAndServe(":8795", nil); err != nil {
        log.Fatal(err)
    }
}
