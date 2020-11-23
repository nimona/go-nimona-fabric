package blob_test

import (
	"testing"

	"github.com/docker/go-units"

	"nimona.io/internal/iotest"
	"nimona.io/pkg/blob"
)

func BenchmarkToBlob1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fr := iotest.ZeroReader(1 * units.MB)
		_, err := blob.ToBlob(fr)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkToBlob100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fr := iotest.ZeroReader(100 * units.MB)
		_, err := blob.ToBlob(fr)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkToBlob1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fr := iotest.ZeroReader(1000 * units.MB)
		_, err := blob.ToBlob(fr)
		if err != nil {
			b.Fail()
		}
	}
}