build:
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o main.out main.go
	go tool objdump -s "main.main" main.out > main.S
	go tool compile -S main.go
	go build -gcflags -S main.go