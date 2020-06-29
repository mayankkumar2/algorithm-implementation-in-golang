package main

import (
	"fmt"
	"math"
	"sort"
)

/*

The following program is implementation of Algorithm to find the shortest distance between 2 points from a set of points.
Time Complexity: O( n Log n )

*/

type Point struct {
	X float64
	Y float64
}

func NotEq(a Point, b Point) bool {
	return !(a.X == b.X && a.Y == b.Y)
}

func MinDistance(p1 Point, p2 Point, p3 Point, p4 Point) (V1 Point, V2 Point, d float64) {
	d = PointDistance(p1, p2)
	V1 = p1
	V2 = p2
	//fmt.Println(p1,p2,p3,p4)
	if NotEq(p1, p2) {
		V1, V2, d = Point{}, Point{}, math.Inf(9)
	}
	if d > PointDistance(p1, p3) && NotEq(p1, p3) {
		d = PointDistance(p1, p3)
		V1 = p1
		V2 = p3
	}
	if d > PointDistance(p1, p4) && NotEq(p1, p4) {
		d = PointDistance(p1, p4)
		V1 = p1
		V2 = p4
	}
	if d > PointDistance(p2, p3) && NotEq(p2, p3) {
		d = PointDistance(p2, p3)
		V1 = p2
		V2 = p3
	}
	if d > PointDistance(p2, p4) && NotEq(p2, p4) {
		d = PointDistance(p2, p4)
		V1 = p2
		V2 = p4
	}
	if d > PointDistance(p3, p4) && NotEq(p3, p4) {
		d = PointDistance(p3, p4)
		V1 = p3
		V2 = p4
	}
	return V1, V2, d
}
func PointDistance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func ClosestPoint(Sx []Point, Sy []Point, a int, b int, v Point, w Point, d float64) (V1 Point, V2 Point, Dist float64) {
	s := b - a + 1
	if s == 1 {
		if d < math.Inf(8) {
			//fmt.Println(w,v,d)
			return w, v, d
		} else {
			//fmt.Println(w,v,math.Inf(8))
			return w, v, math.Inf(8)
		}

	} else if s == 2 {
		V1, V2, Dist = MinDistance(Sx[a-1], Sx[b-1], Sy[a-1], Sy[b-1])
		if Dist < d {
			return V1, V2, Dist
		} else {
			return v, w, d
		}
	} else {
		mid := (a + b) / 2
		Rv1, Rv2, d1 := ClosestPoint(Sx, Sy, mid+1, b, v, w, d)
		Lv1, Lv2, d2 := ClosestPoint(Sx, Sy, a, mid, v, w, d)
		if d1 > d2 {
			V1, V2, Dist = ClosestSplitPoints(Sx, Sy, a, b, Lv1, Lv2, d2)
		} else {
			V1, V2, Dist = ClosestSplitPoints(Sx, Sy, a, b, Rv1, Rv2, d1)
		}
		return V1, V2, Dist
	}
}
func ClosestSplitPoints(Sx []Point, Sy []Point, a int, b int, v Point, w Point, d float64) (v1 Point, v2 Point, D float64) {
	s := b - a + 1
	mid := Sx[s/2]
	Y := make([]Point, 0, b-a+1)
	for i := a - 1; i < b; i++ {
		if Sy[i].X >= mid.X-d && Sy[i].X <= mid.X+d {
			Y = append(Y, Sy[i])
		}
	}
	D = d
	for i, _ := range Y {
		for j := 1; j < int(math.Min(7, float64(len(Y)-i))); j++ {
			if PointDistance(Y[i], Y[i+j]) < D {
				D = PointDistance(Y[i], Y[i+j])
				v1 = Y[i]
				v2 = Y[i+j]
			}
		}
	}
	if D == d {
		return v, w, d
	}
	return v1, v2, D
}

func ShortestDistance() {
	var Px = []Point{
		{X: 91, Y: 1},
		{X: 2, Y: 29},
		{X: 12, Y: 6},
		{X: 9, Y: 54},
		{X: 5, Y: 98},
		{X: 34, Y: 12},
		{X: 53, Y: 23},
		{X: 2, Y: 56},
		{X: 4, Y: 12},
	}
	sort.Slice(Px, func(i, j int) bool {
		return Px[i].X < Px[j].X
	})
	Py := make([]Point, len(Px))
	copy(Py, Px)
	sort.Slice(Px, func(i, j int) bool {
		return Px[i].Y < Px[j].Y
	})
	fmt.Println(ClosestPoint(Px, Py, 1, len(Px), Point{}, Point{}, math.Inf(3)))
}
