// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package fixtures

import (
	object "nimona.io/pkg/object"
)

type (
	CompositeTest struct {
		Metadata                    object.Metadata
		CompositeStringTest         *Composite
		CompositeDataTest           *Composite
		RepeatedCompositeStringTest []*Composite
		RepeatedCompositeDataTest   []*Composite
	}
	TestPolicy struct {
		Metadata   object.Metadata
		Subjects   []string
		Resources  []string
		Conditions []string
		Action     string
	}
	TestStream struct {
		Metadata        object.Metadata
		Nonce           string
		CreatedDateTime string
	}
	TestSubscribed struct {
		Metadata object.Metadata
		Nonce    string
	}
	TestUnsubscribed struct {
		Metadata object.Metadata
		Nonce    string
	}
	TestRequest struct {
		Metadata  object.Metadata
		RequestID string
		Foo       string
	}
	TestResponse struct {
		Metadata  object.Metadata
		RequestID string
		Foo       string
	}
	Parent struct {
		Metadata      object.Metadata
		Foo           string
		Child         *Child
		RepeatedChild []*Child
	}
	Child struct {
		Metadata object.Metadata
		Foo      string
	}
)

func (e *CompositeTest) Type() string {
	return "compositeTest"
}

func (e *CompositeTest) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *CompositeTest) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e CompositeTest) ToObject() *object.Object {
	r := &object.Object{
		Type:     "compositeTest",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	if e.CompositeStringTest != nil {
		if v, err := e.CompositeStringTest.MarshalString(); err == nil {
			r.Data["compositeStringTest"] = object.String(v)
		}
	}
	if e.CompositeDataTest != nil {
		if v, err := e.CompositeDataTest.MarshalBytes(); err == nil {
			r.Data["compositeDataTest"] = object.Data(v)
		}
	}
	if len(e.RepeatedCompositeStringTest) > 0 {
		rv := make(object.StringArray, len(e.RepeatedCompositeStringTest))
		for i, v := range e.RepeatedCompositeStringTest {
			if iv, err := v.MarshalString(); err == nil {
				rv[i] = object.String(iv)
			}
		}
		r.Data["repeatedCompositeStringTest"] = rv
	}
	if len(e.RepeatedCompositeDataTest) > 0 {
		rv := make(object.DataArray, len(e.RepeatedCompositeDataTest))
		for i, v := range e.RepeatedCompositeDataTest {
			if iv, err := v.MarshalBytes(); err == nil {
				rv[i] = object.Data(iv)
			}
		}
		r.Data["repeatedCompositeDataTest"] = rv
	}
	return r
}

func (e *CompositeTest) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *CompositeTest) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *CompositeTest) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["compositeStringTest"]; ok {
		if ev, ok := v.(object.String); ok {
			es := &Composite{}
			if err := es.UnmarshalString(string(ev)); err == nil {
				e.CompositeStringTest = es
			}
		}
	}
	if v, ok := o.Data["compositeDataTest"]; ok {
		if ev, ok := v.(object.Data); ok {
			es := &Composite{}
			if err := es.UnmarshalBytes([]byte(ev)); err == nil {
				e.CompositeDataTest = es
			}
		}
	}
	if v, ok := o.Data["repeatedCompositeStringTest"]; ok {
		if ev, ok := v.(object.StringArray); ok {
			e.RepeatedCompositeStringTest = make([]*Composite, len(ev))
			for i, iv := range ev {
				es := &Composite{}
				if err := es.UnmarshalString(string(iv)); err == nil {
					e.RepeatedCompositeStringTest[i] = es
				}
			}
		}
	}
	if v, ok := o.Data["repeatedCompositeDataTest"]; ok {
		if ev, ok := v.(object.DataArray); ok {
			e.RepeatedCompositeDataTest = make([]*Composite, len(ev))
			for i, iv := range ev {
				es := &Composite{}
				if err := es.UnmarshalBytes([]byte(iv)); err == nil {
					e.RepeatedCompositeDataTest[i] = es
				}
			}
		}
	}
	return nil
}

func (e *TestPolicy) Type() string {
	return "nimona.io/fixtures.TestPolicy"
}

func (e *TestPolicy) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *TestPolicy) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e TestPolicy) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/fixtures.TestPolicy",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	if len(e.Subjects) > 0 {
		rv := make(object.StringArray, len(e.Subjects))
		for i, iv := range e.Subjects {
			rv[i] = object.String(iv)
		}
		r.Data["subjects"] = rv
	}
	if len(e.Resources) > 0 {
		rv := make(object.StringArray, len(e.Resources))
		for i, iv := range e.Resources {
			rv[i] = object.String(iv)
		}
		r.Data["resources"] = rv
	}
	if len(e.Conditions) > 0 {
		rv := make(object.StringArray, len(e.Conditions))
		for i, iv := range e.Conditions {
			rv[i] = object.String(iv)
		}
		r.Data["conditions"] = rv
	}
	r.Data["action"] = object.String(e.Action)
	return r
}

func (e *TestPolicy) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *TestPolicy) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *TestPolicy) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["subjects"]; ok {
		if t, ok := v.(object.StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.Subjects = rv
		}
	}
	if v, ok := o.Data["resources"]; ok {
		if t, ok := v.(object.StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.Resources = rv
		}
	}
	if v, ok := o.Data["conditions"]; ok {
		if t, ok := v.(object.StringArray); ok {
			rv := make([]string, len(t))
			for i, iv := range t {
				rv[i] = string(iv)
			}
			e.Conditions = rv
		}
	}
	if v, ok := o.Data["action"]; ok {
		if t, ok := v.(object.String); ok {
			e.Action = string(t)
		}
	}
	return nil
}

func (e *TestStream) Type() string {
	return "nimona.io/fixtures.TestStream"
}

func (e *TestStream) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *TestStream) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e TestStream) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/fixtures.TestStream",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["nonce"] = object.String(e.Nonce)
	r.Data["createdDateTime"] = object.String(e.CreatedDateTime)
	return r
}

func (e *TestStream) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *TestStream) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *TestStream) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(object.String); ok {
			e.Nonce = string(t)
		}
	}
	if v, ok := o.Data["createdDateTime"]; ok {
		if t, ok := v.(object.String); ok {
			e.CreatedDateTime = string(t)
		}
	}
	return nil
}

func (e *TestSubscribed) Type() string {
	return "nimona.io/fixtures.TestSubscribed"
}

func (e *TestSubscribed) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *TestSubscribed) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e TestSubscribed) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/fixtures.TestSubscribed",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["nonce"] = object.String(e.Nonce)
	return r
}

func (e *TestSubscribed) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *TestSubscribed) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *TestSubscribed) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(object.String); ok {
			e.Nonce = string(t)
		}
	}
	return nil
}

func (e *TestUnsubscribed) Type() string {
	return "nimona.io/fixtures.TestUnsubscribed"
}

func (e *TestUnsubscribed) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *TestUnsubscribed) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e TestUnsubscribed) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/fixtures.TestUnsubscribed",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["nonce"] = object.String(e.Nonce)
	return r
}

func (e *TestUnsubscribed) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *TestUnsubscribed) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *TestUnsubscribed) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["nonce"]; ok {
		if t, ok := v.(object.String); ok {
			e.Nonce = string(t)
		}
	}
	return nil
}

func (e *TestRequest) Type() string {
	return "nimona.io/fixtures.TestRequest"
}

func (e *TestRequest) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *TestRequest) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e TestRequest) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/fixtures.TestRequest",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["requestID"] = object.String(e.RequestID)
	r.Data["foo"] = object.String(e.Foo)
	return r
}

func (e *TestRequest) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *TestRequest) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *TestRequest) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["requestID"]; ok {
		if t, ok := v.(object.String); ok {
			e.RequestID = string(t)
		}
	}
	if v, ok := o.Data["foo"]; ok {
		if t, ok := v.(object.String); ok {
			e.Foo = string(t)
		}
	}
	return nil
}

func (e *TestResponse) Type() string {
	return "nimona.io/fixtures.TestResponse"
}

func (e *TestResponse) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *TestResponse) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e TestResponse) ToObject() *object.Object {
	r := &object.Object{
		Type:     "nimona.io/fixtures.TestResponse",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["requestID"] = object.String(e.RequestID)
	r.Data["foo"] = object.String(e.Foo)
	return r
}

func (e *TestResponse) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *TestResponse) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *TestResponse) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["requestID"]; ok {
		if t, ok := v.(object.String); ok {
			e.RequestID = string(t)
		}
	}
	if v, ok := o.Data["foo"]; ok {
		if t, ok := v.(object.String); ok {
			e.Foo = string(t)
		}
	}
	return nil
}

func (e *Parent) Type() string {
	return "parent"
}

func (e *Parent) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *Parent) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e Parent) ToObject() *object.Object {
	r := &object.Object{
		Type:     "parent",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["foo"] = object.String(e.Foo)
	if e.Child != nil {
		if v, err := e.Child.MarshalObject(); err == nil {
			r.Data["child"] = (v)
		}
	}
	if len(e.RepeatedChild) > 0 {
		rv := make(object.ObjectArray, len(e.RepeatedChild))
		for i, v := range e.RepeatedChild {
			if iv, err := v.MarshalObject(); err == nil {
				rv[i] = (iv)
			}
		}
		r.Data["repeatedChild"] = rv
	}
	return r
}

func (e *Parent) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *Parent) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *Parent) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["foo"]; ok {
		if t, ok := v.(object.String); ok {
			e.Foo = string(t)
		}
	}
	if v, ok := o.Data["child"]; ok {
		if ev, ok := v.(*object.Object); ok {
			es := &Child{}
			if err := es.UnmarshalObject((ev)); err == nil {
				e.Child = es
			}
		}
	}
	if v, ok := o.Data["repeatedChild"]; ok {
		if ev, ok := v.(object.ObjectArray); ok {
			e.RepeatedChild = make([]*Child, len(ev))
			for i, iv := range ev {
				es := &Child{}
				if err := es.UnmarshalObject((iv)); err == nil {
					e.RepeatedChild[i] = es
				}
			}
		}
	}
	return nil
}

func (e *Child) Type() string {
	return "child"
}

func (e *Child) MarshalMap() (object.Map, error) {
	return e.ToObject().Map(), nil
}

func (e *Child) MarshalObject() (*object.Object, error) {
	return e.ToObject(), nil
}

func (e Child) ToObject() *object.Object {
	r := &object.Object{
		Type:     "child",
		Metadata: e.Metadata,
		Data:     object.Map{},
	}
	r.Data["foo"] = object.String(e.Foo)
	return r
}

func (e *Child) UnmarshalMap(m object.Map) error {
	return e.FromObject(object.FromMap(m))
}

func (e *Child) UnmarshalObject(o *object.Object) error {
	return e.FromObject(o)
}

func (e *Child) FromObject(o *object.Object) error {
	e.Metadata = o.Metadata
	if v, ok := o.Data["foo"]; ok {
		if t, ok := v.(object.String); ok {
			e.Foo = string(t)
		}
	}
	return nil
}
