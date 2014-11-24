package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rlouapre/mlxcccli/context"
	"github.com/rlouapre/mlxcccli/digest"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func newClient(authentication string, username string, password string) (*http.Client, error) {
	if authentication == "basic" {
		return &http.Client{}, nil
	}
	if authentication == "digest" {
		return digest.NewTransport(username, password).Client()
	}
	return nil, errors.New("Authentication not supported")
}

func main() {
	context := context.NewContext()

	if context.Authentication != "basic" && context.Authentication != "digest" {
		log.Panicf("Authentication '%s' not supported.", context.Authentication)
	}

	client, err := newClient(context.Authentication, context.Username, context.Password)
	if err != nil {
		log.Fatal(err)
	}

	uri := fmt.Sprintf("http://%s:%v/eval", context.Host, context.Port)

	fmt.Println("About to execute xquery against XCC")
	fmt.Printf("Url: %s\n", uri)
	fmt.Printf("File: %s\n", context.Filename)

	xquery, err := ioutil.ReadFile(context.Filename)
	if err != nil {
		log.Fatal(err)
	}

	// Set data to be posted
	data := make(url.Values)
	data.Set("xquery", string(xquery))

	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("User-Agent", "Go XCC client")
	// req.Header.Set("Accept", "*/*")
	if context.Authentication == "basic" {
		req.SetBasicAuth(context.Username, context.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error: %v\n Response: %v\n", err, resp)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(response))

}
