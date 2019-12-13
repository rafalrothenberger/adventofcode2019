package day

import (
	"fmt"
	"strings"

	"github.com/rafalrothenberger/adventofcode2019/internal/smath"
)

// Moon ...
type Moon struct {
	x, y, z  int
	velocity *Velocity
}

func (m *Moon) String() string {
	return fmt.Sprintf("%d %d %d %s", m.x, m.y, m.z, m.velocity)
}

// StringX ...
func (m Moon) StringX() string {
	return fmt.Sprintf("%d %d", m.x, m.velocity.x)
}

// StringY ...
func (m Moon) StringY() string {
	return fmt.Sprintf("%d %d", m.y, m.velocity.y)
}

// StringZ ...
func (m Moon) StringZ() string {
	return fmt.Sprintf("%d %d", m.z, m.velocity.z)
}

func (m *Moon) potential() int {
	return smath.Abs(m.x) + smath.Abs(m.y) + smath.Abs(m.z)
}

func (m *Moon) kinetic() int {
	return m.velocity.kinetic()
}

func (m *Moon) total() int {
	return m.kinetic() * m.potential()
}

func (m *Moon) applyVelocity() {
	m.x += m.velocity.x
	m.y += m.velocity.y
	m.z += m.velocity.z
}

func (m *Moon) gravity(m1 *Moon) {
	mod := 1
	if m.x == m1.x {
		goto Y
	}
	mod = 1
	if m.x > m1.x {
		mod *= -1
	}
	m.velocity.x += mod
	m1.velocity.x -= mod
Y:
	if m.y == m1.y {
		goto Z
	}
	mod = 1
	if m.y > m1.y {
		mod *= -1
	}
	m.velocity.y += mod
	m1.velocity.y -= mod
Z:
	if m.z == m1.z {
		return
	}
	mod = 1
	if m.z > m1.z {
		mod *= -1
	}
	m.velocity.z += mod
	m1.velocity.z -= mod
}

// Velocity ...
type Velocity struct {
	x, y, z int
}

func (v *Velocity) kinetic() int {
	return smath.Abs(v.x) + smath.Abs(v.y) + smath.Abs(v.z)
}

func (v *Velocity) String() string {
	return fmt.Sprintf("%d %d %d", v.x, v.y, v.z)
}

func data() []*Moon {
	// <x=10, y=15, z=7>
	// <x=15, y=10, z=0>
	// <x=20, y=12, z=3>
	// <x=0, y=-3, z=13>
	return []*Moon{
		{
			x:        10,
			y:        15,
			z:        7,
			velocity: &Velocity{},
		},
		{
			x:        15,
			y:        10,
			z:        0,
			velocity: &Velocity{},
		},
		{
			x:        20,
			y:        12,
			z:        3,
			velocity: &Velocity{},
		},
		{
			x:        0,
			y:        -3,
			z:        13,
			velocity: &Velocity{},
		},
	}
}

// Moons ...
type Moons []*Moon

func (m Moons) String() string {
	res := ""
	for _, moon := range m {
		res += fmt.Sprintf("%s ", moon)
	}

	return strings.Trim(res, " ")
}

// StringX ...
func (m Moons) StringX() string {
	res := ""
	for _, moon := range m {
		res += fmt.Sprintf("%s ", moon.StringX())
	}

	return strings.Trim(res, " ")
}

// StringY ...
func (m Moons) StringY() string {
	res := ""
	for _, moon := range m {
		res += fmt.Sprintf("%s ", moon.StringY())
	}

	return strings.Trim(res, " ")
}

// StringZ ...
func (m Moons) StringZ() string {
	res := ""
	for _, moon := range m {
		res += fmt.Sprintf("%s ", moon.StringZ())
	}

	return strings.Trim(res, " ")
}

// CoordinatesPeriods ...
type CoordinatesPeriods struct {
	x, y, z                   int
	offsetX, offsetY, offsetZ int
}

func (c *CoordinatesPeriods) allData() bool {
	return c.x*c.y*c.z > 0
}

func (c *CoordinatesPeriods) setX(x1, x2 int) {
	if c.x == 0 {
		c.x = x2 - x1
		c.offsetX = x1
	}
}

func (c *CoordinatesPeriods) setY(y1, y2 int) {
	if c.y == 0 {
		c.y = y2 - y1
		c.offsetY = y1
	}
}

func (c *CoordinatesPeriods) setZ(z1, z2 int) {
	if c.z == 0 {
		c.z = z2 - z1
		c.offsetZ = z1
	}
}

// Day12Task2 ...
func Day12Task2() interface{} {
	moons := Moons(data())
	statesX := map[string]int{}
	statesY := map[string]int{}
	statesZ := map[string]int{}
	cp := &CoordinatesPeriods{}
	for i := 0; true; i++ {
		if first, ok := statesX[moons.StringX()]; ok {
			cp.setX(first, i)
		}
		statesX[moons.StringX()] = i

		if first, ok := statesY[moons.StringY()]; ok {
			cp.setY(first, i)
		}
		statesY[moons.StringY()] = i

		if first, ok := statesZ[moons.StringZ()]; ok {
			cp.setZ(first, i)
		}
		statesZ[moons.StringZ()] = i

		if cp.allData() {
			return smath.LCM(cp.x, smath.LCM(cp.y, cp.z))
		}

		updateVelocity(moons)
		for _, m := range moons {
			m.applyVelocity()
		}
	}
	return nil
}

// Day12Task1 ...
func Day12Task1() interface{} {
	moons := data()
	for i := 0; i < 1000; i++ {
		updateVelocity(moons)
		for _, m := range moons {
			m.applyVelocity()
		}
	}

	res := 0
	for _, m := range moons {
		res += m.total()
	}
	return res
}

func updateVelocity(moons []*Moon) []*Moon {
	for i := range moons {
		for j := i; j < len(moons); j++ {
			moons[i].gravity(moons[j])
		}
	}
	return moons
}

func gravity(m1, m2 *Moon) {

}
