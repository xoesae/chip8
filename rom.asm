org $300
square
    db $F8 $F8 $F8 $F8 $F8

org $0

start
    LD    I, #square      ; Carrega endereço do sprite
    LD    V0, 0           ; X = 0
    LD    V1, 13          ; Y = 13
    LD    V2, 59          ; Limite para o quadrado andar, 64 - 5 = 59 % 5
    CLS                   ; Limpa tela

loop
    SE    V0, V2
    CLS
    SE    V0, V2
    DRW   V0, V1, 5       ; Desenha quadrado 5x5 pixels
    SE    V0, V2          ; se chegou no limite, pula o ADD
    ADD   V0, 5
    JP    #loop           ; Loop infinito

end
    JP    #end