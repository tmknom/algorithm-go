package main

import (
	"../../util"
	"math"
	"testing"
)

func TestMathIntMax(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    100,
			y:    10,
			want: 100,
		},
		{
			x:    100,
			y:    1000,
			want: 1000,
		},
		{
			x:    10000,
			y:    10000,
			want: 10000,
		},
		{
			x:    math.MaxInt32,
			y:    math.MaxInt32 << 32,
			want: math.MaxInt32 << 32,
		},
	}

	for _, tc := range cases {
		got := util.IntMax(tc.x, tc.y)
		if got != tc.want {
			t.Errorf("got=%v, tc = %#v", got, tc)
		}
	}
}

func TestMathIntMin(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    100,
			y:    10,
			want: 10,
		},
		{
			x:    100,
			y:    1000,
			want: 100,
		},
		{
			x:    10000,
			y:    10000,
			want: 10000,
		},
		{
			x:    math.MaxInt32,
			y:    math.MaxInt32 << 32,
			want: math.MaxInt32,
		},
	}

	for _, tc := range cases {
		got := util.IntMin(tc.x, tc.y)
		if got != tc.want {
			t.Errorf("got=%v, tc = %#v", got, tc)
		}
	}
}

func TestMathIntAbs(t *testing.T) {
	cases := []struct {
		val  int
		want int
	}{
		{
			val:  0,
			want: 0,
		},
		{
			val:  100,
			want: 100,
		},
		{
			val:  -100,
			want: 100,
		},
		{
			val:  math.MaxInt32,
			want: math.MaxInt32,
		},
		{
			val:  math.MinInt32 + 1,
			want: math.MaxInt32,
		},
	}

	for _, tc := range cases {
		got := util.IntAbs(tc.val)
		if got != tc.want {
			t.Errorf("got=%v, tc = %#v", got, tc)
		}
	}
}

func TestMathIntChooseMax(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    100,
			y:    10,
			want: 100,
		},
		{
			x:    100,
			y:    1000,
			want: 1000,
		},
		{
			x:    10000,
			y:    10000,
			want: 10000,
		},
		{
			x:    math.MaxInt32,
			y:    math.MaxInt32 << 32,
			want: math.MaxInt32 << 32,
		},
	}

	for _, tc := range cases {
		px := tc.x
		util.IntChooseMax(&px, tc.y)
		if px != tc.want {
			t.Errorf("got=%v, tc = %#v", px, tc)
		}
	}
}

func TestMathIntChooseMin(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    100,
			y:    10,
			want: 10,
		},
		{
			x:    100,
			y:    1000,
			want: 100,
		},
		{
			x:    10000,
			y:    10000,
			want: 10000,
		},
		{
			x:    math.MaxInt32,
			y:    math.MaxInt32 << 32,
			want: math.MaxInt32,
		},
	}

	for _, tc := range cases {
		px := tc.x
		util.IntChooseMin(&px, tc.y)
		if px != tc.want {
			t.Errorf("got=%v, tc = %#v", px, tc)
		}
	}
}
