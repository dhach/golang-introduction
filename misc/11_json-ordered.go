package main

import (
	"encoding/json"
	"log"
)

/*
	Reading and processing JSON data of known contents
*/

// define a struct with the fields from the JSON Object
type jsonMessage struct {
	NameLast  string
	NameGiven string
	Animal    string
	/*
		The variable names have to correspond to the key names in the JSON message, not considering case.
		Also note that the variables have to begin with an uppercase letter, meaning they have to be 'exported'.
	*/
}

// define a method on the struct to print its data
func (jm *jsonMessage) printData() {
	log.Printf("%v \n", jm)
	log.Println("Given Name: ", jm.NameGiven)
	log.Println("Last Name: ", jm.NameLast)
	log.Println("Animal: ", jm.Animal)
}

// decode the JSON data
func unmarshalKnownData() {
	printTraceAndGap()

	var message jsonMessage

	err := json.Unmarshal(xData, &message)
	if checkError(err) {
		log.Fatal("Cannot unmarshal JSON data: ", err)
	}

	message.printData()
}

/*


	Writing JSON data of known content

*/

// again, define a struct for our known data
type characterDC struct {
	Name           string `json:"name"` // the literal string does specify the key of the marshalled JSON object
	Affiliation    string `json:"affiliation"`
	CoolnessFactor uint8  `json:"coolness factor"`
}

func marshalKnownData() {
	printTraceAndGap()

	dickGrayson := &characterDC{
		Name:           "Nightwing",
		Affiliation:    "Titans",
		CoolnessFactor: 200,
	}

	jsonData, err := json.Marshal(dickGrayson)
	if checkError(err) {
		log.Fatal("Cannot marshal JSON data: ", err)
	}
	log.Println(string(jsonData))

	/*  This does not compile: */
	// garfield := &characterDC{
	// 	Name: "Beast Boy",
	// 	Affiliation: []string{"Doom Patrol", "Titans"},
	// 	CoolnessFactor: 125,
	// }
}
