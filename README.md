# TSP-Golang
Its an implementation of Travelling Salesman Problem by Branch and Bound Method using Golang. I have implemented it both **sequentially** and **parallely** and then compared the two implementations by calculating **speedup** and **efficiency** using their execution times, for different sizes of cost matrices.

### Pre-Requisites
- Golang, Python should be installed.
- GOPATH should be properly set. For help visit [How to Write Go Code]

### How to Proceed
First use  **generate_di_graph.py** to generate a random cost matrix.
```sh
$ python generate_di_graph.py -n 21 -o output.txt
```
Here replace 21 by the size of matrix required and output.txt by the path for the file in which you want the cost matrix to be printed. Also mention the path for output.txt in **main.go** (line 26). main() function will take the cost matrix input from ouput.txt and will proceed with the calculation of the optimum Hamiltonian Path.






[How to Write Go Code]: <https://golang.org/doc/code.html>
