package main

func main() {
	tasks, results := make(chan int, 45), make(chan int, 45)

	go worker(tasks, results) //Use 5 processes to get fibonacci
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		println(<-results)
	}
}

func worker(tasks <-chan int, results chan<- int) { //Only input and Only output
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
