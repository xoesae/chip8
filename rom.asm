org $300
square
    db $F8 $F8 $F8 $F8 $F8

org $0

start
    LD    I, #square      ; Carrega endereço do sprite
    LD    V0, 30          ; X = 30
    LD    V1, 13          ; Y = 13
    CLS                   ; Limpa tela

loop
    DRW   V0, V1, 5       ; Desenha quadrado 5x5 pixels
    JP    #loop           ; Loop infinito

end
    JP    #end