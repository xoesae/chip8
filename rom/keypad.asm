org $300
sprite
        db $F0 $90 $90 $90 $F0

org $0

start
    CLS

    LD V1, 30 ; X
    LD V2, 15 ; Y

loop
    CLS

    LD V0, K ; waiting for a keypress

    SNE V0, 7   ; execute the next instruction if V0 = 7
    ADD V1, 255 ;

    SNE V0, 9 ; execute the next instruction if V0 = 9
    ADD V1, 1 ;

    LD   I, #sprite ; load sprite addr
    DRW  V1, V2, 1  ;

    JP   #loop