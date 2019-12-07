package day

import (
	"fmt"
	"strconv"
)

// Intersection ...
type Intersection struct {
	x, y       int
	path1Steps int
	path2Steps int
	sumSteps   int
}

func (i Intersection) String() string {
	return fmt.Sprintf("[ %d %d %d %d %d ]", i.x, i.y, i.path1Steps, i.path2Steps, i.sumSteps)
}

var intersections = make([]*Intersection, 0)

// NewIntersection ...
func NewIntersection(x, y int) {
	if x == 0 && y == 0 {
		return
	}
	for _, inter := range intersections {
		if inter.x == x && inter.y == y {
			return
		}
	}
	intersections = append(intersections, &Intersection{x: x, y: y})
}

// On ...
func (i Intersection) On(l Line) (bool, int) {
	if l.startX <= i.x && l.endX >= i.x && l.startY <= i.y && l.endY >= i.y {
		// fmt.Println(l)
		if l.reverseDirection {
			return true, ((l.endX - i.x) + (l.endY - i.y))
		}
		return true, (i.x - l.startX) + (i.y - l.startY)
	}

	return false, (l.endX - l.startX) + (l.endY - l.startY)
}

// Line ...
type Line struct {
	startX           int
	startY           int
	endX             int
	endY             int
	reverseDirection bool
}

// Horizontal ...
func (l Line) Horizontal() bool {
	return l.startY == l.endY
}

// Vertical ...
func (l Line) Vertical() bool {
	return l.startX == l.endX
}

var isPath1 = true

// Intersects ...
func (l Line) Intersects(l2 Line) (bool, int) {
	// Does lines put on themselves count as intersecting?
	if isPath1 {
		return false, 0
	}
	if l.Vertical() {
		if l.startY <= l2.startY && l.endY >= l2.endY {
			if l.startX >= l2.startX && l.startX <= l2.endX {
				NewIntersection(l.startX, l2.startY)
				return true, dst(l.startX, l2.startY)
			}
		}
	}

	if l.Horizontal() {
		if l.startX <= l2.startX && l.endX >= l2.endX {
			if l.startY >= l2.startY && l.startY <= l2.endY {
				NewIntersection(l2.startX, l.startY)
				return true, dst(l2.startX, l.startY)
			}
		}
	}

	return false, 0
}

func newLine(startX, startY, endX, endY int) Line {
	var reverseDirection bool
	if startX > endX {
		h := startX
		startX = endX
		endX = h
		reverseDirection = true
	}
	if startY > endY {
		h := startY
		startY = endY
		endY = h
		reverseDirection = true
	}

	return Line{
		startX:           startX,
		endX:             endX,
		startY:           startY,
		endY:             endY,
		reverseDirection: reverseDirection,
	}
}

// Day3Task2 ...
func Day3Task2() interface{} {
	path1 := []string{"R1009", "D335", "L942", "D733", "L398", "U204", "L521", "D347", "L720", "U586", "R708", "D746", "L292", "U416", "L824", "U20", "R359", "D828", "R716", "U895", "L498", "D671", "L325", "D68", "L667", "U134", "L435", "D44", "R801", "U654", "R188", "U542", "L785", "D318", "L806", "U602", "L465", "U239", "R21", "U571", "R653", "U436", "L52", "U380", "R446", "D960", "R598", "U590", "L47", "U972", "L565", "D281", "R790", "U493", "R864", "D396", "R652", "D775", "L939", "D284", "R554", "U629", "L842", "D837", "R554", "D795", "R880", "D301", "R948", "U974", "L10", "D898", "R588", "D743", "L334", "U59", "L413", "U511", "L132", "U771", "R628", "D805", "R465", "D561", "R18", "D169", "L580", "D99", "L508", "U964", "L870", "D230", "L472", "U897", "L85", "U306", "L103", "U322", "L637", "U464", "R129", "D514", "R454", "U479", "R801", "U18", "R929", "U181", "L113", "D770", "L173", "D124", "L122", "U481", "L666", "D942", "L534", "U608", "R90", "U576", "L641", "U249", "L857", "U197", "R783", "D92", "L938", "D192", "L698", "D862", "R995", "U12", "R766", "D323", "R934", "U315", "R956", "D234", "R983", "D246", "L153", "U26", "L779", "D628", "R174", "D385", "L758", "D486", "R132", "U414", "R915", "D511", "L152", "D309", "L708", "D755", "L679", "D166", "L699", "U734", "R55", "D224", "L582", "U798", "L348", "U219", "L304", "U621", "L788", "D538", "R781", "D509", "R486", "U581", "R759", "D892", "R16", "D552", "L82", "D618", "L309", "D610", "L645", "U146", "L328", "U569", "L307", "D385", "L249", "D231", "R928", "U681", "R384", "D337", "R715", "D798", "L788", "D604", "R517", "U766", "R368", "U430", "L49", "U236", "R621", "U656", "R997", "U268", "L18", "D789", "L935", "D87", "L670", "U35", "R463", "D71", "R268", "U728", "R693", "D863", "R656", "D654", "L350", "U796", "L72", "U562", "R56", "U10", "L651", "D751", "L557", "D518", "R901", "D741", "R787", "D332", "R723", "D980", "R206", "U670", "R645", "D927", "L641", "D863", "R478", "D568", "L858", "D990", "L124", "D864", "L162", "U361", "L407", "U674", "R508", "D284", "L675", "D794", "L138", "U55", "L781", "U37", "R956", "D364", "L111", "U721", "L91", "U559", "L852", "U351", "R994", "U446", "L162", "D345", "R92", "D941", "R572", "U185", "R615", "D590", "R459", "D313", "R127", "D315", "R96", "U751", "R210", "D620", "L790", "U826", "R410", "D652", "R549", "D698", "L805", "U814", "L364", "U905", "L96", "U997", "L689"}
	path2 := []string{"L1008", "D451", "L146", "D628", "R877", "U486", "L464", "U815", "L119", "U208", "R686", "U477", "L510", "D353", "R189", "D437", "R461", "D645", "R639", "U650", "R491", "D744", "L798", "U514", "R598", "U64", "R668", "U771", "R21", "U782", "L564", "U632", "R23", "U112", "R947", "U649", "L205", "D804", "R277", "U683", "L828", "U662", "R890", "U420", "L908", "U484", "R535", "D515", "R390", "U7", "L287", "D967", "R497", "U502", "L893", "D851", "R426", "D656", "R622", "U46", "L106", "U590", "R646", "D29", "R467", "D896", "L155", "U382", "L992", "D189", "L34", "U16", "R132", "U35", "L586", "U812", "L539", "D409", "R776", "D42", "R58", "U323", "R569", "D965", "R648", "D789", "R478", "D587", "R162", "D834", "R979", "D993", "L944", "U84", "R93", "U903", "R491", "U713", "L646", "U235", "R120", "U286", "L919", "U34", "L662", "U834", "L812", "D271", "L73", "U410", "L758", "U210", "R712", "U581", "L520", "D654", "L981", "D516", "R312", "U123", "L153", "U433", "R368", "U606", "L882", "U362", "L261", "U587", "R441", "D691", "L699", "U135", "L825", "D25", "R142", "U191", "L358", "D554", "L487", "D802", "L542", "D266", "R283", "U222", "R113", "D259", "R828", "U182", "R402", "U627", "R769", "D426", "L768", "U571", "R118", "U684", "R803", "D430", "R942", "U514", "R711", "D225", "R299", "U45", "L214", "U712", "L673", "U787", "L164", "D703", "L616", "D587", "R624", "D326", "L614", "D779", "L904", "D563", "L98", "U137", "R687", "U425", "R615", "U671", "L361", "D47", "L767", "D951", "R791", "D116", "R664", "U704", "R291", "U535", "L322", "D989", "R467", "U7", "L974", "D276", "R901", "U51", "L567", "D641", "R112", "U102", "R753", "D127", "R486", "D143", "R259", "U212", "L97", "U505", "R377", "U473", "R514", "D912", "L928", "U401", "R772", "D416", "R695", "U784", "L524", "D341", "R402", "U749", "L1", "U1", "L109", "U921", "L754", "U66", "L927", "U708", "R551", "D687", "R129", "D346", "L408", "D330", "L300", "D920", "R170", "D353", "R97", "D74", "R850", "D511", "R275", "U872", "L748", "U344", "R610", "D391", "R963", "D98", "L89", "U259", "R651", "U651", "L31", "D142", "L104", "U770", "L482", "D677", "R823", "D110", "L606", "U897", "L631", "U437", "L551", "D550", "R301", "D762", "R349", "D824", "R260", "U438", "R249", "D636", "L386", "U926", "R367", "U231", "R752", "U854", "L481", "D764", "R516", "D273", "L726", "D778", "R483", "U513", "R129", "D135", "L224"}
	// path1 = []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	// path2 = []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	// path1 = []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	// path2 = []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}

	verticalLines := []Line{}
	horizontalLines := []Line{}

	// This will generate intersections
	_, horizontalLines, verticalLines = processPath(path1, 1, horizontalLines, verticalLines, false)
	isPath1 = false
	_, _, _ = processPath(path2, 1, horizontalLines, verticalLines, true)

	isPath1 = true
	calculateSteps(path1)
	isPath1 = false
	calculateSteps(path2)

	res := intersections[0].sumSteps
	for _, inter := range intersections {
		if inter.sumSteps < res {
			res = inter.sumSteps
		}
	}
	// fmt.Println(res)
	return res
}

func calculateSteps(path []string) {
	var x, y, newX, newY int

	stepsSum := 0
	for _, p := range path {
		i, err := strconv.Atoi(p[1:])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		switch p[0] {
		case 'R':
			newX = x + i
		case 'L':
			newX = x - i
		case 'U':
			newY = y + i
		case 'D':
			newY = y - i
		}

		line := newLine(x, y, newX, newY)

		for _, intersection := range intersections {
			ok, steps := intersection.On(line)
			if steps < 0 {
				panic(fmt.Errorf("%s", intersection.String()))
			}
			if ok {
				if isPath1 && intersection.path1Steps == 0 {
					intersection.path1Steps = stepsSum + steps
					intersection.sumSteps = intersection.path1Steps + intersection.path2Steps
				} else if intersection.path2Steps == 0 {
					intersection.path2Steps = stepsSum + steps
					intersection.sumSteps = intersection.path1Steps + intersection.path2Steps
				}
			}

		}

		stepsSum += i

		x = newX
		y = newY
	}
}

// Day3Task1 ...
func Day3Task1() interface{} {
	path1 := []string{"R1009", "D335", "L942", "D733", "L398", "U204", "L521", "D347", "L720", "U586", "R708", "D746", "L292", "U416", "L824", "U20", "R359", "D828", "R716", "U895", "L498", "D671", "L325", "D68", "L667", "U134", "L435", "D44", "R801", "U654", "R188", "U542", "L785", "D318", "L806", "U602", "L465", "U239", "R21", "U571", "R653", "U436", "L52", "U380", "R446", "D960", "R598", "U590", "L47", "U972", "L565", "D281", "R790", "U493", "R864", "D396", "R652", "D775", "L939", "D284", "R554", "U629", "L842", "D837", "R554", "D795", "R880", "D301", "R948", "U974", "L10", "D898", "R588", "D743", "L334", "U59", "L413", "U511", "L132", "U771", "R628", "D805", "R465", "D561", "R18", "D169", "L580", "D99", "L508", "U964", "L870", "D230", "L472", "U897", "L85", "U306", "L103", "U322", "L637", "U464", "R129", "D514", "R454", "U479", "R801", "U18", "R929", "U181", "L113", "D770", "L173", "D124", "L122", "U481", "L666", "D942", "L534", "U608", "R90", "U576", "L641", "U249", "L857", "U197", "R783", "D92", "L938", "D192", "L698", "D862", "R995", "U12", "R766", "D323", "R934", "U315", "R956", "D234", "R983", "D246", "L153", "U26", "L779", "D628", "R174", "D385", "L758", "D486", "R132", "U414", "R915", "D511", "L152", "D309", "L708", "D755", "L679", "D166", "L699", "U734", "R55", "D224", "L582", "U798", "L348", "U219", "L304", "U621", "L788", "D538", "R781", "D509", "R486", "U581", "R759", "D892", "R16", "D552", "L82", "D618", "L309", "D610", "L645", "U146", "L328", "U569", "L307", "D385", "L249", "D231", "R928", "U681", "R384", "D337", "R715", "D798", "L788", "D604", "R517", "U766", "R368", "U430", "L49", "U236", "R621", "U656", "R997", "U268", "L18", "D789", "L935", "D87", "L670", "U35", "R463", "D71", "R268", "U728", "R693", "D863", "R656", "D654", "L350", "U796", "L72", "U562", "R56", "U10", "L651", "D751", "L557", "D518", "R901", "D741", "R787", "D332", "R723", "D980", "R206", "U670", "R645", "D927", "L641", "D863", "R478", "D568", "L858", "D990", "L124", "D864", "L162", "U361", "L407", "U674", "R508", "D284", "L675", "D794", "L138", "U55", "L781", "U37", "R956", "D364", "L111", "U721", "L91", "U559", "L852", "U351", "R994", "U446", "L162", "D345", "R92", "D941", "R572", "U185", "R615", "D590", "R459", "D313", "R127", "D315", "R96", "U751", "R210", "D620", "L790", "U826", "R410", "D652", "R549", "D698", "L805", "U814", "L364", "U905", "L96", "U997", "L689"}
	path2 := []string{"L1008", "D451", "L146", "D628", "R877", "U486", "L464", "U815", "L119", "U208", "R686", "U477", "L510", "D353", "R189", "D437", "R461", "D645", "R639", "U650", "R491", "D744", "L798", "U514", "R598", "U64", "R668", "U771", "R21", "U782", "L564", "U632", "R23", "U112", "R947", "U649", "L205", "D804", "R277", "U683", "L828", "U662", "R890", "U420", "L908", "U484", "R535", "D515", "R390", "U7", "L287", "D967", "R497", "U502", "L893", "D851", "R426", "D656", "R622", "U46", "L106", "U590", "R646", "D29", "R467", "D896", "L155", "U382", "L992", "D189", "L34", "U16", "R132", "U35", "L586", "U812", "L539", "D409", "R776", "D42", "R58", "U323", "R569", "D965", "R648", "D789", "R478", "D587", "R162", "D834", "R979", "D993", "L944", "U84", "R93", "U903", "R491", "U713", "L646", "U235", "R120", "U286", "L919", "U34", "L662", "U834", "L812", "D271", "L73", "U410", "L758", "U210", "R712", "U581", "L520", "D654", "L981", "D516", "R312", "U123", "L153", "U433", "R368", "U606", "L882", "U362", "L261", "U587", "R441", "D691", "L699", "U135", "L825", "D25", "R142", "U191", "L358", "D554", "L487", "D802", "L542", "D266", "R283", "U222", "R113", "D259", "R828", "U182", "R402", "U627", "R769", "D426", "L768", "U571", "R118", "U684", "R803", "D430", "R942", "U514", "R711", "D225", "R299", "U45", "L214", "U712", "L673", "U787", "L164", "D703", "L616", "D587", "R624", "D326", "L614", "D779", "L904", "D563", "L98", "U137", "R687", "U425", "R615", "U671", "L361", "D47", "L767", "D951", "R791", "D116", "R664", "U704", "R291", "U535", "L322", "D989", "R467", "U7", "L974", "D276", "R901", "U51", "L567", "D641", "R112", "U102", "R753", "D127", "R486", "D143", "R259", "U212", "L97", "U505", "R377", "U473", "R514", "D912", "L928", "U401", "R772", "D416", "R695", "U784", "L524", "D341", "R402", "U749", "L1", "U1", "L109", "U921", "L754", "U66", "L927", "U708", "R551", "D687", "R129", "D346", "L408", "D330", "L300", "D920", "R170", "D353", "R97", "D74", "R850", "D511", "R275", "U872", "L748", "U344", "R610", "D391", "R963", "D98", "L89", "U259", "R651", "U651", "L31", "D142", "L104", "U770", "L482", "D677", "R823", "D110", "L606", "U897", "L631", "U437", "L551", "D550", "R301", "D762", "R349", "D824", "R260", "U438", "R249", "D636", "L386", "U926", "R367", "U231", "R752", "U854", "L481", "D764", "R516", "D273", "L726", "D778", "R483", "U513", "R129", "D135", "L224"}
	path1 = []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	path2 = []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	path1 = []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	path2 = []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}

	result := 9999999999

	verticalLines := []Line{}
	horizontalLines := []Line{}

	result, horizontalLines, verticalLines = processPath(path1, result, horizontalLines, verticalLines, false)
	isPath1 = false
	result = 9999999999
	result, horizontalLines, verticalLines = processPath(path2, result, horizontalLines, verticalLines, true)

	return result
}

func processPath(steps []string, result int, horizontalLines, verticalLines []Line, second bool) (int, []Line, []Line) {
	var x, y, newX, newY int

	for _, step := range steps {
		i, err := strconv.Atoi(step[1:])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		switch step[0] {
		case 'R':
			newX = x + i
		case 'L':
			newX = x - i
		case 'U':
			newY = y + i
		case 'D':
			newY = y - i
		}

		l := newLine(x, y, newX, newY)
		if l.Horizontal() {
			result = getResult(verticalLines, l, result)
			if !second {
				horizontalLines = append(horizontalLines, l)
			}
		} else {
			result = getResult(horizontalLines, l, result)
			if !second {
				verticalLines = append(verticalLines, l)
			}
		}
		x = newX
		y = newY
	}

	return result, horizontalLines, verticalLines
}

func getResult(lines []Line, line Line, result int) int {
	for _, l := range lines {
		if intersects, r := l.Intersects(line); intersects {
			if result > r {
				result = r
			}
		}
	}

	return result
}

func dst(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x == 0 && y == 0 {
		return 999999999999
	}
	return x + y
}
