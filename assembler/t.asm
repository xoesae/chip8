org $300

; Sprite do quadrado 5x5:
square
    db $F8 $F8 $F8 $F8 $F8

org $0

start
    LD    I, #square      ; I aponta para o sprite do quadrado
    LD    V0, 0           ; X inicial (coluna)
    LD    V1, 16          ; Y inicial (linha)
    LD    VF, 100           ; delay

loop
    CLS                  ; Limpa a tela
    DRW   V0, V1, 5      ; Desenha o quadrado 5 bytes (5x5)
    ADD   V0, 2          ; Move 2 pixels para a direita
    SE    V0, 62         ; Chegou ao lado direito da tela?
    LD    V0, 0          ; Se sim, volta para a esquerda

    LD    DT, VF          ; Pequena pausa

wait_timer
    LD    V2, DT
    SE    V2, 0
    JP    #wait_timer

    JP    #loop          ; Repete o loop

end
    JP    #end           ; Fim


