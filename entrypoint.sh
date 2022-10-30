wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

#watch for .go file change
CompileDeamon --build="go build -o main main.go" --command=./main