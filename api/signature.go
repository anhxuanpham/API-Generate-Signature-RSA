package api

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func readPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("Failed To Parse PEM block containing the key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("Invalid Private Key Type")
	}

	return rsaPrivateKey, nil
}

func createSignature(privateKey *rsa.PrivateKey, data string) (string, error) {
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return signatureBase64, nil
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sign", func(c *gin.Context) {
		type RequestBody struct {
			DataToSign string `json:"data"`
		}

		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid or missing 'data' in request body",
			})
			return
		}

		privateKeyFilePath := "private_key.pem"
		privateKey, err := readPrivateKeyFromFile(privateKeyFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error reading private key",
			})
			return
		}

		signature, err := createSignature(privateKey, requestBody.DataToSign)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error creating signature",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"signature": signature,
		})
	})

	return r
}
