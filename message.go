package respond

import (
	"sync"

	"github.com/mrjosh/respond.go/translations/en"
	"github.com/mrjosh/respond.go/translations/fa"
)

type Messages struct {
	Lang      string
	Success   string
	Failed    string
	Errors    map[string]map[string]interface{}
	Languages map[string]map[string]interface{}
	sync.RWMutex
}

func NewMessages() *Messages {
	return &Messages{
		Lang: "en",
		Languages: map[string]map[string]interface{}{
			"fa": fa.Messages,
			"en": en.Messages,
		},
	}
}

func (m *Messages) AddLanguageTranslation(lang string, messages map[string]interface{}) {
	m.Lock()
	m.Languages[lang] = messages
	m.Unlock()
}

// Load config of response language
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return *Message
func (m *Messages) load() {
	m.RLock()
	translation := m.Languages[m.Lang]
	m.Errors = translation["errors"].(map[string]map[string]interface{})
	m.Success = translation["success"].(string)
	m.Failed = translation["failed"].(string)
	m.RUnlock()
}
