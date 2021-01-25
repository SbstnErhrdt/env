# Env - Work in progress

This is a small go module that helps with environment variables.

## Installation
```
go get github.com/SbstnErhrdt/env
```

## Usage

### Load env files
Loads different env files. Prints warnings if the files not present. 
```go
env.LoadEnvFiles(filenames ...string)
```

### Fallback environment variables
Sets fallback variables for env variables. 
```go
env.FallbackEnvVariable("environmentVariableKey","fallbackValue")
```

### Required variables check
Checks if there are **required** environment variables not set. 
```go
env.CheckPossibleEnvironmentVariables([]string{"environmentVariableKey0","environmentVariableKey1"})
```

### Optional variables check
Checks if there are optional environment variables not set. 
```go
env.CheckPossibleEnvironmentVariables([]string{
  "environmentVariableKey0",
  "environmentVariableKey1"
})
```
