# Hexagonal Arch app with Golang

## Setup

### Add modules

```
go mod init github.com/kaliayev-proton/banking-go-hex

```

### Build

```
go build

```

### Run the app:

```
go run main.go
```

Request to API:

```
curl http://localhost:8000/greet
```

### Inject variables in a Golang app:

```
BANKING_DB_ADDRESS=localhost BANKING_DB_PORT=33066 BANKING_DB_NAME=banking BANKING_DB_USER=root BANKING_DB_PASSWORD=root BANKING_SERVER_ADDRESS=localhost BANKING_SERVER_PORT=8000 go run main.go
```
