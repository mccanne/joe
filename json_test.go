package joe_test

import (
	"encoding/json"
	"testing"

	"github.com/mccanne/joe"
)

func Test_Unmarshal_Undefined(t *testing.T) {
	type test struct {
		Key1 joe.JSON `json:"key1"`
		Key2 joe.JSON `json:"key2"`
	}

	t1 := test{}
	err := json.Unmarshal([]byte(`{"key1": "value1"}`), &t1)
	if err != nil {
		t.Fatal(err)
	}

	if str, _ := t1.Key1.String(); str != "value1" {
		t.Fatalf("Expected key1 to equal 'value1', got '%s'", str)
	}

	if !t1.Key2.IsUndefined() {
		t.Fatalf("Expected Key2 to be undefined, got '%v'", t1.Key2.Value())
	}

	if t1.Key2.IsNull() {
		t.Fatalf("Expected Key2 to not be null, got '%v'", t1.Key2.Value())
	}
}

func Test_Unmarshal_Null(t *testing.T) {
	type test struct {
		Key1 joe.JSON `json:"key1"`
		Key2 joe.JSON `json:"key2"`
	}

	t1 := test{}
	err := json.Unmarshal([]byte(`{"key1": "value1", "key2": null}`), &t1)
	if err != nil {
		t.Fatal(err)
	}

	if str, _ := t1.Key1.String(); str != "value1" {
		t.Fatalf("Expected key1 to equal 'value1', got '%s'", str)
	}

	if !t1.Key2.IsNull() {
		t.Fatalf("Expected Key2 to be undefined, got '%v'", t1.Key2.Value())
	}
}

func Test_Marshal_Null(t *testing.T) {
	rawJSON := `{"key1":"value1","key2":null}`
	type test struct {
		Key1 joe.JSON `json:"key1,omitempty"`
		Key2 joe.JSON `json:"key2,omitemtpy"`
	}

	t1 := test{}

	json.Unmarshal([]byte(rawJSON), &t1)
	data, err := json.Marshal(t1)
	if err != nil {
		t.Fatal(err)
	}

	if rawJSON != string(data) {
		t.Fatalf("Expected json.Marshal to equal '%s', got '%s'", rawJSON, data)
	}
}

func Test_Marshal_Undefined(t *testing.T) {
	rawJSON := `{"key1":"value1"}`
	type test struct {
		Key1 joe.JSON `json:"key1,omitempty"`
		Key2 joe.JSON `json:"key2,omitempty"`
	}

	t1 := test{}

	json.Unmarshal([]byte(rawJSON), &t1)
	_, err := json.Marshal(t1)
	if err == nil {
		t.Fatalf("Expected error to not be '%v', got '%v'", nil, err)
	}
}
