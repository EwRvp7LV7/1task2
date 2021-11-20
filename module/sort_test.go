package module

import (
	"sort"
	"testing"
)

//https://pkg.go.dev/testing
//go test -v -test.bench=. sort_test.go
//go test -v -run BenchmarkSlice sort_test.go
// тестирование времени выполнения
// go test -bench=. видимо скоро устареет
// go test -test.bench=. -test.benchtime 100x крутануть тест 100 раз
// go test -test.bench BenchmarkSlice тастирует все BenchmarkSlice*
// сохрранить в файл
// go test -bench=. -benchmem bench_test.go > new.txt

// Каждый бенчмарк выполняется по многу раз, при этом измеряется время каждого из них. На основе общего количества подсчитывается среднее время выполнения. 
// ns/op означает число наносекунд, затраченных на операцию. Это время, которое потребовалось на вызов метода.  
// добавление флага -benchmem позволит получить дополнительную информацию о том, сколько было выделено байт за операцию (B/op), а также сколько раз за операцию выделялась память (allocs/op).

// Колонки идут в таком порядке:
// название теста - количество процессоров или ядер; BenchmarkPrimeNumbers-4
// количество итераций, за которое произведены измерения;
// среднее время итерации;
// среднее общее количество памяти, выделенное на итерацию;
// среднее количество отдельных выделений памяти на итерацию.


// можно вставлять в код 
// b.StartTimer()  b.StopTimer() и b.ResetTimer().
// вы можете также задать пропускную способность за одну итерацию в байтах при помощи метода b.SetBytes(n int64).

//Performance: Sorting Slice vs Sorting Type (of Slice) with Sort implementation
//https://stackoverflow.com/questions/54276285/performance-sorting-slice-vs-sorting-type-of-slice-with-sort-implementation

// Example of struct we going to sort.

type Point struct {
    X, Y int
}

// --- Struct / Raw Data
var TestCases = []Point{
    {10, 3},
    {10, 4},
    {10, 35},
    {10, 5},
    {10, 51},
    {10, 25},
    {10, 59},
    {10, 15},
    {10, 22},
    {10, 91},
}

// Example One - Sorting Slice Directly
// somehow - slowest way to sort it.
func SortSlice(points []Point) {
    sort.Slice(points, func(i, j int) bool {
        return points[i].Y < points[j].Y
    })
}

func BenchmarkSlice(b *testing.B) {
    tmp := make([]Point, len(TestCases))
    for i := 0; i < b.N; i++ {
        copy(tmp, TestCases)
        SortSlice(tmp)
    }
}

// Example Two - Sorting Slice Directly
// much faster performance
type Points []Point

// Sort interface implementation
func (p Points) Less(i, j int) bool { return p[i].Y < p[j].Y }
func (p Points) Len() int           { return len(p) }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func SortStruct(points []Point) {
    sort.Sort(Points(points))
}

func BenchmarkStruct(b *testing.B) {
    tmp := make([]Point, len(TestCases))
    for i := 0; i < b.N; i++ {
        copy(tmp, TestCases)
        SortStruct(tmp)
    }
}

// // --- Pointers
// var TestCasesPoints = []*Point{
//     &Point{10, 3},
//     &Point{10, 4},
//     &Point{10, 35},
//     &Point{10, 5},
//     &Point{10, 51},
//     &Point{10, 25},
//     &Point{10, 59},
//     &Point{10, 15},
//     &Point{10, 22},
//     &Point{10, 91},
// }
// Например, не избыточная версия следующего составного литерала:
// []*Product{&Product{}, &Product{}}
// будет выглядеть так:
// []*Product{{}, {}}
// --- Pointers
var TestCasesPoints = []*Point{
	{10, 3},
	{10, 4},
	{10, 35},
	{10, 5},
	{10, 51},
	{10, 25},
	{10, 59},
	{10, 15},
	{10, 22},
	{10, 91},
}

// Example Three - Sorting Slice of Pointers
func SortSlicePointers(points []*Point) {
    sort.Slice(points, func(i, j int) bool {
        return points[i].Y < points[j].Y
    })
}

func BenchmarkSlicePointers(b *testing.B) {
    tmp := make([]*Point, len(TestCasesPoints))
    for i := 0; i < b.N; i++ {
        copy(tmp, TestCasesPoints)
        SortSlicePointers(tmp)
    }
}

// Example Four - Sorting Struct (with Slice of pointers beneath it)
type PointsPointer []*Point

func (pp PointsPointer) Less(i, j int) bool { return pp[i].Y < pp[j].Y }
func (pp PointsPointer) Len() int           { return len(pp) }
func (pp PointsPointer) Swap(i, j int)      { pp[i], pp[j] = pp[j], pp[i] }

func SortStructOfSlicePointers(points []*Point) {
    sort.Sort(PointsPointer(points))
}

func BenchmarkStructOfSlicePointers(b *testing.B) {
    tmp := make([]*Point, len(TestCasesPoints))

    for i := 0; i < b.N; i++ {
        copy(tmp, TestCasesPoints)
        SortStructOfSlicePointers(tmp)
    }
}




// var fibTests = []struct {
// 	n        int // input
// 	expected int // expected result
// }{
// 	{1, 1},
// 	{2, 1},
// 	{3, 2},
// 	{4, 3},
// 	{5, 5},
// 	{6, 8},
// 	{7, 13},
// }

// func TestFib(t *testing.T) {
// 	for _, tt := range fibTests {
// 		actual := Fib(tt.n)
// 		if actual != tt.expected {
// 			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
// 			//             t.Error(
// 			//                 "For", pair.values,
// 			//                 "expected", pair.average,
// 			//                 "got", v,
// 			//             )
// 		}
// 	}
// }
