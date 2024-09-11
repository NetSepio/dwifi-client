package callcontract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	contract "chicken/contract" // Assuming you have generated Go bindings for your contract
)

func MintAndPay(client *ethclient.Client, userAddress common.Address) error {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("error loading private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("error getting nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("error getting gas price: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("error getting chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("error creating transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // Adjust as needed
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress("0xA56e3502C8224F9D3AfF5C2D11407E349eaaeBc1")
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		return fmt.Errorf("error creating contract instance: %v", err)
	}

	// Set the payable amount to 0.0001 ether
	payableAmount := new(big.Int).Mul(big.NewInt(100000000000000), big.NewInt(1)) // 0.0001 ether in wei
	auth.Value = payableAmount

	// Set the metadata URI (you may want to make this configurable)
	metadataURI := "https://ipfs.io/ipfs/Qma99ksf4Uo7LpYWH4mYrJ5ShnAx8sSSMNVt5cuDWmhCpR"

	// Call the mint function with the new parameters
	tx, err := instance.Mint(auth, metadataURI)
	if err != nil {
		return fmt.Errorf("error calling Mint function: %v", err)
	}

	fmt.Printf("Payment transaction sent: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("error waiting for transaction to be mined: %v", err)
	}
	fmt.Printf("Payment transaction mined: %s\n", receipt.TxHash.Hex())

	return nil
}
