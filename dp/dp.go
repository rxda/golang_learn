package dp

type work struct{
	start int
	end int
	val int
}


var works = []work{
	{1,4,5},
	{3,5,2},
	{0,6,8},
	{4,7,4},
	{3,8,6},
	{5,9,3},
	{1,4,5},
	{6,10,2},
	{8,11,4},
}



func fib(n int) int{
	x, y := 0,1
	for ;n>0;n--{
		x, y = y, x+y
	}
	return y
}