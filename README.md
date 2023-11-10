# File Encryption/Decryption tool

This is a small Go (Golang) application that provides functionality to encrypt and decrypt files using a specified key.

## Usage

To use the app, run the following command in your terminal:

```bash
go run main.go [flags]
```

### Flags

- **-encrypt**: Make encrypt (Specify either -encrypt or -decrypt, not both)
- **-decrypt**: Make decrypt (Specify either -encrypt or -decrypt, not both)
- **-src**: Source file (Provide the path to the source file)
- **-dest**: Destination file (Provide the path to the destination file)
- **-key**: Private key string

### Examples

#### Encrypting a File

```bash
go run main.go -encrypt -src input.txt -dest encrypted.txt -key 770A8A65DA156D24EE2A093277530142
```

or

```bash
./encryptor -encrypt -src input.txt -dest encrypted.txt -key 770A8A65DA156D24EE2A093277530142
```

#### Decrypting a File

```bash
go run main.go -decrypt -src encrypted.txt -dest decrypted.txt -key 770A8A65DA156D24EE2A093277530142
```

or

```bash
./encryptor -decrypt -src encrypted.txt -dest decrypted.txt -key 770A8A65DA156D24EE2A093277530142
```

## Important Notes

- Please provide both the source and destination file paths.
- **key** flag is mandatory and MUST be a 128-bit.

## Encryption/Decryption Logic

This small tool implements symmetric encryption and decryption using the Advanced Encryption Standard (AES) in Galois/Counter Mode (GCM). GCM is a mode of operation for symmetric key cryptographic block ciphers that provides authenticated encryption. Here's a breakdown of the encryption logic:

### Encryption Logic (Encrypt function):

1. **Key Generation:**
   - The `newGcm` function generates a new Galois/Counter Mode (GCM) cipher using the provided secret key.
   - It uses the Advanced Encryption Standard (AES) block cipher with a key derived from the provided secret key.

2. **Nonce Generation:**
   - A nonce (number used once) is generated with a size equal to the nonce size of the GCM cipher. Nonce is a unique value that should only be used once for a given key.
   - The nonce is created using a secure random number generator (`rand.Reader`).

3. **Encryption:**
   - The `Seal` method of the GCM cipher is used to encrypt the data.
   - The `Seal` method takes the nonce, additional authenticated data (nil in this case), and the plaintext data.
   - It returns the ciphertext along with the nonce as the first part of the ciphertext.

4. **Base64 Encoding:**
   - The resulting ciphertext, including the nonce, is encoded using Base64.URLEncoding.
   - This encoded string is returned as the encrypted data.

### Decryption Logic (Decrypt function):

1. **Key Generation:**
   - Similar to the encryption process, the `newGcm` function generates a new GCM cipher using the provided secret key.

2. **Base64 Decoding:**
   - The input encrypted data (received as a Base64-encoded string) is decoded using Base64.URLEncoding.
   - The decoded data is stored in the `buf` variable.

3. **Nonce Extraction:**
   - The first part of the `buf` is extracted to obtain the nonce.

4. **Ciphertext Extraction:**
   - The remaining part of the `buf` is considered as the ciphertext.

5. **Decryption:**
   - The `Open` method of the GCM cipher is used to decrypt the data.
   - It takes the nonce, additional authenticated data (nil in this case), and the ciphertext.
   - It returns the original plaintext.

### Additional Notes:

- The `newGcm` function initializes the GCM cipher with the AES block cipher and the provided secret key.
- The nonce size is determined by the GCM cipher and is retrieved using `gcm.NonceSize()`.

This encryption scheme provides confidentiality and integrity for the data being encrypted. The use of a nonce ensures that the same plaintext encrypted with the same key will produce different ciphertexts, preventing certain types of attacks.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
