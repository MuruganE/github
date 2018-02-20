// Mock_Server project main.go
package main

import (
	"encoding/json"
	"fmt"
	//"log"
	//"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	contenttypeJSON  = "application/json; charset=utf-8"
	maxpayloadlength = 1 << 12
)

type listValue struct {
	Domains []string `json:"domains"`
}

type GTI struct {
	URL       string `json:"url"`
	ID        string `json:"id"`
	Password  string `json:"password"`
	ClientID  string `json:"clientid"`
	Product   string `json:"product"`
	Certcheck bool   `json:"certcheck"`
	Timeout   string `json:"timeout"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	//	logger, err := zap.NewProduction()
	//	if err != nil {
	//		log.Fatalf("can't initialize zap logger: %v", err)
	//	}
	//	defer logger.Sync()

	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println("Requesting for getting gti cred from server.....!")
	gti := GTI{
		URL:       "https://preprod.rest.gti.mcafee.com/1",
		ID:        "is-safeconnect-beta",
		Password:  "vA1c1(92a%_H{y6793n3e@<e8x-BbL",
		ClientID:  "67d8d1082c2f2f821f438b2359b7a5b4",
		Product:   "Safe Connect",
		Certcheck: false,
		Timeout:   "5s",
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	logger.Info("Now logs should be colored") //log.Printf("Now log system is stared......!")
	b3, _ := json.Marshal(gti)
	logger.Info("testing...........................! ", zap.ByteString("newStr", b3))

	w.Header().Set("Content-Type", contenttypeJSON)
	json.NewEncoder(w).Encode(gti)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "TEST 1 Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println("New request for getting white list")
	listVal := new(listValue)

	//TODO reading white list and update
	defaultConfigFile := "white_list.json"
	file, err := os.Open(defaultConfigFile)
	if err != nil {
		fmt.Println("Error during opening configuration file: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&listVal); err != nil {
		fmt.Println("Error during decoding configuration file: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contenttypeJSON)
	json.NewEncoder(w).Encode(listVal)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "TEST 2 Hi there, I love %s!", r.URL.Path[1:])
	//fmt.Fprintf(w, "TEST 1 Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println("New request for getting black list")
	listVal := new(listValue)

	//TODO reading white list and update
	defaultConfigFile := "black_list.json"
	file, err := os.Open(defaultConfigFile)
	if err != nil {
		fmt.Println("Error during opening configuration file: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&listVal); err != nil {
		fmt.Println("Error during decoding configuration file: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contenttypeJSON)
	json.NewEncoder(w).Encode(listVal)
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/gti", handler)
	http.HandleFunc("/whitelist", handler1)
	http.HandleFunc("/blacklist", handler2)
	http.ListenAndServe(":8082", nil)

}
