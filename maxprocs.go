package main

import "runtime"

func main() {
	// this statement allow go to utilize max number of logical processer / cpu
	runtime.GOMAXPROCS(runtime.NumCPU())

}
