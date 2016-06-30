package prq

import(
	"tsp3/myNode"
	"container/heap"
)

type PriorityQueue []*myNode.Node

// constructor
func Pqc() PriorityQueue {			// NEW STUFF
	pq := make(PriorityQueue,0)
	heap.Init(&pq)
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) size() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// return pq[i].priority > pq[j].priority
	node1 := pq[i]								
	node2 := pq[j]								

	if node1==nil || node2==nil {
		return false
	}
	if node1.Size() < node2.Size() {
        return false
    }else{
    	if node1.Size() > node2.Size() {
	        return true
	    }else{ 
	    	if node1.Size() == node2.Size() {
		        if node1.LowerBound() < node2.LowerBound() {
		            return true
		        }else{ 
		        	if node1.LowerBound() > node2.LowerBound() {
			            return false
			        }else{ 
			        	if node1.LowerBound() == node2.LowerBound() {
				            // Add up the sum of the Cities()
				            sumThis := 0
				            for i := 1; i <= node1.Size(); i++ {
				                sumThis += int(node1.Cities()[i]);
				            }
				            sumOther := 0
				            for i := 1; i <= node1.Size(); i++ {
				                sumOther += int(node2.Cities()[i]);
				            }

				            if sumThis <= sumOther {
				                return true
				            }else{ 
				            	if sumThis > sumOther {
				                	return false
				                }
				            } 
				        }
			        }
		        }
		    }
    	}
    }
    return true
}


func (pq PriorityQueue) Swap(i, j int) {
	if pq.Len()<1 {
		return
	}
	if i>=pq.Len() || j>=pq.Len() {
		return
	}
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*myNode.Node)
	node.Index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

func (pq *PriorityQueue) Add(node *myNode.Node) {
	if(pq.Len()<1){
		*pq = append(*pq, node)
	}else{
		heap.Push(pq, node)
	}
	
}


func (pq *PriorityQueue) First() myNode.Node {		
	node := heap.Pop(pq).(*myNode.Node)
	return *node
}


