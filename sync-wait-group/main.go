package main

import (
	"fmt"
	"sync"

)

// it is similar to async and await keyword used in many programming languages


func worker(i int, wg * sync.WaitGroup){
	defer wg.Done()
	fmt.Println("workker started " , i);
	fmt.Println("worker  ended " , i);
}

func main() {
	fmt.Println("exploring the sync waitgroups")
	var wg sync.WaitGroup

	for i := 1; i <= 4 ;i++{
		wg.Add(1);
		go worker(i , &wg);
	}

	wg.Wait()
	fmt.Println("all worker ended");
		
}
