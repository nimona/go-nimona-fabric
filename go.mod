module nimona.io

require (
	github.com/DataDog/zstd v1.4.0 // indirect
	github.com/Sereal/Sereal v0.0.0-20190614071512-cf1bab6c7a3a // indirect
	github.com/asdine/storm v2.2.1+incompatible
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/cayleygraph/cayley v0.7.5
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/d4l3k/messagediff v1.2.1 // indirect
	github.com/dlclark/regexp2 v1.1.6 // indirect
	github.com/dop251/goja v0.0.0-20190614071512-1a71e42e74ec // indirect
	github.com/emersion/go-upnp-igd v0.0.0-20170924120501-6fb51d2a2a53
	github.com/fatih/structs v1.1.0
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sourcemap/sourcemap v2.1.2+incompatible // indirect
	github.com/go-test/deep v1.0.1 // indirect
	github.com/gobwas/glob v0.2.3
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gonum/floats v0.0.0-20181209220543-c233463c7e82 // indirect
	github.com/gonum/internal v0.0.0-20181124074243-f884aa714029 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/james-bowman/sparse v0.0.0-20190515093507-80c6877364c7
	github.com/jinzhu/copier v0.0.0-20180308034124-7e38e58719c3
	github.com/joeycumines/go-dotnotation v0.0.0-20180131115956-2d3612e36c5d
	github.com/mitchellh/mapstructure v1.1.2
	github.com/mr-tron/base58 v1.1.2
	github.com/pkg/errors v0.8.1
	github.com/remyoudompheng/bigfft v0.0.0-20190515093507-babf20351dd7 // indirect
	github.com/spaolacci/murmur3 v1.1.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/tylertreat/BoomFilters v0.0.0-20181028192813-611b3dbe80e8 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522 // indirect
	golang.org/x/net v0.0.0-20190614071512-d28f0bde5980
	golang.org/x/sys v0.0.0-20190614071512-5ed2794edfdc // indirect
	gonum.org/v1/gonum v0.0.0-20190515092121-7e53d113562e
	gonum.org/v1/netlib v0.0.0-20190331212654-76723241ea4e // indirect
	google.golang.org/appengine v1.6.1 // indirect
)

replace (
	nimona.io/cmd/nimona => ./cmd/nimona
	nimona.io/tools/community => ./tools/community
	nimona.io/tools/objectify => ./tools/objectify
	nimona.io/tools/proxy => ./tools/proxy
	nimona.io/tools/vanity => ./tools/vanity
)

replace (
	github.com/ugorji/go/codec => github.com/ugorji/go v1.1.2
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
