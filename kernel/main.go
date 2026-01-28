package main

func main() {}

//go:nosplit
//export KernelMain
func KernelMain() {
	Print("Hello from Go Kernel!\n")
	Print("Initializing Interrupts...\n")
	InitInterrupts()
	Print("Interrupts Initialized. Press any key...\n")

	for {
	}
}
