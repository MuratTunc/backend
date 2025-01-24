#!/bin/bash

# Check if Docker is installed
if ! command -v docker &> /dev/null
then
    echo "Docker is not installed. Please install Docker first."
    exit 1
fi

# Define container name and database credentials
CONTAINER_NAME="depixen-postgres"
DB_PASSWORD="depixen-pass"
DB_NAME="postgres"
DB_PORT="5439"

# Check if the container name is already in use
if [ "$(docker ps -aq -f name=${CONTAINER_NAME})" ]; then
    echo "Container name '${CONTAINER_NAME}' is already in use."
    echo "Stopping and removing the existing container..."
    docker stop "${CONTAINER_NAME}"
    docker rm "${CONTAINER_NAME}"
fi

# Pull the latest PostgreSQL Docker image
echo "Pulling the latest PostgreSQL Docker image..."
docker pull postgres:latest

# Run the PostgreSQL container with the specified configuration
echo "Starting the PostgreSQL container..."
docker run -d \
  --name "${CONTAINER_NAME}" \
  -e POSTGRES_PASSWORD="${DB_PASSWORD}" \
  -e POSTGRES_DB="${DB_NAME}" \
  -p "${DB_PORT}":5432 \
  -v depixen-volume:/var/lib/postgresql/data \
  postgres:latest

# Wait for the database to be ready
echo "Waiting for PostgreSQL to start..."
sleep 15

# Create the table in the database
echo "Creating the table 'tb_casestudy'..."

# Execute SQL commands to create the table
docker exec -it "${CONTAINER_NAME}" psql -U postgres -d "${DB_NAME}" -c "
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
