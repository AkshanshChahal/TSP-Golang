package mycost

type Cost struct{
	costMatrix	[][]int16
}

// constructor
func Costc(n_rows, n_cols int)	Cost {
	c := new(Cost)
	c.costMatrix = make([][]int16, n_rows + 1)
	for i := 0; i <= n_cols; i++ {
		c.costMatrix[i] = make([]int16, n_cols + 1)
	}
	return *c
}


func (cost *Cost)	AssignCost(val int16, row, col int){
	cost.costMatrix[row][col] = val
}

func (cost *Cost)	Getcost(row, col int) int16 {
	return cost.costMatrix[row][col]
} 