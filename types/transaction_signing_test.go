package types

import (
	"math/big"
	"testing"

	"github.com/dappledger/ann-go-sdk/common"
	"github.com/dappledger/ann-go-sdk/crypto"
)

const (
	accPriv = "48deaa73f328f38d5fcb29d076b2b639c8491f97d245fc22e95a86366687903a"
	accAddr = "28112ca022224ae7757bcd559666be5340ff109a"
)

func TestSigner(t *testing.T) {

	tx := NewTransaction(1, common.HexToAddress(accAddr), big.NewInt(0), 100, big.NewInt(0), []byte("123"))

	signer := new(AnnsteadSigner)

	privkey, err := crypto.ToECDSA(common.Hex2Bytes(accPriv))
	if err != nil {
		panic(err)
	}
	sig, err := crypto.Sign(signer.Hash(tx).Bytes(), privkey)
	if err != nil {
		panic(err)
	}
	sigTx, err := tx.WithSignature(signer, sig, true)
	if err != nil {
		panic(err)
	}
	addr, err := signer.Sender(sigTx)
	t.Log(addr.Hex(), signer.Hash(tx), err)

}
