package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/patiee/avs-go-operator/contract"
	"github.com/patiee/avs-go-operator/eigen"
)

func main() {
	logger := log.Default()

	logger.Print("Starting go-operator")
	defer logger.Print("go-operator exited")

	env, err := godotenv.Read(".env")
	if err != nil {
		logger.Fatalf("Error while reading .env file: %v\n", err)
	}

	client, err := ethclient.Dial(fmt.Sprintf("ws://%s", env["RPC_URL"]))
	if err != nil {
		logger.Fatalf("Error while connecting to Ethereum client: %v\n", err)
	}

	gasLimit, err := strconv.Atoi(env["GAS_LIMIT"])
	if err != nil {
		logger.Fatalf("Error while parsing gas limit: %v\n", err)
	}

	gasPrice, err := strconv.Atoi(env["GAS_PRICE"])
	if err != nil {
		logger.Fatalf("Error while parsing gas limit: %v\n", err)
	}
	gasPriceInt := big.NewInt(int64(gasPrice))

	contractService, err := contract.New(client, logger, uint64(gasLimit), gasPriceInt, env["HELLO_WORLD_ADDRESS"])
	if err != nil {
		logger.Fatalf("Error while creating smart contract service: %v\n", err)
	}

	eigenService, err := eigen.New(uint64(gasLimit), gasPriceInt, env["HOLESKY_DELEGATION_MANAGER_ADDRESS"], client, logger)
	if err != nil {
		logger.Fatalf("Error while creating eigen smart contract service: %v\n", err)
	}

	privateKey, err := crypto.HexToECDSA(env["WALLET_KEY"])
	if err != nil {
		logger.Fatalf("Failed to load private key: %v\n", err)
	}

	if err := eigenService.RegisterAsOperator(privateKey); err != nil {
		logger.Fatalf("Error registering as operator: %v\n", err)
	}

	if err := contractService.StartListeningForEvents(privateKey); err != nil {
		logger.Fatalf("Error while listening for smart contract events: %v\n", err)
	}

}
