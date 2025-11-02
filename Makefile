compile:
	go run assembler/main.go rom.asm rom.ch8

emulate:
	go run emulator/main.go rom.ch8

test:
	go test ./...