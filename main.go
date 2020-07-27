package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/lanceplarsen/test/vault"
)

func main() {
	vault, err := vault.InitializeVault()
	if err != nil {
		panic(err)
	}

	path, _ := os.Executable()
	b, _ := ioutil.ReadFile(path + "/vault/data/cred.json")

	input := make(map[string]interface{})

	json.Unmarshal(b, &input)

	vault.Create(input, "initialize")

	vault.Read("initialize")
	input = make(map[string]interface{})
	input["ping2"] = "pong"
	err = vault.Append(input, "initialize")
	if err != nil {
		panic(err)
	}

	input["ping2"] = "pong2"
	err = vault.Update(input, "initialize")
	if err != nil {
		panic(err)
	}

	r, _ := vault.Read("initialize")
	log.Println(r["ping2"])

	err = vault.Delete("initialize")
	if err != nil {
		panic(err)
	}
}
