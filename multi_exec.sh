#!/bin/bash

# Fixed list of directories (modify these as needed)
# Supports directory names with spaces (enclose in quotes)
dirs=(
    "go"
    "packages/core/go"
    "packages/db/go"
    "packages/document/go"
    "packages/geom/go"
    #"packages/graph/go"
    "packages/http/go"
    "packages/lang/go"
    #"packages/net/go"
    "packages/openapi/go"
    "packages/protobuf/go"
    "packages/rpc/go"
    "packages/yaml/go"
)

# Check if a command to execute was provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 command [arguments]"
    echo "Example: $0 ls -l"
    echo "Example: $0 grep 'pattern' *.txt"
    echo "Example: $0 git pull origin main"
    echo "Example: $0 find . -name '*.log' -delete"
    exit 1
fi

# Command to execute (takes all input arguments)
command_args="$@"

# Iterate through all fixed directories and execute the command
for dir in "${dirs[@]}"; do
    # Check if the directory exists
    if [ ! -d "$dir" ]; then
        echo "Warning: Directory '$dir' does not exist, skipped"
        continue
    fi

    # Display execution information
    echo "====================================="
    echo "Current directory: $dir"
    echo "Executing command: $command_args"

    # Execute command in target directory (runs in subshell, doesn't affect script's working directory)
    (
        cd "$dir" || exit  # Exit subshell if cd fails
        eval "$command_args"  # Execute command (supports pipes, redirections, etc.)
    )

    # Display command execution status
    exit_code=$?
    if [ $exit_code -eq 0 ]; then
        echo "-------------------------------------"
        echo "Command executed successfully in '$dir'"
    else
        echo "-------------------------------------"
        echo "Command failed in '$dir' with exit code: $exit_code"
    fi
done

echo "====================================="
echo "Command execution completed for all directories"