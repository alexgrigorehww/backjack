package description

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Description struct {
	Title     string    `json:"title"`
	Version   string    `json:"version"`
	Licensed  string    `jsons:"licensed"`
	Poweredby string    `json:"poweredby"`
	About     string    `json:"about"`
	Creators  []Creator `json:"creators"`
}

type Creator struct {
	Name string `json:"name"`
	Area string `json:"area"`
}

func GetDescription() string {
	fmt.Println("Start")
	// Open our jsonFile
	jsonFile, err := os.Open("description.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err.Error()
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	var description Description
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &description)
	returnedString := description.Title + " " + description.Version +
		"\n" + description.Licensed + "\n" + description.Poweredby + "\n" + description.About + "\nAuthors"
	for _, a := range description.Creators {
		returnedString += "\n" + a.Name + " : " + a.Area
	}
	return returnedString
}
