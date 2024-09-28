# Backend

## Prerequisites
- [Go 1.23.1](https://go.dev/dl/)
- [sqlite3](https://www.sqlite.org/download.html)

## Getting Started
### Running Without Make
```bash
go mod tidy
go run ./cmd/api
```
Flags can be passed to the `go run ./cmd/api` command. For example, to run the server on port 8080:
```bash
go run ./cmd/api -port 8080
```
For all available flags, run `go run ./cmd/api -help`.

### Running With Make (UNIX only)
```bash
make tidy
make run
```
Flags can be managed in a `.envrc` file. For example, to run the server on port 8080, create a `.envrc` file with the following content:
```bash
export API_PORT==8080
```
All available flags can be found in the `Makefile`.

## Other Commands
All available commands can be found in the `Makefile`. To see all available commands, run `make help`.

## Exposed Endpoints
| Method | Path | Description |
| --- | --- | --- |
| GET | /v1/healthcheck | Healthcheck endpoint |
| GET | /v1/results | Get all results |
| GET | /v1/results/{id} | Get a result by ID |
| POST | /v1/results | Create a new result |

### Examples
Add a new result:
```bash
curl -X POST http://localhost:8080/v1/results -d '{"name": "John Doe", "score": 100}'
# Output: 
# {"result":{"id":1,"name":"John Doe","score":100}}
```
Get all results:
```bash
curl http://localhost:8080/v1/results
# Output:
# {"results":[{"id":2,"name":"John Doe","score":200},{"id":1,"name":"John Doe","score":100}]}
```
