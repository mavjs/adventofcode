#!/bin/bash

# Check if an argument (year) has been provided
if [ -z "$1" ]; then
    echo "Please provide the year as an argument."
    exit 1
fi

# Create the initial folder based on the argument given
mkdir "$1"

# Change directory to the newly created year folder
cd "$1" || exit

# Loop from 01 to 25 and create sub folders for each day of Advent of Code challenges
for i in {1..25}; do
    # Format the day number with leading zero if necessary
    day=$(printf "%02d" "$i")
    
    mkdir "day$day"
    
    # Change directory into each day's folder using a subshell to avoid changing back
    (
        cd "day$day" || exit
        
        # Create subfolders for Python, Rust, and Go
        mkdir py rs go
    )
done

echo "Folders have been successfully created!"
