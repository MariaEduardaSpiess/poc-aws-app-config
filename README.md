# POC AWS AppConfig

This project uses gin framework to create an API with the endpoint `/<cpf>`. This endpoint receives a CPF, goes to AWS AppConfig to get the feature flags, validates if the flag is enabled and if the cpf is valid against the regex.

## How to run

Create an `.env` file with the environment variables `AWS_ID` and `AWS_SECRET` and run the application with `go run main.go` on the terminal.

### Requirements

Go 1.19
AWS AppConfig application created

### AppConfig Environment

Application: `Teste`  
Environment: `Dev`  
Feature Flag: `Teste Feature Flags`  
Flag: `feature_xpto`  
Flag Attribute: `regex_cpfs_rollout`  

## Test

With the application running, call the endpoint with the cURL:
```

```