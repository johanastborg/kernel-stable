# Go Kernel Project

This project implements a basic x86 kernel written in Go and Assembly. It demonstrates:
- Multiboot compliant bootloader stub (written in Go Assembly).
- VGA text mode driver.
- Interrupt Descriptor Table (IDT) setup.
- Basic keyboard interrupt handling.

## Prerequisites

To build and run this kernel on Kubuntu (or any Debian-based Linux), you need the following tools:

1.  **Go (Golang)**: Version 1.22 or higher.
2.  **Make**: Build automation tool.
3.  **QEMU**: Emulator to run the kernel.

### Installation

```bash
sudo apt update
sudo apt install build-essential qemu-system-x86
```

Ensure Go is installed. If not, follow instructions at [go.dev/doc/install](https://go.dev/doc/install).

## Building

To compile the kernel, simply run:

```bash
make
```

This will produce a `kernel.bin` file in the root directory.

## Running

You can run the kernel using QEMU. The kernel is multiboot-compliant, so you can load it directly using the `-kernel` flag.

```bash
qemu-system-i386 -kernel kernel.bin
```

Alternatively, you can use the Makefile target:

```bash
make run
```

### Expected Output

1.  A QEMU window will open.
2.  You should see the text:
    ```
    Hello from Go Kernel!
    Initializing Interrupts...
    Interrupts Initialized. Press any key...
    ```
3.  Type on your keyboard. You should see "K" printed for every key press (as per the current simple interrupt handler logic).

## Project Structure

*   `kernel/boot.s`: Multiboot header and kernel entry point (`_start`).
*   `kernel/main.go`: Kernel main loop.
*   `kernel/vga.go`: VGA driver implementation.
*   `kernel/interrupts.go`: IDT setup and interrupt handling logic.
*   `kernel/ops.s`: Low-level assembly helpers for I/O ports (`inb`, `outb`) and interrupts.
*   `kernel/traps.s`: ISR wrappers.

## Notes

This kernel uses Go's internal linker with the `-buildmode=exe` (default) but tweaked flags (`-T 0x100000 -E _start`) to create a flat binary suitable for booting. It runs in 32-bit protected mode as set up by the bootloader (QEMU's internal bootloader or GRUB).
