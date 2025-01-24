Steps to start the backend service

After downloading the zip file to your computer, follow the steps below

## Sorting of folders and files
```bash
.
└── main-service
    ├── cmd
    │   ├── main.go
    │   ├── models
    │   │   └── models.go
    │   ├── service
    │   │   ├── endpoints.go
    │   │   └── service.go
    │   └── transport
    │       └── transport.go
    ├── go.mod
    ├── go.sum
    ├── Readme.md
    └── scripts
        ├── free_port_8080.sh
        └── start_postgres.sh

7 directories, 10 files
```

## Step 1: Installing Docker postgres database

```bash
sudo ./scripts/start_postgres.sh
```

Before pulling the image and running the container, the script checks if a container with the same name exists. 
If it does, it stops and removes the container and starts to make new database.


## Step 2: Checks if GO modules are loaded

```bash
go mod tidy
```


## Step 3: Free Port 8080 (Optional)
If port 8080 is being used by another process, you can run your script free_port_8080.sh to free the port. Run this command:
sudo ./scripts/free_port_8080.sh

## Step 4: Run Your Go Service
```bash
mutu@mutu:~/projects/backend/main-service$ go run cmd/main.go 
2025/01/24 09:22:27 INFO: Starting server on port :8080...
```

This will start your service on the specified port (8080 in this case), and you should see a log message 
indicating that the server is running.


