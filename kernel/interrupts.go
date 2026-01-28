package main

import "unsafe"

// IDT Entry
type IDTEntry struct {
	OffsetLow  uint16
	Selector   uint16
	Zero       uint8
	TypeAttr   uint8
	OffsetHigh uint16
}

// IDT Pointer - packed
// We need to ensure no padding. Go structs are usually aligned.
// However, IDTPtr is 6 bytes.
// We should pack it. But Go doesn't have __attribute__((packed)).
// We can use an array of bytes or construct it carefully.
// But LoadIDT takes a pointer.
// asm expects 6 bytes at the pointer.
// If struct is padded, it might be 8 bytes.
// offset 0: Limit (2 bytes)
// offset 2: Base (4 bytes)
// offset 6: Padding (2 bytes)
// 'lidt' instruction takes 6 bytes memory operand. It will read 6 bytes.
// If padding is at the end, it's fine.
// If padding is in between, it fails.
// uint16 is 2 bytes aligned. uint32 is 4 bytes aligned.
// So there will be 2 bytes padding between Limit and Base.
// Limit (0-1), Padding (2-3), Base (4-7). Total 8 bytes.
// This is BAD for 'lidt'.
// Solution: Use a byte array or pass Limit and Base separately to ASM and construct there.

// Defined in asm.S
var idt [256]IDTEntry

func AsmOutb(port uint16, data uint8)
func AsmInb(port uint16) uint8
func LoadIDT(ptr unsafe.Pointer) // Changed to unsafe.Pointer to pass byte array
func GetKeyboardHandlerAddr() uint32
func EnableInterrupts()

//go:nosplit
//export KeyboardHandler
func KeyboardHandler() {
	// Read scan code
	scancode := AsmInb(0x60)

	// Send EOI to PIC (Master)
	AsmOutb(0x20, 0x20)

	// Simple output
	if scancode < 0x80 { // Press event
		Print("K")
	}
}

//go:nosplit
func InitInterrupts() {
	// Remap PIC
	AsmOutb(0x20, 0x11)
	AsmOutb(0xA0, 0x11)
	AsmOutb(0x21, 0x20)
	AsmOutb(0xA1, 0x28)
	AsmOutb(0x21, 0x04)
	AsmOutb(0xA1, 0x02)
	AsmOutb(0x21, 0x01)
	AsmOutb(0xA1, 0x01)
	AsmOutb(0x21, 0xFD) // Unmask IRQ1
	AsmOutb(0xA1, 0xFF)

	keyboardAddr := GetKeyboardHandlerAddr()

	// IRQ1 is INT 0x21
	idt[0x21] = IDTEntry{
		OffsetLow:  uint16(keyboardAddr & 0xFFFF),
		Selector:   0x08,
		Zero:       0,
		TypeAttr:   0x8E,
		OffsetHigh: uint16((keyboardAddr >> 16) & 0xFFFF),
	}

    // Construct IDTPtr manually to avoid padding
    // 6 bytes: Limit (2), Base (4)
    var idtPtr [6]byte
    limit := uint16(unsafe.Sizeof(idt)) - 1
    base := uint32(uintptr(unsafe.Pointer(&idt)))

    idtPtr[0] = byte(limit & 0xFF)
    idtPtr[1] = byte((limit >> 8) & 0xFF)
    idtPtr[2] = byte(base & 0xFF)
    idtPtr[3] = byte((base >> 8) & 0xFF)
    idtPtr[4] = byte((base >> 16) & 0xFF)
    idtPtr[5] = byte((base >> 24) & 0xFF)

	LoadIDT(unsafe.Pointer(&idtPtr[0]))

    EnableInterrupts()
}
