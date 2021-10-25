package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type hasher256 struct {
	secretKey string
}

func NewHasher256(secretKey string) *hasher256 {
	return &hasher256{secretKey: secretKey}
}

func (h *hasher256) HashPassword(pass string) (string, error) {
	hm := hmac.New(sha256.New, []byte(h.secretKey))
	_, err := hm.Write([]byte(pass))
	if err != nil {
		return "", err
	}
	signedPass := base64.StdEncoding.EncodeToString(hm.Sum(nil))
	return signedPass, err
}

func (h *hasher256) Equal(pass string, hashedPass string) (bool, error) {
	hashPass2, err := h.HashPassword(pass)
	if err != nil {
		return false, err
	}
	return hashPass2 == hashedPass, nil
}
