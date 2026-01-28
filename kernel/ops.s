#include "textflag.h"

// func AsmOutb(port uint16, data uint8)
// Stack layout (Go 386 ABI):
// port: offset 0 (2 bytes)
// data: offset 2 (1 byte)
// (Verified via go tool compile -S)
TEXT ·AsmOutb(SB), NOSPLIT, $0
	MOVW port+0(FP), DX
	MOVB data+2(FP), AX
	OUTB
	RET

// func AsmInb(port uint16) uint8
// Stack layout (Go 386 ABI):
// port: offset 0 (2 bytes)
// padding: offset 2 (2 bytes)
// ret: offset 4 (1 byte)
// (Verified via go tool compile -S: caller reads return from 4(SP))
TEXT ·AsmInb(SB), NOSPLIT, $0
	MOVW port+0(FP), DX
	INB
	MOVB AL, ret+4(FP)
	RET

// func LoadIDT(ptr unsafe.Pointer)
TEXT ·LoadIDT(SB), NOSPLIT, $0
	MOVL ptr+0(FP), AX
	LIDT (AX)
	RET

// func EnableInterrupts()
TEXT ·EnableInterrupts(SB), NOSPLIT, $0
	STI
	RET

// func GetKeyboardHandlerAddr() uint32
TEXT ·GetKeyboardHandlerAddr(SB), NOSPLIT, $0
	MOVL $·isr_keyboard_wrapper(SB), AX
	MOVL AX, ret+0(FP)
	RET
