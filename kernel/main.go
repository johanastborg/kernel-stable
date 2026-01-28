package main

func main() {}

//go:nosplit
//export KernelMain
func KernelMain() {
	Print("Hello from Go Kernel!\n")
	Print("Initializing Interrupts...\n")
	InitInterrupts()
	Print("Interrupts Initialized. Press any key...\n")

	AddTask("Task 1", Task1)
	AddTask("Task 2", Task2)

	Print("Starting Scheduler...\n")
	Schedule()
}

func Task1() {
	Print("1")
}

func Task2() {
	Print("2")
}
