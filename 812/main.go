package main

// 812. Largest Triangle Area
// Given an array of points on the X-Y plane points where points[i] = [xi, yi], return the area of the largest triangle
// that can be formed by any three different points. Answers within 10-5 of the actual answer will be accepted.

func largestTriangleArea(points [][]int) float64 {
	var s float64 = 0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				s = max(s, area(points[i], points[j], points[k]))
			}
		}
	}
	return s
}

func area(a []int, b []int, c []int) float64 {
	s := float64(0.5) * float64((a[0]-c[0])*(b[1]-a[1])-(a[0]-b[0])*(c[1]-a[1]))
	if s < 0 {
		return -s
	} else {
		return s
	}
}
