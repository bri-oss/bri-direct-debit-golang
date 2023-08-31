![BRI API](./images/logo.svg)

# BRI Inquiry API in Golang

Welcome to the BRI Inquiry API workspace.
This workspace contains the Golang implementation of the BRI Inquiry API.

## Quick Start

1. Run the application using `go run main.go` from the terminal
2. Or, run the application using VSCode's debugger feature

## Configuration

The workspace is configured to use the sandbox environment by default. To change the environment, set following environment variables:

| Variable          | Description                   |
| ----------------- | ----------------------------- |
| BRI_CLIENT_ID     | Client ID provided by BRI     |
| BRI_CLIENT_SECRET | Client Secret provided by BRI |

To check pre-configured environment variables, run the following command:

```bash
printenv | grep BRI
```

## Documentation

For more information, please refer to the [BRI Inquiry API documentation](https://developers.bri.co.id/id/docs/informasi-rekening).
