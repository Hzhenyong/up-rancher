package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type api struct {
	Certs string `json:"certs"`
	Key   string `json:"key"`
	RancherUrl string `json:"rancher_url"`
	AK string	`json:"username"`
	SK string	`json:"password"`
	CD string	`json:"cd"`
}

type domain struct {
	Web1 string `json:"web1"`
	Web2 string `json:"web2"`
}

func main() {

	//https域名地址
	do := domain{
		Web1: os.Getenv("WEB1"),
		Web2: os.Getenv("WEB2"),
	}
	certs, err := ioutil.ReadFile("/acme.sh/"+do.Web1+"/"+do.Web1+".cer")
	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := ioutil.ReadFile("/acme.sh/"+do.Web1+"/"+do.Web1+".key")
	if err != nil {
		fmt.Println(err)
		return
	}
	//ssl证书和秘钥，
	//ak 、sk 是rancher的授权
	//RANCHER_URL 是要更新的证书的api地址
	instance := api{
		Certs: string(certs),
		Key:   string(key),
		AK: os.Getenv("AK"),
		SK: os.Getenv("SK"),
		RancherUrl: os.Getenv("RANCHER_URL"),
		CD: os.Getenv("CD"),
	}


	bs, err := json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
		return
	}
	req, err := http.NewRequest("PUT", instance.RancherUrl, bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
		return
	}
	//跳过ssl证书验证，
	t := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: t}

	req.SetBasicAuth(instance.AK, instance.SK)
	req.Header.Add("Content-Type", "application/json")
	//response, err := client.Do(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("使用跳过ssl验证")
		resp, err = client.Do(req)
		if err != nil {
			log.Printf("跳过也无法更新，请验证rancher地址是否正常")
			return
		}
	}
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