package module

import (
	"fmt"

)
func Main1() {

	var clues = [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5}}

			fmt.Println(clues[1][2])

}

func Snail(snaipMap [][]int) (res []int) {

	for i := 0; i < 9; i++ {
		switch {
		case i < 3:
			res = append(res, snaipMap[0][i])
		case i == 3:
			res = append(res, snaipMap[1][2])
		case i == 4:
			res = append(res, snaipMap[2][2])
		case i == 5:
			res = append(res, snaipMap[2][1])
		case i == 6:
			res = append(res, snaipMap[2][0])
		case i == 7:
			res = append(res, snaipMap[1][0])
		case i == 8:
			res = append(res, snaipMap[1][1])
		}
	}

	return res
}

// func main() {

// var clues = []int{
//       0, 0, 1, 2,
//       0, 2, 0, 0,
//       0, 3, 0, 0,
//       0, 1, 0, 0, }

// 	fmt.Println(SolvePuzzle(clues))

// }
// func GetIndexes(cl int) (res [2]int) {

//         switch {
//         case cl <4:
//             return [...]int {1, cl}
//         case cl <8:
//             return [...]int {cl-4, 3}
//         case cl <12:
//             return [...]int {3, cl-8}
// 		default:
//             return [...]int {cl-12, 1}
//         }
// }

// func SolvePuzzle(clues []int) (res [][]int) {

// 		for i := 1; i <= 4; i++ { //row
// 		for j := 1; i <= 4; j++ { //column
// 		for k := 1; k <= 4; k++ {
// 		for t, w := range clues {

// 		}
// 		}
// 		}
// 	}

// }

// func stringSum(s string) int { // calculates sum of str's digits
//   sum := 0
//   for _, v := range s {
//     sum += int(v) - '0'
//   }
//   return sum
// }

// func OrderWeight(s string) string {
//   // convert to arr
//   arr := strings.Fields(s)

//   // sort
//   sort.SliceStable(arr, func(i, j int) bool {
//     if diff := stringSum(arr[i]) - stringSum(arr[j]); diff == 0  { // if same "weight"
//       return arr[i] < arr[j] // just compare the strings directly
//     } else {
//       return diff < 0 // otherwise compare using weight diff
//     }
//   })

//   // convert back to string
//   return strings.Join(arr, " ")
// }

// func wordWeight(str string) int {
//   s := 0
//   for _, r := range str {
//     v := int(r - '0')
//     s += v
//   }
//   return s
// }

// func OrderWeight(strng string) string {
//   weights := strings.Split(strng, " ")
//   sort.Slice(weights, func(i, j int) bool {
//     s1 := wordWeight(weights[i])
//     s2 := wordWeight(weights[j])
//     if s1 == s2 {
//       return weights[i] < weights[j]
//     }

//     return s1 < s2
//   })

//   return strings.Join(weights, " ")
// }

// type Weight1 struct {
// 	Summ int
// 	Str  string
// 	Used bool
// }

// func OrderWeight(strng string) (res string) {
// 	var strSABOder []string
// 	for _, sn := range strings.Split(strng, " ") {
// 		if len(sn) > 0 {
// 			strSABOder = append(strSABOder, sn)
// 		}
// 	}

// 	sort.Strings(strSABOder) //alphabetical ordering

// 	var nSl []Weight1
// 	for _, sn := range strSABOder {
// 		summ := 0
// 		for _, od := range sn {
// 			summ += int(od - 48) //get summ 12 =>1+2 =3
// 		}
// 		// if summ > 0 {
// 			nSl = append(nSl, Weight1{summ, sn, false})
// 		// }
// 	}

// 	var summS []int
// 	for _, ws := range nSl {
// 		summS = append(summS, ws.Summ)
// 	}

// 	sort.Ints(summS) //sort by summ
// 	fmt.Println(summS)

// 	for _, n := range summS {
// 		for i, sn := range nSl {
// 			if sn.Summ == n && !sn.Used {
// 				res += sn.Str + " "
// 				nSl[i] = Weight1{sn.Summ, sn.Str, true}
// 				//sn =  WeightS{sn.Summ, sn.Str, true} не пишет в массив
// 				continue
// 			}
// 		}
// 	}
// 	return strings.TrimRight(res, " ")
// }

// func OrderWeight(strng string) (res string) {
// 	nSl := make(map[int]string)

// 	for _, sn := range strings.Split(strng, " ") {
// 		summ := 0
// 		for _, od := range sn {
// 			summ += int(od - 48)
// 		}
// 		if summ > 0 {
// 			nSl[summ] = sn
// 		}
// 	}

// 	var keys []int
// 	for k := range nSl {
// 		keys = append(keys, k)
// 	}
// 	sort.Ints(keys)

// 	for _, k := range keys {
// 		res += nSl[k] + " "
// 	}
// 	return res
// }

// func RemovNb(m uint64) (res [][2]uint64) {
// 	//var res [][2]uint64
// 	summ := m * (m + 1) / 2

// 	for i := uint64(1); i <= m; i++ {
// 		for j := uint64(1); j < i; j++ {
// 			if i*j == summ-i-j {
// 				//row := make([]int, 2)

// 				//res = append(res, [2]uint64{0: i, 1: j})
// 				res = append(res, [2]uint64{j, i})
// 				res = append(res, [2]uint64{i, j})
// 			}
// 		}
// 	}
// 	return res
// }

// 	dataJson := "[1, 2, 3, 4, 5, 6, 7, 8, 9, 0]"

// 	fmt.Println(createPhoneNumber(dataJson))

// func createPhoneNumber(numbers [10]uint) string {
// 	// var arr []int
// 	// _ = json.Unmarshal([]byte(dataJson), &arr)
// 	// //return strings.Join(arr, "")
// 	// //return fmt.Sprint(arr)
// 	var str string
// 	for i := range numbers {
// 		str += strconv.Itoa(int(numbers[i]))
// 	}

// 	return fmt.Sprintf("(%s) %s-%s", str[:3], str[3:6], str[6:])
// }

// 	s := ")(())((()())())"
// 	// // fmt.Println(ValidParentheses(s))
// 	fmt.Println(ValidParentheses(s))
// func ValidParentheses(parens string) bool {
// 	for {
// 		strRepl := strings.Replace(parens, "()", "", -1)
// 		if parens == strRepl {
// 			return len(parens) == 0
// 		}
// 		parens = strRepl
// 	}

// }

// func ValidParentheses1(parens string) bool {

// 	for strRepl := strings.Replace(parens, "()", "", -1); !(parens == strRepl); parens = strRepl {

// 		fmt.Println(parens, " ", strRepl)
// 	}

// 	return len(parens) == 0

// }

// s := "RqaEzty"
// sl := strings.Split(s,"")
// for i, letter := range sl  {
//    sl[i] = strings.Title(strings.Repeat(strings.ToLower(letter), i + 1))
// }

// fmt.Println(strings.Join(sl, "-"))

// 	a := "xyaabbbccccdefww"
// b := "xxxxyyyyabklmopq"
// //  var sb strings.Builder
// // 	sb.WriteString(a)
// // 	sb.WriteString(b)
//   str :=  a+b
//     chars :=  make(map[string]struct{})
//     for i := 0; i < len(str); i++ {
// 		if _, ok := chars["foo"]; !ok {
// 		chars[string(str[i])] = struct{}{}
// }
//     }
// 	keys := make([]string, 0, len(chars))
//     for k := range chars {
//         keys = append(keys, k)
//     }
//     sort.Strings(keys)

// 	fmt.Println(strings.Join(keys, ""))

// import (
// 	"fmt"
// 	//"sort"
// 	"math"
//     "math/big"
// 	// "strings"
// 	// "strconv"
// )

// func FindNb(m int) int {

// 	sumBildNum := big.NewInt(0)

// 	for i := 1; i < math.MaxInt64; i++ {

// 		sumBildNum.Add( sumBildNum, new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(3), nil))

// 		// 		if i==2022 {
// 		// 	fmt.Println(sumBildNum, "more than")
// 		// }

// 		if sumBildNum.Cmp(big.NewInt(int64(m))) == 0 {
// 			return i
// 		}
// 		if sumBildNum.Cmp(big.NewInt(int64(m))) > 0 || sumBildNum.Cmp(big.NewInt(math.MaxInt64)) > 0 {
// 			// fmt.Println(sumBildNum, "\n", int64(m))
// 			return -1
// 		}
// 	}
// 	return -1 //for compiler, but never used
// }

// strIP := "12.255.56.1"

// fmt.Println(is_valid_ip(strIP))
// func is_valid_ip(ip string) bool {

// 	//its may be in one string, but not understandeble
// 	sRegEx1 :=`([1-9]\d{1,2}|\d)\.` // 0. but not 05.
// 	sRepeatRegex :=strings.Repeat(sRegEx1, 4)
// 	sTrimLast := strings.TrimRight(sRepeatRegex, `\.`)
// 	sRE := fmt.Sprintf(`^%s$`, sTrimLast) // only get "1.2.3.4", not get "foo (1.2.3.4).5."

// 	//matched, _ := regexp.MatchString(sRE, ip)
// 	res := regexp.MustCompile(sRE).FindStringSubmatch(ip)

// 	if len(res) == 0 {
// 		return false
// 	}

// 	for i := 1; i< 5; i++ {
// 		numIP, _ := strconv.ParseInt(res[i], 10, 32) //десятичная
// 		if numIP > 255 {
// 			return false
// 		}

// 		//     for i, name := range r.SubexpNames() {
//         // if i != 0 {
//         //     subMatchMap[name] = match[i]
//         // }
//     }

// 	return true
// }

// //array
// var myArray [10] int
// myArray := [3]string{"arp", "live", "strong"}
// myArray := [...]string{"arp", "live", "strong"}

// //slice рекомендуют использовать в основном
// var mySlice []float64 EOF
// var mySlice  = []string{} []
// mySlice  := []string{}
// mySlice := make([]string, 5), make([]string, 5),
// x := make([]float64, 5, 10) здесь 10 capacity При увеличении слайса его длина увеличивается с оезервом capacity помогает заранее выделить память
// var letters = []string{“a”, “b”, “c”}
// mySlice1 := myArray[1:5], myArray[:5], myArray[1:]
// chars = append(chars, i)
// maps
// var m map[string]int
// m = make(map[string]int)
// m = map[string]int{}
// seen := make(map[string]struct{}) //struct{} заглушка
// seen[s1] = struct{}{}

// a1 := [...]string{"live", "arp", "arp", "strong"}
// a2 := [...]string{"lively", "alive", "harp", "sharp", "armstrong"}
// a1 := [...]string{"tarp", "mice", "bull"}
// a2 := [...]string{"lively", "alive", "harp", "sharp", "armstrong"}
// var a1 = []string{"cod", "code", "wars", "ewar", "ar"}
//var a1 = []string{}

// fmt.Print(InArray(a1[:], a2[:]))
//fmt.Println(InArray(a1,a2))
// func InArray(array1 []string, array2 []string) []string {
// 	seen := make(map[string]struct{}) //struct{} заглушка
// 	result := []string{}
// 	for _, s1 := range array1 {
// 		if _, ok := seen[s1]; ok {
// 			continue
// 		}
// 		seen[s1] = struct{}{}
// 		fmt.Println(seen)
// 		for _, s2 := range array2 {
// 			if strings.Contains(s2, s1) {
// 				result = append(result, s1)
// 				break
// 			}
// 		}
// 	}
// 	sort.Strings(result)
// 	return result
// }

// func InArray(array1 []string, array2 []string) []string {

// 	sCross := []string{}
// 	sWithoutDublicates := []string{}

// 	for _, str1 := range array1 {
// 		bExist := false
// 		for _, str2 := range sWithoutDublicates {
// 			if strings.Compare(str2, str1) == 0 {
// 				bExist = true
// 				break
// 			}
// 		}
// 		if !bExist {

// 			sWithoutDublicates = append(sWithoutDublicates, str1)
// 		}

// 	}

// 	sort.Strings(sWithoutDublicates)

// 	for _, str1 := range sWithoutDublicates {
// 		for _, str2 := range array2 {
// 			if strings.Contains(str2, str1) {
// 				sCross = append(sCross, str1)
// 				break
// 			}
// 		}
// 	}

// 	return sCross
// }

// func ToCamelCase(s string) string {

// fmt.Println(ToCamelCase(""))
//   //if
// 	   var x = regexp.MustCompile(`[_-](\w)`)
//       res := regex.FindAllStringSubmatch(s, -1)

//         for i := range res {
//            s= strings.Replace(s, res[i][0], strings.ToUpper(res[i][1]), -1)
//         }

// 	return s
// }

/* var romanNumeralDict = func() map[string]int {
	return map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
}
func Decode(roman string) int {

	sumRN := 0
	maxRN := 0

	for i := len(roman) - 1; i >= 0; i-- {

		curN := romanNumeralDict()[string(roman[i])]

		if maxRN <= curN {
			sumRN += curN
			maxRN = curN
		} else {
			sumRN -= curN
		}
	}
  return sumRN
}
*/

/*  // Печать серии
 fmt.Printf ("Printf:% s \ n", "Распечатать в соответствии с форматом, нужно добавить пробел и разрыв строки")
 fmt.Print ("Print:", "Формат по умолчанию распечатывается, несимвольные строки автоматически добавляются с пробелами, а разрывы строк нужно добавлять самостоятельно \ n")
 fmt.Println ("Println:", "Печать в формате по умолчанию, все элементы автоматически добавляются с пробелами, а разрывы строк добавляются автоматически")

 // Серия спринтов
 str: = fmt.Sprintf ("Sprintf:% s \ n", "Печать в соответствии с форматом, вам нужно добавить пробел и разрыв строки")
fmt.Print(str)
 str = fmt.Sprint ("Sprint:", "Формат по умолчанию печатается, в несимвольных строках автоматически добавляются пробелы, а разрывы строк нужно добавлять самостоятельно \ n")
fmt.Print(str)
 str = fmt.Sprintln ("Sprintln:", "Формат по умолчанию распечатывается, все элементы автоматически добавляют пробелы и автоматически добавляют новые строки")
fmt.Print(str)

 // Серия Fprint
 fmt.Fprintf (os.Stdout, "Fprintf:% s \ n", "Печать в соответствии с форматом, вам нужно добавить пробел и обернуть самостоятельно")
 fmt.Fprint (os.Stdout, "Fprint:", "Печатать в формате по умолчанию, автоматически добавлять пробелы для нестроковых символов и их нужно добавлять для разрывов строк \ n")
 fmt.Fprintln (os.Stdout, "Fprintln:", "Печать в формате по умолчанию, все элементы автоматически добавляются с пробелами и автоматически добавляются разрывы строк")


 // Серия Scanf
var a, b, c int

 fmt.Scanf ("% d% d% d", & a, & b, & c) // Используйте любой пробел или табуляцию для разделения при вводе, без разрыва строки
fmt.Println(a, b, c)
 fmt.Scanf ("% d,% d,% d", & a, & b, & c) // Используйте запятые для разделения ввода, оставьте то же самое с кавычками, другие символы допустимы
fmt.Println(a, b, c)
 fmt.Scan (& a, & b, & c) // отделяется пробелом, клавишей TAB или возвратом каретки при вводе, пока не будет достигнуто количество параметров
fmt.Println(a, b, c)
 fmt.Scanln (& a, & b, & c) // Используйте любой пробел, TAB для разделения при вводе, нельзя использовать новую строку, новая строка завершается
fmt.Println(a, b, c)

input := "12\n\n34\n56"
 fmt.Sscanf (input, "% d \ n \ n% d \ n% d", & a, & b, & c) // ввод соответствует строке% d, если между% d нет символа, то ввод Может быть любым пробелом или TAB
fmt.Println(a, b, c)
 fmt.Sscan (input, & a, & b, & c) // Любой пробел, табуляция, интервал возврата каретки можно использовать во вводе до тех пор, пока не будет достигнуто количество параметров
fmt.Println(a, b, c)
input2 := "23 45 67"
 fmt.Sscanln (input2, & a, & b, & c) // ввод может использовать только произвольные пробелы и интервалы табуляции, без разрывов строк, разрывы строк завершаются
fmt.Println(a, b, c)


 fmt.Fscanf (os.Stdin, "% d,% d,% d", & a, & b, & c) // При вводе используйте запятые для разделения, оставьте то же самое с кавычками, другие символы допустимы. Когда нет интервала, ввод по умолчанию разделяется любым пробелом или табуляцией и не может быть перенесен на строку
fmt.Println(a, b, c)
 fmt.Fscan (os.Stdin, & a, & b, & c) // Используйте любой пробел, TAB или возврат каретки для разделения при вводе, пока не будет достигнуто количество параметров
fmt.Println(a, b, c)
 fmt.Fscanln (os.Stdin, & a, & b, & c) // Используйте любой пробел и TAB для разделения при вводе. Вы не можете использовать новую строку, и новая строка завершается
fmt.Println(a, b, c)

 // Серия Errorf
 err: = fmt.Errorf ("Ошибка:% s", "Ошибка теста!") // Генерировать тип ошибки
fmt.Println(err) */
