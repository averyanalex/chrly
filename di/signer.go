package di

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/elyby/chrly/http"
	. "github.com/elyby/chrly/signer"
	"strings"

	"github.com/goava/di"
	"github.com/spf13/viper"
)

var signer = di.Options(
	di.Provide(newTexturesSigner,
		di.As(new(http.TexturesSigner)),
	),
)

func newTexturesSigner(config *viper.Viper) (*Signer, error) {
	// TODO: add CHANGELOG and README entries about this variable
	// TODO: rename param variable
	keyStr := config.GetString("textures.signer.pem")
	if keyStr == "" {
		return nil, errors.New("texturesSigner.pem must be set in order to sign textures")
	}

	var keyBytes []byte
	if strings.HasPrefix(keyStr, "base64:") {
		base64Value := keyStr[7:]
		decodedKey, err := base64.URLEncoding.DecodeString(base64Value)
		if err != nil {
			return nil, err
		}

		keyBytes = decodedKey
	} else {
		keyBytes = []byte(keyStr)
	}

	rawPem, _ := pem.Decode(keyBytes)
	key, err := x509.ParsePKCS1PrivateKey(rawPem.Bytes)
	if err != nil {
		return nil, err
	}

	return &Signer{Key: key}, nil
}
