package seq_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

func ExampleSeq_step() {
	// seq 2 2 10
	yup.MustRun(
		Seq("2", "2", "10"),
	)
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}

