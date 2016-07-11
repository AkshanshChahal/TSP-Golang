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

#### Now to build and run

First make sure that all the directories (one per package) for parallel code are in the src folder.
To give you an idea of how the workspace will look like:
```sh
bin/
    hello                          # command executable
    main                           # command executable
pkg/
  darwin_amd64/
      tsp_parallel/
          myTSP.a           # package object
          myNode.a          # package object
          mycost.a          # package object
          prq.a             # package object
src/
  hello/
	    hello.go                    # command source
	tsp_parallel/
	    main/
	        main.go                 # command source
	    myTSP/
	        tsp.go                  # command source
	        processnodes.go         # command source
	    myNode/
	        node.go                 # command source
	    mycost/
	        cost.go                 # command source
	    prq/
	        priority_queue.go       # command source
	        
... (many more repositories and packages omitted) ...    
```	    
Now we can build and install the program with the go tool:
```sh
$ go install tsp_parallel/main
```
We can run this command from anywhere on our system. This command produces an executable binary and installs this binary to the workspace's bin directory as **main**.
We can now run the program by typing its full path at the command line:
```sh
$ $GOPATH/bin/main
```

Similarly we can do for the sequential program.


### To Do
  coming soon ...

### Sources
  coming soon ...
  
### Liscence
[MIT Liscence]





[How to Write Go Code]: <https://golang.org/doc/code.html>
[MIT Liscence]: <https://github.com/AkshanshChahal/TSP-Golang/blob/master/LICENSE>
