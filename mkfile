all:V:	y.go
	go test -v

y.go:	dbg.y
	go tool yacc $prereq

clean:V:
	rm -f y.go y.output

