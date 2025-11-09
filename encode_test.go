package base64_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/base64"
)

func ExampleBase64_encode() {
	// echo "Hello World" | base64
	gloo.MustRun(
		Base64(strings.NewReader("Hello World")),
	)
	// Output:
	// SGVsbG8gV29ybGQ=
}
