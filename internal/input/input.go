package input

// Get ...
func Get(day, task, testCode int) interface{} {
	switch day*100 + task*10 + testCode {
	case 100, 110, 120:
		return day1
	case 200, 210, 220:
		return []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 19, 10, 23, 2, 10, 23, 27, 1, 27, 6, 31, 1, 13, 31, 35, 1, 13, 35, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 9, 51, 2, 51, 13, 55, 1, 5, 55, 59, 2, 59, 9, 63, 1, 13, 63, 67, 2, 13, 67, 71, 1, 71, 5, 75, 2, 75, 13, 79, 1, 79, 6, 83, 1, 83, 5, 87, 2, 87, 6, 91, 1, 5, 91, 95, 1, 95, 13, 99, 2, 99, 6, 103, 1, 5, 103, 107, 1, 107, 9, 111, 2, 6, 111, 115, 1, 5, 115, 119, 1, 119, 2, 123, 1, 6, 123, 0, 99, 2, 14, 0, 0}
	case 500, 510, 520:
		return []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 61, 45, 225, 102, 94, 66, 224, 101, -3854, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1101, 31, 30, 225, 1102, 39, 44, 224, 1001, 224, -1716, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 92, 41, 225, 101, 90, 40, 224, 1001, 224, -120, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 223, 224, 223, 1101, 51, 78, 224, 101, -129, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 1, 170, 13, 224, 101, -140, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1101, 14, 58, 225, 1102, 58, 29, 225, 1102, 68, 70, 225, 1002, 217, 87, 224, 101, -783, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1101, 19, 79, 225, 1001, 135, 42, 224, 1001, 224, -56, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 2, 139, 144, 224, 1001, 224, -4060, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 9, 51, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 329, 101, 1, 223, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 344, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 359, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 374, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 419, 1001, 223, 1, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 1001, 223, 1, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 494, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 509, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 524, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 554, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 629, 1001, 223, 1, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 644, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}
	case 600, 610, 620:
		return day6
	case 700, 710, 720:
		return []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 55, 76, 89, 106, 187, 268, 349, 430, 99999, 3, 9, 101, 4, 9, 9, 1002, 9, 2, 9, 101, 5, 9, 9, 1002, 9, 2, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 1002, 9, 4, 9, 101, 2, 9, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 1001, 9, 4, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
	case 701, 711, 721:
		return []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	case 702, 712, 722:
		return []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}
	case 800, 810, 820:
		return "221222001022120222222222212222220200222222222222122222222220222222212222222122022222202222222222212212222222222220222022222021022022222222202222222220222222102222020222222222212222220222222222222222122222222220222222222222222122222221222222222222202222222222222220222222222021022122222222212222222222221222200022022222222222222222221221222222222222122222222222222222222222222022122220202222222222202222222222222221222222222122022122222222102222222222222222112022220222222222202222221201222222222222222222222222222222202222222022022221212222222222222202222222222222222022222022222222222222012222222221222222122222121222022222212222220200222222222222022222222220222222222222222221122222212222222222222222222222222222222222222122222122222222002222222220220222021022122222022222202222220220222222222222022222222220222222202222022020022221222222222222212222222202222222222222222121222022222222122222222221222222020122021222222222212222222212222222222222222222222221222222222222122022122222212222222222212202222212222220222222222220122222222222022222222222222222012122021222122222202222222212222222222222022222222221222222222222022021022222202222222222212222222222222222222122222222122022222222012222222222221222001022122222222222212222221221222222222222222222222220222222202222022120022220212222222222202212222212222221222022222020122122222222012222222221221222211022022222222222212222222220222222222222022222222222222222222222122022122220222222222222212212222212222221222022222021122022222222212222222221222222010222122222022222212222220201222222222222121222222221222222202222122021022221202222222222212212222202221220222122222021122122222222212222022220221222110122220222222222222222221200222222222222022222222220222222222222022222122220202222222222202212222212222220222222222221022022222222022222122220220222002022121222022222202222221211222222222222021222222220222222222222222221122221222222222222202202222212222222222122222120222122222222002212222221222222002222121222122222202222221210222222222222220222222221222222202222022121122222222222222222202222222212221221222122222122222222222222202222122022221222101022021222022222202222220201222222222222021222222220222222222222222120122220202222222222222212222212221222222022222220122022222222202200222220222222200122020222022222202222200201022222222222020222222221222222212212022121222220212222222222212022222202222222222122222222022222222222112221022122222222200022121222222222212222211210122222222222021222222021222222212212222022022221222222222222212102222212220220222222222021222122222222002200022021221222011222121222222222212220211221022222222222120222222122222222212222022020122220212222222222202002222212220220222122222022121222222222112201222020221222010222022222222222212220220201222222222222122222222122222222222202222222022222222222222222212012222202220221222222222121121022222222202221222020220222112022022222222222202221212211222222222222122222222222222222202212022112222222222222222222202212222212222220222022222121020022222222002212022022221222202022021222222222202222212222122222222222221222222221222222212212022011222201222222202222202012122202220222222122222021020222222222222210122221222222200022222222122222222220211222122222222222122222222022222222202212122112122222202222200222222102022222120222222222222122121022222222002211022122220222201222121222222222222221221222222222222222222222222020222222212202122201122200222222201222222102222222021222222022222121221222222222012221122222221222021222221222222222012221210221222222222222020222222020222222222222222201022212212222212222212012122222222220222222222222022222222222012200122220222222011022121222022222222221221221222222220222122222222020222222212212122200222210212222200222202012022202220220222222222222222022222222012220122222221222010022022222222222222220222222222222222222021221222222222222212222022102022200222222201222212212022222021222222122222022022022222222002212022222220222220022020222222222002220221221222222220222121220222120222222222222222020222220202222210222202102222212122222222022222022122002222222212211122021221222222022022222122222012222222210022222221222021201222221222222212212022121222222222222200222212212222222221221222022222120220022202222102222222122221222102022222222022222202220212200122222220222020221222021222222222222122110122201222222222222222102222212022220222122222000121112212222202200222020221222001222221222222222222222210200122222222222122210222020222222202222012102222222212222222222222122022202221220222222222112120112222222212200122222221222211122121222122222112222202221122222220222220200222020222222202222122101222202222222202222202212022212021221222222222210022112222222202201022121220222021112121222222222102220210221022222222222021222222222222222202202002011122220222222211222220220022212120222222022222101121122202222012210122022220222022122020222022212112221201211122222220222121222222221222222212212012022222221212222221222211202022220222220222222222111122012202222102222222021222222201222022222122212022222200221122222221222122212222120222222222212102210122200202222221222211121022211121221222222222121122012222222112202122020221222201002222222122222212221221202022222220222122201222221222222212222122002222201202222210222202210222221020222222022222110020222222222102210022220221222220222020222222222112221210210022222222222121210222122222222202212012122222221202222200222210012022221121221222122222201020022202222002200222220222222212202120222122222122220201202022222221222121201222122221222202212022122122210212222221222212002120220122220222022222012122122222222022202122021221222020012220122122222102222212222022222220222221202222120220222222222022110222222212222210222210210021111221221222122222101120102222222102200020220222222201222022222122202202220200210122022220222222201222221222222212202112210122201202222221222200021222210220222222022220111222212222222102210120020220221022012221022022202022220210221122222222222121222222120222222222222202121122212212222221222210202022112121221222222220021022022102222212201222020221221002102222122122222222222210202022122220222122212222121220222222202222002122210212222202222210200121012021222222122222101220122012222102202022220222220000022122222022212002222212210022222221222221211222022222222202202102122022202202222201222221022122111020221222122221220222112022222102211121121221222200222121022022202102220201221122222220222120220222221222222202222122121022210212222201222211021120110220221222122220002022122202222102200121020221221002022120222122222212220212211022222222222220202222222221222212222012202022220222022201202212100122120222221222122220121122102122222012220120121221222021002022222022212212220220202222022220222121212222220221222222212122212222220212022210212221220121220222222222222220102122202022222112221121121220220002002121222122202022220221211222022220222021211222222221222222212022122222221222122200202210200120210120220222022222200221002022222212211220020222221021112022022222222012220221212122222220222121221222122222222222202112211122222222122201212221001022102020220222122221111120112202222222200111022222222022112121022022212112220200200022022222222122221222021222222202222100210122221202222220212202102221001021222222122221202120211112222222202002022222222021022112022022202202222210220122222221222221221222121222222210222001020022201202122200222202010220210120221222222221101222010222222212221221120221222001112201122022212002221221200122222220222022221022220222202202212112121122220212022200212202121022110120220222122220201021012212222202210210222220222211102012222222212102222221200122022222222221200222020222202211202200002022220212222202212200100120122022222222122222011121002122222022220211022220220011212212222122212102221200222022012222222122201222121221212211202110022222221222122200202221022221000122221222122220211221101002222202202010121220221011112121022022222012220202202122012222222120222022121220222202202120200222220222122200222221100222001020221222022222120022210012222122200121122222222220022012122022202222221220211002102220222222210222121222212200212101122022200222122201202211200120201021220222022222102021112112222222201002221222220010212212222022212102220221210222122220222022200122022222212220222111210122200202122201202200011122122022221222222221020220122022222122220112022222221201222101222122212212221220220122022220222121220222120222212200212101000022200202222221212202200020110222220222122220221221010212222222211111022221221111102211222122222212221211200012012222222122221222220221212200212110020222201212220212222212210022012121220222122221020020020022222202212102222220220010112221222222212022220200202002112222222222202222220222222201212001121122220212121222212202001021010020222222022221202122202122222212210001120220220002012021122222202212220220200102222222122020210022220221222221212122002022210212022201212222012222021221221222022221221222200122222012201121222222220110002220122222202202220202201102222222122121201022122221202221202020121022221022021220222211021022020121222222222222120222000202222012220020020222221210202110022022222222222201212002222222122022200122220221202210212200021222210102122212202201201120111101221222022221000221102212222022212222221221222210222011222122212202221212212222122222022221211222022220212211202022101022202122020200202202212021122012220222122222202222111102222102212020122221222100122201222122222022222221210112012220022121211222121222212220222002221022212002120202222220202122201200222222022222010222211212222012221012120220221101202211022122212112220222220112112221022222202020222220202211200011011022210022122202222220212221220021222222122220110120001222222102201221121222222021102010222102222102222202201222002220222220220120022222222101212002201222211212121222202201112222022211220222222220001220220002222122222012022221221120102220222102222112222201211022212221122020200020221220202202222220220222211122221200202222110020101111220222022222002221202202222112200012120222221001222111022112212112222201212111012222122120211221222221212121222022220022220222020221222222210222000000221222122220012020001212222022210222021222220121212011122202212002220222211012012222222020221021022220222220212202221022220002121200212222002121212022222222022221101221120102122112212120021221221102202102122122222112220200211020122220222021211020222022222212200011011122212122020211222202020022121101222222022200001120120212022022210011022221221121022000222122222012222210220100022220022222202011121000222112201211110122201212122201212210020220121001222222222220012221122222122022210222221220222112222222022212202202221211220020212221122020201001220022212020211220022122201122020200202220011220021120222222022210200121012002022112221122020222220211122011222202212212220110211221012220022022211221122110222012221222220022210102221212222220000221221021222222122222202120210102112212220110221221220201202002122102202122221111200200202221222120202202022220212211200102200022212222020202222202002120121110221222022211112121010112102002200002021221222202012001122122202102222212222210022212022021202022122111202222212101022222202212222220122221012022211112220222022220000121022002122012212120221221222222102000022022202222221202221222212201122020210000220111222211212022211222210012121222002202111021112022220222022222111220000222212211221221021222221120012022122112222002221021202212002210022220210010120020222212202202112122201122020202021211102221012112222222222201001122110102112012221212121220220112202201022202222212222101220001202210222120201112021011202110202000012022212012220212011202112120011202222222122210011222120012102102212002020220222201102111122012212002212121220111222201022121221010020110222101220101001122201222020200102200000122200011222222022200211022002022012021202212120222222001112212222202212212210012200222222222122220200020121000202000210011201122202202120200211210020121120010222222022210021101111222212020200121122220021201122210012102202122220000212200002202222120212120121121202112200020000222211002221222101221110120201102222222022210022222011112102110220211120221121201122011222102222222200021222222112201222022220200222110222001202001222022022012222220211210101122021110221222022210200021002012212102210020122222020222002220022002222102202120222220122201122022201212021202212110222121122122110102020210211202012220010101222222122222021101122112112220200210122222120221102202122222222102221120221211202222122022000122220200202020210201221122220012001212121200120022010101222222122220021102220202002122211211220222021101122110022112222222202001212002112201222221112201120120212011210210012022202222100222120211000022220111221022222220212120202122002102202120221222120011202211202002222112222212210202012220222122010012222000202011212021102122211122210200120220211222202110220122122222120212020202102220201122021220222202102122022212212012200002220111122201022220021200020211212202200100200222202122202201122211002220120021221222022221012102220102212021201121220220122001002120212112212012222001222010012210022022210100222220202111210212010222112112110110200222120121000112221022122202022101220222112020200222100220221211122122001012212222211110212112102210122121002100121112222022200002211222111012010110012201001121021200222022022201112211111212022010212012122220021221122121120112222122211012210110102221122121201112222102202202201101021122101012222022020200102120011110221122122220201202021122122120210122201222221002202120220022222102201210200000112220122221002202021001222102220202221002100201110102212210100122202221221122222212101011212002202001211220121220020022102021211022202202211201202220222222222122001021222011222122211021102102002010101121020221212221211101020122122200222220012222112102202020010220021101102220001222212202200100200002102212122122010211122200202111200020120102212112012120001201200020012012220022222212200020111102022002210120101220220221212011011112222212202110222212212222122121012112222020222222210222021201020212101222022211102020100120220122222212000001101202122101201102122220202010012111210102212112221010222211102202022022021202222101212000202110220121010010202102200201120222222020120022122201012212121122022002210120211222200112002001020222202100201110102200212201222222212112220011202020211121200022012222212200200222010022101000022122122221120212222002012002211112011222001022102220101002212022201102110111122211022222020011120022212110200000000221101021101001101212122220211211020022122211002000122022022002221002112221222110122122202002202122220122212121022211222021100202020022202120201022222122000011112001022201102122210022121122122220110100210022222111200202101220112122022022220012202022001121201222102210022021122112121101202001222211020201201200002101000222220021100121120222122210212222200012002221212020211220211220002122122112222222222012022002202211222120000011022110212102200202001212021102201221111210110020010100222022122210222101121122122122212112001111221120201000002201120222022201211022010022200102200021212112011212002220201001111220212121200111110202222211000010110100100100000020021122111101202"
	case 821:
		return "0222112222120000"
	}
	panic("No input!")
}

var day1 = []int{
	121656,
	110933,
	80850,
	137398,
	76307,
	50450,
	124691,
	86449,
	145386,
	148648,
	68909,
	134697,
	109636,
	115718,
	134485,
	89267,
	64829,
	109070,
	84257,
	109010,
	97574,
	98363,
	123029,
	105568,
	114500,
	92041,
	128869,
	148350,
	144605,
	91862,
	134417,
	54710,
	147843,
	121914,
	127855,
	74545,
	89596,
	106562,
	69863,
	147082,
	135724,
	111637,
	68869,
	103685,
	99453,
	80908,
	136020,
	64974,
	125159,
	87504,
	62499,
	73294,
	128811,
	121567,
	54673,
	66647,
	66871,
	71228,
	101622,
	130675,
	69025,
	146118,
	79970,
	118267,
	122279,
	89523,
	62965,
	148036,
	119625,
	127056,
	54980,
	143581,
	103274,
	83064,
	125131,
	54362,
	115851,
	139103,
	140674,
	69616,
	81353,
	116441,
	73898,
	51403,
	137019,
	93146,
	67273,
	138182,
	126680,
	148683,
	127805,
	111741,
	102219,
	99603,
	90453,
	147581,
	102136,
	109913,
	144899,
	140572,
}

var day6 = []string{
	"HX5)C21",
	"VZD)VW7",
	"BTP)JHL",
	"699)GSD",
	"2FB)ZM1",
	"QS3)7XJ",
	"59Z)N3Z",
	"HS2)BKY",
	"7SN)YLN",
	"NJV)YC5",
	"YJP)D32",
	"5C8)KZ2",
	"3D2)FWP",
	"LN3)JHZ",
	"2CQ)4Q6",
	"NLJ)KWV",
	"7MZ)PDW",
	"Q2N)LP8",
	"2VJ)BRY",
	"V6D)CNL",
	"BBT)7QW",
	"Y9S)T8J",
	"F3Y)1X3",
	"G6V)1MG",
	"Q1C)ZL3",
	"8YV)F2H",
	"F8B)DRS",
	"D3H)R95",
	"GG5)WRZ",
	"KZ2)T52",
	"3Z7)X96",
	"9JC)B98",
	"VVK)96T",
	"1V9)52M",
	"F56)VKQ",
	"TV7)HQG",
	"TYL)NH7",
	"WQM)C3P",
	"ZQK)HCT",
	"LR4)P5H",
	"PKR)D9J",
	"MC9)4PC",
	"WKR)PHX",
	"8CS)B5P",
	"RBX)TV7",
	"CQT)J48",
	"XRB)CFS",
	"K32)57Y",
	"97B)Y3Q",
	"7XJ)41B",
	"59S)PDL",
	"N32)CW1",
	"Y8J)8NY",
	"R9B)S4Y",
	"13L)61J",
	"3MC)WK9",
	"NB9)XPY",
	"R8V)LHB",
	"1S5)2SL",
	"8K2)LST",
	"2NW)6H1",
	"B27)Q2B",
	"14L)W35",
	"41H)Z5M",
	"Y8G)8VG",
	"X5S)CFP",
	"FRK)6PM",
	"HTR)QRC",
	"BJ2)5Y3",
	"CS3)R9W",
	"ZL3)GTG",
	"FKX)NF5",
	"HC6)6FG",
	"QG7)85C",
	"Y5Y)1VG",
	"FLJ)DJH",
	"28P)K4Z",
	"S2N)J1C",
	"GY6)GW9",
	"CMW)QC7",
	"L7F)VJ2",
	"YG1)FKX",
	"GQJ)GY7",
	"WC8)KTZ",
	"6JS)M79",
	"L7J)TQD",
	"YVQ)ZYV",
	"R6X)1HB",
	"GTH)TM5",
	"JJJ)GJS",
	"CKH)NQK",
	"2NG)ZVK",
	"564)3RC",
	"VJ2)L7Z",
	"VW7)7SD",
	"6F7)SAN",
	"YVH)6JN",
	"T8C)CKH",
	"K8Y)F6G",
	"953)HQ9",
	"FBW)HWJ",
	"SGS)65F",
	"GXR)6SL",
	"M6K)K3X",
	"TJH)QKG",
	"DN7)YS4",
	"2B3)B2P",
	"ZQ1)WFM",
	"YF1)KS3",
	"NNS)NDD",
	"VTS)PF1",
	"9T4)H7Q",
	"GJS)8VP",
	"D5Q)NR3",
	"LHB)3RK",
	"H32)5DQ",
	"3GZ)1WV",
	"H48)8CN",
	"H5S)WZH",
	"SBH)LFH",
	"BCX)6M8",
	"T46)9P6",
	"Q1X)4L9",
	"72F)ZC2",
	"PCB)SC1",
	"YQ6)S3M",
	"Q2B)PMF",
	"V4S)YVQ",
	"NFK)5C8",
	"43S)9CQ",
	"CNC)YPF",
	"HVS)41H",
	"FBW)4GX",
	"NHF)45D",
	"QRB)CZQ",
	"L79)7MY",
	"W47)4XS",
	"GMS)8BJ",
	"6K1)5NX",
	"K53)JY1",
	"583)GZJ",
	"XHD)MRD",
	"NWM)DFD",
	"G46)JX6",
	"8BJ)1FW",
	"1MJ)692",
	"HTQ)F9B",
	"6PZ)92W",
	"1Q1)8MN",
	"9HS)772",
	"BN2)29B",
	"9Q7)VJ3",
	"C5D)C4R",
	"JS6)9F9",
	"PF1)J3T",
	"ZP4)BCX",
	"3X3)VKL",
	"V6L)SYJ",
	"3VP)Z3Q",
	"P1L)HXV",
	"1DS)8YV",
	"9P6)C16",
	"RMW)Y34",
	"W1D)72Q",
	"FC7)B3Z",
	"BD8)GMH",
	"38M)PCB",
	"6DT)LN3",
	"KBY)9W8",
	"3JS)R85",
	"C46)RBX",
	"NHF)9K5",
	"DHR)STK",
	"SSP)Z6F",
	"WYW)H7R",
	"NDX)NX9",
	"M19)TH9",
	"P3R)RTW",
	"BTS)R8V",
	"5GC)DSN",
	"D9J)3LM",
	"KCB)LWT",
	"TTQ)FH4",
	"72Q)FHY",
	"Z6F)F3Y",
	"XBB)VFV",
	"QW5)VVK",
	"YNC)26J",
	"5LK)2LR",
	"H2V)GCW",
	"BC5)GY6",
	"1P3)D8V",
	"TL2)465",
	"65F)S9V",
	"B5P)K32",
	"27Z)YVH",
	"VQ2)G6N",
	"Z2W)LHX",
	"D8K)JN4",
	"1FK)RK1",
	"HQP)2N8",
	"COM)Q1H",
	"CHK)VZD",
	"Q6V)Q1X",
	"1NX)C5D",
	"CC6)J37",
	"L3G)KKP",
	"KST)QG7",
	"6FG)T8C",
	"YC5)56M",
	"43X)HMM",
	"D45)2RD",
	"TH9)7FY",
	"SM5)S81",
	"9MC)MVR",
	"J48)JH9",
	"VQT)1D2",
	"2TV)SP7",
	"Q9N)658",
	"J1V)26X",
	"8L5)XWZ",
	"DC2)R89",
	"B3Z)7WL",
	"DJH)CHK",
	"LST)X22",
	"F9B)L4V",
	"RGX)K85",
	"68T)VG7",
	"YM3)SJN",
	"JSF)97X",
	"B98)RZZ",
	"959)PKP",
	"6M8)2VZ",
	"MT3)CMM",
	"K2Q)QYF",
	"RTK)HPL",
	"RWY)2PT",
	"7QW)LGX",
	"LXZ)X8T",
	"T43)KV9",
	"XSM)2ND",
	"X2B)SDQ",
	"HZC)KZG",
	"2RS)C6L",
	"N3Z)BR9",
	"2B3)P3D",
	"ZST)GL1",
	"QQ8)TYL",
	"7TD)BT1",
	"2VZ)8YW",
	"KS3)V4T",
	"Q1H)MJY",
	"S9V)ZMG",
	"MRK)9V8",
	"JVS)LL3",
	"8VG)BJ2",
	"SCM)72S",
	"GP9)RZB",
	"KS5)DN7",
	"Y34)Y3W",
	"LX1)8VF",
	"T6B)D8K",
	"RPB)564",
	"41B)364",
	"Q7Q)1FT",
	"TBG)479",
	"XBX)XP9",
	"JLZ)F2F",
	"XHK)SZQ",
	"P1P)ZDH",
	"583)Z7L",
	"CMM)DK2",
	"KP1)CFW",
	"69X)2VJ",
	"SYJ)BQQ",
	"147)2XS",
	"GQJ)FLJ",
	"ZCK)Q2N",
	"NZS)1BD",
	"3RC)3Z7",
	"TCC)BLQ",
	"FH4)SCM",
	"VM6)FQD",
	"G8V)ZXS",
	"8WX)2CQ",
	"HV4)J3F",
	"72L)MF6",
	"W66)953",
	"C3R)1KK",
	"4JC)YQQ",
	"Q7S)ZCT",
	"BJY)R9B",
	"2SL)XC1",
	"XRM)J1V",
	"SK3)NP7",
	"2TV)G1V",
	"HX9)J7C",
	"J37)6LZ",
	"GMH)NPR",
	"5T6)YRR",
	"7RB)89P",
	"T9L)Q7Q",
	"4GX)2DP",
	"6LZ)79K",
	"NHQ)YQ6",
	"Y3W)C3R",
	"6SL)K9L",
	"C46)P1P",
	"8NY)8RM",
	"966)HXQ",
	"B6S)C46",
	"T39)DDV",
	"XV7)2Q3",
	"YJB)ZX2",
	"QC7)3NF",
	"Y42)N28",
	"B1J)1SG",
	"PHZ)38M",
	"48G)GVK",
	"ZLY)X5H",
	"YKK)NB9",
	"WK9)4GG",
	"R5G)BTP",
	"J4C)QS3",
	"45D)LGP",
	"658)BN2",
	"B91)Z6D",
	"XPY)YF1",
	"FQT)MN9",
	"4ZQ)PJG",
	"W86)K8Y",
	"8VF)B91",
	"P3D)J31",
	"MTT)HH7",
	"GVK)RV9",
	"GY7)6R6",
	"72S)PB4",
	"XP9)FBW",
	"5NR)72L",
	"YM9)BBT",
	"C4G)Z5C",
	"987)G95",
	"CDG)YYB",
	"SZQ)J39",
	"BZY)YZC",
	"SKT)6PZ",
	"ZVK)5VD",
	"BHG)TBG",
	"X1X)9JC",
	"MST)959",
	"RYV)MQQ",
	"GL1)V26",
	"FQD)JLZ",
	"N5S)GWD",
	"B2H)L3G",
	"QYF)3JP",
	"WWB)B27",
	"PHX)VM6",
	"MSZ)CMW",
	"HH7)MKP",
	"9XZ)VNR",
	"3NF)H48",
	"YQM)CVF",
	"B98)H2V",
	"K91)2NZ",
	"8MN)PF7",
	"RZB)C1F",
	"YRR)3PR",
	"2VR)9P2",
	"4V5)VVC",
	"2DP)M6K",
	"KWV)8FW",
	"YPF)28L",
	"5LG)SSP",
	"3JP)TSR",
	"JW9)7WZ",
	"M71)P5J",
	"Y34)MRS",
	"GSD)XSM",
	"Q59)59S",
	"71Y)958",
	"QZC)FDZ",
	"QRC)RYV",
	"RDF)CS3",
	"3T3)LYZ",
	"5NX)PL2",
	"5L1)JJ6",
	"F1X)9T4",
	"RV9)MJJ",
	"ZHD)7JQ",
	"CZ9)3D2",
	"92W)C34",
	"VHH)QQ8",
	"364)RRK",
	"J3G)TZ5",
	"8VP)TZG",
	"3PR)VQ2",
	"ZM1)ZCK",
	"D7F)43S",
	"TSR)DKK",
	"F7C)BV2",
	"DFD)92T",
	"84M)HMX",
	"M34)RN9",
	"49F)KDG",
	"P9T)966",
	"6RG)TL2",
	"QCF)K6Q",
	"KKP)B7H",
	"JQF)XHK",
	"2XS)W1L",
	"XJR)XRB",
	"4GC)KB5",
	"BX8)G6V",
	"VKL)ZLY",
	"RBX)KZJ",
	"3RK)489",
	"V26)QRB",
	"2Q3)7TW",
	"GNF)WPF",
	"YJT)7K6",
	"PL2)71Y",
	"2ND)2NG",
	"HWJ)V9X",
	"VF5)VQT",
	"2WD)WD9",
	"MF6)GKK",
	"787)J4C",
	"G8Q)358",
	"HVH)FNN",
	"XQG)Z11",
	"JG9)PBB",
	"Y5Q)GXR",
	"R19)N99",
	"P5R)DLL",
	"246)SFG",
	"1SP)Y5Z",
	"16N)NFK",
	"2NZ)HW8",
	"B7M)LVQ",
	"GZ8)F7C",
	"L39)WK8",
	"2LR)YWY",
	"BNT)3Q8",
	"T46)98H",
	"VQT)5T6",
	"JX6)T43",
	"9FZ)P9T",
	"QRL)Q1C",
	"J3T)7H1",
	"PJG)2RS",
	"YMG)1NX",
	"YYB)XP3",
	"5LH)TY6",
	"NX9)VHH",
	"XBR)PF5",
	"89P)TQY",
	"254)L7F",
	"J6L)4LK",
	"FLW)QCF",
	"8FW)84M",
	"P7C)1V9",
	"LFL)3ZN",
	"7TW)JNM",
	"RZZ)GCT",
	"D33)4L2",
	"KK9)V63",
	"K3X)GG5",
	"B3F)FVX",
	"XW6)GMC",
	"4L9)ZRN",
	"6PM)BNT",
	"V54)XJR",
	"8ZK)MV7",
	"XW4)NJV",
	"QYM)SPV",
	"BPB)G46",
	"BQQ)DDN",
	"H7Q)4WC",
	"B58)Y9S",
	"9P2)SQX",
	"729)C4G",
	"ZFX)QLH",
	"Z5C)3B7",
	"M79)73S",
	"RN9)W9V",
	"F2X)H5S",
	"2HW)9FZ",
	"669)7PQ",
	"SPN)SKT",
	"QCV)MQ8",
	"Z5T)ZQK",
	"4HJ)3B6",
	"WZH)246",
	"1KK)2R9",
	"WD9)JFP",
	"YWY)8NZ",
	"CFW)R5K",
	"Y42)3KH",
	"LF8)8JX",
	"4WC)5T5",
	"X8K)W66",
	"T3V)1DS",
	"ZMG)NYX",
	"VVD)KBY",
	"N4B)49J",
	"TQ1)7RL",
	"WYW)3XW",
	"P4G)V6D",
	"2N8)GTH",
	"C16)YC6",
	"ZM1)TQX",
	"PBY)Y1Y",
	"X8T)BD8",
	"NQK)NLJ",
	"3KH)PBY",
	"4LF)MT3",
	"K85)69W",
	"SYD)72J",
	"CVK)HTR",
	"55S)M4C",
	"R95)XV7",
	"YLC)WC8",
	"FRK)D7F",
	"DRS)ZKJ",
	"7H1)F7Y",
	"F7Y)HQS",
	"HD2)CWS",
	"GZ8)P7C",
	"Q8B)HYP",
	"QCV)KXK",
	"YLN)MS4",
	"RYV)HTQ",
	"V4T)CBX",
	"VKQ)T6Y",
	"26X)YNC",
	"2KX)VQQ",
	"RX6)KRR",
	"P5J)CD1",
	"K46)2NW",
	"CWS)61S",
	"6QX)KFZ",
	"S4Y)NKG",
	"GW9)Y4L",
	"1W7)94J",
	"XQG)L39",
	"S3Z)365",
	"Y3Q)GZ8",
	"XWZ)X8K",
	"RRR)28P",
	"Y4L)7C2",
	"26J)SM5",
	"96T)Z9K",
	"QKG)N4B",
	"97X)XRM",
	"592)B58",
	"YS6)L7J",
	"ZS4)DZM",
	"Z7L)F3Q",
	"1D2)5TM",
	"LKM)P98",
	"K6Q)ZHD",
	"SSQ)HC6",
	"ZSH)MRK",
	"MQ8)BJK",
	"JHL)CVB",
	"HYP)YLC",
	"1SG)2VR",
	"ZC2)ZQ1",
	"JPZ)FRK",
	"P55)GX2",
	"Z6D)3JS",
	"8RM)C3J",
	"35P)JPZ",
	"PKP)JG9",
	"9BV)5NR",
	"7PL)CQT",
	"7MY)QW5",
	"TQX)XVY",
	"LR4)HQP",
	"29B)PN8",
	"7YM)P5R",
	"QYF)7Y7",
	"W1L)6RG",
	"7SC)NZS",
	"2RD)YT6",
	"BPZ)Y42",
	"3HR)5FP",
	"27Z)S3Z",
	"4LK)B2H",
	"C5T)RX6",
	"LF8)5CF",
	"RRK)LK1",
	"YQQ)R5G",
	"BLQ)J5Q",
	"KFZ)6B3",
	"Q3P)KVC",
	"HXV)BTS",
	"NYX)HK1",
	"J5Q)T6B",
	"YC6)H3Z",
	"NP7)3T3",
	"KR4)T5R",
	"SJN)JGB",
	"PN8)X2B",
	"Y7T)5LK",
	"G5C)8G2",
	"JN4)611",
	"NR3)CVK",
	"N99)WC3",
	"SPV)QRL",
	"ZYV)K91",
	"ZXS)ZSH",
	"V6D)LR4",
	"Z5M)YM9",
	"VPT)JFZ",
	"FDZ)5MR",
	"2VR)WFF",
	"BRK)NWM",
	"8YW)D33",
	"1VG)RTL",
	"7ZZ)5LS",
	"4GC)7RB",
	"FDC)96K",
	"GJP)VDS",
	"DKK)K53",
	"Q1Z)2VC",
	"PDL)RGX",
	"9V8)XBB",
	"7Y7)MPF",
	"C21)XHD",
	"LL3)CRZ",
	"7FV)YPC",
	"T3D)GM1",
	"G7H)LXZ",
	"TSW)1FK",
	"R5K)JKZ",
	"8P7)TQ1",
	"KV9)V54",
	"TM4)G8V",
	"K3X)SFH",
	"1TX)C8B",
	"L6R)T9L",
	"CFR)NCP",
	"YKK)LX1",
	"7K6)P55",
	"R85)9ZJ",
	"NGV)NHQ",
	"F2F)HX9",
	"HMX)3SB",
	"HQG)JL8",
	"V18)PGR",
	"YZC)HV4",
	"VCP)2F8",
	"4LY)FZ5",
	"F7Z)CZ9",
	"M63)XW3",
	"9W8)7YM",
	"QSX)D3H",
	"LVQ)B1J",
	"GX2)DC2",
	"6QC)T71",
	"LGX)D9T",
	"S81)9K1",
	"C1F)M34",
	"JBW)W9M",
	"PZK)973",
	"1LW)WZX",
	"VLH)QCV",
	"358)KMC",
	"XP3)SYD",
	"S7X)1LW",
	"61J)1P3",
	"H7R)TJH",
	"4C3)8K2",
	"1BD)GJP",
	"3B7)GQJ",
	"CRZ)X4L",
	"7PQ)N57",
	"DDN)LFL",
	"28L)9RZ",
	"LWT)J6L",
	"692)LF8",
	"BT1)4YS",
	"C3J)TCC",
	"Z3Q)G7H",
	"MRD)YJP",
	"D32)SK3",
	"HPL)2HW",
	"P5H)R19",
	"3JP)V4S",
	"GM1)H9L",
	"Q7Q)699",
	"ZRF)KGW",
	"K9L)2WD",
	"Q5Y)Q8B",
	"GY3)FQT",
	"4RN)D5Q",
	"SC1)ZP4",
	"W35)2FB",
	"5D9)8L5",
	"59C)592",
	"ZDH)4Q7",
	"RTL)2KX",
	"GCT)GMS",
	"HMM)TM4",
	"RGP)C55",
	"YZW)1W7",
	"6H1)H32",
	"SK3)XJ1",
	"J4T)59Z",
	"Y1Y)D4L",
	"F6G)YSC",
	"YVP)3X3",
	"QTF)13L",
	"CVF)G8Q",
	"96K)HZC",
	"G95)NHF",
	"BV2)9BY",
	"5LQ)XBX",
	"SFH)V18",
	"3RC)MSZ",
	"WFF)YOU",
	"CNL)HK8",
	"ZX2)787",
	"NKX)YKK",
	"C3P)WB3",
	"HQ9)WQM",
	"RK1)T3V",
	"9BY)MST",
	"489)YZW",
	"LP8)JQF",
	"71Q)CTP",
	"S3M)P3R",
	"N57)FDC",
	"TQY)9J6",
	"3SB)729",
	"3ZY)B3F",
	"HGL)97B",
	"BQQ)4GC",
	"KLD)F7Z",
	"BQB)N1Q",
	"WD9)XW6",
	"HQG)P1L",
	"VC1)KS5",
	"P67)BC5",
	"3B6)RPB",
	"X4L)XB6",
	"NKG)5GC",
	"GCW)HSC",
	"4GG)MSS",
	"2F8)HD2",
	"611)5DS",
	"W9M)PHZ",
	"6SL)72F",
	"8CN)C4F",
	"Y6C)K2Q",
	"3LW)9GD",
	"MFC)J9X",
	"B2P)X5S",
	"N28)Q3P",
	"CW1)VF5",
	"CFP)R6X",
	"K91)GQV",
	"3Q8)TNM",
	"KZJ)1MJ",
	"SDQ)PLC",
	"8JC)5YG",
	"3ZN)3HR",
	"4J6)N5S",
	"HK8)BZY",
	"FQM)TTQ",
	"TM5)526",
	"SQX)FQM",
	"772)BRK",
	"V63)FC7",
	"DZM)2B3",
	"BRY)YS6",
	"LCZ)YG1",
	"C4F)W47",
	"WFM)6QX",
	"N28)TVJ",
	"L7Z)468",
	"MVR)254",
	"ZCT)9BV",
	"SX1)583",
	"52M)TSW",
	"2PT)1Q1",
	"7FY)3BF",
	"YVP)LCZ",
	"2CQ)7TD",
	"D8V)5L1",
	"Q3F)ZST",
	"L7Z)YDP",
	"4LF)SWH",
	"MQQ)PKR",
	"WC8)FLW",
	"JL8)DF1",
	"5Y3)2TV",
	"5DQ)8X6",
	"M5Q)6D7",
	"KGW)QTF",
	"T2T)RDF",
	"4Q7)K46",
	"L4K)JSF",
	"FVX)HVH",
	"49J)XWX",
	"8NZ)3MC",
	"F6N)KP1",
	"WZX)ZFX",
	"H9L)F8B",
	"1MG)W86",
	"KXK)SS3",
	"X2B)RWY",
	"N1Q)669",
	"43R)Q59",
	"JH9)458",
	"C87)T2T",
	"XW3)N1Y",
	"GZJ)MC9",
	"2DP)3ZY",
	"5DS)JS6",
	"MRS)L79",
	"YP2)71Q",
	"MJJ)VC1",
	"7SD)MFC",
	"13L)997",
	"4YS)4LY",
	"479)CFR",
	"V9X)7MZ",
	"LP1)NNS",
	"B7H)ZS4",
	"5CF)VCP",
	"FZ5)3JN",
	"T6Y)R3T",
	"J53)VPT",
	"VVC)7ZZ",
	"776)Q1Z",
	"94J)LP1",
	"4PC)4ZQ",
	"T5R)GP9",
	"F2F)4V5",
	"PND)KST",
	"Z11)N7V",
	"TX7)M5Q",
	"VDS)X1X",
	"6D7)YP2",
	"YS4)5LQ",
	"FHY)7PL",
	"HD5)VTS",
	"BKY)WWN",
	"4HJ)4C3",
	"ZX2)V6L",
	"BR9)3LW",
	"DX4)BPB",
	"TY6)GY3",
	"XBX)Y5Q",
	"4XS)94T",
	"32F)8P7",
	"2S1)B7M",
	"D9J)YT7",
	"54D)Z5T",
	"KMC)7SC",
	"NJK)1KF",
	"CBX)SPN",
	"KDG)M7P",
	"PB4)147",
	"L4V)YPD",
	"JHZ)YJB",
	"J31)KK9",
	"JKZ)Q6V",
	"LWX)ZRF",
	"3LM)LKM",
	"4TW)4HC",
	"N1Y)4J6",
	"LYZ)TX7",
	"JFZ)4LF",
	"1FT)8CS",
	"8JX)LWX",
	"W9V)5LH",
	"73S)9Q7",
	"HFK)YVP",
	"RGV)HFK",
	"5MR)3MP",
	"8JC)987",
	"BJY)SGS",
	"NYS)DXM",
	"X1X)6K1",
	"JJ6)KLD",
	"ZRN)VVD",
	"6RG)GNF",
	"4Q6)NJK",
	"H3Z)7SN",
	"R3T)RRR",
	"57Y)KNF",
	"DF1)JYX",
	"XWX)6QC",
	"FWP)4HJ",
	"3B6)1SP",
	"9F9)9HS",
	"Y21)WVQ",
	"N57)1S5",
	"TVJ)Y8G",
	"TZG)T46",
	"1X3)BN6",
	"DXM)CC6",
	"T3D)SFB",
	"XQC)NDX",
	"X22)JBW",
	"365)TJL",
	"J7C)XQG",
	"VFV)S2N",
	"4PZ)PZK",
	"4HC)YQM",
	"NPR)Y8J",
	"VG7)WYW",
	"C6L)M63",
	"P98)JMK",
	"Z9K)1RM",
	"SFG)4RN",
	"VJ3)ZJH",
	"WZX)WKR",
	"PGR)NGV",
	"SYB)B6S",
	"J39)VQM",
	"N7V)RMW",
	"22F)59C",
	"H3Z)NKX",
	"MJY)HX5",
	"TJL)5LX",
	"GWD)S7X",
	"123)Y5Y",
	"XJ1)4PZ",
	"DK2)J3G",
	"V11)L4K",
	"958)F2X",
	"6R6)V11",
	"VW7)Y6C",
	"CD1)8ZK",
	"96K)JG2",
	"YT6)C5T",
	"KZ2)9X2",
	"JFP)PND",
	"WB3)9MC",
	"98H)BJY",
	"C8B)DX4",
	"465)HS2",
	"WC3)QZC",
	"9J6)XBR",
	"LHX)68T",
	"B2B)L6R",
	"468)XW4",
	"MN9)RTK",
	"TSF)6F7",
	"XB6)F56",
	"C55)HGL",
	"526)MTT",
	"9CQ)KR4",
	"699)9XZ",
	"ZF1)5D9",
	"NDD)BQB",
	"Y5Z)43R",
	"M7P)Y7T",
	"GQV)K1J",
	"HQS)4TW",
	"1WV)XL6",
	"LGP)QSX",
	"6JN)P67",
	"X5H)NYS",
	"DDV)74T",
	"PNT)ZF1",
	"V54)Q9N",
	"KB5)SX1",
	"1FW)3GZ",
	"CHK)G5C",
	"T8J)ZMN",
	"CZQ)VML",
	"4L2)CDG",
	"TNM)JVS",
	"JG2)D45",
	"D4L)Y21",
	"VJ3)RGP",
	"K1J)T39",
	"SWH)BX8",
	"ZMN)35P",
	"458)F6N",
	"43S)JW9",
	"KVC)Q5Y",
	"F3Q)69X",
	"S5H)8JC",
	"JY1)J4T",
	"VML)BHG",
	"NCP)43X",
	"2VC)4JC",
	"MSZ)YM3",
	"9K1)P4G",
	"SP7)XM6",
	"J1C)3VP",
	"973)PNT",
	"79K)YJT",
	"5YG)SBH",
	"56M)T3D",
	"SCM)Y7X",
	"R89)2S1",
	"YDP)C87",
	"VQM)32F",
	"J3F)W1D",
	"PDW)16N",
	"QLH)49F",
	"G1V)WWB",
	"ZJH)YMG",
	"VTS)B2B",
	"KTZ)48G",
	"MSS)FJV",
	"HXQ)27Z",
	"C34)CNC",
	"NH7)JJJ",
	"SKT)14L",
	"VLH)8WX",
	"7WL)SSQ",
	"D9T)BPZ",
	"YT7)54D",
	"72J)M19",
	"92T)HVS",
	"DLL)RGV",
	"S9V)J53",
	"BQB)QRH",
	"NF5)DHR",
	"85C)RYH",
	"XW4)XQC",
	"JGB)F1X",
	"2R9)PB2",
	"KZG)M71",
	"SS3)QYM",
	"T52)S5H",
	"7C2)HD5",
	"4TW)Q7S",
	"HW8)123",
	"JX6)Q3F",
	"HK1)22F",
	"9K5)7FV",
	"9X2)TSF",
	"3BF)6DT",
	"JLZ)1TX",
	"Q1Z)VLH",
	"5LX)S9H",
	"7WZ)N32",
	"X96)55S",
	"PF7)9Q6",
	"59Z)6Y3",
	"K4Z)SYB",
	"P7C)KCB",
	"997)776",
	"SFB)Z2W",
	"HSC)5LG",
	"9GD)PY5",
	"YPD)6JS",
}
