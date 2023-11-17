# Usage

```sh
# Build the binary
go build cmd/overlap.go

# Usage: ./overlap <ip> <ip>
# <ip> format must be in CIDR format. E.g: 192.168.1.1/24

# Execute it
./overlap 192.168.1.1/24 192.168.1.10/24
same
```