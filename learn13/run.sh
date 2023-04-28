#!/bin/bash

go build -o bookings cmd/web/*.go
# ./bookings
./bookings -dbname=bookings -dbuser=dbmaster -cache=false -production=false