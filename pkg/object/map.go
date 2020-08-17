package object

import (
	"sort"
)

type Map struct {
	m mapIterator
}

type mapIterator interface {
	value(k string) Value
	iterate(func(k string, v Value) bool) bool
}

type mapPair struct {
	k      string
	v      Value
	parent mapIterator
}

func (p mapPair) value(k string) Value {
	if k == p.k {
		return p.v
	}
	if p.parent == nil {
		return nil
	}
	return p.parent.value(k)
}

func (p mapPair) iterate(f func(k string, v Value) bool) bool {
	if !f(p.k, p.v) {
		return false
	}
	if p.parent == nil {
		return true
	}
	return p.parent.iterate(f)
}

func (m Map) iterate(f func(k string, v Value) bool) bool {
	if m.m == nil {
		return true
	}
	seen := make(map[string]bool)
	return m.m.iterate(func(k string, v Value) bool {
		cont := true
		if !seen[k] {
			cont = f(k, v)
			seen[k] = true
		}
		return cont
	})
}

func (m Map) keys() []string {
	keys := []string{}
	m.m.iterate(func(k string, v Value) bool {
		keys = append(keys, k)
		return true
	})
	sort.Strings(keys)
	return keys
}

func (m Map) iterateSorted(f func(k string, v Value) bool) bool {
	if m.m == nil {
		return true
	}
	keys := m.keys()
	for _, k := range keys {
		cont := f(k, m.Value(k))
		if !cont {
			return false
		}
	}
	return true
}

func (m Map) Iterate(f func(k string, v Value) bool) {
	m.iterate(f)
}

func (m Map) Value(path string) Value {
	v, err := resolvePath(path, m)
	if err != nil {
		return nil
	}
	return v
}

func (m Map) set(k string, v Value) Map {
	return Map{
		mapPair{
			k:      k,
			v:      v,
			parent: m.m,
		},
	}
}

func (m Map) Set(path string, v Value) Map {
	nm, err := setPath(m, path, v)
	if err != nil {
		panic(err)
	}
	return nm.(Map)
}

func (m Map) PrimitiveHinted() interface{} {
	if m.m == nil {
		return nil
	}
	p := map[string]interface{}{}
	m.Iterate(func(k string, v Value) bool {
		if v != nil {
			p[k] = v.PrimitiveHinted()
		}
		return true
	})
	return p
}

func (m Map) IsEmpty() bool {
	return m.m == nil
}
