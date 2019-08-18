package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	afterFunc()
}

func parse(){

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value
}

func format(){
	t:=time.Now().Format("2006-02-02 15:04:05")
	fmt.Println(t)
}

func tick(){
	ticker := time.Tick(3* time.Second)
	for v := range ticker{
		fmt.Println(v)
	}
}

func after(){
	aft := time.After(5 * time.Second)
	go func() {
		<-aft
		os.Exit(0)
	}()
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}

func afterFunc(){
	time.AfterFunc(5 * time.Second, func() {
		os.Exit(0)
	})
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}

func timer(){
	t := time.NewTimer(5 * time.Second)
	t.Reset(5 * time.Second)
	t.Stop()
}

func duration(){
	h,_ := time.ParseDuration("12h38m17s")
	fmt.Println(h.String())
	fmt.Println(h.Hours())
	fmt.Println(h.Minutes())
	fmt.Println(h.Seconds())
}

func round()  {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	}
}

func date(){
	bjtime := time.FixedZone("Beijing Time", int((8*time.Hour).Seconds()))
	t1 := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, time.February, 16, 0, 0, 0, 0,
		bjtime)
	fmt.Println(t1,t2,t1.In(bjtime))
}
