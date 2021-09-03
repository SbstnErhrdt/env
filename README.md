# Env

This is a small go module that helps with environment variables.

* Load environment files `.env`
* Provides fallback variables

## Installation
```
go get github.com/SbstnErhrdt/env
```

## Usage

Put a `.env` file in the working directory.
This might look like this
```
KEY=value1234
```

In your `main()` funcion:
```go
func main() {
	// load end
	env.LoadEnvFiles() // reads the .env file in the working directory
  ...
}
```
If no filename is provided, the file `.env` in the working directory is used.



### Load env files
Loads different env files. Prints warnings if the files not present. 
```go
...
env.LoadEnvFiles(filenames ...string)
...
```

Specify filename:
```go
...
env.LoadEnvFiles("production.enbv)
...
```

### Fallback environment variables
Sets fallback variables for env variables. 
```go
...
env.FallbackEnvVariable("environmentVariableKey","fallbackValue")
...
```

### Required variables check
Checks if there are **required** environment variables not set. 
```go
...
env.CheckRequiredEnvironmentVariables("environmentVariableKey0","environmentVariableKey1")
...
```

### Optional variables check
Checks if there are optional environment variables not set. 
```go
...
env.CheckOptionalEnvironmentVariables("environmentVariableKey0","environmentVariableKey1")
...
```
