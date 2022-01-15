package notes
// https://play.golang.org/
import (
	"fmt"
	"strings"
	//"encoding/json"
	// "sort"
	// "strconv"
	"fmt"
	"strings"
	"sort"
)

func notes() {


var a, b, c string0
a, b, c := 80, 80, 80
a[i], a[j] = a[j], a[i] //tuple assignment
defer
fmt.Printf("%3d, %5d\n", ai[i], rnd)
Pointers
p:= &v
v= *p
p:= new(string)

//format
fmt.Sprintf("Binary: %b\\%b", 4, 5) //`Binary: 100\101`
// []int64{0, 1}
// [0 1]	        %v	Default format
// []int64{0, 1}	%#v	Go-syntax format
// []int64	        %T	The type of the value
// +15	            %+d	Always show sign
// ␣␣café	        %6s	Width 6, right justify
// café␣␣	        %-6s	Width 6, left justify
// "café"	        %q	Quoted string
// \b	            U+0008 backspace
// \t	            U+0009 horizontal tab
// \n	            U+000A line feed or newline
// 1.234560e+02	%e	Scientific notation
// 123.456000	    %f	Decimal point, no exponent
// 123.46	        %.2f	Default width, precision 2
// ␣␣123.46	    %8.2f	Width 8, precision 2
// uuid  len=36    a5550ca8-8dfe-4704-94e0-853c1f87c3ee



//switch
switch hour := time.Now().Hour(); { // missing expression means "true"
case hour < 12:
    fmt.Println("Good morning!")
case ' ', '\t', '\n', '\f', '\r':
default:

//arrays
var myArray [10] int
myArray := [3]string{"arp", "live", "strong"}
myArray := [...]string{"arp", "live", "strong"}

//slice рекомендуют использовать в основном
var mySlice []float64
var mySlice  = []string{} []
mySlice  := []string{}
var clues = [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5}}
mySlice := make([]string, 5), make([]string, 5),
x := make([]float64, 5, 10) //здесь 10 capacity При увеличении слайса его длина увеличивается с оезервом capacity помогает заранее выделить память
var letters = []string{“a”, “b”, “c”}
mySlice1 := myArray[1:5], myArray[:5], myArray[1:]
chars = append(chars, i)
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"


//slice 2 array
arr[:]  // arr is an array; arr[:] is the slice of all elements
slice := []byte("abcdefgh")
var arr [4]byte
copy(arr[:], slice)


//maps
var m map[string]int
m = make(map[string]int)
m = map[string]int{}
seen := make(map[string]struct{}) //struct{} заглушка
seen[s1] = struct{}{}
sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

 




//loops
for i := 0; i <= 4; i++ {
for ; sum < 1000; {
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
for {
    if someCondition {
        break
    }
    // do action here
}

for i, s1 := range array1 { //i=0...
	if _
, ok := seen[s1]; ok {
		continue
			break
	}
}
for range sharks {
    sharks = append(sharks, "shark")
}
Nested For Loops
goto small
small:







//sort
sort.Strings(keys)

sort.Slice(family, func(i, j int) bool {
sort.SliceStable(family, func(i, j int) bool {
    return family[i].Age < family[j].Age
})

//===================
m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
//===================



//strings, runes
sl := strings.Split(s,"")
strings.ToUpper(isbn)
sTrimLast := strings.TrimRight(sRepeatRegex, `\.`)
strings.Fields("  foo bar  baz   ") //["foo" "bar" "baz"]
strings.Join(s, ", ")
repeated:=strings.Repeat(str, 4)
strings.Contains("0123456789", str)
int2str
strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A)), ", "), "[]")
ra :=[]rune(tab)
tab=string(ra)

strings.FieldsFunc(“/a/b/c”, func(c rune) bool { return c == ‘/’ })


func stringSum(s string) int { // calculates sum of str's digits
sum := 0
for _, v := range s {
    sum += int(v) - '0'
    }
    return sum
}

str := "ab£"
chars := []rune(str) //{'a', 'b'}
for i := 0; i < len(chars); i++ {
    char := string(chars[i])
    println(char)
}




//regex
 re0 := regexp.MustCompile(`\b+`)
  for _, str1 := range re.FindAllString(strArr[0], -1){
}
re.MatchString//true
re.ReplaceAllString("-ab-axxb-", "$1W")
re.Split("banana", -1)
(?i) ignore case



//big
sumBildNum := big.NewInt(0)
sumBildNum.Add( sumBildNum, new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(3), nil))
lightSpeed := big.NewInt(299792)
secondsPerDay := big.NewInt(86400)
distance := new(big.Int)
distance.SetString("24000000000000000000", 10)
 seconds := new(big.Int)
c.Add(a, b) //c=a+b
c.Sub(a, b) //c=a-b
c.Mul(a, b) //c=a*b
seconds.Div(distance, lightSpeed)
c.Exp(a, b) // z = x**y
x.Cmp(y) //-1 if x <  y
Use Float.Quo for big.Float division:
z := new(big.Float).Quo(x, y)


//math
//Округление например до 5ти
p := math.Round(x/5)*5
math.Floor(1.51)Floor returns the greatest integer value less!! than or equal to x.
Ceil(x float64) float64 Ceil returns the least integer value greater!!! than or equal to x.

math.Sqrt(x float64) float64





//time
https://pkg.go.dev/time
time.Parse(clockLayout, clock)
// clockLayout -шаблон на основе 2006-01-02T15:04:05-0700
time.UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
const clockLayout = "15:04:05"
var clockZero, _ = time.Parse(clockLayout, "00:00:00")
const dataLayout = "2006-Jan-02"
t, _ = time.Parse(dataLayout, "2013-Feb-03")

 t1.Sub(t0).Minutes() //t1-t0

h, m, s := c.Clock()
d := time.Duration(h)*time.Hour +
	time.Duration(m)*time.Minute +
	time.Duration(s)*time.Second

t.Format(time.UnixDate)
t.UTC().Format(time.UnixDate)

// Подчеркивание представляет пробел, если дата имеет только одну цифру.
do("Пробелы", "<_2>", "< 7>")
// "0" указывает заполнение нулями для однозначных значений.
do("Нули", "<02>", "<07>")

//24h
layout1 := "03:04:05PM"
layout2 := "15:04:05"
t, err := time.Parse(layout1, "07:05:45PM")
fmt.Println(t.Format(layout1))
fmt.Println(t.Format(layout2))
currentTime.Format("2006-01-02 15:04:05"))

//add
t.Add(time.Hour * 1)
AddDate(years int, months int, days int)
t.AddDate(0, 0, 1) //add day





//Go routines and multithreading
go func(){fmt.Println("Hello inline")}()
// ослать что-то в канал, синтаксис будет 
ch <- 1
// Что-то получить из канала можно так: 
var := <- ch
// A send to a nil channel blocks forever
// A receive from a nil channel blocks forever
// A send to a closed channel panics
// A receive from a closed channel returns the zero value immediately

time.Sleep(time.Second)





go func(msg string) {
        fmt.Println(msg)
    }("going")o
 go f("goroutine")



/***
Функции GoLang
***/

// Неопределенные параметры означают, что количество параметров, передаваемых функцией, является неопределенным числом.
func myfunc(args ...int) {
for _, arg := range args {
fmt.Println(arg)
}
}
// Анонимная функция
// может передаваться в качестве параметра в другую функцию
f := func(x, y int) int {
return x + y
}

func action(n1 int, n2 int, operation func(int, int) int){
    result := operation(n1, n2)
    fmt.Println(result)
}

// Функция как результат другой функции
func multiply(x int, y int) int{ return x * y}
func selectFn(n int) (func(int, int) int){
        return multiply
}


// Преимуществом анонимных функций является то, что они имеют доступ к окружению, в котором они вызываются
// замыкания (closures)
// Замыкание - это функция, которая ссылается к переменным вне ее тела. 
f := square()
fmt.Println(f())        // 9


/***
побитовые операции
***/
var a uint = 60	/* 60 = 0011 1100 */  
var b uint = 13	/* 13 = 0000 1101 */
var c uint = 0          

c = a & b       /* 12 = 0000 1100 */ 
c = a | b       /* 61 = 0011 1101 */
xor
c = a ^ b       /* 49 = 0011 0001 */
invert
^0101 =1010	Bitwise NOT (same as 1111 ^ 0101
// shift left
// Сдвиг влево смещает биты влево на указанное вами количество разрядов, заполняя созданное пространство нулевыми битами.
// Из-за того, как работает двоичный код, где каждый столбец представляет собой последовательную степень 2 (1, 2, 4, 8, 16…), сдвиг влево на два разряда аналогичен умножению на 4.
c = a << 2     /* 240 = 1111 0000 */
c = a >> 2     /* 15 = 0000 1111 */

// math/bits
bits.UintSize	32 or 64	Size of a uint in bits
bits.Reverse8(00101110)	01110100	Bits in reversed order

// Маска - это просто кусок единиц и нулей той же длины, что и ваш входной двоичный файл. Все, что вы хотите изменить в своем вводе, вы устанавливаете на 0, а все остальное оставляете прежним. Таким образом, при запуске input & mask

~0 << n
// По-английски это гласит: «Создайте поле из единиц, достаточно большое, чтобы заполнить пространство, с которым я работаю, а затем сдвиньте его влево на n мест».



//struct
type Vertex struct {
	X int
	Y int
}
v := Vertex{1, 2}
	v.X = 4
	
Verbose bool   `json:"verbose,omitempty"`
// тег omitempty, указанный для всех полей. Это означает, что поля с нулевыми значениями не будут внесены в JSON.


}