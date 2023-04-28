#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings
# ./bookings -dbname=bookings -dbuser=dbmaser -cache=false -production=false