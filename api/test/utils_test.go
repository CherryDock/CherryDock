package test

import (
	"github.com/CherryDock/CherryDock/api/jsonutils"
	"testing"
)

func TestFormatToJson(t *testing.T) {
	type TestStruct struct {
		Val1 float64
		Val2 string
	}

	test := TestStruct{1.0, "test"}

	formated := jsonutils.FormatToJson(test)
	var expected = `{"Val1":1,"Val2":"test"}`

	if string(formated) != expected {
		t.Fatalf("Error, output should be %s", expected)
	}
}

func TestParseJson(t *testing.T) {
	type TestStruct struct {
		Val1 float64
		Val2 string
	}

	test := TestStruct{1.0, "test"}

	formated := jsonutils.FormatToJson(test)

	parsed := jsonutils.ParseJson(formated)

	if parsed["Val1"] != 1.0 || parsed["Val2"] != "test" {
		t.Fatal("Fail to parse to json")
	}

}
