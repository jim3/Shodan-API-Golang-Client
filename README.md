An exercise in order to learn more about Go, the Shodan API and it's CVEDB, using the `flag` package to create a simple CLI application, and have fun.

## Installation & Usage
```bash
git clone https://github.com/jim3/Shodan-API-Golang-Client.git
cd Shodan-API-Golang-Client
go run . -iplookup "8.8.8.8" # Host information
go run . -search "product:Apache" # Shodan search example
```

### Pre-requisites
- A Shodan account and API key.
- Go installed on your machine. [https://golang.org/dl](https://golang.org/dl)
- Set your Shodan API key as an environment variable inside of your `.bashrc` or `.zshrc` file.

```bash
 export SHODAN_API_KEY="MY_API_KEY"
```

### Flags
- `-iplookup` : Lookup information about a specific IP address.
- `-search` : Search Shodan for a specific query.
- `-cve` : Lookup information about a specific CVE.

### Example Output:
```bash
go run . -iplookup "8.8.8.8"
IP              City                 Ports
8.8.8.8         Mountain View        443, 53

go run . -iplookup "45.79.192.191"
IP              City                 Ports
45.79.192.191   Atlanta              80, 8080, 22
```
