package test

import (
	jsonv1 "encoding/json"
	"path"
	"testing"

	"github.com/bytedance/sonic"
	jsonv2 "github.com/go-json-experiment/json"
	goccy "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/minio/simdjson-go"
	segment "github.com/segmentio/encoding/json"
	"github.com/sugawarayuuta/sonnet"
	"github.com/valyala/fastjson"
	"github.com/wI2L/jettison"
)

func BenchmarkMarshal(b *testing.B) {
	for _, data := range testsdata {
		data := data
		b.Run(path.Join(data.name, "encoding/json"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := jsonv1.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := jsonv1.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "go-json-experiment/json"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := jsonv2.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := jsonv2.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "json-iterator/go"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := jsoniter.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := jsoniter.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "segmentio/encodinng"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := segment.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := segment.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "goccy/go-json"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := goccy.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := goccy.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "bytedance/sonic"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := sonic.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := sonic.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "sugawarayuuta/sonnet"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := sonnet.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := sonnet.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "mailru/easyjson"), func(b *testing.B) {
			fnc := easyJSONTab[data.name]
			val := fnc()
			err := easyjson.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := easyjson.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "wI2L/jettison"), func(b *testing.B) {
			fnc := compatTab[data.name]
			val := fnc()
			err := jsonv1.Unmarshal(data.cont, val)
			if err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := jettison.Marshal(val)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for _, data := range testsdata {
		data := data
		b.Run(path.Join(data.name, "encoding/json"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := jsonv1.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "go-json-experiment/json"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := jsonv2.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "json-iterator/go"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := jsoniter.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "segmentio/encodinng"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := segment.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "goccy/go-json"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := goccy.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "bytedance/sonic"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := sonic.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "sugawarayuuta/sonnet"), func(b *testing.B) {
			fnc := compatTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := sonnet.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "mailru/easyjson"), func(b *testing.B) {
			fnc := easyJSONTab[data.name]
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				err := easyjson.Unmarshal(data.cont, fnc())
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "minio/simdjson-go"), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := simdjson.Parse(data.cont, nil)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(path.Join(data.name, "valyala/fastjson"), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(data.cont)))
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				_, err := fastjson.ParseBytes(data.cont)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
