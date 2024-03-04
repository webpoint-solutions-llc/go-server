# Go Server Boilerplate

Go server boilerplate with PostgreSQL

## Getting Started

> NOTE:
>
> We use Makefile to run all our scripts so install it first if you don't have it

### Installing Make

- On Windows
```sh
winget install GnuWin32.Make
```

- On Mac
```sh
brew install make
```

- On Linux

use your distro's package manager to install

### Initializing the project

1. Clone the repository

2. Go inside the directory
```sh
cd <dir-name>
```
3. Initialize go module
```sh
go mod init <your-project-name>
```

4. Organize imports
```sh
make init
```

## Running the server

- Start the server in Devlopment mode
```sh
make watch
```

- Start the server
```sh
make run
```

- Build the application
```sh
make build
```

## Database

- Install tern
```sh
go install github.com/jackc/tern/v2@latest
```

- Make new migration

```sh
tern new -m internal/database/sql/migrations/ name_of_migration
```

- Run the migration

```sh
make migrate
```

- Write sql queries insite `internal/database/sql/migrations/` direcotry.

to reference how to properly write the query go [here](https://docs.sqlc.dev/en/latest/howto/select.html)

- Generate fully type-safe idiomatic Go code from SQL.

```sh
make generate
```
