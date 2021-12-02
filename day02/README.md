# Elm

This project executes in the elm repl 

## How to build
```bash
docker build -t aoc01 . 
```

## How to run
The environment variable "part" specifies which part of the solution to run. Please adjust your code accordingly (see example in App.java).
```bash
docker run -e part=part1 aoc01
```