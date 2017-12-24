#include "textflag.h"

TEXT ·programExchange(SB), NOSPLIT, $24-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ d+16(FP), CX
    MOVQ 0(CX)(AX*8), DX
    MOVQ 0(CX)(BX*8), R9
    MOVQ BX, 0x80(CX)(DX*8)
    MOVQ AX, 0x80(CX)(R9*8)
    MOVQ R9, 0(CX)(AX*8)
    MOVQ DX, 0(CX)(BX*8)

TEXT ·programExchang2e(SB), 3, $0-24
    MOVQ 0(CX)(R8*8), SI
    MOVQ 0(CX)(DI*8), R9
    MOVQ DI, 0x80(CX)(SI*8)
    MOVQ R8, 0x80(CX)(R9*8)
    MOVQ R9, 0(CX)(R8*8)
    MOVQ SI, 0(CX)(DI*8)

