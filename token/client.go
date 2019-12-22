package token

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"

	"github.com/lazerdye/go-eth/wallet"
)

const (
	DGD         = "dgd"
	DGDContract = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a"
	DGDDecimals = 9

	STORJ         = "storj"
	STORJContrace = "0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac"
	STORJDecimals = 8
)

var (
	tokenRepo map[string]tokenInfo
)

type tokenInfo struct {
	contract common.Address
	decimals int
}

func init() {
	tokenRepo = map[string]tokenInfo{
		DGD:   {common.HexToAddress(DGDContract), DGDDecimals},
		STORJ: {common.HexToAddress(STORJContrace), STORJDecimals},
	}
}

func Tokens() []string {
	ret := make([]string, len(tokenRepo))
	i := 0
	for n, _ := range tokenRepo {
		ret[i] = n
		i++
	}
	return ret
}

type Client struct {
	client   *ethclient.Client
	instance *Token
	info     tokenInfo
}

func NewClient(tokenName string, client *ethclient.Client) (*Client, error) {
	token, ok := tokenRepo[tokenName]
	if !ok {
		return nil, errors.Errorf("Token not registered: %s", tokenName)
	}
	instance, err := NewToken(token.contract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance, token}, nil
}

func (c *Client) BalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := c.instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}

	balanceFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(c.info.decimals)))
	return balanceFloat, nil
}

func (c *Client) Transfer(sourceAccount *wallet.Account, destAccount string, amount float64, transmit bool) error {
	nonce, err := c.client.PendingNonceAt(context.Background(), sourceAccount.Account.Address)
	if err != nil {
		return err
	}
	toAddress := common.HexToAddress(destAccount)
	transferFnSignature := []byte("transfer(address,uint256)")

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Infof("Hash: %s", hexutil.Encode(methodID))
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	log.Infof("Padded hash: %s", hexutil.Encode(paddedAddress))

	fAmount := new(big.Float).Mul(big.NewFloat(amount), big.NewFloat(math.Pow10(c.info.decimals)))
	iAmount, _ := fAmount.Int(nil)
	paddedAmount := common.LeftPadBytes(iAmount.Bytes(), 32)
	log.Infof("Padded amount: %s", hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gasLimit, err := c.client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	return err
	//}
	gasLimit := uint64(200000)
	log.Infof("Gas limit: %d", gasLimit)

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	log.Infof("Gas price: %d", gasPrice)

	value := big.NewInt(0) // no ethereum transferred
	tx := types.NewTransaction(nonce, c.info.contract, value, gasLimit, gasPrice, data)
	log.Infof("Transaction: %+v", tx)

	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("ChainID: %+v", chainID)

	txSigned, err := sourceAccount.SignTx(tx, chainID)
	if err != nil {
		return err
	}
	log.Infof("Signed transaction: %+v", txSigned)

	fmt.Printf("Transaction: %s\n", txSigned.Hash().Hex())

	if transmit {
		err = c.client.SendTransaction(context.Background(), txSigned)
		if err != nil {
			return err
		}
	}

	return nil
}
