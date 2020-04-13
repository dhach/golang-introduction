package main

import (
	"encoding/json"
	"log"
)

var (
	xData = []byte(`
	{
		"nameGiven": "Jean",
		"nameLast": "Grey",
		"animal": "phoenix"
	}`)
)

/*
	Reading and processing JSON of unknown contents
*/
func unmarshalUnknownData() {
	printTraceAndGap()

	var decodedData interface{}
	err := json.Unmarshal(xData, &decodedData)
	if checkError(err) {
		log.Fatal("Cannot unmarshal JSON data: ", err)
	}
	// log.Println(decodedData)
	// log.Print("\n\n")

	/*
		What to do with the decoded data?
		This is straight outta the official Golang blog...
	*/
	decodedMessage := decodedData.(map[string]interface{})
	for k, v := range decodedMessage {
		switch vv := v.(type) {
		case string:
			log.Println(k, "is string", vv)
		case float64:
			log.Println(k, "is float64", vv)
		case []interface{}:
			log.Println(k, "is an array:")
			for i, u := range vv {
				log.Println(i, u)
			}
		default:
			log.Println(k, "is of a type I don't know how to handle")
		}
	}
}

/*

	Writing JSON data of unknown content
*/
func marshalUnknownData() {
	printTraceAndGap()

	/*
		To encode arbitrary data into JSON a map is the best solution.
		Note that the keys have to be strings.
	*/
	mcuData := map[string]interface{}{
		"Frank Castle": map[string]interface{}{
			"alias": "Punisher",
			"affiliations": []string{
				"None",
			},
		},
		"editions count": 9999,
	}

	jsonData, err := json.Marshal(mcuData)
	if checkError(err) {
		log.Fatal("Cannot marshal JSON data: ", err)
	}

	log.Println(string(jsonData))

}

/*
	interfaces: "named collections of method signatures"

	Empty interface 'interface{}' can hold values of any type.
	Especially useful if dealing with data of unknown type
*/
