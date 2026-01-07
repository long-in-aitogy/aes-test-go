# Makefile for AES-128 Server
PROGNAME = client
SRC = main.go aes128.go

all: build

build:
	go build -o $(CURDIR)/build/$(PROGNAME).exe $(SRC)

clean:
	powershell -c "Remove-Item -Path "$(CURDIR)/build" -Recurse"