package seq_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

func ExampleSeq_basic() {
	// seq 5
	yup.MustRun(
		Seq("5"),
	)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

