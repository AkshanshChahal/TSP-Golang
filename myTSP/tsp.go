package myTSP

import(
	"fmt"
	"tsp2/mycost"
	"tsp2/myNode"
	"tsp2/prq"
	"sync"
)

const(
	MaxInt8   =  1<<7  - 1
    MinInt8   = -1<<7
    MaxInt16  =  1<<15 - 1
    MinInt16  = -1<<15
    MaxInt32  =  1<<31 - 1
    MinInt32  = -1<<31
    MaxInt64  =  1<<63 - 1
    MinInt64  = -1<<63
    MaxUint8  =  1<<8  - 1
    MaxUint16 =  1<<16 - 1
    MaxUint32 =  1<<32 - 1
    MaxUint64 =  1<<64 - 1
)

type TSP struct{
	numRows				int
	bestTour			int
	bestNode			myNode.Node
	c 					mycost.Cost
	totalNodeCount		int64
	stop				bool
	queue 				prq.PriorityQueue
	numberThreads		int
	threads 			[]ProcNodes
	numberStopped		int
    done                bool
}

// Constructor
func TSPc(costMatrix [][]int16, size int) *TSP {
	tsp := new(TSP)
	tsp.numRows = size
	tsp.bestTour = MaxUint32/2
	tsp.totalNodeCount = 0
	tsp.stop = false
    tsp.queue = prq.Pqc()
	tsp.numberThreads = size-1
	tsp.numberStopped = 0
    tsp.done = false
	tsp.threads = make([]ProcNodes, tsp.numberThreads)	
	tsp.c = mycost.Costc(size, size)
	for row := 1; row <= size; row++ {
	    for col := 1; col <= size; col++ {
	        tsp.c.AssignCost(costMatrix[row-1][col-1], row, col)
	    }
	}
	return tsp
}


func (tsp *TSP) output1(threadNumber int, queue prq.PriorityQueue, totalNodeCount int64) {
    fmt.Println()
    fmt.Println("Thread number: ", threadNumber)
    fmt.Println("Nodes generated: " , totalNodeCount / 1000000 , " million nodes.")
    fmt.Println("queue.Len(): " , queue.Len())
    if tsp.bestTour != MaxUint32/2 {                     
        fmt.Println("Best tour cost: " , tsp.bestTour)          
        fmt.Println("Best tour: " , tsp.bestNode)               
        fmt.Println()
    } 
}

func (tsp *TSP) output2(next myNode.Node, threadNumber int, queue prq.PriorityQueue, totalNodeCount *int64) {
    bestTourx := next.LowerBound()
    bestNodex := next
    if bestTourx < tsp.bestTour {     
        tsp.setBestTour(bestTourx)
        tsp.setBestNode(bestNodex)
        fmt.Println()
        fmt.Println("\nThread number: " , threadNumber , " improved the best tour.")
        fmt.Println("Nodes generated till now: " , *totalNodeCount)
        fmt.Println("Best tour cost till now: " , bestTourx)
        fmt.Println("Best tour till now: " , bestNodex.ToString(tsp.numRows))
        fmt.Println()
    } 
}


func (tsp *TSP) Stop(forced bool, threadNumber int, ch chan bool) {

    var mutex1 = &sync.Mutex{}
    if forced && tsp.stop {
        return
    }
    if tsp.queue.Len() == 0 {         
        tsp.numberStopped++
    }else { 
        if !forced {
            t := prq.Pqc()  // New Priority Queue
            mutex1.Lock()
            n := tsp.queue.First()
            mutex1.Unlock()

            t.Add(&n)
            totalNodeCount := tsp.threads[threadNumber - 1].TotalNodeCount()
            tsp.threads[threadNumber - 1] = procNodesc(tsp.numRows, threadNumber, t, tsp.c, tsp, totalNodeCount)

            go tsp.threads[threadNumber - 1].run(&tsp.bestTour, tsp.numRows, mutex1, &tsp.c, ch)
        }
    }
    if tsp.numberStopped == (tsp.numberThreads) || forced {
        tsp.stop = true
        for i := 0; i < tsp.numberThreads; i++ {
            tsp.threads[i].setStop()
            ch <- true
        }
        // Count the total number of nodes generated from the threads
        var nodesGenerated int64 = 0
        for i := 0; i < tsp.numberThreads; i++ {
            nodesGenerated += tsp.threads[i].TotalNodeCount()
        }
        tsp.totalNodeCount += nodesGenerated
        if !forced {
            fmt.Println("OPTIMUM SOLUTION OBTAINED !!")
        }else {
            fmt.Println("Solution forced to stop prematurely and may not be optimum.")
        }    
        fmt.Println("The total number of nodes generated: " , tsp.totalNodeCount)
        fmt.Println("Minimum Tour cost is : " , tsp.bestTour)
        fmt.Println("Best Tour of cities  : " , tsp.bestNode.ToString(tsp.numRows))
        tsp.done = true
        ch <- true
    }
}



func (tsp *TSP) setBestTour( bestTourx int ) {
    if bestTourx < tsp.bestTour {
        tsp.bestTour = bestTourx
    }
}

func (tsp *TSP) setBestNode( bestNode myNode.Node ) {
    tsp.bestNode = bestNode
}

func (tsp *TSP) GenerateSolution(ongoing bool) {

    var mutex = &sync.Mutex{}   // for threads synchronization
    if (!ongoing) {
        // Create root node
        cities := make([]int8, 2)
        cities[1] = 1
        root := myNode.Nodec(cities, 1, tsp.numRows)
        root.SetLevel(1)
        tsp.totalNodeCount++
        root.ComputeLowerBound(tsp.numRows, &tsp.c)
        fmt.Println("The lower bound for root node (no constraints): " , root.LowerBound())
        tsp.queue.Add(root)
        
        mutex.Lock()
        next := tsp.queue.First()
        mutex.Unlock()
        
        newLevel := next.Level() + 1;
        nextCities := next.Cities()
        size := next.Size()
        // Preparing a priority queue of nodes of level one
        for city := 2; !tsp.stop && city <= tsp.numRows; city++ {   
            if !tsp.present( int8(city), nextCities) {
                newTour := make([]int8, size+2)
                for index := 1; index <= size; index++ {
                    newTour[index] = nextCities[index]
                }
                newTour[size + 1] = int8 (city)
                newNode := myNode.Nodec(newTour, size + 1, tsp.numRows)
                newNode.SetLevel(newLevel)
                tsp.totalNodeCount++
                newNode.ComputeLowerBound(tsp.numRows, &tsp.c)
                tsp.queue.Add(newNode)
            } 
        }

        // Spawn the threads and start the process going
        for i := 0; i < tsp.numberThreads; i++ {
            t := prq.Pqc()
            mutex.Lock()
            n := tsp.queue.First()  // Popping from the priority queue
            mutex.Unlock()
            sumThis := 0
            for ix := 1; ix <= n.Size(); ix++ {
                sumThis += int(n.Cities()[ix]);
            }
            t.Add(&n)
            tsp.threads[i] = procNodesc( tsp.numRows, i+1, t, tsp.c, tsp, 0)    
        } 
    }

    ch := make(chan bool, tsp.numRows-1)    // buffered channel to ensure all goroutines have finished 
    
    for i := 0; i < tsp.numberThreads; i++ {
        go tsp.threads[i].run(&tsp.bestTour, tsp.numRows, mutex, &tsp.c, ch)   // goroutines
    }

    for i := 0; i < tsp.numRows; i++ {  // waiting for the goroutines to execute completely
        <-ch   
    }

}

func (tsp *TSP) nodesGenerated() int64 {
    return tsp.totalNodeCount
}
func (tsp *TSP) present( city int8, cities []int8) bool {
    for i := 1; i <= len(cities) - 1; i++ {
        if cities[i] == city {
            return true
        }
    }
    return false
}


