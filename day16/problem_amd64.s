#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT day16·programExchange<>(SB),NOSPLIT,$0
    // x := p.data[0 + m.A]
    MOVQ 0(CX)(R8*8), SI
    // y := p.data[0 + m.B]
    MOVQ 0(CX)(DI*8), R9
    // p.data[16 + x] = m.B
    MOVQ DI, 0x80(CX)(SI*8)
    // p.data[16 + y] = m.A
    MOVQ R8, 0x80(CX)(R9*8)
    // p.data[0 + m.A] = y
    MOVQ R9, 0(CX)(R8*8)
    // p.data[0 + m.B] = x
    MOVQ SI, 0(CX)(DI*8)
    INCQ BX
    CMPQ BX, AX
    JGE exit
    LEAQ 0(BX)(BX*2), SI
    MOVQ 0x10(DX)(SI*8), DI
    MOVQ 0x8(DX)(SI*8), R8
    MOVQ 0(DX)(SI*8), SI
    LEAQ day16·programPointers<>(SB), R9
    LEAQ (R9)(SI*8), R9
    JMP 0(R9)
exit:
    RET

TEXT day16·programSpin<>(SB),NOSPLIT,$0
    MOVQ 0x100(CX), SI
    SUBQ R8, SI
    ANDQ $0xf, SI
    MOVQ SI, 0x100(CX)
    INCQ BX
    CMPQ BX, AX
    JGE exit
    LEAQ 0(BX)(BX*2), SI
    MOVQ 0x10(DX)(SI*8), DI
    MOVQ 0x8(DX)(SI*8), R8
    MOVQ 0(DX)(SI*8), SI
    LEAQ day16·programPointers<>(SB), R9
    LEAQ (R9)(SI*8), R9
    JMP 0(R9)
exit:
    RET

TEXT day16·programPartner<>(SB),NOSPLIT,$0
    // x := p.data[16 + m.A]
    MOVQ 0x80(CX)(R8*8), SI
    // y := p.data[16 + m.B]
    MOVQ 0x80(CX)(DI*8), R9
    // p.data[0 + x] = m.B
    MOVQ DI, 0(CX)(SI*8)
    // p.data[0 + y] = m.A
    MOVQ R8, 0(CX)(R9*8)
    // p.data[16 + m.A] = y
    MOVQ R9, 0x80(CX)(R8*8)
    // p.data[16 + m.B] = x
    MOVQ SI, 0x80(CX)(DI*8)
    INCQ BX
    CMPQ BX, AX
    JGE exit
    LEAQ 0(BX)(BX*2), SI
    MOVQ 0x10(DX)(SI*8), DI
    MOVQ 0x8(DX)(SI*8), R8
    MOVQ 0(DX)(SI*8), SI
    LEAQ day16·programPointers<>(SB), R9
    LEAQ (R9)(SI*8), R9
    JMP 0(R9)
exit:
    RET

TEXT day16·programLoop<>(SB),NOSPLIT,$0
    INCQ BX
    CMPQ BX, AX
    JGE exit
    LEAQ 0(BX)(BX*2), SI
    MOVQ 0x10(DX)(SI*8), DI
    MOVQ 0x8(DX)(SI*8), R8
    MOVQ 0(DX)(SI*8), SI
    LEAQ day16·programPointers<>(SB), R9
    LEAQ (R9)(SI*8), R9
    JMP 0(R9)
exit:
    RET


DATA day16·programPointers<>+0(SB)/8,$day16·programExchange<>(SB)
DATA day16·programPointers<>+8(SB)/8,$day16·programPartner<>(SB)
DATA day16·programPointers<>+16(SB)/8,$day16·programSpin<>(SB)
GLOBL day16·programPointers<>(SB),RODATA+NOPTR,$24

TEXT ·programDance(SB),NOSPLIT,$0-32
    MOVQ pr+0(FP), CX
    MOVQ moves+8(FP), DX
    MOVQ moves_len+16(FP), AX
    XORQ BX, BX
    DECQ BX
    JMP day16·programLoop<>(SB)



TEXT ·programDance2(SB), $0-32
    MOVQ pr+0(FP), CX
    MOVQ moves+8(FP), DX
    MOVQ moves_len+16(FP), AX
    XORQ BX, BX
    JMP check
loop:
    INCQ BX
check:
    CMPQ BX, AX
    JGE exit
    LEAQ 0(BX)(BX*2), SI
    MOVQ 0x10(DX)(SI*8), DI
    MOVQ 0x8(DX)(SI*8), R8
    SHLQ $0x03, SI
    MOVBQZX 0x1(DX)(SI*1), R9
    MOVBQZX 0(DX)(SI*1), SI
    TESTL R9, R9
    JE not_exchange
    // x := p.data[0 + m.A]
    MOVQ 0(CX)(R8*8), SI
    // y := p.data[0 + m.B]
    MOVQ 0(CX)(DI*8), R9
    // p.data[16 + x] = m.B
    MOVQ DI, 0x80(CX)(SI*8)
    // p.data[16 + y] = m.A
    MOVQ R8, 0x80(CX)(R9*8)
    // p.data[0 + m.A] = y
    MOVQ R9, 0(CX)(R8*8)
    // p.data[0 + m.B] = x
    MOVQ SI, 0(CX)(DI*8)
    JMP loop
not_exchange:
    TESTL SI, SI
    JE not_spin
    MOVQ 0x100(CX), SI
    SUBQ R8, SI
    ANDQ $0xf, SI
    MOVQ SI, 0x100(CX)
    JMP loop
not_spin:
    // x := p.data[16 + m.A]
    MOVQ 0x80(CX)(R8*8), SI
    // y := p.data[16 + m.B]
    MOVQ 0x80(CX)(DI*8), R9
    // p.data[0 + x] = m.B
    MOVQ DI, 0(CX)(SI*8)
    // p.data[0 + y] = m.A
    MOVQ R8, 0(CX)(R9*8)
    // p.data[16 + m.A] = y
    MOVQ R9, 0x80(CX)(R8*8)
    // p.data[16 + m.B] = x
    MOVQ SI, 0x80(CX)(DI*8)
    JMP loop
exit:
    RET
