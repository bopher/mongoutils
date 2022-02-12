package mongoutils_test

import (
	"fmt"
	"testing"

	"github.com/bopher/mongoutils"
)

func TestIn(t *testing.T) {
	if r := mongoutils.In("name", "a", "b", "c"); fmt.Sprint(r) != "map[name:map[$in:[a b c]]]" {
		t.Log(r)
		t.Fatal("Fail In")
	}
}

func TestSet(t *testing.T) {
	if r := mongoutils.Set("John"); fmt.Sprint(r) != "map[$set:John]" {
		t.Log(r)
		t.Fatal("Fail Set")
	}
}

func TestSetNested(t *testing.T) {
	if r := mongoutils.SetNested("name", "John"); fmt.Sprint(r) != "map[$set:map[name:John]]" {
		t.Log(r)
		t.Fatal("Fail SetNested")
	}
}

func TestMatch(t *testing.T) {
	if r := mongoutils.Match("John"); fmt.Sprint(r) != "map[$match:John]" {
		t.Log(r)
		t.Fatal("Fail Match")
	}
}

func TestRegex(t *testing.T) {
	if r := mongoutils.Regex("name", "John", "i"); fmt.Sprint(r) != `map[name:{"pattern": "John", "options": "i"}]` {
		t.Log(r)
		t.Fatal("Fail Regex")
	}
}
