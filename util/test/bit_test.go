package main

import (
	"../../util"
	"testing"
)

func TestBitToBinary8(t *testing.T) {
	cases := []struct {
		src  uint
		want string
	}{
		{
			src:  0,
			want: "0000 0000",
		},
		{
			src:  1,
			want: "0000 0001",
		},
		{
			src:  1 << 7,
			want: "1000 0000",
		},
		{
			src:  (1 << 8) - 1,
			want: "1111 1111",
		},
		{
			src:  1 << 8,
			want: "0000 0000",
		},
	}

	for _, tc := range cases {
		got := util.ToBinary8(tc.src)
		if got != tc.want {
			t.Errorf("got=%v, want = %v, src = %v", got, tc.want, tc.src)
		}
	}
}

func TestBitToBinary16(t *testing.T) {
	cases := []struct {
		src  uint
		want string
	}{
		{
			src:  0,
			want: "0000 0000 0000 0000",
		},
		{
			src:  1,
			want: "0000 0000 0000 0001",
		},
		{
			src:  1 << 15,
			want: "1000 0000 0000 0000",
		},
		{
			src:  (1 << 16) - 1,
			want: "1111 1111 1111 1111",
		},
		{
			src:  1 << 16,
			want: "0000 0000 0000 0000",
		},
	}

	for _, tc := range cases {
		got := util.ToBinary16(tc.src)
		if got != tc.want {
			t.Errorf("got=%v, want = %v, src = %v", got, tc.want, tc.src)
		}
	}
}

func TestBitToBinary32(t *testing.T) {
	cases := []struct {
		src  uint
		want string
	}{
		{
			src:  0,
			want: "0000 0000 0000 0000 0000 0000 0000 0000",
		},
		{
			src:  1,
			want: "0000 0000 0000 0000 0000 0000 0000 0001",
		},
		{
			src:  1 << 31,
			want: "1000 0000 0000 0000 0000 0000 0000 0000",
		},
		{
			src:  (1 << 32) - 1,
			want: "1111 1111 1111 1111 1111 1111 1111 1111",
		},
		{
			src:  1 << 32,
			want: "0000 0000 0000 0000 0000 0000 0000 0000",
		},
	}

	for _, tc := range cases {
		got := util.ToBinary32(tc.src)
		if got != tc.want {
			t.Errorf("got=%v, want = %v, src = %v", got, tc.want, tc.src)
		}
	}
}

func TestBitToBinary64(t *testing.T) {
	cases := []struct {
		src  uint
		want string
	}{
		{
			src:  0,
			want: "0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000",
		},
		{
			src:  1,
			want: "0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0001",
		},
		{
			src:  1 << 63,
			want: "1000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000",
		},
		{
			src:  (1 << 64) - 1,
			want: "1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111",
		},
	}

	for _, tc := range cases {
		got := util.ToBinary64(tc.src)
		if got != tc.want {
			t.Errorf("got=%v, want = %v, src = %v", got, tc.want, tc.src)
		}
	}
}
