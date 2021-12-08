# Conway's Game of Life simulation
[![Build Passing](https://github.com/heartwilltell/gameoflife/actions/workflows/master.yml/badge.svg)](https://github.com/heartwilltell/gameoflife/actions/workflows/master.yml)

This repo contains implementation of Conway's Game of Life simulator.

## Build

**Prerequisite:**
Tho Go programming language < v1.16 should be installed on the system.

To build the game executable run the following command:

```shell
go build -o game ./cmd
```

## Test
**Prerequisite:**
Tho Go programming language < v1.16 should be installed on the system.

```shell
go test -race -cover ./...
```

## Lint

**Prerequisite:**
[golangci-lint](https://golangci-lint.run) should be installed on the system.

```shell
golangci-lint run ./...
```

## Run

To run the game simulation run the following command:

```shell
./game run
```

The game will print generations to the stdout.