package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	mathrand "math/rand"
	"os"
	"strings"
	"time"

	"github.com/mit-dci/dlc-oracle-go"
)

func main() {
	privateKey, err := getOrCreateKey()
	if err != nil {
		panic(err)
	}

	// Print out the public key for the oracle
	pubKey := dlcoracle.PublicKeyFromPrivateKey(privateKey)
	fmt.Printf("Oracle public key: %x\n", pubKey)

	for {
		// Generate a new one-time signing key
		privPoint, err := dlcoracle.GenerateOneTimeSigningKey()
		if err != nil {
			panic(err)
		}

		// Generate the public key to the one-time signing key (R-point) and print it out
		rPoint := dlcoracle.PublicKeyFromPrivateKey(privPoint)
		fmt.Printf("R-Point for next publication: %x\n", rPoint)

		// Sleep for 1 minute
		time.Sleep(time.Second * 60)

		// Generate random value between 10000 and 20000
		randomValue := uint64(mathrand.Int31n(10000) + 10000)

		// Generate message to sign. Uses the same encoding as expected by LIT when settling the contract
		message := dlcoracle.GenerateNumericMessage(randomValue)

		// Sign the message
		signature, err := dlcoracle.ComputeSignature(privateKey, privPoint, message)
		if err != nil {
			panic(err)
		}

		// Print out the value and signature
		fmt.Printf("Signed message. Value: %d\nSignature: %x\n", randomValue, signature)
	}
}

func getOrCreateKey() ([32]byte, error) {
	// Initialize the byte array that will hold the generated key
	var priv [32]byte

	// Check if the privatekey.hex file exists
	_, err := os.Stat("privatekey.hex")
	if err != nil {
		if os.IsNotExist(err) {
			// If not, generate a new private key by reading 32 random bytes
			rand.Read(priv[:])

			// Convert the key in to a hexadecimal format
			keyhex := fmt.Sprintf("%x\n", priv[:])

			// Save the hexadecimal value into the file
			err := ioutil.WriteFile("privatekey.hex", []byte(keyhex), 0600)

			if err != nil {
				// Unable the save the key file, return the error
				return priv, err
			}
		} else {
			// Some other error occurred while checking the file's existence, return the error
			return priv, err
		}
	}

	// At this point, the file either existed or is created. Read the private key from the file
	keyhex, err := ioutil.ReadFile("privatekey.hex")
	if err != nil {
		// Unable to read the key file, return the error
		return priv, err
	}

	// Trim any whitespace from the file's contents
	keyhex = []byte(strings.TrimSpace(string(keyhex)))

	// Decode the hexadecimal format into a byte array
	key, err := hex.DecodeString(string(keyhex))
	if err != nil {
		// Unable to decode the hexadecimal format, return the error
		return priv, err
	}

	// Copy the variable-width byte array key into priv ([32]byte)
	copy(priv[:], key[:])

	// Return the key
	return priv, nil
}
