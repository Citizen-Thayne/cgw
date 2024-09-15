all: board main

board:
	go build pkg/board/board.go

main:
	go build main.go
