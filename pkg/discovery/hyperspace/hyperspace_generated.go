// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package hyperspace

import (
	json "encoding/json"

	crypto "nimona.io/pkg/crypto"
	object "nimona.io/pkg/object"
)

type (
	Request struct {
		QueryContentBloom []int64             `json:"queryContentBloom:ai"`
		Nonce             string              `json:"nonce:s"`
		Signature         *crypto.Signature   `json:"@signature:o"`
		Authors           []*crypto.PublicKey `json:"@authors:ao"`
	}
	Announced struct {
		AvailableContentBloom []int64             `json:"availableContentBloom:ai"`
		Nonce                 string              `json:"nonce:s"`
		Signature             *crypto.Signature   `json:"@signature:o"`
		Authors               []*crypto.PublicKey `json:"@authors:ao"`
	}
)

func (e *Request) EventName() string {
	return "Request"
}

func (e *Request) GetType() string {
	return "ContentProvider.Request"
}

func (e *Request) ToObject() object.Object {
	m := map[string]interface{}{
		"@ctx:s":    "ContentProvider.Request",
		"@domain:s": "ContentProvider",
		"@event:s":  "Request",
	}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &m)
	return object.Object(m)
}

func (e *Request) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}

func (e *Announced) EventName() string {
	return "Announced"
}

func (e *Announced) GetType() string {
	return "ContentProvider.Announced"
}

func (e *Announced) ToObject() object.Object {
	m := map[string]interface{}{
		"@ctx:s":    "ContentProvider.Announced",
		"@domain:s": "ContentProvider",
		"@event:s":  "Announced",
	}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &m)
	return object.Object(m)
}

func (e *Announced) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}