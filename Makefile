all:
	go build -o beelde cmd/beelde/main.go
clean:
	rm -rf ./beelde
