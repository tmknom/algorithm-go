package main

import (
	"../../util"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSliceInitInfIntSlice(t *testing.T) {
	cases := []struct {
		size int
		want util.SingleIntSlice
	}{
		{
			size: 0,
			want: []int{},
		},
		{
			size: 3,
			want: []int{util.PositiveInf, util.PositiveInf, util.PositiveInf},
		},
	}

	for _, tc := range cases {
		got := util.InitSingleIntSlice(tc.size, util.PositiveInf)
		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("tc: %#v: diff (-got +want):\n%s", tc, diff)
		}
	}
}
