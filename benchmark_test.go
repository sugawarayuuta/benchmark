package benchmark

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/bytedance/sonic"
	decoder "github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	exp "github.com/go-json-experiment/json"
	goccy "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	segment "github.com/segmentio/encoding/json"
	"github.com/sugawarayuuta/sonnet"
)

type (
	tester struct {
		name string
		new  func() any
		data []byte
	}
	library struct {
		name      string
		decode    func(io.Reader, any) error
		unmarshal func([]byte, any) error
		encode    func(io.Writer, any) error
		marshal   func(any) ([]byte, error)
		noStream  bool
	}
)

var (
	testers = []tester{
		{name: "canada_geometry", new: func() any { return new(canadaRoot) }, data: gunzip("testdata/canada_geometry.json.gz")},
		{name: "citm_catalog", new: func() any { return new(citmRoot) }, data: gunzip("testdata/citm_catalog.json.gz")},
		{name: "synthea_fhir", new: func() any { return new(syntheaRoot) }, data: gunzip("testdata/synthea_fhir.json.gz")},
		{name: "twitter_status", new: func() any { return new(twitterRoot) }, data: gunzip("testdata/twitter_status.json.gz")},
		{name: "golang_source", new: func() any { return new(golangRoot) }, data: gunzip("testdata/golang_source.json.gz")},
		{name: "string_unicode", new: func() any { return new(stringRoot) }, data: gunzip("testdata/string_unicode.json.gz")},
	}
	libraries = []library{
		{
			name:      "encoding/json",
			decode:    func(reader io.Reader, val any) error { return json.NewDecoder(reader).Decode(val) },
			unmarshal: json.Unmarshal,
			encode:    func(writer io.Writer, val any) error { return json.NewEncoder(writer).Encode(val) },
			marshal:   json.Marshal,
		},
		{
			name:      "json-iterator/go",
			decode:    func(reader io.Reader, val any) error { return jsoniter.NewDecoder(reader).Decode(val) },
			unmarshal: jsoniter.Unmarshal,
			encode:    func(writer io.Writer, val any) error { return jsoniter.NewEncoder(writer).Encode(val) },
			marshal:   jsoniter.Marshal,
		},
		{
			name:      "goccy/go-json",
			decode:    func(reader io.Reader, val any) error { return goccy.NewDecoder(reader).Decode(val) },
			unmarshal: goccy.Unmarshal,
			encode:    func(writer io.Writer, val any) error { return goccy.NewEncoder(writer).Encode(val) },
			marshal:   goccy.Marshal,
		},
		{
			name:      "segmentio/encoding",
			decode:    func(reader io.Reader, val any) error { return segment.NewDecoder(reader).Decode(val) },
			unmarshal: segment.Unmarshal,
			encode:    func(writer io.Writer, val any) error { return segment.NewEncoder(writer).Encode(val) },
			marshal:   segment.Marshal,
		},
		{
			name:      "bytedance/sonic",
			decode:    func(reader io.Reader, val any) error { return decoder.NewStreamDecoder(reader).Decode(val) },
			unmarshal: sonic.Unmarshal,
			encode:    func(writer io.Writer, val any) error { return encoder.NewStreamEncoder(writer).Encode(val) },
			marshal:   sonic.Marshal,
		},
		{
			name:      "sugawarayuuta/sonnet",
			decode:    func(reader io.Reader, val any) error { return sonnet.NewDecoder(reader).Decode(val) },
			unmarshal: sonnet.Unmarshal,
			encode:    func(writer io.Writer, val any) error { return sonnet.NewEncoder(writer).Encode(val) },
			marshal:   sonnet.Marshal,
		},
		{
			name:      "go-json-experiment/json",
			unmarshal: exp.Unmarshal,
			marshal:   exp.Marshal,
			noStream:  true,
		},
	}
)

func gunzip(path string) []byte {
	fl, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fl.Close()

	gz, err := gzip.NewReader(fl)
	if err != nil {
		panic(err)
	}
	defer gz.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func Benchmark(bench *testing.B) {
	for _, tst := range testers {
		for _, lib := range libraries {
			val := tst.new()
			err := lib.unmarshal(tst.data, val)
			if err != nil {
				bench.Fatal(err)
			}

			bench.Run(tst.name+"/"+lib.name+"/marshal", func(bench *testing.B) {
				bench.ReportAllocs()
				for idx := 0; idx < bench.N; idx++ {
					_, err := lib.marshal(val)
					if err != nil {
						bench.Fatal(err)
					}
				}
			})

			bench.Run(tst.name+"/"+lib.name+"/unmarshal", func(bench *testing.B) {
				bench.ReportAllocs()
				for idx := 0; idx < bench.N; idx++ {
					err := lib.unmarshal(tst.data, tst.new())
					if err != nil {
						bench.Fatal(err)
					}
				}
			})

			if lib.noStream {
				continue
			}

			bench.Run(tst.name+"/"+lib.name+"/encoder", func(bench *testing.B) {
				bench.ReportAllocs()
				for idx := 0; idx < bench.N; idx++ {
					err := lib.encode(io.Discard, val)
					if err != nil {
						bench.Fatal(err)
					}
				}
			})

			bench.Run(tst.name+"/"+lib.name+"/decoder", func(bench *testing.B) {
				bench.ReportAllocs()
				var reader bytes.Reader
				for idx := 0; idx < bench.N; idx++ {
					reader.Reset(tst.data)
					err := lib.decode(&reader, val)
					if err != nil {
						bench.Fatal(err)
					}
				}
			})
		}
	}
}
