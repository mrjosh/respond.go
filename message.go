package josh

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Message struct {

	// language of response
	Lang   string

	// success field of response
	Success string `yml:"success"`

	// failed field of response
	Failed string `yml:"failed"`

	// respond error messages
	Respond map[string]map[interface{}]interface{} `json:"respond"`
}

// Load config of response language
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return *Message
func (message *Message) LoadConfig() *Message {

	YmlFile, err := ioutil.ReadFile("./errors/" + message.Lang + ".yml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(YmlFile, &message)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return message
}