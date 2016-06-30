package myTSP

import(
    "fmt"
	"tsp2/myNode"
	"tsp2/prq"
	"tsp2/mycost"
	"sync"
)

type ProcNodes struct{
	queue 				prq.PriorityQueue
	numRows				int
	c 					mycost.Cost
	totalNodeCount		int64
	stop				bool
	tsp					*TSP	
	threadNumber		int
}


func procNodesc(numRows, threadNumber int, queue prq.PriorityQueue, c mycost.Cost, tsp *TSP, totalNodeCount int64) ProcNodes{
	pn := new(ProcNodes)
	pn.stop = false
	pn.numRows = numRows
	pn.threadNumber = threadNumber
	pn.queue = queue
	pn.c = c
	pn.tsp = tsp
	pn.totalNodeCount = totalNodeCount
	return *pn
}

func (pn *ProcNodes) setQueue(queue prq.PriorityQueue) {
    pn.queue = queue
}
func (pn *ProcNodes) setStop() {
    pn.stop = true
}



func (pn *ProcNodes) run(bestTour *int, tnumRows int, mutex *sync.Mutex, c *mycost.Cost, ch chan bool) {
    x := pn.threadNumber

    for !pn.stop  && pn.queue.Len() > 0 {
        mutex.Lock()
        next := pn.queue.First()  
        mutex.Unlock()
        
        if next.Size() == tnumRows - 1 && next.LowerBound() < *bestTour {			// 	TAKE CARE
        	mutex.Lock()
            pn.tsp.output2(next, pn.threadNumber, pn.queue, &pn.totalNodeCount)
            mutex.Unlock()
        }

        if next.LowerBound() < *bestTour {		
            newLevel := next.Level() + 1
            nextCities := pn.Copy(next.Cities())	
            size := next.Size()

            for city := 2; !pn.stop && city <= tnumRows; city++ {

                if !pn.present( int8(city), nextCities) {
                	newTour := make([]int8, size+2)
                    for index := 1; index <= size; index++ {
                        newTour[index] = nextCities[index]
                    }
                    newTour[size + 1] = int8(city)
                    newNode := myNode.Nodec(newTour, size + 1, tnumRows)
                    newNode.SetLevel(newLevel)
                    pn.totalNodeCount++
                    newNode.ComputeLowerBound(tnumRows, c)
                
                    if newNode.LowerBound() < *bestTour {	
                        mutex.Lock()
                        pn.queue.Add(newNode)
                        mutex.Unlock()
                    }else{
                        newNode = nil
                    }
                } 
            }
        }else {
            // next = nil 
        }
    }   // while loop ends

    fmt.Println("================================================================================================For loop ended for thread ", x)

    if !pn.stop {                            // while loop finished due to   QUEUE.SIZE() == 0
    	mutex.Lock()
        pn.tsp.Stop(false, pn.threadNumber, ch)
        mutex.Unlock()
    }
}

func (pn *ProcNodes) Queue()  prq.PriorityQueue {
    return pn.queue
}

func (pn *ProcNodes) TotalNodeCount() int64 {
    return pn.totalNodeCount
}

func (pn *ProcNodes) present(city int8, cities []int8) bool {
    for i := 1; i <= len(cities) - 1; i++ {
        if (cities[i] == city) {
            return true
        }
    }
    return false
}

func (pn *ProcNodes) Copy( cities []int8) []int8 {
    l := len(cities)
    n := make([]int8, l)
    for i := 0; i < l; i++ {
        n[i] = cities[i]
    }
    return n
}




