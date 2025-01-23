#!/bin/bash

# Check if Docker is installed
if ! command -v docker &> /dev/null
then
    echo "Docker is not installed. Please install Docker first."
    exit 1
fi

# Pull the latest PostgreSQL Docker image
echo "Pulling the latest PostgreSQL Docker image..."
docker pull postgres:latest

# Run the PostgreSQL container with the specified configuration
echo "Starting the PostgreSQL container..."

docker run -d \
  --name depixen-postgres \
  -e POSTGRES_PASSWORD=depixen-pass \
  -e POSTGRES_DB=postgres \
  -p 5439:5432 \
  -v depixen-volume:/var/lib/postgresql/data \
  postgres:latest

# Wait for the database to be ready
echo "Waiting for PostgreSQL to start..."
sleep 15

# Create the table in the database
echo "Creating the table 'tb_casestudy'..."

# Execute SQL commands to create the table
docker exec -it depixen-postgres psql -U postgres -d postgres -c "
CREATE TABLE IF NOT EXISTS tb_casestudy (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    imageUrl VARCHAR(255),
    creationtime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
"

# Confirm the table creation
echo "Table 'tb_casestudy' created successfully!"

# End of script
