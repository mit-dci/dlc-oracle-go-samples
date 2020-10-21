package crypto

import (
	"github.com/btcsuite/btcd/btcec"

	"github.com/awnumar/memguard"
)

var safeKey *memguard.LockedBuffer

func StoreKey(key *[32]byte) error {
	newA := memguard.NewBufferFromBytes(key[:])
	safeKey = newA

	return nil
}

func RetrieveKey() *[32]byte {
	key := new([32]byte)
	copy(key[:], safeKey.String())
	return key
}

func GetPubKey() (*[33]byte, error) {
	result := new([33]byte)
	key := RetrieveKey()
	_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), key[:])
	key = nil
	copy(result[:], pubKey.SerializeCompressed()[:])
	return result, nil
}
