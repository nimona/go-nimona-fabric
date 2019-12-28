// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package chat

import (
	json "encoding/json"

	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
	schema "nimona.io/pkg/schema"
)

type (
	ConversationCreated struct {
		Name      string            `json:"name:s,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
	ConversationTopicUpdated struct {
		Topic     string            `json:"topic:s,omitempty"`
		DependsOn []object.Hash     `json:"dependsOn:ar,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
	ConversationMessageAdded struct {
		Body      string            `json:"body:s,omitempty"`
		DependsOn []object.Hash     `json:"dependsOn:ar,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
	ConversationMessageRemoved struct {
		Removes   object.Hash       `json:"removes:r,omitempty"`
		DependsOn []object.Hash     `json:"dependsOn:ar,omitempty"`
		Signature *crypto.Signature `json:"@signature:o,omitempty"`
		Identity  crypto.PublicKey  `json:"@identity:s,omitempty"`
	}
)

func (e ConversationCreated) GetType() string {
	return "mochi.io/conversation.Created"
}

func (e ConversationCreated) GetSchema() *schema.Object {
	return &schema.Object{
		Properties: []*schema.Property{
			&schema.Property{
				Name:       "name",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@signature",
				Type:       "nimona.io/crypto.Signature",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@identity",
				Type:       "nimona.io/crypto.PublicKey",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e ConversationCreated) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "mochi.io/conversation.Created"
	if e.Name != "" {
		m["name:s"] = e.Name
	}
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if e.Identity != "" {
		m["@identity:s"] = e.Identity
	}
	if schema := e.GetSchema(); schema != nil {
		m["$schema:o"] = schema.ToObject().ToMap()
	}
	return object.Object(m)
}

func (e *ConversationCreated) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e ConversationTopicUpdated) GetType() string {
	return "mochi.io/conversation.TopicUpdated"
}

func (e ConversationTopicUpdated) GetSchema() *schema.Object {
	return &schema.Object{
		Properties: []*schema.Property{
			&schema.Property{
				Name:       "topic",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "dependsOn",
				Type:       "relationship",
				Hint:       "r",
				IsRepeated: true,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@signature",
				Type:       "nimona.io/crypto.Signature",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@identity",
				Type:       "nimona.io/crypto.PublicKey",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e ConversationTopicUpdated) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "mochi.io/conversation.TopicUpdated"
	if e.Topic != "" {
		m["topic:s"] = e.Topic
	}
	if len(e.DependsOn) > 0 {
		m["dependsOn:ar"] = e.DependsOn
	}
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if e.Identity != "" {
		m["@identity:s"] = e.Identity
	}
	if schema := e.GetSchema(); schema != nil {
		m["$schema:o"] = schema.ToObject().ToMap()
	}
	return object.Object(m)
}

func (e *ConversationTopicUpdated) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e ConversationMessageAdded) GetType() string {
	return "mochi.io/conversation.MessageAdded"
}

func (e ConversationMessageAdded) GetSchema() *schema.Object {
	return &schema.Object{
		Properties: []*schema.Property{
			&schema.Property{
				Name:       "body",
				Type:       "string",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "dependsOn",
				Type:       "relationship",
				Hint:       "r",
				IsRepeated: true,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@signature",
				Type:       "nimona.io/crypto.Signature",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@identity",
				Type:       "nimona.io/crypto.PublicKey",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e ConversationMessageAdded) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "mochi.io/conversation.MessageAdded"
	if e.Body != "" {
		m["body:s"] = e.Body
	}
	if len(e.DependsOn) > 0 {
		m["dependsOn:ar"] = e.DependsOn
	}
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if e.Identity != "" {
		m["@identity:s"] = e.Identity
	}
	if schema := e.GetSchema(); schema != nil {
		m["$schema:o"] = schema.ToObject().ToMap()
	}
	return object.Object(m)
}

func (e *ConversationMessageAdded) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e ConversationMessageRemoved) GetType() string {
	return "mochi.io/conversation.MessageRemoved"
}

func (e ConversationMessageRemoved) GetSchema() *schema.Object {
	return &schema.Object{
		Properties: []*schema.Property{
			&schema.Property{
				Name:       "removes",
				Type:       "relationship",
				Hint:       "r",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "dependsOn",
				Type:       "relationship",
				Hint:       "r",
				IsRepeated: true,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@signature",
				Type:       "nimona.io/crypto.Signature",
				Hint:       "o",
				IsRepeated: false,
				IsOptional: false,
			},
			&schema.Property{
				Name:       "@identity",
				Type:       "nimona.io/crypto.PublicKey",
				Hint:       "s",
				IsRepeated: false,
				IsOptional: false,
			},
		},
	}
}

func (e ConversationMessageRemoved) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@type:s"] = "mochi.io/conversation.MessageRemoved"
	m["removes:r"] = e.Removes
	if len(e.DependsOn) > 0 {
		m["dependsOn:ar"] = e.DependsOn
	}
	if e.Signature != nil {
		m["@signature:o"] = e.Signature.ToObject().ToMap()
	}
	if e.Identity != "" {
		m["@identity:s"] = e.Identity
	}
	if schema := e.GetSchema(); schema != nil {
		m["$schema:o"] = schema.ToObject().ToMap()
	}
	return object.Object(m)
}

func (e *ConversationMessageRemoved) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
