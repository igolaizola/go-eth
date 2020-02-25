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
	WETH         = "weth"
	WETHContract = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	WETHDecimals = 18

	ANT         = "ant"
	ANTContract = "0x960b236A07cf122663c4303350609A66A7B288C0"
	ANTDecimals = 18

	BAT         = "bat"
	BATContract = "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"
	BATDecimals = 18

	CUSD         = "cusd"
	CUSDContract = "0x1410d4eC3D276C0eBbf16ccBE88A4383aE734eD0"
	CUSDDecimals = 18

	DAI         = "dai"
	DAIContract = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	DAIDecimals = 18

	DGD         = "dgd"
	DGDContract = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a"
	DGDDecimals = 9

	EXMR         = "exmr"
	EXMRContract = "0x331fA6C97c64e47475164b9fC8143b533c5eF529"
	EXMRDecimals = 18

	FOAM         = "foam"
	FOAMContract = "0x4946Fcea7C692606e8908002e55A582af44AC121"
	FOAMDecimals = 18

	GNO         = "gno"
	GNOContract = "0x6810e776880c02933d47db1b9fc05908e5386b96"
	GNODecimals = 18

	GWIT         = "gwit"
	GWITContract = "0x55D0Bb8D7e7fBf5B863C7923c4645FF83c3D0033"
	GWITDecimals = 18

	KNC         = "knc"
	KNCContract = "0xdd974d5c2e2928dea5f71b9825b8b646686bd200"
	KNCDecimals = 18

	LABX         = "labx"
	LABXContract = "0xF0daeC652dD7C9f779e7C0F3ff5610fa3B61f61F"
	LABXDecimals = 18

	LINK         = "link"
	LINKContract = "0x514910771AF9Ca656af840dff83E8264EcF986CA"
	LINKDecimals = 18

	LPT         = "lpt"
	LPTContract = "0x58b6a8a3302369daec383334672404ee733ab239"
	LPTDecimals = 18

	MLN         = "mln"
	MLNContract = "0xec67005c4e498ec7f55e092bd1d35cbc47c91892"
	MLNDecimals = 18

	MKR         = "mkr"
	MKRContract = "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"
	MKRDecimals = 18

	OHDAI         = "ohdai"
	OHDAIContract = "0x64a03cE1E52B4e579f0A1cf44cF95C0D7898B5A3"
	OHDAIDecimals = 18

	REP         = "rep"
	REPContract = "0x1985365e9f78359a9B6AD760e32412f4a445E862"
	REPDecimals = 18

	RIGO         = "rigo"
	RIGOContract = "0x4FbB350052Bca5417566f188eB2EBCE5b19BC964"
	RIGODecimals = 18

	SAI         = "sai"
	SAIContract = "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359"
	SAIDecimals = 18

	SPANK         = "spank"
	SPANKContract = "0x42d6622deCe394b54999Fbd73D108123806f6a18"
	SPANKDecimals = 18

	STORJ         = "storj"
	STORJContract = "0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac"
	STORJDecimals = 8

	TUSD         = "tusd"
	TUSDContract = "0x0000000000085d4780B73119b644AE5ecd22b376"
	TUSDDecimals = 18

	USDC         = "usdc"
	USDCContract = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	USDCDecimals = 6

	ZRX         = "zrx"
	ZRXContract = "0xE41d2489571d322189246DaFA5ebDe1F4699F498"
	ZRXDecimals = 18
)

var (
	gasLimit *big.Int
	gasPrice uint64

	tokenRepo map[string]tokenInfo
)

type tokenInfo struct {
	contract common.Address
	decimals int
}

func init() {
	gasLimit = big.NewInt(5000000000)
	gasPrice = 780000
	tokenRepo = map[string]tokenInfo{
		WETH:  {common.HexToAddress(WETHContract), WETHDecimals},
		ANT:   {common.HexToAddress(ANTContract), ANTDecimals},
		BAT:   {common.HexToAddress(BATContract), BATDecimals},
		CUSD:  {common.HexToAddress(CUSDContract), CUSDDecimals},
		DAI:   {common.HexToAddress(DAIContract), DAIDecimals},
		DGD:   {common.HexToAddress(DGDContract), DGDDecimals},
		EXMR:  {common.HexToAddress(EXMRContract), EXMRDecimals},
		FOAM:  {common.HexToAddress(FOAMContract), FOAMDecimals},
		GNO:   {common.HexToAddress(GNOContract), GNODecimals},
		GWIT:  {common.HexToAddress(GWITContract), GWITDecimals},
		KNC:   {common.HexToAddress(KNCContract), KNCDecimals},
		LABX:  {common.HexToAddress(LABXContract), LABXDecimals},
		LINK:  {common.HexToAddress(LINKContract), LINKDecimals},
		LPT:   {common.HexToAddress(LPTContract), LPTDecimals},
		MKR:   {common.HexToAddress(MKRContract), MKRDecimals},
		MLN:   {common.HexToAddress(MLNContract), MLNDecimals},
		OHDAI: {common.HexToAddress(OHDAIContract), OHDAIDecimals},
		REP:   {common.HexToAddress(REPContract), REPDecimals},
		RIGO:  {common.HexToAddress(RIGOContract), RIGODecimals},
		SAI:   {common.HexToAddress(SAIContract), SAIDecimals},
		SPANK: {common.HexToAddress(SPANKContract), SPANKDecimals},
		STORJ: {common.HexToAddress(STORJContract), STORJDecimals},
		TUSD:  {common.HexToAddress(TUSDContract), TUSDDecimals},
		USDC:  {common.HexToAddress(USDCContract), USDCDecimals},
		ZRX:   {common.HexToAddress(ZRXContract), ZRXDecimals},
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

func TokenByName(name string) (common.Address, int, bool) {
	token, ok := tokenRepo[name]
	if !ok {
		return common.Address{}, 0, false
	}
	return token.contract, token.decimals, true
}

func TokenByAddress(address common.Address) (string, int, bool) {
	for n, t := range tokenRepo {
		if t.contract.Hex() == address.Hex() {
			return n, t.decimals, true
		}
	}
	return "", 0, false
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

func (c *Client) ContractAddress() common.Address {
	return c.info.contract
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
	// TODO: Why did I do this?
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
	gasLimit := uint64(200000) // TODO: Gas limit.
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

	// TODO: Let the sender do it.
	if transmit {
		err = c.client.SendTransaction(context.Background(), txSigned)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Approve(ctx context.Context, from *wallet.Account, contract common.Address, value *big.Float) (*types.Transaction, error) {
	fAmount := new(big.Float).Mul(value, big.NewFloat(math.Pow10(c.info.decimals)))
	iAmount, _ := fAmount.Int(nil)
	opts, err := from.NewTransactor(gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	return c.instance.Approve(opts, contract, iAmount)
}

func (c *Client) Allowance(ctx context.Context, address, contract common.Address) (*big.Int, error) {
	opts := bind.CallOpts{Context: ctx}
	return c.instance.Allowance(&opts, address, contract)
}
