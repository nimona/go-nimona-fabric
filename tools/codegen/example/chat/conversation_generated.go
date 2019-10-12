// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package chat

import (
	json "encoding/json"
	crypto "example/crypto"

	object "nimona.io/pkg/object"
	stream "nimona.io/pkg/stream"
)

type (
	ConversationTopicSet struct {
		Stream    *crypto.Hash      `json:"stream:o,omitempty"`
		Topic     string            `json:"topic:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Authors   []*stream.Author  `json:"authors:ao,omitempty"`
	}
	ConversationNameSet struct {
		Stream    *crypto.Hash      `json:"stream:o,omitempty"`
		Name      string            `json:"name:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Authors   []*stream.Author  `json:"authors:ao,omitempty"`
	}
	ConversationMessageAdded struct {
		Stream    *crypto.Hash      `json:"stream:o,omitempty"`
		Parents   []*crypto.Hash    `json:"parents:ao,omitempty"`
		Body      string            `json:"body:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Authors   []*stream.Author  `json:"authors:ao,omitempty"`
	}
)

func (e *ConversationTopicSet) GetType() string {
	return "example/conversation.ConversationTopicSet"
}

func (e *ConversationTopicSet) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "example/conversation.ConversationTopicSet"
	if e.Stream != nil {
		m["stream:o"] = e.Stream.ToObject().ToMap()
	}
	m["topic:s"] = e.Topic
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if len(e.Authors) > 0 {
		m["authors:ao"] = func() []interface{} {
			a := make([]interface{}, len(e.Authors))
			for i, v := range e.Authors {
				a[i] = v.ToObject().ToMap()
			}
			return a
		}()
	}
	return object.Object(m)
}

func (e *ConversationTopicSet) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e *ConversationNameSet) GetType() string {
	return "example/conversation.ConversationNameSet"
}

func (e *ConversationNameSet) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "example/conversation.ConversationNameSet"
	if e.Stream != nil {
		m["stream:o"] = e.Stream.ToObject().ToMap()
	}
	m["name:s"] = e.Name
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if len(e.Authors) > 0 {
		m["authors:ao"] = func() []interface{} {
			a := make([]interface{}, len(e.Authors))
			for i, v := range e.Authors {
				a[i] = v.ToObject().ToMap()
			}
			return a
		}()
	}
	return object.Object(m)
}

func (e *ConversationNameSet) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e *ConversationMessageAdded) GetType() string {
	return "example/conversation.ConversationMessageAdded"
}

func (e *ConversationMessageAdded) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "example/conversation.ConversationMessageAdded"
	if e.Stream != nil {
		m["stream:o"] = e.Stream.ToObject().ToMap()
	}
	if len(e.Parents) > 0 {
		m["parents:ao"] = func() []interface{} {
			a := make([]interface{}, len(e.Parents))
			for i, v := range e.Parents {
				a[i] = v.ToObject().ToMap()
			}
			return a
		}()
	}
	m["body:s"] = e.Body
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if len(e.Authors) > 0 {
		m["authors:ao"] = func() []interface{} {
			a := make([]interface{}, len(e.Authors))
			for i, v := range e.Authors {
				a[i] = v.ToObject().ToMap()
			}
			return a
		}()
	}
	return object.Object(m)
}

func (e *ConversationMessageAdded) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
