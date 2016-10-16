package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Credential struct {
	Mail     string `json:"email"`
	Password string `json:"password"`
}

func ReadCredentials() Credential {
	file, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		fmt.Printf("File config read error: %v\n", err)
		os.Exit(1)
	}

	var credentials Credential
	json.Unmarshal(file, &credentials)
	return credentials
}
