#!/bin/sh
go test -v -coverprofile _cover.out .
go tool cover -html=_cover.out -o _cover.html
open _cover.html
