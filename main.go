package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
	loggly "loggly"
	//"net/http"
	//"os"
)

func main() {

	// url := "http://api.football-data.org/v2/competitions/2003"

	// var api_Token = "0d5fb989868f4421bce51517a5bbb62d"

	// //resp, err := http.Get(url)

	// // Create a new request using http
	// req, err := http.NewRequest("GET", url, nil)

	// // add authorization header to the req
	// req.Header.Add("X-Auth-Token", api_Token)

	// // Send req using http Client
	// client := &http.Client{}
	// resp, err := client.Do(req)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = os.Stdout.Write(body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	var tag string
	tag = "My-Go-Demo"

	// Instantiate the client
	client := loggly.New(tag)

	// Valid EchoSend (message echoed to console and no error returned)
	err := client.EchoSend("info", "Good morning!")
	fmt.Println("err:", err)

	// Valid Send (no error returned)
	err = client.Send("error", "Good morning! No echo.")
	fmt.Println("err:", err)

	// Invalid EchoSend -- message level error
	err = client.EchoSend("blah", "blah")
	fmt.Println("err:", err)
}
