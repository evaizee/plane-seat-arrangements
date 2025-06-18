#!/bin/sh
set -e

echo "Running database migrations..."
echo "Current directory: $(pwd)"
echo "Listing /app directory:"
ls -la /app

echo "Checking for migrations directory:"
ls -la /app/migrations || echo "Migrations directory not found!"

echo "Environment variables:"
echo "CONFIG_PATH=$CONFIG_PATH"
echo "MIGRATIONS_PATH=$MIGRATIONS_PATH"
echo "DATABASE_HOST=$DATABASE_HOST"
echo "DATABASE_PORT=$DATABASE_PORT"
echo "DATABASE_USER=$DATABASE_USER"
echo "DATABASE_NAME=$DATABASE_NAME"
echo "DATABASE_SSLMODE=$DATABASE_SSLMODE"

# Test database connection
echo "Testing database connection..."
APK_ADD=$(which apk 2>/dev/null)
if [ -n "$APK_ADD" ]; then
  apk --no-cache add postgresql-client
fi

PGPASSWORD=$DATABASE_PASSWORD psql -h $DATABASE_HOST -p $DATABASE_PORT -U $DATABASE_USER -d $DATABASE_NAME -c "SELECT 1;" || echo "Database connection failed!"

# Check if we need to wait for the database to be ready
if [ "$WAIT_FOR_DB" = "true" ]; then
  echo "Waiting for database to be ready..."
  # Simple retry mechanism
  for i in $(seq 1 4); do
    echo "Attempt $i of 4..."
    cd /app && /app/migrate up -verbose && break
    echo "Migration attempt $i failed, error code: $?"
    echo "Database connection details: $DATABASE_USER@$DATABASE_HOST:$DATABASE_PORT/$DATABASE_NAME"
    echo "Retrying in 2 seconds..."
    sleep 2
  done
else
  # Run migrations directly
  cd /app && /app/migrate up -verbose
fi

echo "Migrations completed successfully!"
