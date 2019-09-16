package private

import (
	"testing"
)

func TestBench(t *testing.T) {

	rp := NewReplacePayload([]string{"0xffff"}, []byte("123456"))

	byt, _ := rp.Encode()

	rp.Decode(byt)

	t.Log(rp.Payload, rp.PrivateMembers)
}
