# go-i2p-modproxy

A simple Go modules proxy with support for working in anonymous networks.

## Overview

go-i2p-modproxy acts as an intermediary for Go module requests, allowing developers to fetch Go modules through anonymous networks. It listens for incoming requests via I2P and routes outgoing requests appropriately based on the target.

## Network Operations

- **Incoming Connections**: Listens on I2P for incoming module requests
- **Outgoing Connections**:
  - Routes requests to I2P services through the I2P network
  - Routes requests to Tor onion services through the Tor network
  - Routes requests to clearnet (regular internet) through Tor for anonymity

## Prerequisites

- Go 1.19 or later
- Running I2P router
- Running Tor service

## Installation

```bash
git clone https://github.com/yourusername/go-i2p-modproxy.git
cd go-i2p-modproxy
go build
```

## Usage

1. Start your I2P router and Tor service
2. Run the proxy:
   ```bash
   ./go-i2p-modproxy
   ```
3. Configure Go to use your proxy by setting `GOPROXY` environment variable to your I2P address
4. **Important**: Set HTTP_PROXY and HTTPS_PROXY environment variables to access the I2P address:
   ```bash
   export HTTP_PROXY=http://127.0.0.1:4444
   export HTTPS_PROXY=http://127.0.0.1:4444
   ```
5. Use `go get` and other Go commands as normal

## Example

```bash
# Set the GOPROXY to your I2P address
export GOPROXY=http://your-i2p-address.b32.i2p

# Set proxies to access I2P network
export HTTP_PROXY=http://127.0.0.1:4444
export HTTPS_PROXY=http://127.0.0.1:4444

# Now you can use go commands
go get github.com/example/module
```

## Security Considerations

This proxy routes traffic through anonymous networks but does not guarantee complete anonymity. Properly configured I2P and Tor services are required for enhanced privacy.

## License

[MIT License](LICENSE)