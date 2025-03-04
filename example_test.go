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
	// [who  on first]

}

/*
Test discovery:
For every file ending with _test.go, run every function that matches either:
- Example[A-Z_].*, body must include // Output: Comment
- Test[A-Z_].*
*/
