package main

import (
	"github.com/imroc/req"
	"io/ioutil"
	"log" // now using the log package for output
	"net/http"
)

var (
	targetPage string = "https://gitea.full-spectrum-cyber.space/"
)

/*
	Using the net/http package from the standard lib.
*/
func netHTTPRequest() {
	printTraceAndGap()

	resp, err := http.Get(targetPage)

	if checkError(err) {
		log.Fatal("Cannot get web page: ", err)
	}
	defer resp.Body.Close()
	log.Println("Does this get logged?")

	body, err := ioutil.ReadAll(resp.Body)
	if checkError(err) {
		log.Fatal("Cannot read page body: ", err)
	}

	bodyString := string(body)
	log.Println(bodyString)
}

/*
	Using the third-party package "github.com/imroc/req".
	It's a bit more concise and user-friendly
*/
func imrocReq() {
	printTraceAndGap()

	resp, err := req.Get(targetPage)
	if checkError(err) {
		log.Fatal("Cannot get web page: ", err)
	}

	respString := resp.String()

	log.Println(respString)
}

/*
	Let's POST some stuff
*/
func imrocPost() {
	printTraceAndGap()

	targetPagePost := "http://docker-machine:8080/post"

	data := req.Param{
		"foo":  "bar",
		"user": "dhs",
	}

	resp, err := req.Post(targetPagePost, data)
	if checkError(err) {
		log.Fatal("Cannot POST message: ", err)
	}

	respString := resp.String()
	log.Println(respString)

}
