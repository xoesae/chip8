# Opcodes

| Instruction        |                       | Opcode | OK  |
|--------------------|-----------------------|--------|-----|
| CLS                | CLEAR                 | 00E0   | [X] |
| RET                | Return                | 00EE   | [ ] |
| JP addr            | JUMP addr             | 1NNN   | [X] |
| CALL addr          | CALL                  | 2NNN   | [ ] |
| SE Vx, byte        | Vx == byte            | 3XNN   | [X] |
| SNE Vx, byte       | Vx != byte            | 4XNN   | [X] |
| SE Vx, Vy          | Vx == Vy              | 5XY0   | [X] |
| LD Vx, byte        | Vx = byte             | 6XNN   | [X] |
| ADD Vx, byte       | Vx += byte            | 7XNN   | [X] |
| LD Vx, Vy          | Vx = Vy               | 8XY0   | [X] |
| OR Vx, Vy          | Vx OR Vy              | 8XY1   | [X] |
| AND Vx, Vy         | Vx &= Vy              | 8XY2   | [X] |
| XOR Vx, Vy         | Vx ^= Vy              | 8XY3   | [X] |
| ADD Vx, Vy         | Vx += Vy              | 8XY4   | [X] |
| SUB Vx, Vy         | Vx -= Vy              | 8XY5   | [X] |
| SHR Vx {, Vy}      | Vx >>= 1              | 8XY6   | [X] |
| SUBN Vx, Vy        | Vx = Vy - Vx          | 8XY7   | [X] |
| SHL Vx {, Vy}      | Vx <<= 1              | 8XYE   | [X] |
| SNE Vx, Vy         | Vx != Vy              | 9XY0   | [X] |
| LD I, addr         | I = addr              | ANNN   | [X] |
| JP V0, addr        | Jump offset           | BNNN   | [ ] |
| RND Vx, byte       |                       | CXNN   | [ ] |
| DRW Vx, Vy, nibble | DRAW                  | DXYN   | [X] |
| SKP Vx             | Is a key not pressed? | EX9E   | [ ] |
| SKNP Vx            | Is a key pressed?     | EXA1   | [ ] |
| LD Vx, DT          | delay → Vx            | FX07   | [X] |
| LD Vx, K           | wait for a keypress   | FX0A   | [X] |
| LD DT, Vx          | Vx → delay            | FX15   | [X] |
| LD ST, Vx          | Vx → sound tmr        | FX18   | [X] |
| ADD I, Vx          | I += Vx               | FX1E   | [X] |
| LD F, Vx           | sprite addr→I         | FX29   | [X] |
| LD B, Vx           | BCD → mem             | FX33   | [X] |
| LD [I], Vx         | regs → mem            | FX55   | [X] |
| LD Vx, [I]         | mem → regs            | FX65   | [X] |

# DRAW

Hex is converted to binary and the 1's is a ON pixels, and 0's is OFF pixels

| Hex | Binary   | Pixels   |
|-----|----------|----------|
| F0  | 11110000 | ####.... |
| 90  | 10010000 | #..#.... |
| 90  | 10010000 | #..#.... |
| 90  | 10010000 | #..#.... |
| F0  | 11110000 | ####.... |

# Keypad

╔═══╦═══╦═══╦═══╗
║ 1 ║ 2 ║ 3 ║ C ║
╠═══╬═══╬═══╬═══╣
║ 4 ║ 5 ║ 6 ║ D ║
╠═══╬═══╬═══╬═══╣
║ 7 ║ 8 ║ 9 ║ E ║
╠═══╬═══╬═══╬═══╣
║ A ║ 0 ║ B ║ F ║
╚═══╩═══╩═══╩═══╝

╔═══╦═══╦═══╦═══╗
║ 1 ║ 2 ║ 3 ║ 4 ║
╠═══╬═══╬═══╬═══╣
║ q ║ w ║ e ║ r ║
╠═══╬═══╬═══╬═══╣
║ a ║ s ║ d ║ f ║
╠═══╬═══╬═══╬═══╣
║ z ║ x ║ c ║ v ║
╚═══╩═══╩═══╩═══╝


# References

- https://github.com/mattmikolay/chip-8/wiki/Mastering-CHIP%E2%80%908
- https://github.com/stianeklund/chip8