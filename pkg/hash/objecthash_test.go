package hash

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"nimona.io/pkg/object"
)

func TestObjectHash(t *testing.T) {
	v := map[string]interface{}{
		"str:s": "foo",
	}

	kh := hash(object.HintString, []byte("str:s"))
	vh := hash(object.HintString, []byte("foo"))
	ob := append(kh, vh...)
	oh := hash(object.HintObject, ob)

	o := object.FromMap(v)
	h := New(o)
	assert.Equal(t, oh, h.D)
}

func TestObjectHashWithSignature(t *testing.T) {
	v := map[string]interface{}{
		"str:s": "foo",
		"@signature:o": map[string]string{
			"foo:s": "bar",
		},
	}

	kh := hash(object.HintString, []byte("str:s"))
	vh := hash(object.HintString, []byte("foo"))
	ob := append(kh, vh...)
	oh := hash(object.HintObject, ob)

	o := object.FromMap(v)
	h := New(o)
	assert.Equal(t, oh, h.D)
}

func TestObjectHashDocs(t *testing.T) {
	v := map[string]interface{}{
		"some-string": "bar",
		"nested-object": map[string]interface{}{
			"unsigned-number-one": 1,
			"array-of-ints:ai":    []int{-1, 0, 1},
		},
	}

	o := object.FromMap(v)
	h := New(o)
	assert.NotNil(t, h)
}

func TestLongObjectHash(t *testing.T) {
	v := map[string]interface{}{
		"i:i":     int(math.MaxInt32),
		"i8:i":    int8(math.MaxInt8),
		"i16:i":   int16(math.MaxInt16),
		"i32:i":   int32(math.MaxInt32),
		"i64:i":   int64(math.MaxInt64),
		"u:u":     uint(math.MaxUint32),
		"u8:u":    uint8(math.MaxUint8),
		"u16:u":   uint16(math.MaxUint16),
		"u32:u":   uint32(math.MaxUint32),
		"f32:f":   float32(math.MaxFloat32),
		"f64:f":   float64(math.MaxFloat64),
		"Ai8:ai":  []int8{math.MaxInt8, math.MaxInt8 - 1},
		"Ai16:ai": []int16{math.MaxInt16, math.MaxInt16 - 1},
		"Ai32:ai": []int32{math.MaxInt32, math.MaxInt32 - 1},
		"Ai64:ai": []int64{math.MaxInt64, math.MaxInt64 - 1},
		"Au16:au": []uint16{math.MaxUint16, math.MaxUint16 - 1},
		"Au32:au": []uint32{math.MaxUint32, math.MaxUint32 - 1},
		"Af32:af": []float32{math.MaxFloat32, math.MaxFloat32 - 1},
		"Af64:af": []float64{math.MaxFloat64, math.MaxFloat64 - 1},
		"AAi:aai": [][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		"AAf:aaf": [][]float32{
			[]float32{math.MaxFloat32, math.MaxFloat32 - 1},
			[]float32{math.MaxFloat32, math.MaxFloat32 - 1},
		},
		"O:o": map[string]interface{}{
			"s:s": "foo",
			"u:u": uint64(12),
		},
		"bool:b": true,
	}

	o := object.FromMap(v)
	h := New(o)
	assert.NotNil(t, h)
}

func TestLongObjectHashInterfaces(t *testing.T) {
	v := map[string]interface{}{
		"I:i":   1,
		"Ai:ai": []interface{}{1, 2},
		"S:s":   "a",
		"As:as": []interface{}{"a", "b"},
	}

	o := object.FromMap(v)
	h := New(o)
	assert.NotNil(t, h)

	b := `{"I:i":1,"Ai:ai":[1,2],"S:s":"a","As:as":["a","b"]}` // nolint
	nv := map[string]interface{}{}
	json.Unmarshal([]byte(b), &nv) // nolint

	no := object.FromMap(nv)
	nh := New(no)
	assert.NotNil(t, nh)

	assert.Equal(t, h, nh)
}