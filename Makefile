all: bench

bench:
	go test -run=XX -bench=. -benchtime=10s
