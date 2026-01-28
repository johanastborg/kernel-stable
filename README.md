# Go Kernel Project

This project is a simple x86 kernel written in Go and Assembly, designed to be booted by a Multiboot-compliant bootloader (like QEMU's `-kernel` option or GRUB).

## Prerequisites

To build and run this kernel on Ubuntu Linux, you will need the following tools:

1.  **Go (version 1.22+)**: The Go programming language.
2.  **Make**: Build automation tool.
3.  **QEMU**: A generic and open source machine emulator and virtualizer (specifically `qemu-system-i386`).

### Installation on Ubuntu

```bash
sudo apt update
sudo apt install golang-go make qemu-system-x86
```

Ensure that Go is in your PATH. You can verify the installation with:

```bash
go version
make --version
qemu-system-i386 --version
```

## Building the Kernel

To build the kernel executable (`kernel.bin`), simply run `make` in the root directory of the project:

```bash
make
```

This command will compile the Go code and Assembly files, linking them into a Multiboot-compliant ELF binary named `kernel.bin`.

## Running the Kernel

You can run the kernel using QEMU. The `Makefile` includes a convenience target for this.

To build and run the kernel in one step:

```bash
make run
```

Alternatively, if you want to run `kernel.bin` manually using QEMU:

```bash
qemu-system-i386 -kernel kernel.bin
```

## Cleaning Up

To remove the generated `kernel.bin` file:

```bash
make clean
```
