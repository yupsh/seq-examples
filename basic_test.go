package seq_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

func ExampleSeq_basic() {
	// seq 5
	gloo.MustRun(
		Seq(5.0),
	)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
