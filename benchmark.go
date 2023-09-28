package test

import easyjson "github.com/mailru/easyjson"

type testdata struct {
	name string
	cont []byte
}

var testsdata = []testdata{
	{name: "canada_geometry", cont: gunzip("testdata/canada_geometry.json.gz")},
	{name: "citm_catalog", cont: gunzip("testdata/citm_catalog.json.gz")},
	{name: "golang_source", cont: gunzip("testdata/golang_source.json.gz")},
	{name: "string_unicode", cont: gunzip("testdata/string_unicode.json.gz")},
	{name: "synthea_fhir", cont: gunzip("testdata/synthea_fhir.json.gz")},
	{name: "twitter_status", cont: gunzip("testdata/twitter_status.json.gz")},
}

var compatTab = map[string]func() any{
	"canada_geometry": func() any {
		return new(canadaRoot)
	},
	"citm_catalog": func() any {
		return new(citmRoot)
	},
	"golang_source": func() any {
		return new(golangRoot)
	},
	"string_unicode": func() any {
		return new(stringRoot)
	},
	"synthea_fhir": func() any {
		return new(syntheaRoot)
	},
	"twitter_status": func() any {
		return new(twitterRoot)
	},
}

var easyJSONTab = map[string]func() easyjson.MarshalerUnmarshaler{
	"canada_geometry": func() easyjson.MarshalerUnmarshaler {
		return new(canadaRootEasyJSON)
	},
	"citm_catalog": func() easyjson.MarshalerUnmarshaler {
		return new(citmRootEasyJSON)
	},
	"golang_source": func() easyjson.MarshalerUnmarshaler {
		return new(golangRootEasyJSON)
	},
	"string_unicode": func() easyjson.MarshalerUnmarshaler {
		return new(stringRootEasyJSON)
	},
	"synthea_fhir": func() easyjson.MarshalerUnmarshaler {
		return new(syntheaRootEasyJSON)
	},
	"twitter_status": func() easyjson.MarshalerUnmarshaler {
		return new(twitterRootEasyJSON)
	},
}
