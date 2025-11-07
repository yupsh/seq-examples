package seq_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

func ExampleSeq_range() {
	// seq 3 7
	yup.MustRun(
		Seq("3", "7"),
	)
	// Output:
	// 3
	// 4
	// 5
	// 6
	// 7
}

