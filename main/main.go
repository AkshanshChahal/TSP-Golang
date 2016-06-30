package main

import(
	"fmt"
	"tsp2/myTSP"
	"time"
	"io/ioutil"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {

	start := time.Now()

	/**
		Taking the cost matrix input from a file	
	*/
	dat, err := ioutil.ReadFile("/Users/akshanshchahal/Desktop/mfile.txt")
    check(err)
    s := string(dat)
    x := strings.Fields(s)	// splitting a string into slice of strings about whitespaces 
    size, err := strconv.Atoi(x[0])
    costMatrix := make([][]int16, size)
    k := 1
    var y int
    for i := 0; i < size; i++ {
        costMatrix[i] = make([]int16, size)
        for j := 0; j < size; j++ {
            y, err = strconv.Atoi(x[k])
            costMatrix[i][j] = int16(y)
            k++
        }
    }

	mytsp := myTSP.TSPc(costMatrix, size)
	mytsp.GenerateSolution(false)	
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println("Time taken for ", size , " CITIES by the parallel program : ", elapsed)
	fmt.Println()		

}
