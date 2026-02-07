org $300
csprite
    db $F0 $80 $80 $80 $F0
asprite
    db $F0 $90 $F0 $90 $90
rsprite
    db $E0 $90 $E0 $A0 $90
lsprite
    db $80 $80 $80 $80 $F0
osprite
    db $F0 $90 $90 $90 $F0
ssprite
    db $F0 $80 $F0 $10 $F0

org $0

start
    LD    I, #csprite    ;
    LD    V0, 0          ; X = 0
    LD    V1, 13         ; Y = 13
    LD    V2, 28         ; Limite: 64 - (6 * 6) = 28
    LD    V3, 6          ;
    LD    V4, 30         ;
    LD    V5, 1          ;
    CLS                  ;

loop
    SE    V0, V2
    CLS

    ; C
    LD    I, #csprite
    DRW   V0, V1, 5

    ; A
    LD    I, #asprite
    ADD   V0, V3
    DRW   V0, V1, 5

    ; R
    LD    I, #rsprite
    ADD   V0, V3
    DRW   V0, V1, 5

    ; L
    LD    I, #lsprite
    ADD   V0, V3
    DRW   V0, V1, 5

    ; O
    LD    I, #osprite
    ADD   V0, V3
    DRW   V0, V1, 5

    ; S
    LD    I, #ssprite
    ADD   V0, V3
    DRW   V0, V1, 5


    SUB   V0, V4         ; Volta V0 para a posição inicial

    SE    V0, V2         ; Se chegou no limite, pula o ADD
    ADD   V0, V5

    JP    #loop

end
    JP    #end