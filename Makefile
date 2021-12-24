input: 
	vasm65_std -Fbin misc/asm/input.asm -o input ; hexdump  -C input

install:
	go install golang.org/x/tools/cmd/stringer@latest

generate:
	go generate ./...

build: generate
	go build -o bin/6502 cmd/6502.go

run: generate input
	go run cmd/6502.go
