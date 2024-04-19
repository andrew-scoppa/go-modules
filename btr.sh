# This bash script is used to automate the following steps for a Go program:
# Steps:
# 1. Build the program. If the build is unsuccessful, then the script will stop and show an error message.
# 2. Run the test cases in all folders and files with the suffix "_test.go". If the any test case is unsuccessful, then the script will stop and show an error message.
# 3. Run the program.

# Usage:
# ./btr.sh


# Build the program.
go build
if [ $? -ne 0 ]; then
    echo "The program did not build successfully."
    exit 1
fi

# Run the test cases in all folders and files with the suffix "_test.go"
go test -v ./...
if [ $? -ne 0 ]; then
    echo "The test cases did not run successfully."
    exit 1
fi

# Run the program.
./$(basename $(pwd))




