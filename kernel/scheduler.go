package main

const MAX_TASKS = 10

type Task struct {
	Name string
	Run  func()
}

var (
	tasks    [MAX_TASKS]Task
	numTasks int
)

func AddTask(name string, run func()) {
	if numTasks < MAX_TASKS {
		tasks[numTasks] = Task{Name: name, Run: run}
		numTasks++
	}
}

func Schedule() {
	for {
		for i := 0; i < numTasks; i++ {
			tasks[i].Run()
		}
		// Simple delay loop
		for i := 0; i < 5000000; i++ {
		}
	}
}
