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
	PublicKey               *paseto.V4AsymmetricPublicKey
	JTI                     *uuid.UUID
	Signed, Email, Whatsapp *string
}
type SuperAdminPayload struct {
	JTI                                uuid.UUID
	PublicKey, Signed, Email, Whatsapp string
	Status                             string
}

func New() *Paseto {
	return &Paseto{}
}

func (p *Paseto) SetSuperAdminToken(payload SuperAdminPayload) (*Paseto, error) {
	p.JTI = &payload.JTI
	p.Email = &payload.Email
	p.Whatsapp = &payload.Whatsapp
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	err := token.Set(JTI_KEY, p.JTI)
	if err != nil {
		return nil, errors.New("error jti id not set")
	}
	err = token.Set(EMAIL, p.Email)
	if err != nil {
		return nil, errors.New("error jti e not set")
	}
	err = token.Set(WHATSAPP, p.Whatsapp)
	if err != nil {
		return nil, errors.New("error jti w not set")
	}

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public()
	signed := token.V4Sign(secretKey, nil)
	return &Paseto{
		PublicKey: &publicKey,
		Signed:    &signed,
	}, nil
}

func (p *Paseto) ClaimSuperAdmin(payload SuperAdminPayload) error {
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
	if payload.Email != mapVal[EMAIL].(string) {
		return errors.New("email invalid")
	}
	if payload.Whatsapp != mapVal[WHATSAPP].(string) {
		return errors.New("whatsapp invalid")
	}

	return nil
}
