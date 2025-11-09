package seq_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

func ExampleSeq_range() {
	// seq 3 7
	gloo.MustRun(
		Seq(3.0, 7.0),
	)
	// Output:
	// 3
	// 4
	// 5
	// 6
	// 7
}
