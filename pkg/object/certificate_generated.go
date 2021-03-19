// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package object

type (
	Certificate struct {
		Metadata Metadata
		Nonce    string
		Created  string
		Expires  string
	}
	CertificateRequest struct {
		Metadata               Metadata
		ApplicationName        string
		ApplicationDescription string
		ApplicationURL         string
		Subject                string
		Resources              []string
		Actions                []string
		Nonce                  string
	}
)

func (e *Certificate) Type() string {
	return "nimona.io/Certificate"
}

func (e *Certificate) MarshalMap() (Map, error) {
	return e.ToObject().Map(), nil
}

func (e Certificate) ToObject() *Object {
	r := &Object{
		Type:     "nimona.io/Certificate",
		Metadata: e.Metadata,
		Data:     Map{},
	}
	r.Data["nonce"] = String(e.Nonce)
	r.Data["created"] = String(e.Created)
	r.Data["expires"] = String(e.Expires)
	return r
}

func (e *Certificate) UnmarshalMap(m Map) error {
	return e.FromObject(FromMap(m))
}

func (e *Certificate) FromObject(o *Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(String); ok {
			e.Nonce = string(t)
		}
	}
	if v, ok := o.Data["created"]; ok {
		if t, ok := v.(String); ok {
			e.Created = string(t)
		}
	}
	if v, ok := o.Data["expires"]; ok {
		if t, ok := v.(String); ok {
			e.Expires = string(t)
		}
	}
	return nil
}

func (e *CertificateRequest) Type() string {
	return "nimona.io/CertificateRequest"
}

func (e *CertificateRequest) MarshalMap() (Map, error) {
	return e.ToObject().Map(), nil
}

func (e CertificateRequest) ToObject() *Object {
	r := &Object{
		Type:     "nimona.io/CertificateRequest",
		Metadata: e.Metadata,
		Data:     Map{},
	}
	r.Data["applicationName"] = String(e.ApplicationName)
	r.Data["applicationDescription"] = String(e.ApplicationDescription)
	r.Data["applicationURL"] = String(e.ApplicationURL)
	r.Data["subject"] = String(e.Subject)
	if len(e.Resources) > 0 {
		rv := make(StringArray, len(e.Resources))
		for i, iv := range e.Resources {
			rv[i] = String(iv)
		}
		r.Data["resources"] = rv
	}
	if len(e.Actions) > 0 {
		rv := make(StringArray, len(e.Actions))
		for i, iv := range e.Actions {
			rv[i] = String(iv)
		}
		r.Data["actions"] = rv
	}
	r.Data["nonce"] = String(e.Nonce)
	return r
}

func (e *CertificateRequest) UnmarshalMap(m Map) error {
	return e.FromObject(FromMap(m))
}

func (e *CertificateRequest) FromObject(o *Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["applicationName"]; ok {
		if t, ok := v.(String); ok {
			e.ApplicationName = string(t)
		}
	}
	if v, ok := o.Data["applicationDescription"]; ok {
		if t, ok := v.(String); ok {
			e.ApplicationDescription = string(t)
		}
	}
	if v, ok := o.Data["applicationURL"]; ok {
		if t, ok := v.(String); ok {
			e.ApplicationURL = string(t)
		}
	}
	if v, ok := o.Data["subject"]; ok {
		if t, ok := v.(String); ok {
			e.Subject = string(t)
		}
	}
	if v, ok := o.Data["resources"]; ok {
		if t, ok := v.(StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.Resources = rv
		}
	}
	if v, ok := o.Data["actions"]; ok {
		if t, ok := v.(StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.Actions = rv
		}
	}
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(String); ok {
			e.Nonce = string(t)
		}
	}
	return nil
}
