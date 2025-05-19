#!/bin/sh

until nc -z -v -w30 $DB_HOST $DB_PORT
do
  echo "⏳ Aguardando o MySQL ($DB_HOST:$DB_PORT)..."
  sleep 2
done

echo "✅ MySQL está pronto"