package nlp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	// if tokens != expected{ // Can't compare slices with == in go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
}
