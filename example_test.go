package nlp_test

import (
	"fmt"

	"github.com/parjinderpannu/nlp"
)

func ExampleTokenize() {
	text := "Who's on first"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// output:
	// [who s on first]

}
