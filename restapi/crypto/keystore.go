package crypto

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec"

	"github.com/awnumar/memguard"
)

var safeKey *memguard.LockedBuffer

func StoreKey(key *[32]byte) error {
	newA, err := memguard.NewImmutableFromBytes(key[:])
	if err != nil {
		fmt.Println(err)
		memguard.SafeExit(1)
		return err
	}
	safeKey = newA

	return nil
}

func RetrieveKey() *[32]byte {
	key := new([32]byte)
	copy(key[:], safeKey.Buffer())
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
