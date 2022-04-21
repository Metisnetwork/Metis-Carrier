package metispay

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Metisnetwork/Metis-Carrier/ach/metispay/contracts"
	"github.com/Metisnetwork/Metis-Carrier/ach/metispay/kms"
	"github.com/Metisnetwork/Metis-Carrier/common/hexutil"
	"github.com/Metisnetwork/Metis-Carrier/core"
	"github.com/Metisnetwork/Metis-Carrier/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethereumtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"strings"
)

const keystoreFile = ".keystore"

//"task:0x81050ea1c64ab0ed96e50b151c36bcef180eea18d3b3245e7c4a42aa08638c58"
var (
	contractMetisPayAddress          = common.HexToAddress("0x3979cA71ea6B4c0A7CF23a8Bf216fD9FC37a4dF9")
	defaultDataTokenPrepaymentAmount = big.NewInt(1)

	mockTaskID, _ = hexutil.DecodeBig("0x81050ea1c64ab0ed96e50b151c36bcef180eea18d3b3245e7c4a42aa08638c58")
)

func LoadCarrierKey() (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatalf("cannot load private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, fromAddress
}

type Config struct {
	URL           string `json:"chain"`
	WalletPwd     string
	privateKey    *ecdsa.PrivateKey
	walletAddress common.Address
}

type MetisPayManager struct {
	dataCenter               core.CarrierDB
	Config                   *Config
	Kms                      kms.KmsService
	client                   *ethclient.Client
	chainID                  *big.Int
	abi                      abi.ABI
	contractMetisPayInstance *contracts.MetisPay
}

func (metisPay *MetisPayManager) loadPrivateKey() {
	if wallet, err := metisPay.dataCenter.QueryOrgWallet(); err != nil {
		log.Errorf("load organization wallet error: %v", err)
		return
	} else if wallet == nil {
		log.Warnf("organization wallet not found, please generate it.")
		return
	} else {
		metisPay.Config.walletAddress = wallet.Address
		if metisPay.Kms != nil {
			if key, err := metisPay.Kms.Decrypt(wallet.PriKey); err != nil {
				log.Errorf("decrypt organization wallet private key error: %v", err)
				return
			} else {
				priKey, err := crypto.ToECDSA([]byte(key))
				if err != nil {
					log.Errorf("convert organization wallet private key to ECDSA error: %v", err)
					return
				}
				metisPay.Config.privateKey = priKey

			}
		} else {
			key, err := crypto.HexToECDSA(wallet.PriKey)
			if err != nil {
				log.Errorf("not a valid private key hex: %v", err)
				return
			} else {
				metisPay.Config.privateKey = key
			}
		}
	}
}

func (metisPay *MetisPayManager) getPrivateKey() *ecdsa.PrivateKey {
	if metisPay.Config.privateKey == nil {
		metisPay.loadPrivateKey()
	}
	return metisPay.Config.privateKey
}

func (metisPay *MetisPayManager) loadKeystore() {
	var content string
	keyBytes, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		log.Fatalf("read keystore error, %v", err)
	}
	content = string(keyBytes)

	if metisPay.Kms != nil {
		content, err = metisPay.Kms.Decrypt(content)
		if err != nil {
			log.Fatalf("KMS decrypt keystore error, %v", err)
		}
	}

	key, err := keystore.DecryptKey([]byte(content), metisPay.Config.WalletPwd)
	if err != nil {
		log.Fatalf("decrypt keystore error, %v", err)
	}
	metisPay.Config.privateKey = key.PrivateKey
	metisPay.Config.walletAddress = key.Address
}

func NewMetisPayManager(db core.CarrierDB, config *Config, kmsConfig *kms.Config) *MetisPayManager {
	log.Info("Init MetisPay manager ...")
	metisPay := new(MetisPayManager)
	metisPay.dataCenter = db

	if config != nil && len(config.URL) > 0 {
		metisPay.Config = config

		client, err := ethclient.Dial(config.URL)
		if err != nil {
			log.Fatal(err)
		}
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		metisPay.client = client
		metisPay.chainID = chainID

		instance, err := contracts.NewMetisPay(contractMetisPayAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		metisPay.contractMetisPayInstance = instance

		if kmsConfig != nil {
			metisPay.Kms = &kms.AliKms{Config: kmsConfig}
		}
		metisPay.loadPrivateKey()
	}

	abiCode, err := abi.JSON(strings.NewReader(contracts.MetisPayABI))
	if err != nil {
		log.Fatal(err)
	}
	metisPay.abi = abiCode
	return metisPay
}

func (metisPay *MetisPayManager) buildTxOpts() (*bind.TransactOpts, error) {
	nonce, err := metisPay.client.PendingNonceAt(context.Background(), metisPay.Config.walletAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := metisPay.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	//txOpts := bind.NewKeyedTransactor(metisPay.Config.privateKey)
	txOpts, err := bind.NewKeyedTransactorWithChainID(metisPay.getPrivateKey(), metisPay.chainID)
	if err != nil {
		return nil, err
	}

	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)     // in wei
	txOpts.GasLimit = uint64(300000) // in units
	txOpts.GasPrice = gasPrice
	return txOpts, nil
}

func convert(dataTokenAddressList []string) ([]common.Address, []*big.Int) {
	tokenAddressList := make([]common.Address, len(dataTokenAddressList))
	dataTokenAmountList := make([]*big.Int, len(dataTokenAddressList))
	for idx, addrHex := range dataTokenAddressList {
		tokenAddressList[idx] = common.HexToAddress(addrHex)
		dataTokenAmountList[idx] = defaultDataTokenPrepaymentAmount
	}
	return tokenAddressList, dataTokenAmountList
}

// EstimateTaskGas estimates gas fee for a task's sponsor.
// EstimateTaskGas returns estimated gas and suggested gas price.
func (metisPay *MetisPayManager) EstimateTaskGas(dataTokenAddressList []string) (uint64, *big.Int, error) {
	tokenAddressList, tokenAmountList := convert(dataTokenAddressList)

	gasLimit1, err := metisPay.estimateGas("prepay", mockTaskID, metisPay.Config.walletAddress, big.NewInt(1), tokenAddressList, tokenAmountList)
	if err != nil {
		log.Errorf("call EstimateTaskGas error: %v", err)
		return 0, nil, err
	}

	gasLimit2, err := metisPay.estimateGas("settle", mockTaskID, new(big.Int).SetUint64(1))
	if err != nil {
		log.Errorf("call EstimateTaskGas error: %v", err)
		return 0, nil, err
	}

	gasPrice, err := metisPay.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("call SuggestGasPrice error: %v", err)
		return 0, nil, err
	}
	return gasLimit1 + gasLimit2, gasPrice, nil
}

// estimateGas call EstimateGas() by RPC
func (metisPay *MetisPayManager) estimateGas(method string, params ...interface{}) (uint64, error) {
	input, err := metisPay.abi.Pack(method, params...)
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{From: metisPay.Config.walletAddress, To: &contractMetisPayAddress, Data: input, Gas: 0, GasPrice: big.NewInt(0)}
	estimatedGas, err := metisPay.client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, err
	}
	return estimatedGas, nil
}

// QueryOrgWallet returns the organization wallet address
func (metisPay *MetisPayManager) QueryOrgWallet() (common.Address, error) {
	wallet, err := metisPay.dataCenter.QueryOrgWallet()

	if nil != err {
		log.WithError(err).Error("failed to query organization wallet. ", err)
		return common.Address{}, errors.New("failed to query organization wallet")
	}

	if wallet != nil {
		return wallet.Address, nil
	}

	return common.Address{}, nil
}

// GenerateOrgWallet generates a new wallet if there's no wallet for current organization, if there is an organization wallet already, just reuse it.
func (metisPay *MetisPayManager) GenerateOrgWallet() (common.Address, error) {
	walletAddr, err := metisPay.QueryOrgWallet()
	if err != nil {
		return common.Address{}, err
	}

	if walletAddr != (common.Address{}) {
		log.Warnf("organization wallet exists, just retuens current wallet: %s", walletAddr)
		return walletAddr, nil
	}

	key, _ := crypto.GenerateKey()
	keyHex := hex.EncodeToString(crypto.FromECDSA(key))
	addr := crypto.PubkeyToAddress(key.PublicKey)

	if metisPay.Kms != nil {
		if cipher, err := metisPay.Kms.Encrypt(keyHex); err != nil {
			return common.Address{}, errors.New("cannot encrypt organization wallet private key")
		} else {
			keyHex = cipher
		}
	}

	wallet := new(types.OrgWallet)
	wallet.PriKey = keyHex
	wallet.Address = addr
	if err := metisPay.dataCenter.StoreOrgWallet(wallet); err != nil {
		log.WithError(err).Error("Failed to store organization wallet")
		return common.Address{}, errors.New("failed to store organization wallet")
	} else {
		metisPay.Config.privateKey = key
		metisPay.Config.walletAddress = addr
		return addr, nil
	}
}

// Prepay transfers funds from task sponsor to MetisPay.
// Prepay returns hx.Hash, estimate gas for calling Prepay() and error.
func (metisPay *MetisPayManager) Prepay(taskID *big.Int, taskSponsorAccount common.Address, dataTokenAddressList []common.Address) (common.Hash, uint64, error) {
	if metisPay.getPrivateKey() == nil {
		log.Errorf("cannot send Prepay transaction cause organization wallet missing")
		return common.Hash{}, 0, errors.New("organization private key is missing")
	}

	dataTokenAmountList := make([]*big.Int, len(dataTokenAddressList))
	for idx, _ := range dataTokenAddressList {
		dataTokenAmountList[idx] = defaultDataTokenPrepaymentAmount
	}
	gasLimit, err := metisPay.estimateGas("prepay", taskID, new(big.Int).SetUint64(1), dataTokenAddressList, dataTokenAmountList)
	if err != nil {
		log.Errorf("failed to estimate gas for MetisPay.Prepay() error: %v", err)
		return common.Hash{}, 0, errors.New("failed to estimate gas for MetisPay.Prepay()")
	}

	//估算gas, +30%
	//gasLimit = uint64(float64(gasLimit) * 1.30)

	opts, err := metisPay.buildTxOpts()
	if err != nil {
		log.Errorf("failed to build transact options to call MetisPay.Prepay(): %v", err)
		return common.Hash{}, 0, errors.New("failed to build transact options to call MetisPay.Prepay()")
	}
	tx, err := metisPay.contractMetisPayInstance.Prepay(opts, taskID, taskSponsorAccount, new(big.Int).SetUint64(gasLimit), dataTokenAddressList, dataTokenAmountList)
	if err != nil {
		log.Errorf("failed to call MetisPay.Prepay(): %v", err)
		return common.Hash{}, 0, errors.New("failed to call MetisPay.Prepay()")
	}
	log.Debugf("call MetisPay.Prepay() txHash:%v, taskID:%s ", tx.Hash().Hex(), hexutil.EncodeBig(taskID))

	return tx.Hash(), gasLimit, nil
}

// Settle get funds clearing, transfers funds from MetisPay to DataToken owner.
// Settle returns hx.Hash, estimate gas for calling Settle() and error.
func (metisPay *MetisPayManager) Settle(taskID *big.Int, gasRefundPrepayment int64) (common.Hash, uint64, error) {
	if metisPay.getPrivateKey() == nil {
		log.Errorf("cannot send Settle transaction cause organization wallet missing")
		return common.Hash{}, 0, errors.New("organization private key is missing")
	}

	gasLimit, err := metisPay.estimateGas("settle", taskID, new(big.Int).SetUint64(1))
	if err != nil {
		log.Errorf("failed to estimate gas for MetisPay.Settle()r: %v", err)
		return common.Hash{}, 0, errors.New("failed to estimate gas for MetisPay.Settle()")
	}
	//估算gas, +30%
	//gasLimit = uint64(float64(gasLimit) * 1.30)

	if int64(gasLimit) > gasRefundPrepayment {
		gasLimit = uint64(int64(gasLimit) - gasRefundPrepayment)
	}

	opts, err := metisPay.buildTxOpts()
	if err != nil {
		log.Errorf("failed to build transact options: %v", err)
	}

	tx, err := metisPay.contractMetisPayInstance.Settle(opts, taskID, new(big.Int).SetUint64(gasLimit))
	if err != nil {
		log.Errorf("failed to call MetisPay.Settle(): %v", err)
		return common.Hash{}, 0, errors.New("failed to call MetisPay.Settle()")
	}
	log.Debugf("call MetisPay.Settle() txHash:%v, taskID:%s ", tx.Hash().Hex(), hexutil.EncodeBig(taskID))

	return tx.Hash(), gasLimit, nil
}



// GetReceipt returns the tx receipt. The caller goroutine will be blocked, and the caller could receive the receipt by channel.
func (metisPay *MetisPayManager) GetReceipt(ctx context.Context, txHash common.Hash, period time.Duration) *ethereumtypes.Receipt {

	if period < 0 { // do once only
		receipt, err := metisPay.client.TransactionReceipt(context.Background(), txHash)
		if nil != err {
			//including NotFound
			log.Errorf("query prepay transaction receipt failed, %s", err)
			return nil
		} else {
			return receipt
		}

	} else {
		ticker := time.NewTicker(period)

		for {
			select {
			case <- ctx.Done():
				return nil
			case <-ticker.C:
				receipt, err := metisPay.client.TransactionReceipt(context.Background(), txHash)
				if nil != err {
					//including NotFound
					log.Errorf("query prepay transaction receipt failed, %s", err)
				} else {
					return receipt
				}
			}
		}
	}
}

// GetTaskState returns the task payment state.
// -1 : task is not existing in PayMetis.
// 1 : task has prepaid
func (metisPay *MetisPayManager) GetTaskState(taskID *big.Int) (int, error) {
	if state, err := metisPay.contractMetisPayInstance.TaskState(&bind.CallOpts{}, taskID); err != nil {
		return -1, err
	} else {
		return int(state), nil
	}
}
