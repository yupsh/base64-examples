package base64_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/base64"
)

func ExampleBase64_decode() {
	// echo "SGVsbG8gV29ybGQ=" | base64 -d
	yup.MustRun(
		Base64(Decode, strings.NewReader("SGVsbG8gV29ybGQ=")),
	)
	// Output:
	// Hello World
}

