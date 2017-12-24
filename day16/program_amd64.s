TEXT ·programSpin(SB), 2, $24-24
    MOVOU control+0(FP), X1
    MOVQ data+16(FP), AX
    MOVOU 0(AX), X0
    PSHUFB X1, X0
    MOVOU X0, 0(AX)
    RET

TEXT ·programPartner(SB), 3, $32-32
    MOVB a+8(FP), AL
    MOVB b+9(FP), BL
    MOVQ data+0(FP), CX

    MOVOU 0(CX), X0

    MOVD AX, X1
    PUNPCKLBW X1, X1
    PUNPCKLBW X1, X1
    PSHUFL $0, X1, X1

    MOVD BX, X2
    PUNPCKLBW X2, X2
    PUNPCKLBW X2, X2
    PSHUFL $0, X2, X2

    PCMPEQB X0, X1
    PMOVMSKB X1, AX
    BSFL AX, AX

    PCMPEQB X0, X2
    PMOVMSKB X2, BX
    BSFL BX, BX

    MOVQ AX, pa+16(FP)
    MOVQ BX, pb+24(FP)
    RET
