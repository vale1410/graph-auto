
all: 
	gcc -o graph graph.c
	go build convert.go

clean: 
	rm -fr convert
	rm -fr graph

