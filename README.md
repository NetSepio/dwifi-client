# Erebrus DWiFi Client

## Description

Erebrus DWiFi Client is a Go-based application that allows users to connect to decentralized WiFi networks. It scans for nearby WiFi networks, fetches information about Erebrus-registered networks from an API, and enables users to connect to these networks with a pay-per-minute model using cryptocurrency.

## Features

- Scan for nearby WiFi networks
- Fetch and display Erebrus-registered DWiFi networks
- Connect to selected DWiFi networks
- Real-time connection monitoring
- Automatic disconnection and payment processing

## Prerequisites

- Go 1.15 or higher
- Windows operating system (for WiFi scanning and connection features)
- Ethereum wallet

## Installation

1. Clone the repository:https://github.com/NetSepio/dwifi-client.git
2. Navigate to the project directory:cd ./dwifi-client

## Configuration

1. Create a `.env` file in the project root directory with the following content:

Dwifi_NODE_URL=
PRIVATE_KEY=
CONTRACT_ADDRESS=

## Usage

1. Run the application:go run main.go


2. Follow the on-screen prompts to:
- Enter your Ethereum wallet address
- Select an available Erebrus DWiFi network
- Connect to the chosen network
- Disconnect and process payment

## Dependencies

- github.com/ethereum/go-ethereum: Ethereum client
- github.com/fatih/color: Colored console output
- github.com/joho/godotenv: Environment variable management
