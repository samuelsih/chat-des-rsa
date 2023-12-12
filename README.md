# chat-des-rsa

### How To Run (With Taskfile)

1. Make sure you have Makefile
2. For server, run `make serverx`
3. For client, run `make clientx`

### How to Run (Without Taskfile)

Note: **Replace 12345678 with your 64bit key**

1. For server, run `DES_KEY=12345678 go run server/main.go server/tools.go`
2. For client, run `DES_KEY=12345678 go run client/main.go`