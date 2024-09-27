# Backend

## Prerequisites
- [Go 1.23.1](https://go.dev/dl/)
- [sqlite3](https://www.sqlite.org/download.html)
- [make](https://gnuwin32.sourceforge.net/packages/make.htm) (optional)

## Getting Started
### Running Without Make
```bash
go mod tidy
go run cmd/api
```
Flags can be passed to the `go run cmd/api` command. For example, to run the server on port 8080:
```bash
go run cmd/api -port 8080
```
For all available flags, run `go run cmd/api -help`.

### Running With Make
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
