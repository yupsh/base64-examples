package base64_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/base64"
)

func ExampleBase64_decode() {
	// echo "SGVsbG8gV29ybGQ=" | base64 -d
	gloo.MustRun(
		Base64(Decode, strings.NewReader("SGVsbG8gV29ybGQ=")),
	)
	// Output:
	// Hello World
}
