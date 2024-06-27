package eigen

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// Service for Eigen smart contracts
type Service struct {
	chainID           *big.Int
	gasLimit          uint64
	gasPrice          *big.Int
	logger            *log.Logger
	client            *ethclient.Client
	delegationAddress common.Address
	delegation        *delegationmanager.ContractDelegationManager
}

// New returns a new Eigen service
func New(gasLimit uint64, gasPrice *big.Int, delegationAddress string, client *ethclient.Client, logger *log.Logger) (*Service, error) {
	delegationContractAddress := common.HexToAddress(delegationAddress)
	contractDelegation, err := delegationmanager.NewContractDelegationManager(delegationContractAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "Error while creating delegation manager")
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error while getting network id")
	}

	return &Service{
		chainID:           chainID,
		gasLimit:          gasLimit,
		gasPrice:          gasPrice,
		logger:            logger,
		client:            client,
		delegationAddress: delegationContractAddress,
		delegation:        contractDelegation,
	}, nil
}

// RegisterAsOperator executes RegisterAsOperator smart contract function
func (s *Service) RegisterAsOperator(pk *ecdsa.PrivateKey) error {
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("Could not create public key from private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return errors.Wrap(err, "Error while getting nonce")
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pk, s.chainID)
	if err != nil {
		return errors.Wrap(err, "Error while creating transactor")
	}
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)
	transactor.GasLimit = s.gasLimit
	transactor.GasPrice = big.NewInt(21000)

	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		// DeprecatedEarningsReceiver: common.HexToAddress(operator.EarningsReceiverAddress),
		// StakerOptOutWindowBlocks:   operator.StakerOptOutWindowBlocks,
		// DelegationApprover: common.HexToAddress(operator.DelegationApproverAddress),
	}

	tx, err := s.delegation.RegisterAsOperator(transactor, opDetails, "")
	if err != nil {
		return errors.Wrap(err, "Error while registering as operator")
	}

	if err = s.client.SendTransaction(context.Background(), tx); err != nil {
		return errors.Wrap(err, "Error while sending transaction")
	}

	s.logger.Printf("Registered as operator, tx hash: %s\n", tx.Hash().Hex())
	return nil
}
