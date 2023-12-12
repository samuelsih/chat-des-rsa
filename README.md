# chat-des-rsa

### How To Run (With Taskfile)

0. Create .env with key `DES_KEY=...`
1. Make sure you have Taskfile installed
2. For server, run `task server`
3. For client, run `task client`

### How to Run (Without Taskfile)

Note: **Replace 12345678 with your 64bit key**

1. For server, run `go run server/main.go server/tools.go`
2. For client, run `DES_KEY=12345678 go run client/main.go`