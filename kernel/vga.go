package main

import "unsafe"

const (
	VGA_WIDTH  = 80
	VGA_HEIGHT = 25
	VGA_ADDR   = 0xB8000
)

var (
	cursorX int
	cursorY int
	color   byte = 0x0F // White on Black
)

//go:nosplit
func putChar(c byte) {
	if c == '\n' {
		cursorX = 0
		cursorY++
	} else {
		offset := uintptr(cursorY*VGA_WIDTH + cursorX)
		addr := VGA_ADDR + offset*2
		*(*byte)(unsafe.Pointer(addr)) = c
		*(*byte)(unsafe.Pointer(addr + 1)) = color
		cursorX++
	}

	if cursorX >= VGA_WIDTH {
		cursorX = 0
		cursorY++
	}
	if cursorY >= VGA_HEIGHT {
		cursorY = 0 // Simple wrap around for now
	}
}

//go:nosplit
func Print(s string) {
	for i := 0; i < len(s); i++ {
		putChar(s[i])
	}
}
