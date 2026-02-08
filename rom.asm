org $0

start


loop
    LD    V0, K      ; Espera uma tecla e armazena em V0
    JP    #loop      ; Volta para esperar outra tecla

end
    JP    #end

; F00A 1000 1004