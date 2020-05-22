package token2

import (
	"context"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/lazerdye/go-eth/client"
)

type TokenInfo struct {
	Address  string `json:"address"`
	Decimals uint8  `json:"decimals"`
}

type Tokens struct {
	Tokens map[string]TokenInfo `json:"tokens"`
}

type Client struct {
	*client.Client

	instance *Erc20
	Address  common.Address
	Decimals uint8
}

func NewClient(etherClient *client.Client, address common.Address, decimals uint8) (*Client, error) {
	instance, err := NewErc20(address, etherClient)
	if err != nil {
		return nil, err
	}
	return &Client{etherClient, instance, address, decimals}, nil
}

func ByAddress(ctx context.Context, client *client.Client, address common.Address) (*Client, error) {
	opts := bind.CallOpts{Context: ctx}
	instance, err := NewErc20(address, client)
	if err != nil {
		return nil, err
	}
	decimals, err := instance.Decimals(&opts)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance, address, decimals}, nil
}

func (c *Client) Symbol(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{Context: ctx}
	return c.instance.Symbol(opts)
}

type Registry struct {
	tokens map[string]*Client
}

func RegistryFromFile(etherClient *client.Client, filename string) (*Registry, error) {
	var tokens Tokens
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(yamlFile, &tokens); err != nil {
		return nil, err
	}
	registry := &Registry{}
	registry.tokens = make(map[string]*Client, len(tokens.Tokens))
	for name, tokenInfo := range tokens.Tokens {
		client, err := NewClient(etherClient, common.HexToAddress(tokenInfo.Address), tokenInfo.Decimals)
		if err != nil {
			return nil, err
		}
		if err := registry.Register(name, client); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func (r *Registry) Register(name string, client *Client) error {
	if _, ok := r.tokens[name]; ok != false {
		return errors.Errorf("Token %s already defined", name)
	}
	r.tokens[name] = client
	return nil
}

func (r *Registry) ByName(name string) (*Client, error) {
	client, ok := r.tokens[name]
	if !ok {
		return nil, errors.Errorf("Token %s not found", name)
	}
	return client, nil
}

func (r *Registry) ByAddress(address common.Address) (string, *Client, error) {
	for name, client := range r.tokens {
		if client.Address == address {
			return name, client, nil
		}
	}
	return "", nil, errors.Errorf("Client with address %s not found", address.String())
}

func (r *Registry) Client(etherClient *client.Client, tokenName string) (*Client, error) {
	token, ok := r.tokens[tokenName]
	if !ok {
		return nil, errors.Errorf("Unknown token: %s", tokenName)
	}
	return token, nil
}