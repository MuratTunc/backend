Step 1: Ensure PostgreSQL is Running
sudo ./scripts/start_postgres.sh
Before pulling the image and running the container, the script checks if a container with the same name exists. 
If it does, it stops and removes the container and starts to make new database.


Step 2: Make Sure Your Go Modules Are Installed
go mod tidy


Step 3: Free Port 8080 (Optional)
If port 8080 is being used by another process, you can run your script free_port_8080.sh to free the port. Run this command:
sudo ./scripts/free_port_8080.sh

Step 4: Run Your Go Service
go run cmd/main.go

This will start your service on the specified port (8080 in this case), and you should see a log message 
indicating that the server is running.


