package router

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zchat-team/zchat-server/api/passport"
	"github.com/zchat-team/zchat-server/internal/middleware"
	sPassport "github.com/zchat-team/zchat-server/internal/service/passport"
	"github.com/zmicro-team/zmicro/core/log"
	"io/ioutil"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

func parseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}

func RegisterAPI(r *gin.Engine) {
	Swagger(r)

	skipPath := []string{
		"/api/v1/passport/login",
		"/api/v1/passport/smsLogin",
		"/api/v1/passport/sms",
		"/api/v1/passport/refreshToken",
		"/api/v1/test",
	}

	key, err := ioutil.ReadFile("conf/auth_key")
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	privKey, err := parseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Fatal(err)
	}

	_ = privKey

	g := r.Group("/api/v1",
		middleware.CheckLogin(skipPath...),
		//middleware.CheckSign(middleware.PrivKey(privKey)),
	)

	passport.RegisterPassportHTTPServer(g, sPassport.GetService())
}
