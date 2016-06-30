package myNode

import(
	"tsp2/mycost"
    "strconv"
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

type Node struct {
	lowerBound 		int
	size			int 		// Number of cities in partial tour
	cities			[]int8  	// stores partial tour
	blocked 		[]bool		// Used to Compute smallest and nextSmallest
	level 			int 		// The level in the tree
    Index           int         // The index of the item in the heap.
    // The index is needed by update and is maintained by the heap.Interface methods.
}

// constructor
func Nodec(cities []int8, size , numRows int) *Node {	
	nd := new(Node)							
	nd.size = size
	nd.cities = cities
	nd.blocked = make([]bool, numRows+1)
	return nd
}



// Commands
func (node *Node) ComputeLowerBound(numRows int, c *mycost.Cost) {
    node.lowerBound = 0
    if node.size == 1 {
        for i := 1; i <= numRows; i++ {
            node.lowerBound += node.minimum(i, numRows, c)
        }
    }else {
        // Obtain fixed portion of bound
        for i := 2; i <= node.size; i++ {
            node.blocked[node.cities[i]] = true
            node.lowerBound += int(c.Getcost(int(node.cities[i - 1]), int(node.cities[i])))
        }
        node.blocked[1] = true
        node.lowerBound += node.minimum(int(node.cities[node.size]), numRows, c)
        node.blocked[1] = false
        node.blocked[node.cities[node.size]] = true
        for i := 2; i <= numRows; i++ {
            if !node.blocked[i] {
                node.lowerBound += node.minimum(i, numRows, c)
            } 
        }
    }
}


func (node *Node) SetLevel(level int) { 
    node.level = level
}
func (node *Node) setCities(cities []int8) { 
    node.cities = cities
}

// Queries  
func (node *Node) Size() int {
    return node.size
}
func (node *Node) Cities() []int8 { 
    return node.cities    
}
func (node *Node) Level() int { 
    return node.level
}
func (node *Node) LowerBound() int { 
    return node.lowerBound
}


func (node *Node) ToString(numRows int) string {
    result := "<";
    for i := 1; i < len(node.cities); i++ {
            result += strconv.Itoa(int(node.cities[i])) + " "
    }
    if len(node.cities) == numRows {
        for i := 2; i <= numRows; i++ {
            if !node.present(int8(i), node.cities) {
                result += strconv.Itoa(i) + " "
                break
            }
        }
        result += "1>"
    }else {
        result += ">"
    }
    return result
}


func (node *Node) minimum(index, numRows int, c *mycost.Cost) int {
    smallest := MaxInt64;
    for col := 1; col <= numRows; col++ {
        if !node.blocked[col] && col != index && int(c.Getcost(index, col)) < smallest {
            smallest = int(c.Getcost(index, col))
        }
    }
    return smallest
}


func (node *Node) present(city int8, cities []int8) bool {
    for i := 1; i <= len(cities) - 1; i++ {
        if cities[i] == city {
            return true
        } 
    }
    return false
}

