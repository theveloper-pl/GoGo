package main

import (
	"fmt"
	"time"
	"sync"
)

type Philosopher struct{
	name string
	rightFork int
	leftFork int
}

var philosophers = []Philosopher{
	{name: "First", rightFork:4, leftFork:0},
	{name: "Second", rightFork:0, leftFork:1},
	{name: "Third", rightFork:1, leftFork:2},
	{name: "Fourth", rightFork:2, leftFork:3},
	{name: "Fifth", rightFork:3, leftFork:4},
}

var order = []string{}
var oderMutax sync.Mutex

var hunger = 3
var eatTime = 0 * time.Second
var thinkTime = 0 * time.Second
var sleepTime = 0 * time.Second

func dine(){
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers) ; i++ {
		forks[i] = &sync.Mutex{}
	}

	for i :=0; i < len(philosophers); i++ {
		go dinningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

}

func dinningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n",philosopher.name)
	seated.Done()

	seated.Wait()



	for i := hunger; i>0; i--{

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n",philosopher.name)		
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n",philosopher.name)
		} else{
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n",philosopher.name)			
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n",philosopher.name)		
		}

		fmt.Printf("%s has both forks and is eating.\n",philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking.\n",philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("%s put down the forks.\n",philosopher.name)
		

	}
	oderMutax.Lock()
	order = append(order, philosopher.name)
	oderMutax.Unlock()
	fmt.Printf("%s is satisfied.\n", philosopher.name)
	fmt.Printf("%s left the table.\n", philosopher.name)
}

func main() {

	fmt.Println("Dinning Philosophers problem")
	fmt.Println("----------------------------")
	fmt.Println("Table is empty")

	dine()

	fmt.Println("----------------------------")
	fmt.Println("Table is empty")
	fmt.Println(order)
}
