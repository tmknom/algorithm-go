package main

import (
	"../../util"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSliceInitInfIntSlice(t *testing.T) {
	cases := []struct {
		size int
		want []int
	}{
		{
			size: 0,
			want: []int{},
		},
		{
			size: 3,
			want: []int{util.Inf, util.Inf, util.Inf},
		},
	}

	for _, tc := range cases {
		got := util.InitInfIntSlice(tc.size)
		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("tc: %#v: diff (-got +want):\n%s", tc, diff)
		}
	}
}
