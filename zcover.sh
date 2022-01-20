#!/bin/sh
mkdir -p ./_cover
go test -v -coverprofile ./_cover/cover.out .
go tool cover -html=./_cover/cover.out -o ./_cover/cover.html
open ./_cover/cover.html
