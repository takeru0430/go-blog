#!/bin/bash

go build -ldflags "-s -w" -o public/app.cgi app.go
