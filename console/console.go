package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"

	"github.com/mit-dci/dlc-oracle-go"
)

var Log = log.New(os.Stdout,
	"INFO: ",
	log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	// Create app folder in user's homedir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dataDir := path.Join(usr.HomeDir, ".dlcoracle")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, 0700)
	}

	// Read or create a keyfile
	keyFilePath := path.Join(dataDir, "privkey.hex")
	key, err := dlcoracle.ReadKeyFile(keyFilePath)
	if err != nil {
		Log.Fatal("Could not open or create keyfile:", err)
		os.Exit(1)
	}

	// Print out the public key for the oracle
	pubKey := dlcoracle.PublicKeyFromPrivateKey(*key)
	Log.Printf("Oracle public key: %x\n", pubKey)

	for {
		// Generate one-time signing private scalar
		privPoint, err := dlcoracle.GenerateOneTimeSigningKey()
		if err != nil {
			Log.Fatal("Could not generate OTS scalar:", err)
			os.Exit(1)
		}

		// Print out the R-Point (public key to the private scalar)
		rPoint := dlcoracle.PublicKeyFromPrivateKey(privPoint)
		Log.Printf("R-Point for next publication: %x\n", rPoint)

		// Get numeric value to sign from console
		fmt.Print("Enter number to publish (-1 to exit): ")
		var input string
		fmt.Scanln(&input)

		// Convert value to numeric
		i, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			Log.Fatal("Invalid value: ", err)
			os.Exit(1)
		}

		// Break when user entered -1
		if i == -1 {
			Log.Println("-1 entered, exiting...")
			break
		}

		// Convert numeric value to a byte array
		var buf bytes.Buffer
		binary.Write(&buf, binary.BigEndian, i)
		message := buf.Bytes()

		// Sign the message
		sig, err := dlcoracle.ComputeSignature(*key, privPoint, message)
		if err != nil {
			Log.Fatal("Error signing value: ", err)
			os.Exit(1)
		}

		Log.Printf("Signature: %x\n\n", sig)

	}

}
