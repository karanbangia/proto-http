# Protocol Buffer HTTP Example

This project demonstrates a simple HTTP server that accepts JSON requests and returns Protocol Buffer responses.

## Prerequisites

1. Go 1.21 or later
2. Protocol Buffer Compiler (protoc)

### Installing Prerequisites

#### macOS

```bash
# Install Go
brew install go

# Install Protocol Buffer Compiler
brew install protobuf
```
## Setup

1. Clone the repository:

```bash
git clone https://github.com/your-username/proto-http.git
cd proto-http
```

2. Install the Protocol Buffer compiler plugin for Go:

```bash
export PATH=$PATH:$(go env GOPATH)/bin && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

```

3. Generate Protocol Buffer code:

```bash
chmod +x generate.sh
./generate.sh
```

4. Install dependencies:

```bash
go mod tidy
```

## Running the Server

Start the server:

```bash
go run main.go
```

The server will start on port 8080.

## Testing the API

### Using curl

1. Get a user (returns protobuf):

```bash
curl http://localhost:8080/user?id=1
```

2. Create a user (JSON input, protobuf output):

```bash
curl -X POST \
  http://localhost:8080/user \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "2",
    "name": "Jane Doe",
    "email": "jane@example.com"
}'
```

### Using Postman

1. GET Request:

   - URL: `http://localhost:8080/user?id=1`
   - Method: GET
   - Response will be in protobuf format

2. POST Request:
   - URL: `http://localhost:8080/user`
   - Method: POST
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "id": "2",
       "name": "Jane Doe",
       "email": "jane@example.com"
     }
     ```
   - Response will be in protobuf format

## Project Structure

- `api/user.proto`: Protocol Buffer definition
- `server/server.go`: HTTP server implementation
- `main.go`: Server startup code
- `generate.sh`: Script to generate Go code from proto files
