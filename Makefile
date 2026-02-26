.PHONY: compile emulate all test

emulate:
	go run emulator/main.go rom/keypad.ch8

compile:
	go run assembler/main.go rom/keypad.asm rom/keypad.ch8

all: compile emulate

test:
	go test ./...