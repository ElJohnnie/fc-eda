#!/bin/sh

while ! nc -z wallet-mysql 3306; do
  echo "Aguardando o MySQL do wallet..."
  sleep 2
done

exec "$@"