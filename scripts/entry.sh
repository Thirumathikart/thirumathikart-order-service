#!/bin/sh

until nc -z -v -w30 ordersdb 8003
do
  echo "Waiting for database connection..."
  # wait for 5 seconds before check again
  sleep 5
done

echo -e "\e[34m >>> Starting the server \e[97m"
$1
