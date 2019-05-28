package crypto

import (
	"bytes"
	"fmt"

	"github.com/ZZMarquis/gm/sm2"

	"github.com/dappledger/AnnChain-go-sdk/ed25519"
	"github.com/dappledger/AnnChain-go-sdk/ed25519/extra25519"
	ghash "github.com/dappledger/AnnChain-go-sdk/go-hash"
	"github.com/dappledger/AnnChain-go-sdk/wire"
)

const (
	PubKeyLenEd25519 = 32
	PubKeyLenGmsm2   = 64
)

// PubKey is part of Account and Validator.
type PubKey interface {
	Address() []byte
	Bytes() []byte
	KeyString() string
	VerifyBytes(msg []byte, sig Signature) bool
	Equals(PubKey) bool
}

// Types of PubKey implementations
const (
	PubKeyTypeEd25519 = byte(0x01)
	PubKeyTypeGmsm2   = byte(0x03)
)

type PubKeyEd25519 [32]byte

func (pubKey PubKeyEd25519) Address() []byte {
	w, n, err := new(bytes.Buffer), new(int), new(error)
	wire.WriteBinary(pubKey[:], w, n, err)
	if *err != nil {
		panic(*err)
	}
	// append type byte
	encodedPubkey := append([]byte{PubKeyTypeEd25519}, w.Bytes()...)
	return ghash.DoHash(encodedPubkey)
}

func (pubKey PubKeyEd25519) Bytes() []byte {
	return wire.BinaryBytes(struct{ PubKey }{pubKey})
}

func (pubKey PubKeyEd25519) VerifyBytes(msg []byte, sig_ Signature) bool {
	sig, ok := sig_.(SignatureEd25519)
	if !ok {
		return false
	}
	pubKeyBytes := [32]byte(pubKey)
	sigBytes := [64]byte(sig)
	return ed25519.Verify(&pubKeyBytes, msg, &sigBytes)
}

// For use with golang/crypto/nacl/box
// If error, returns nil.
func (pubKey PubKeyEd25519) ToCurve25519() *[32]byte {
	keyCurve25519, pubKeyBytes := new([32]byte), [32]byte(pubKey)
	ok := extra25519.PublicKeyToCurve25519(keyCurve25519, &pubKeyBytes)
	if !ok {
		return nil
	}
	return keyCurve25519
}

func (pubKey PubKeyEd25519) String() string {
	return fmt.Sprintf("PubKeyEd25519{%X}", pubKey[:])
}

// Must return the full bytes in hex.
// Used for map keying, etc.
func (pubKey PubKeyEd25519) KeyString() string {
	return fmt.Sprintf("%X", pubKey[:])
}

func (pubKey PubKeyEd25519) Equals(other PubKey) bool {
	if otherEd, ok := other.(PubKeyEd25519); ok {
		return bytes.Equal(pubKey[:], otherEd[:])
	} else {
		return false
	}
}

//-------------------------------------

// PubKeyGmsm2 Implements PubKey
type PubKeyGmsm2 [64]byte

func (pubKey PubKeyGmsm2) Address() []byte {
	w, n, err := new(bytes.Buffer), new(int), new(error)
	wire.WriteBinary(pubKey[:], w, n, err)
	if *err != nil {
		panic(*err)
	}
	// append type byte
	encodedPubkey := append([]byte{PubKeyTypeGmsm2}, w.Bytes()...)
	return ghash.DoHash(encodedPubkey)
}

func (pubKey PubKeyGmsm2) Bytes() []byte {
	return wire.BinaryBytes(struct{ PubKey }{pubKey})
}

func (pubKey PubKeyGmsm2) VerifyBytes(msg []byte, sig_ Signature) bool {
	pub, err := sm2.RawBytesToPublicKey(pubKey[:])
	if err != nil {
		panic(err)
	}
	sig, ok := sig_.(SignatureGmsm2)
	if !ok {
		return false
	}
	return sm2.Verify(pub, nil, msg, SignatureGmsm2ToSM2Sign(sig))
}

func (pubKey PubKeyGmsm2) String() string {
	return fmt.Sprintf("PubKeyGmsm2{%X}", pubKey[:])
}

// Must return the full bytes in hex.
// Used for map keying, etc.
func (pubKey PubKeyGmsm2) KeyString() string {
	return fmt.Sprintf("%X", pubKey[:])
}

func (pubKey PubKeyGmsm2) Equals(other PubKey) bool {
	if otherSecp, ok := other.(PubKeyGmsm2); ok {
		return bytes.Equal(pubKey[:], otherSecp[:])
	} else {
		return false
	}
}
