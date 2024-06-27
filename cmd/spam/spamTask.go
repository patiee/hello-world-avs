package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/patiee/avs-go-operator/contract"
	"golang.org/x/exp/rand"
)

func generateRandomName() string {
	adjectives := []string{"Quick", "Lazy", "Sleepy", "Noisy", "Hungry"}
	nouns := []string{"Fox", "Dog", "Cat", "Mouse", "Bear"}

	rand.Seed(uint64(time.Now().UnixNano()))

	adjIndex := rand.Intn(len(adjectives))
	nounIndex := rand.Intn(len(nouns))
	number := rand.Intn(1000)

	return fmt.Sprintf("%s%s%d", adjectives[adjIndex], nouns[nounIndex], number)
}

func main() {
	logger := log.Default()

	logger.Print("Starting go-operator")
	defer logger.Print("go-operator exited")

	env, err := godotenv.Read(".env")
	if err != nil {
		logger.Fatalf("Error while reading .env file: %w", err)
	}

	client, err := ethclient.Dial(fmt.Sprintf("ws://%s", env["RPC_URL"]))
	if err != nil {
		logger.Fatalf("Error while connecting to Ethereum client: %w", err)
	}

	gasLimit, err := strconv.Atoi(env["GAS_LIMIT"])
	if err != nil {
		logger.Fatalf("Error while parsing gas limit: %w", err)
	}

	gasPrice, err := strconv.Atoi(env["GAS_PRICE"])
	if err != nil {
		logger.Fatalf("Error while parsing gas limit: %v\n", err)
	}
	gasPriceInt := big.NewInt(int64(gasPrice))

	contractService, err := contract.New(client, logger, uint64(gasLimit), gasPriceInt, env["HELLO_WORLD_ADDRESS"])
	if err != nil {
		logger.Fatalf("Error while creating smart contract service: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(env["WALLET_KEY"])
	if err != nil {
		logger.Fatalf("Failed to load private key: %w", err)
	}

	// Create a new task every 15 seconds
	for {
		if err := contractService.CreateNewTask(privateKey, generateRandomName()); err != nil {
			logger.Fatalf("Failed to create a new task: %w", err)
		}

		time.Sleep(15 * time.Second)
	}
}
