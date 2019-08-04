package sort

import (
	"fmt"
	"sort"
)

//sortFunction
func sortNumbers() bool {
	//int, float64, string
	nums := []int{2, 1, 5, 7, 3, 4, 6, 8, 0, -1}
	var nums1 = make([]int, len(nums))
	_ = copy(nums1, nums)
	sort.Ints(nums1)

	fmt.Println("sort int slice：", nums1)
	//fmt.Println(sort.IntsAreSorted(nums1))

	//float
	floats := []float64{-0.5, 0.3, 0.998, 1.5, 0.72, 0.00}
	floats1 := make([]float64, len(floats))
	_ = copy(floats1, floats)
	sort.Float64s(floats1)
	fmt.Println("sort float slice：", floats1)
	//fmt.Println(sort.Float64sAreSorted(floats1))

	//IntSlice Type or Float64Slice Type
	intSlice := sort.IntSlice(nums)
	fmt.Println("slice[0] < slice[1] before sort :", intSlice.Less(0, 1))
	intSlice.Sort()
	fmt.Println("sort IntSlice：", intSlice)
	fmt.Println("IntSlice length：", intSlice.Len())
	fmt.Println("slice[0] < slice[1] after sort :", intSlice.Less(0, 1))

	//if x in slice, return the index, if not, return the index if value insert into sorted slice
	fmt.Println("search :", intSlice.Search(-1))
	return true
}

type personSlice []Person

type Person struct {
	Name string
	Age  int
}

func (p personSlice) Len() int {
	return len([]Person(p))
}

func (p personSlice) Less(i, j int) bool {
	return (p)[i].Age < (p)[j].Age
}

//can't be a pointer reviver
//func (p *personSlice) Swap(i, j int) {
//	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
//}
func (p personSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//notice that this method is not in the interface
func (p *personSlice) Sort()  {
	sort.Sort(*p)
}

func structSort() bool {
	ps := personSlice{
		{
			"han",
			17,
		},
		{
			"li",
			16,
		},
		{
			"zhao",
			19,
		},
	}
	//use sort.Sort() or implement a Sort func
	//sort.Sort(ps)
	ps.Sort()
	fmt.Println(ps)
	return true
}
