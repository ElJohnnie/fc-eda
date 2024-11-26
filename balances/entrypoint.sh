#!/bin/sh

while ! nc -z balances-mysql 3306; do
  echo "Aguardando o MySQL do balances..."
  sleep 2
done

exec "$@"