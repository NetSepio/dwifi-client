# Erebrus DWiFi Client

![Erebrus DWiFi Client](https://github.com/user-attachments/assets/521f54b2-6b4b-446f-9dd0-8f6f300d200d)

## Description

Erebrus DWiFi Client is a Go-based application that enables users to connect to decentralized WiFi networks. It scans for nearby WiFi networks, retrieves information about Erebrus-registered networks from an API, and allows users to connect to these networks using a pay-per-minute model with AGNG on the Peaq network.

## Features

- Scan for nearby WiFi networks
- Fetch and display Erebrus-registered DWiFi networks
- Connect to selected DWiFi networks
- Real-time connection monitoring
- Automatic disconnection and NFT minting

## Prerequisites

- Go 1.15 or higher
- Windows operating system (for WiFi scanning and connection features)
- Peaq wallet

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/NetSepio/dwifi-client.git
   ```
2. Navigate to the project directory:
   ```
   cd dwifi-client
   ```

## Configuration

1. Create a `.env` file in the project root directory with the following content:
   ```
   AGUNG_NODE_URL=<Your peaq(Agung) node URL>
   PRIVATE_KEY=<Your Peaq wallet private key>
   ```
   Replace `<Your peaq(Agung) node URL>` with a valid Peaq (Agung) node URL and `<Your Peaq wallet private key>` with your wallet's private key.

## Usage

1. Run the application:
   ```
   go run main.go
   ```

2. Follow the on-screen prompts to:
   - Enter your Peaq wallet address
   - View available Erebrus DWiFi networks in your vicinity
   - Select a network to connect
   - Mint an NFT (this happens automatically before connecting)
   - Connect to the chosen network
   - Monitor your connection time
   - Disconnect when desired

3. To disconnect, press 'd' when prompted.

## Smart Contract Information

The Erebrus DWiFi Client interacts with a verified smart contract deployed on the Peaq network. Here are the details:

- Contract Address: `0x5940445e1e8A419ebea10B45c5d1C0F603926F41`
- Function Used: `mint`

The contract is verified, and you can find more information about its functions and interactions at:
[https://agung-testnet.subscan.io/account/0x5940445e1e8a419ebea10b45c5d1c0f603926f41?tab=transaction&evm_contract_tab=read](https://agung-testnet.subscan.io/account/0x5940445e1e8a419ebea10b45c5d1c0f603926f41?tab=transaction&evm_contract_tab=read)

## Important Notes

- Ensure you have sufficient AGNG in your wallet to cover the connection costs and gas fees for NFT minting.
- The application will mint an NFT before connecting to a network. This is a required step for using the DWiFi service.
- Connection costs are displayed per minute for each available network.
- Be cautious when entering your private key in the `.env` file. Never share this file or your private key with others.

## Dependencies

- github.com/ethereum/go-ethereum: Ethereum client (used for Peaq network interactions)
- github.com/fatih/color: Colored console output
- github.com/joho/godotenv: Environment variable management

## Troubleshooting

- If you encounter issues connecting to the Peaq network, ensure your `AGUNG_NODE_URL` is correct and accessible.
- If NFT minting fails, check that your wallet has sufficient AGNG for gas fees.
- For WiFi connection issues, ensure you have the necessary permissions on your Windows system.

For more detailed information, please refer to the [official Erebrus documentation](https://docs.netsepio.com/latest/v/erebrus/netsepio/erebrus/setup/setup-dwifi-client).

## Support

If you encounter any issues or have questions, please open an issue on the [GitHub repository](https://github.com/NetSepio/dwifi-client/issues).

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- The Erebrus team for their decentralized WiFi network concept
- The Go community for their excellent libraries and tools
- The Peaq network for providing the blockchain infrastructure
