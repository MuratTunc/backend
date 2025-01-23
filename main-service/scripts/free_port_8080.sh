#!/bin/bash

# Check if port 8080 is in use
PID=$(sudo lsof -t -i :8080)

if [ -n "$PID" ]; then
    echo "Port 8080 is in use by process PID: $PID. Killing the process..."
    
    # Kill the process using the port
    sudo kill -9 $PID
    
    # Check if the process was successfully killed
    if [ $? -eq 0 ]; then
        echo "Port 8080 is now free."
    else
        echo "Failed to kill the process using port 8080."
    fi
else
    echo "Port 8080 is not in use."
fi
