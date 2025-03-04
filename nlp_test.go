package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tokenizeCases = []struct { // anonymous struct
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "s", "on", "first"}},
	{"", nil},
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range tokenizeCases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	/* // Before testify
	// if tokens != expected{ // Can't compare slices with == in go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}
