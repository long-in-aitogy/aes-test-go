# Makefile for AES-128 Server
PROGNAME = server
SRC = main.go aes128_decrypt.go

all: build

build:
	go build -o build/$(PROGNAME).exe $(SRC)

clean:
	Remove-Item -Path -Recurse build