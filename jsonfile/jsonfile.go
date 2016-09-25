package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"
)

// FileExists - function to test the existance of a file
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

//ProcessAPIDesc - Entry point
func ProcessAPIDesc(filePath string, globalConc int) error {
	jsonFileReader, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening the file")
		return err
	}
	jsonParser := json.NewDecoder(jsonFileReader)
	var APIDescrition APIDesc
	err = jsonParser.Decode(&APIDescrition)
	if err != nil {
		fmt.Println("Error in decoding the json")
		return err
	}
	return nil
}
