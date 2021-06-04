package respond

import (
	"github.com/mrjosh/respond.go/translations/en"
	"github.com/mrjosh/respond.go/translations/fa"
)

type Message struct {

	// language of response
	Lang string

	// success field of response
	Success string

	// failed field of response
	Failed string

	// respond error messages
	Errors map[string]map[string]interface{}
}

// Load config of response language
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return *Message
func (message *Message) LoadConfig() *Message {

	var translation map[string]interface{}

	switch message.Lang {
	case "fa":
		translation = fa.Errors
	default:
		translation = en.Errors
	}

	message.Errors = translation["errors"].(map[string]map[string]interface{})
	message.Success = translation["success"].(string)
	message.Failed = translation["failed"].(string)

	return message
}
