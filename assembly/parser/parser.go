package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parser struct{}

func (p *Parser) ParseFile(filename, output string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var bin []byte
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		opcode := p.parseLine(line)

		high := byte(opcode >> 8)
		low := byte(opcode & 0xFF)

		bin = append(bin, high, low)
	}

	err = os.WriteFile(output, bin, 0644)
	if err != nil {
		panic(err)
	}
}

func (p *Parser) parseLine(line string) uint16 {
	parts := strings.Fields(line)

	switch parts[0] {
	case "CLS":
		return 0x00E0
	case "RET":
		return 0x00EE
	case "LD":
		x := parts[1][1] - '0'
		if strings.HasPrefix(parts[2], "V") {
			y := parts[2][1] - '0'
			return 0x8000 | (uint16(x) << 8) | uint16(y<<4) | 0x0
		} else {
			val, _ := strconv.Atoi(parts[2])
			return 0x6000 | (uint16(x) << 8) | uint16(val)
		}
	case "ADD":
		x := parts[1][1] - '0'
		if strings.HasPrefix(parts[2], "V") {
			y := parts[2][1] - '0'
			return 0x8004 | (uint16(x) << 8) | uint16(y<<4)
		} else {
			val, _ := strconv.Atoi(parts[2])
			return 0x7000 | (uint16(x) << 8) | uint16(val)
		}
	case "JP":
		addr, _ := strconv.ParseUint(parts[1], 16, 16)
		return 0x1000 | uint16(addr)
	case "CALL":
		addr, _ := strconv.ParseUint(parts[1], 16, 16)
		return 0x2000 | uint16(addr)
	case "SE":
		x := parts[1][1] - '0'
		if strings.HasPrefix(parts[2], "V") {
			y := parts[2][1] - '0'
			return 0x5000 | (uint16(x) << 8) | uint16(y<<4)
		} else {
			val, _ := strconv.Atoi(parts[2])
			return 0x3000 | (uint16(x) << 8) | uint16(val)
		}
	case "SAVE":
		x := parts[1][1] - '0'
		return 0xF055 | (uint16(x) << 8)
	case "LOAD":
		x := parts[1][1] - '0'
		return 0xF065 | (uint16(x) << 8)
	default:
		panic("error")
	}
}
