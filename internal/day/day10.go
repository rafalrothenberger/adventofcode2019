package day

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/rafalrothenberger/adventofcode2019/internal/input"
	"github.com/rafalrothenberger/adventofcode2019/internal/smath"
)

// Point ...
type Point struct {
	x, y          int
	visiblePoints int
	orgPoint      *Point
	movedPoint    *Point
	destroyed     int
}

func (p Point) alpha() float64 {
	return p.arg() + math.Pi/2.0
}

func (p Point) arg() float64 {
	a := p.x
	b := p.y
	if a == 0 {
		if b == 0 {
			panic("Trying to calculate arg of point (0,0)")
		}
		if b > 0 {
			return math.Pi / 2.0
		}
		return math.Pi / -2.0
	}
	if a > 0 {
		return math.Atan(float64(b) / float64(a))
	}
	return math.Atan(float64(b)/float64(a)) + math.Pi
}

func (p Point) distance() int {
	// As we are going to compare only points on the same line we don't care about actual distance,
	// just relative to other points on this line, so taxi metric is good enough
	x := p.movedPoint.x
	y := p.movedPoint.y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func (p Point) String() string {
	return fmt.Sprintf("{%d,%d}", p.x, p.y)
}

func (p Point) relativeTo(p2 Point) Point {
	return Point{
		x:        p.x - p2.x,
		y:        p.y - p2.y,
		orgPoint: &p,
	}
}

func (p Point) scaleCoordinates() Point {
	if p.x == 0 && p.y == 0 {
		return p
	}
	r := smath.GCD(p.x, p.y)
	return Point{
		x:          p.x / r,
		y:          p.y / r,
		orgPoint:   p.orgPoint,
		movedPoint: &p,
	}
}

// VisPoints ...
type VisPoints map[string]Point

func (vp VisPoints) String() string {
	res := []string{}
	for _, s := range vp {
		res = append(res, s.String())
	}

	return strings.Join(res, ",")
}

// Points ...
type Points []Point

func (p Points) Len() int           { return len(p) }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Points) Less(i, j int) bool { return p[i].alpha() < p[j].alpha() }

// Day10Task2 ...
func Day10Task2() interface{} {
	const (
		day      = 10
		task     = 2
		testCode = 0
	)
	station := Point{x: 17, y: 22}
	// station := Point{x: 11, y: 13}

	input := input.Get(day, task, testCode).(string)
	input = strings.Trim(input, "\n")

	lines := strings.Split(input, "\n")

	points := make([]Point, 0)
	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				points = append(points, Point{x: x, y: y})
			}
		}
	}

	relPoints := MakePointsRelative(station, points)
	visiblePoints := VisPoints{}

	// For my test data it will destroy 200th asteroid in first run, as 288 are visible from the station
	for _, p := range relPoints {
		if p.x == 0 && p.y == 0 {
			continue
		}
		pp, ok := visiblePoints[p.String()]
		if !ok {
			visiblePoints[p.String()] = p
			continue
		}
		if pp.distance() > p.distance() {
			visiblePoints[p.String()] = p
		}
	}

	destroyed := make([]Point, len(visiblePoints))
	i := 0
	for _, p := range visiblePoints {
		destroyed[i] = p
		i++
	}

	sort.Sort(Points(destroyed))

	orgPoints := make([]Point, len(destroyed))
	for i, p := range destroyed {
		orgPoints[i] = *p.orgPoint
	}

	return orgPoints[199]
}

// Day10Task1 ...
func Day10Task1() interface{} {
	const (
		day      = 10
		task     = 1
		testCode = 0
	)

	input := input.Get(day, task, testCode).(string)
	input = strings.Trim(input, "\n")

	lines := strings.Split(input, "\n")

	points := make([]Point, 0)
	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				points = append(points, Point{x: x, y: y})
			}
		}
	}

	res := 0
	resPoint := Point{}

	for _, point := range points {
		relPoints := MakePointsRelative(point, points)
		visiblePoints := VisPoints{}
		for _, p := range relPoints {
			if p.x == 0 && p.y == 0 {
				continue
			}
			visiblePoints[p.String()] = p

		}
		if len(visiblePoints) > res {
			res = len(visiblePoints)
			resPoint = point
		}
	}
	resPoint.visiblePoints = res

	return resPoint
}

// MakePointsRelative ...
func MakePointsRelative(point Point, points []Point) []Point {
	newPoints := make([]Point, len(points)-1)
	i := 0
	for _, p := range points {
		if p.x == point.x && p.y == point.y {
			continue
		}
		newPoints[i] = p.relativeTo(point).scaleCoordinates()
		i++
	}
	return newPoints
}
