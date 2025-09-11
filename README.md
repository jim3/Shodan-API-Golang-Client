# Shodan API Golang Client

## Description
An exercise in order to learn more about using the Shodan API (making calls, parsing JSON responses, etc.), Golang structs & methods, the `flag` package for command-line arguments, and parsing JSON using Go.

## Installation & Usage
```bash
git clone https://github.com/jim3/Shodan-API-Golang-Client.git
cd Shodan-API-Golang-Client
go run . -iplookup "8.8.8.8" # Host information
go run . -search "product:Apache" # Shodan search example
```

### Pre-requisites
- A Shodan account and API key.
- Go installed on your machine. [https://golang.org/dl/](https://golang.org/dl/).
- Set your Shodan API key as an environment variable inside of your `.bashrc` or `.zshrc` file.
- 
```bash
 export SHODAN_API_KEY="MY_API_KEY"
```
