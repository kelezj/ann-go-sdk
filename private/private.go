package private

import (
	"errors"

	"github.com/dappledger/ann-go-sdk/rlp"
)

type ReplacePayload struct {
	PrivateMembers []string
	Payload        []byte
}

func NewReplacePayload(privateMembers []string, payload []byte) *ReplacePayload {
	return &ReplacePayload{
		PrivateMembers: privateMembers,
		Payload:        payload,
	}
}

func (p *ReplacePayload) Encode() ([]byte, error) {
	payload, err := rlp.EncodeToBytes(p)
	if err != nil {
		return nil, errors.New("private replace payload encode error:" + err.Error())
	}
	return payload, err
}

func (p *ReplacePayload) Decode(data []byte) error {
	err := rlp.DecodeBytes(data, p)
	if err != nil {
		return errors.New("private replace payload decode error:" + err.Error())
	}
	return nil
}
