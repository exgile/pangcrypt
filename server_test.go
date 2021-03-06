package pangcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerCryptoGood(t *testing.T) {
	tests := []struct {
		key           uint8
		plain, cipher []byte
	}{
		{
			key: 0,
			plain: []byte{
				0x96, 0x00, 0xFF, 0xE0, 0xF5, 0x05, 0x00, 0x00,
				0x00, 0x00,
			},
			cipher: []byte{
				0xEE, 0x13, 0x00, 0xFB, 0x00, 0x00, 0x00, 0x36,
				0x1B, 0x96, 0x00, 0xF5, 0xFB, 0x63, 0x05, 0xFF,
				0xE0, 0xF5, 0x05, 0x11, 0x00, 0x00,
			},
		},
		{
			key: 7,
			plain: []byte{
				0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			},
			cipher: []byte{
				0x49, 0x10, 0x00, 0x20, 0x00, 0x00, 0x00, 0x88,
				0x18, 0x01, 0x00, 0x06, 0x18, 0x01, 0x00, 0x01,
				0x11, 0x00, 0x00,
			},
		},
		{
			key: 5,
			plain: []byte{
				0x06, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x50, 0x61, 0x6e, 0x67, 0x79, 0x61,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00,
			},
			cipher: []byte{
				0x23, 0x2b, 0x00, 0xaa, 0x00, 0x00, 0x02, 0x13,
				0x07, 0x06, 0x02, 0x14, 0x66, 0x68, 0x67, 0x29,
				0x00, 0x4f, 0x67, 0x59, 0x76, 0x23, 0x00, 0x70,
				0x76, 0x22, 0x00, 0xce, 0x9d, 0x20, 0x20, 0x89,
				0x20, 0x07, 0x25, 0x17, 0xdc, 0x07, 0x05, 0x00,
				0x00, 0x00, 0x00, 0x11, 0x00, 0x00,
			},
		},
		{
			key: 5,
			plain: []byte{
				0x01, 0x00, 0x00, 0x04, 0x00, 0x6a, 0x6f, 0x68,
				0x6e, 0xb4, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x08, 0x00, 0x6a, 0x6f, 0x68,
				0x6e, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00,
			},
			cipher: []byte{
				0x84, 0x25, 0x00, 0x45, 0x00, 0x00, 0x00, 0x58,
				0x09, 0x01, 0x00, 0x27, 0x0d, 0x01, 0x6a, 0x6f,
				0x6c, 0x6e, 0xde, 0x7f, 0x68, 0x43, 0xb4, 0x10,
				0x09, 0x25, 0x00, 0x6a, 0x66, 0x60, 0x6e, 0x1e,
				0x0a, 0x1b, 0x1a, 0x74, 0x65, 0x62, 0x74, 0x00,
			},
		},
		{
			key: 5,
			plain: []byte{
				0x01, 0x00, 0x00, 0x04, 0x00, 0x6a, 0x6f, 0x68,
				0x6e, 0xb4, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x08, 0x00, 0x6a, 0x6f, 0x68,
				0x6e, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00,
			},
			cipher: []byte{
				0x84, 0x25, 0x00, 0x45, 0x00, 0x00, 0x00, 0x58,
				0x09, 0x01, 0x00, 0x27, 0x0d, 0x01, 0x6a, 0x6f,
				0x6c, 0x6e, 0xde, 0x7f, 0x68, 0x43, 0xb4, 0x10,
				0x09, 0x25, 0x00, 0x6a, 0x66, 0x60, 0x6e, 0x1e,
				0x0a, 0x1b, 0x1a, 0x74, 0x65, 0x62, 0x74, 0x00,
			},
		},
	}

	for _, test := range tests {
		decrypted, err := ServerDecrypt(test.cipher, test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.plain, decrypted, "server decrypt")

		encrypted, err := ServerEncrypt(test.plain, test.key, test.cipher[0])
		assert.Nil(t, err)
		assert.Equal(t, test.cipher, encrypted, "server encrypt")
	}
}

func TestInvalidServerKey(t *testing.T) {
	var err error

	_, err = ServerEncrypt([]byte{}, 0x10, 0x00)
	assert.EqualError(t, err, "key 0x10 is too large (maximum key is 0x0f)")

	_, err = ServerDecrypt([]byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, 0x10)
	assert.EqualError(t, err, "key 0x10 is too large (maximum key is 0x0f)")
}

func TestInvalidServerBuffer(t *testing.T) {
	var err error

	_, err = ServerDecrypt([]byte{}, 0x00)
	assert.EqualError(t, err, "buffer too small (have 0 bytes, need at least 8.)")

	_, err = ServerDecrypt([]byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xff}, 0x00)
	assert.EqualError(t, err, "EOF")
}
