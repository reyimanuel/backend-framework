package token

import (
	"backend/config"
	"crypto/rsa"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var jwtConfig *jwtStruct

type jwtStruct struct {
	jwtLifeTime        uint
	jwtRefreshLifeTime uint
	privateKey        *rsa.PrivateKey
	publicKey         *rsa.PublicKey
}

func Load() {
	cfg := config.Get()

	// Gunakan langsung string PEM, bukan file path
	publicKeyPEM := []byte(cfg.PublicKeyPEM)
	privateKeyPEM := []byte(cfg.PrivateKeyPEM)

	// Parse Public Key
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}

	// Parse Private Key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	jwtConfig = &jwtStruct{
		jwtLifeTime:        cfg.AccessTokenLifeTime,
		jwtRefreshLifeTime: cfg.RefreshTokenLifeTime,
		publicKey:         publicKey,
		privateKey:        privateKey,
	}
}
