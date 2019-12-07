package day

import (
	"fmt"
	"strings"

	"github.com/rafalrothenberger/adventofcode2019/internal/input"
)

// CosmicEntity ...
type CosmicEntity struct {
	orbits         int
	name           string
	parent         *CosmicEntity
	children       []*CosmicEntity
	distanceFromMe int
}

// String ...
func (e *CosmicEntity) String() string {
	return fmt.Sprintf("{%s %d}", e.name, e.distanceFromMe)
}

// Day6Task1 ...
func Day6Task1() interface{} {
	const (
		day      = 6
		task     = 1
		testCode = 0
	)

	input := input.Get(day, task, testCode).([]string)
	sum, _ := orbits(input)
	return sum
}

// Day6Task2 ...
func Day6Task2() interface{} {
	const (
		day      = 6
		task     = 2
		testCode = 0
	)

	input := input.Get(day, task, testCode).([]string)

	_, entities := orbits(input)
	findMeAndCalculateDistances(entities)
	return findSantaAndCalculateDistanceFromMe(entities)
}

func reverse(a []*CosmicEntity) []*CosmicEntity {
	b := make([]*CosmicEntity, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		b[len(a)-i-1] = a[i]
	}
	return b
}

func findSantaAndCalculateDistanceFromMe(entities []*CosmicEntity) int {
	// sum := 0
	var santa *CosmicEntity
	for _, e := range entities {
		if e.name == "SAN" {
			santa = e
			break
		}
	}

	var help *CosmicEntity
	help = santa.parent
	dst := 0
	for true {
		if help == nil {
			break
		}
		if help.distanceFromMe != -1 {
			break
		}
		help = help.parent
		dst++
	}

	return dst + help.distanceFromMe
}

func findMeAndCalculateDistances(entities []*CosmicEntity) {
	var me *CosmicEntity
	for _, e := range entities {
		if e.name == "YOU" {
			me = e
			break
		}
	}
	var help *CosmicEntity
	help = me.parent
	dst := 0
	for true {
		help.distanceFromMe = dst
		if help.parent == nil {
			break
		}
		help = help.parent
		dst++
	}
}

func orbits(input []string) (int, []*CosmicEntity) {
	entities := []*CosmicEntity{
		{
			name:   "COM",
			orbits: 0,
		},
	}
	sum := 0
	newInput := []string{}
	for true {
		for _, i := range input {
			parent := strings.Split(i, ")")[0]
			entity := strings.Split(i, ")")[1]
			found := false
			for _, e := range entities {
				if e.name == parent {
					en := &CosmicEntity{
						name:           entity,
						orbits:         e.orbits + 1,
						parent:         e,
						distanceFromMe: -1,
					}
					entities = append(entities, en)
					e.children = append(e.children, en)
					sum += e.orbits + 1
					found = true
					break
				}
			}
			if !found {
				newInput = append(newInput, i)
			}
		}
		if len(newInput) == 0 {
			break
		}
		input = newInput
		newInput = make([]string, 0)
	}

	// return GetOrbits(entities[0], 0)
	return sum, entities
}
