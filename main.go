package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type api struct {
	Certs string `json:"certs"`
	Key   string `json:"key"`
}

func main() {
	certs, err := ioutil.ReadFile("./https.cer")
	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := ioutil.ReadFile("./https.key")
	if err != nil {
		fmt.Println(err)
		return
	}
	instance := api{
		Certs: string(certs),
		Key:   string(key),
	}

	bs, err := json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("PUT", "https://rancher.node3.autops.xyz/v3/project/c-9qc8v:p-ptb52/certificates/p-ptb52:hzy", bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
		return
	}

	req.SetBasicAuth("token-nkvlm", "z6b8fbsdxp7qmq9rcns6sp6rblbkztxzsbjrwpqrgksrpmwnzt95p8")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	respBS, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(respBS))
}