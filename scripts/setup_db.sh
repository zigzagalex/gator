#!/bin/bash

set -e

echo "🔍 Checking Postgres installation..."

pg_version=$(psql --version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+' || echo "")

if [ -z "$pg_version" ]; then
  echo "📦 Installing Postgres 15..."
  brew install postgresql@15
  brew services start postgresql@15
elif [[ $(echo "$pg_version < 15" | bc) -eq 1 ]]; then
  echo "⚠️  Postgres version too old: $pg_version"
  echo "📦 Installing Postgres 15..."
  brew install postgresql@15
  brew services start postgresql@15
else
  echo "✅ Postgres $pg_version installed"
fi

echo "🧪 Checking if 'gator' database exists..."
if ! psql -lqt | cut -d \| -f 1 | grep -qw gator; then
  echo "📦 Creating 'gator' database..."
  createdb gator
else
  echo "✅ Database 'gator' already exists"
fi

echo "📈 Running migrations from gator/sql/schema..."
goose -dir ./gator/sql/schema postgres "postgres://localhost/gator?sslmode=disable" up

echo "✅ Done. Database is ready."
