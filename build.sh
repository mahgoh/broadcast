#! /bin/bash

go build -ldflags "-X main.Version=$(git describe --tags --dirty)" -o bin/broadcast ./cmd