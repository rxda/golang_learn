package sortfunc

//冒泡排序

func IntSort(in []int) []int {
	length := len(in)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if in[i] > in[j] {
				in[i], in[j] = in[j], in[i]
			}
		}
	}
	return in
}

func findShortestSubArray(nums []int) int {
	var imap = make(map[int][]int)
	for k,v :=range nums{
		imap[v] = append(imap[v],k)
	}
	maxlen :=0
	for _,v :=range imap{
		if maxlen< len(v){
			maxlen =len(v)
		}
	}
	maxk :=make([]int, len(imap))

	for k,v :=range imap{
		if maxlen== len(v){
			maxk =append(maxk,k)
		}
	}
	minlength :=maxlen
	for _,v :=range maxk{
		length :=imap[v][len(imap[v])-1]-imap[v][0]
		if  length < minlength{
			minlength = length
		}
	}
	return minlength
}
