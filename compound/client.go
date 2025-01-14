package compound

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

const ()

var (
	CethContractAddress     = common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
	CErc20ContractAddresses = map[string]common.Address{
		"bat":  common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"),
		"dai":  common.HexToAddress("0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643"),
		"zrx":  common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"),
		"uni":  common.HexToAddress("0x35A18000230DA775CAc24873d00Ff85BccdeD550"),
		"usdc": common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563"),
		"usdt": common.HexToAddress("0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9"),
	}

	mintGasSpeed = gasoracle.Fast
	mintGasLimit = uint64(300000)

	mintLogHash = common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f")
)

type Client interface {
	FromWei(*big.Int) decimal.Decimal
	Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error)
	Redeem(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error)
	ParseLog(log *types.Log) (interface{}, error)
}

type cethClient struct {
	*client.Client

	ceth *Ceth
}

func NewCEthClient(client *client.Client) (Client, error) {
	ceth, err := NewCeth(CethContractAddress, client)
	if err != nil {
		return nil, err
	}
	return &cethClient{client, ceth}, nil
}

func CompileCEthABI() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(CethABI))
}

func (c *cethClient) FromWei(wei *big.Int) decimal.Decimal {
	return client.EthFromWei(wei)
}

func (c *cethClient) Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, client.EthToWei(amount), gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}

	return c.ceth.Mint(opts)
}

func (c *cethClient) Redeem(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	opts := &bind.CallOpts{Context: ctx}
	decimals, err := c.ceth.Decimals(opts)
	if err != nil {
		return nil, err
	}
	amountWei := amount.Shift(int32(decimals.Int64())).BigInt()
	gasPrice, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	trans, err := account.NewTransactor(ctx, nil, gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}
	return c.ceth.Redeem(trans, amountWei)
}

func (c *cethClient) ParseLog(log *types.Log) (interface{}, error) {
	return nil, errors.New("Not Implemented")
}

type cerc20Client struct {
	*token2.Client

	cerc20 *CErc20
}

func NewCErc20Client(ctx context.Context, client *client.Client, address common.Address) (Client, error) {
	tok, err := token2.ByAddress(ctx, client, address)
	if err != nil {
		return nil, err
	}
	cerc20, err := NewCErc20(address, client)
	if err != nil {
		return nil, err
	}
	return &cerc20Client{tok, cerc20}, nil
}

func CompileCErc20ABI() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(CErc20ABI))
}

func (c *cerc20Client) Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}
	amountWei, err := c.ToWeiCapped(ctx, amount, account)
	if err != nil {
		return nil, err
	}

	return c.cerc20.Mint(opts, amountWei)
}

func (c *cerc20Client) Redeem(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	opts := &bind.CallOpts{Context: ctx}
	decimals, err := c.cerc20.Decimals(opts)
	if err != nil {
		return nil, err
	}
	amountWei := amount.Shift(int32(decimals.Int64())).BigInt()
	gasPrice, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	trans, err := account.NewTransactor(ctx, nil, gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}
	return c.cerc20.Redeem(trans, amountWei)
}

func (c *cerc20Client) ParseLog(logObj *types.Log) (interface{}, error) {
	switch logObj.Topics[0] {
	case mintLogHash:
		log.Infof("Found CethMint")
		return c.cerc20.ParseMint(*logObj)
	}
	return nil, nil
}
