package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/tanx-libs/tanx-connector-go/contract"
)

func getTransactionOpt(ctx context.Context, client simulated.Client, ethPrivateKey string, signerFn bind.SignerFn, amount float64, decimal int) (*bind.TransactOpts, error) {
	if ethPrivateKey[:2] == "0x" {
		ethPrivateKey = ethPrivateKey[2:]
	}

	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	amountInWei := ToWei(amount, decimal)

	tipCap, _ := client.SuggestGasTipCap(context.Background())
	feeCap, _ := client.SuggestGasPrice(context.Background())

	return &bind.TransactOpts{
		From:      fromAddress,
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    signerFn,
		Value:     amountInWei,
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
		Context:   ctx,
	}, nil
}

func getTokenBalance(ctx context.Context, client simulated.Client, ethAddress string, currency Currency, decimal string, contractAddress string) (*big.Float, error) {
	var balance *big.Int

	ethAdr := common.HexToAddress(ethAddress)

	d, err := strconv.Atoi(decimal)
	if err != nil {
		return nil, err
	}

	switch currency {
	case ETH, MATIC:
		balance, err = client.BalanceAt(ctx, ethAdr, nil)
		if err != nil {
			return nil, err
		}

	default:
		tokenAddr := common.HexToAddress(contractAddress)
		ctr, err := contract.NewErc20(tokenAddr, client)
		if err != nil {
			return nil, err
		}

		balance, err = ctr.BalanceOf(nil, ethAdr)
		if err != nil {
			return nil, err
		}
	}

	return ToDecimal(balance, d), nil
}

func setAllowance(ctx context.Context, client simulated.Client, tokenContractAddress string, txOpts *bind.TransactOpts, senderAddr common.Address, blockchainDecimal int, allowance float64) error {
	tokenContractAddr := common.HexToAddress(tokenContractAddress)

	erc20Contract, err := contract.NewErc20(tokenContractAddr, client)
	if err != nil {
		return err
	}

	txOpts.GasLimit, err = client.EstimateGas(ctx, ethereum.CallMsg{
		From: tokenContractAddr,
	})
	if err != nil {
		return err
	}

	amountInWei := ToWei(allowance, blockchainDecimal)

	_, err = erc20Contract.Approve(txOpts, senderAddr, amountInWei)
	if err != nil {
		return err
	}

	return nil
}


func  getAllowance(
	client simulated.Client,
	ethAddress string,
	tokenContract string,
	decimal int,
	spenderAddress common.Address,
) (float64, error) {
	ctr, err := contract.NewErc20(common.HexToAddress(tokenContract), client)
	if err != nil {
		return 0, err
	}

	allowance, err := ctr.Allowance(nil, common.HexToAddress(ethAddress), spenderAddress)
	if err != nil {
		return 0, err
	}

	res, _ := ToDecimal(allowance, decimal).Float64()

	return res, nil
}

