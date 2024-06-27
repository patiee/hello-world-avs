package contract

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	helloworld "github.com/patiee/avs-go-operator/abis"
)

// Service for smart contract events
type Service struct {
	gasLimit          uint64
	gasPrice          *big.Int
	chainID           *big.Int
	helloWorld        *helloworld.HelloWorld
	helloWorldAddress common.Address
	client            *ethclient.Client
	logger            *log.Logger
}

// New returns a new Service for smart contract events
func New(client *ethclient.Client, logger *log.Logger, gasLimit uint64, gasPrice *big.Int, smartContractAddress string) (*Service, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error while getting network id")
	}

	helloWorldAddress := common.HexToAddress(smartContractAddress)
	contract, err := helloworld.NewHelloWorld(helloWorldAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "Error while creating smar")
	}
	return &Service{
		gasLimit:          gasLimit,
		chainID:           chainID,
		helloWorld:        contract,
		helloWorldAddress: helloWorldAddress,
		client:            client,
		logger:            logger,
	}, nil
}

// CreateNewTask calls create_task on smart contract
func (s *Service) CreateNewTask(pk *ecdsa.PrivateKey, name string) error {
	transactor, err := s.transactor(pk)
	if err != nil {
		return err
	}

	// Call the contract function
	tx, err := s.helloWorld.CreateNewTask(transactor, name)
	if err != nil {
		return errors.Wrap(err, "Error while calling create_task")
	}

	if err = s.client.SendTransaction(context.Background(), tx); err != nil {
		return errors.Wrap(err, "Error while sending transaction")
	}

	s.logger.Printf("New task created, tx hash: %s\n", tx.Hash().Hex())
	return nil
}

func (s *Service) transactor(pk *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Could not create public key from private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "Error while getting nonce")
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pk, s.chainID)
	if err != nil {
		return nil, errors.Wrap(err, "Error while creating transactor")
	}
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)
	transactor.GasLimit = s.gasLimit
	transactor.GasPrice = s.gasPrice

	return transactor, nil
}

// StartListeningForEvents is watching smart contract events
func (s *Service) StartListeningForEvents(pk *ecdsa.PrivateKey) error {
	tasks := make(chan *helloworld.HelloWorldNewTaskCreated)
	sub, err := s.helloWorld.WatchNewTaskCreated(nil, tasks, nil)
	if err != nil {
		return errors.Wrap(err, "Error while subscribing for logs")
	}

	for {
		select {
		case err := <-sub.Err():
			return errors.Wrap(err, "Subscription error")
		case task := <-tasks:
			log.Printf("Received task: %+v", task)
			transactor, err := s.transactor(pk)
			if err != nil {
				return err
			}

			msgEip191 := eip191Hash(fmt.Sprintf("Hello %s", task.Task.Name))
			msg := []byte(msgEip191)

			sig, err := signMessage(pk, msg)
			if err != nil {
				return errors.Wrap(err, "Error while signing message")
			}
			s.helloWorld.RespondToTask(transactor, task.Task, task.TaskIndex, sig)
		}
	}
}

func signMessage(pk *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)

	r, s, err := ecdsa.Sign(rand.Reader, pk, hash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	return append(r.Bytes(), s.Bytes()...), nil
}

func eip191Hash(msg string) string {
	hash := crypto.Keccak256Hash([]byte(msg))
	return fmt.Sprintf("\x19\x01%s", hash.Hex())
}
