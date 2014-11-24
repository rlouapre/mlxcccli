package context

import (
	_ "errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Context struct {
	Host           string
	Port           int
	Authentication string
	Username       string
	Password       string
	Filename       string
}

var context = &Context{}

func NewContext() *Context {
	if err := parseYaml(); err != nil {
		parseCommandLine()
	}
	return context
}

func parseYaml() error {
	_, err := os.Stat("xcc.yml")
	if err != nil {
		return err
	}
	conf, err := ioutil.ReadFile("xcc.yml")
	err = yaml.Unmarshal(conf, &context)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	fmt.Printf("context from yaml:\n%#v\n\n", context)
	return nil
}

func parseCommandLine() {
	flag.IntVar(&context.Port, "port", 0, "XCC port")
	flag.StringVar(&context.Host, "host", "localhost", "XCC host")
	flag.StringVar(&context.Username, "username", "admin", "Username")
	flag.StringVar(&context.Password, "password", "password", "Password")
	flag.StringVar(&context.Filename, "filename", "", "XQuery filename")
	flag.StringVar(&context.Authentication, "authentication", "basic", "Authentication type 'basic' or 'digest' supported.")
	flag.Parse()
}
