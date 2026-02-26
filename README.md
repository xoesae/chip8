# Opcodes
 
- [x] 00E0 
- [ ] 00EE
- [x] 1NNN
- [ ] 2NNN
- [x] 3XNN
- [x] 4XNN 
- [x] 5XY0 
- [x] 6XNN 
- [x] 7XNN 
- [x] 8XY0 
- [x] 8XY1 
- [x] 8XY2 
- [x] 8XY3 
- [x] 8XY4 
- [x] 8XY5 
- [x] 8XY6 
- [x] 8XY7 
- [x] 8XYE 
- [ ] 9XY0 
- [x] ANNN 
- [ ] BNNN 
- [ ] CXNN 
- [x] DXYN 
- [ ] EX9E 
- [ ] EXA1 
- [ ] FX07 
- [x] FX0A 
- [ ] FX15 
- [ ] FX18 
- [ ] FX1E 
- [x] FX29 
- [ ] FX33 
- [x] FX55 
- [x] FX65 

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