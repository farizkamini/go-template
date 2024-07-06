package pst

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/gofrs/uuid/v5"
)

const (
	JTI_KEY  = "jti"
	EMAIL    = "email"
	WHATSAPP = "whatsapp"
)

type Paseto struct {
	PublicKey *paseto.V4AsymmetricPublicKey
	JTI       *uuid.UUID
	Signed    *string
}
type Payload struct {
	JTI               uuid.UUID
	PublicKey, Signed string
}

func New() *Paseto {
	return &Paseto{}
}

func (p *Paseto) SetToken(payload Payload) (*Paseto, error) {
	p.JTI = &payload.JTI
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	err := token.Set(JTI_KEY, p.JTI)
	if err != nil {
		return nil, errors.New("error jti id not set")
	}

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public()
	signed := token.V4Sign(secretKey, nil)
	return &Paseto{
		PublicKey: &publicKey,
		Signed:    &signed,
	}, nil
}

func (p *Paseto) Claim(payload Payload) error {
	parser := paseto.NewParser()
	parser.AddRule(paseto.IdentifiedBy(payload.JTI.String()))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))
	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(payload.PublicKey)
	if err != nil {
		return errors.New("public key not match")
	}

	parsedToken, errToken := parser.ParseV4Public(publicKey, payload.Signed, nil)
	if errToken != nil {
		return errors.New("token not match")
	}
	mapVal := parsedToken.Claims()
	if string(payload.JTI.String()) != mapVal[JTI_KEY].(string) {
		return errors.New("jti invalid")
	}

	return nil
}
