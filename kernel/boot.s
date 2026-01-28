#include "textflag.h"

#define ALIGN (1<<0)
#define MEMINFO (1<<1)
#define FLAGS (ALIGN | MEMINFO)
#define MAGIC 0x1BADB002
#define CHECKSUM -(MAGIC + FLAGS)

TEXT _start(SB), NOSPLIT, $0
	JMP entry

	// Alignment padding to ensure Multiboot header is 4-byte aligned.
	// We assume JMP entry is 2 bytes (short jump).
	BYTE $0x90
	BYTE $0x90

	// Multiboot Header
	LONG $MAGIC
	LONG $FLAGS
	LONG $CHECKSUM

entry:
	// Initialize Stack
	MOVL $stack_bottom(SB), SP
	ADDL $16384, SP

	CALL ·KernelMain(SB)

loop:
	HLT
	JMP loop

GLOBL stack_bottom(SB), NOPTR, $16384
