#include "textflag.h"

TEXT ·isr_keyboard_wrapper(SB), NOSPLIT, $0
	PUSHAL
	CALL ·KeyboardHandler(SB)
	POPAL
	IRETL
