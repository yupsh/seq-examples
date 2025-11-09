package seq_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

func ExampleSeq_step() {
	// seq 2 2 10
	gloo.MustRun(
		Seq(2.0, 2.0, 10.0),
	)
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}
