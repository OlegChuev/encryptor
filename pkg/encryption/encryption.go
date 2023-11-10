package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

/*
Encrypt encrypts the given data using the specified secret key.

Parameters:
- data: Data to be encrypted.
- secretKey: Secret key used for encryption.

Returns:
- encoded: Encrypted data.
- err: Error, if any, encountered during the encryption process.

Example Usage:
encryptedData, err := Encrypt(originalData, "yourSecretKey")

	if err != nil {
	    fmt.Println("Encryption failed:", err)
	    return
	}

fmt.Println("Encryption successful. Encrypted Data:", encryptedData)
*/
func Encrypt(data []byte, secretKey string) (encoded []byte, err error) {
	gcm, err := newGcm(secretKey)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	encodedData := base64.URLEncoding.EncodeToString(cipherText)
	encoded = []byte(encodedData)

	return
}

/*
Decrypt decrypts the given encrypted data using the specified secret key.

Parameters:
- data: Encrypted data to be decrypted.
- secretKey: Secret key used for decryption.

Returns:
- decoded: Decrypted data.
- err: Error, if any, encountered during the decryption process.

Example Usage:
decryptedData, err := Decrypt(encryptedData, "yourSecretKey")

	if err != nil {
	    fmt.Println("Decryption failed:", err)
	    return
	}

fmt.Println("Decryption successful. Decrypted Data:", decryptedData)
*/
func Decrypt(data []byte, secretKey string) (decoded []byte, err error) {
	gcm, err := newGcm(secretKey)
	if err != nil {
		return
	}

	buf, err := base64.URLEncoding.DecodeString(string(data))
	if err != nil {
		return
	}

	nonce := buf[:gcm.NonceSize()]
	cipherText := buf[gcm.NonceSize():]
	decoded, err = gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return
	}

	return
}

/*
newGcm creates a new Galois/Counter Mode (GCM) cipher for symmetric encryption using the Advanced Encryption Standard (AES).

Parameters:
- secretKey: Secret key used to derive the AES block cipher.

Returns:
- gcm: GCM cipher for encryption and decryption.
- err: Error, if any, encountered during the cipher creation process.

Implementation Details:
- It initializes an AES cipher block with the provided secret key.
- It then creates a GCM cipher using the AES block cipher.
- The resulting GCM cipher is used for secure encryption and decryption.

Example Usage:
gcmCipher, err := newGcm("yourSecretKey")

	if err != nil {
	    fmt.Println("GCM creation failed:", err)
	    return
	}
*/
func newGcm(secretKey string) (gcm cipher.AEAD, err error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}

	gcm, err = cipher.NewGCM(block)
	if err != nil {
		return
	}

	return
}
