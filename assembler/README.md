# Chip8 Assembler

- [x] 00E0
- [x] 00EE
- [x] 1NNN
- [x] 2NNN
- [x] 3XNN
- [ ] 4XNN
- [x] 5XY0
- [x] 6XNN
- [x] 7XNN
- [x] 8XY0
- [ ] 8XY1
- [ ] 8XY2
- [ ] 8XY3
- [x] 8XY4
- [x] 8XY5
- [ ] 8XY6
- [x] 8XY7
- [ ] 8XYE
- [ ] 9XY0
- [x] ANNN
- [ ] BNNN
- [ ] CXNN
- [ ] DXYN
- [ ] EX9E
- [ ] EXA1
- [x] FX07
- [ ] FX0A
- [x] FX15
- [x] FX18
- [x] FX1E
- [x] FX29
- [ ] FX33
- [x] FX55
- [x] FX65

| Instruction        |                       | Opcode |
|--------------------|-----------------------|--------|
| CLS                | CLEAR                 | 00E0   |
| RET                | Return                | 00EE   |
| JP addr            | JUMP addr             | 1NNN   |
| CALL addr          | CALL                  | 2NNN   |
| SE Vx, byte        | Vx == byte            | 3XNN   |
| SNE Vx, byte       | Vx != byte            | 4XNN   |
| SE Vx, Vy          | Vx == Vy              | 5XY0   |
| LD Vx, byte        | Vx = byte             | 6XNN   |
| ADD Vx, byte       | Vx += byte            | 7XNN   |
| LD Vx, Vy          | Vx = Vy               | 8XY0   |
| OR Vx, Vy          | Vx                    | = Vy   |
| AND Vx, Vy         | Vx &= Vy              | 8XY2   |
| XOR Vx, Vy         | Vx ^= Vy              | 8XY3   |
| ADD Vx, Vy         | Vx += Vy              | 8XY4   |
| SUB Vx, Vy         | Vx -= Vy              | 8XY5   |
| SHR Vx {, Vy}      | Vx >>= 1              | 8XY6   |
| SUBN Vx, Vy        | Vx = Vy - Vx          | 8XY7   |
| SHL Vx {, Vy}      | Vx <<= 1              | 8XYE   |
| SNE Vx, Vy         | Vx != Vy              | 9XY0   |
| LD I, addr         | I = addr              | ANNN   |
| JP V0, addr        | Jump offset           | BNNN   |
| RND Vx, byte       |                       | CXNN   |
| DRW Vx, Vy, nibble | DRAW                  | DXYN   |
| SKP Vx             | Is a key not pressed? | EX9E   |
| SKNP Vx            | Is a key pressed?     | EXA1   |
| LD Vx, DT          | delay → Vx            | FX07   |
| LD Vx, K           | wait for a keypress   | FX0A   |
| LD DT, Vx          | Vx → delay            | FX15   |
| LD ST, Vx          | Vx → sound tmr        | FX18   |
| ADD I, Vx          | I += Vx               | FX1E   |
| LD F, Vx           | sprite addr→I         | FX29   |
| LD B, Vx           | BCD → mem             | FX33   |
| LD [I], Vx         | regs → mem            | FX55   |
| LD Vx, [I]         | mem → regs            | FX65   |

## References
- https://nnarain.github.io/2017/09/30/Chip8-Assembler-Introduction.html
- https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go