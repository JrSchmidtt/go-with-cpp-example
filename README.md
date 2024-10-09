# golang_c

## Requirements
- Go 1.22 or higher
- C++ compiler (g++)
- Task runner (optional but recommended)

## Installation Requirements on Windows
#### Run Windows PowerShell as Administrator

Install golang using chocolatey package manager.
```
choco install golang
```
Check the version of golang installed.
```
go version
```

Install g++ compiler using chocolatey package manager.
```
choco install mingw
```
Check the version of g++ installed.
```
g++ --version
```

Install task runner using chocolatey package manager.
```
choco install go-task
```
Check the version of task runner installed.
```
task --version
```

## Running the code

To run the code, use the following command in the terminal.
```
task run
```

To only compile the code, use the following command in the terminal.
```
task build
```

To build and run the code, use the following command in the terminal.
```
task build-and-run
```

## Warning 
run using task runner to avoid copilation errors in c++ files.