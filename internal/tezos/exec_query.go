package tezos

import (
	"context"

	"github.com/hyperledger/firefly-transaction-manager/pkg/ffcapi"
)

// QueryInvoke executes a method on a blockchain smart contract, which might execute Smart Contract code, but does not affect the blockchain state.
func (c *tezosConnector) QueryInvoke(_ context.Context, req *ffcapi.QueryInvokeRequest) (*ffcapi.QueryInvokeResponse, ffcapi.ErrorReason, error) {
	// TODO: to implement
	return nil, "", nil
}
