// this file was auto-generated by internal/cmd/gentypes/main.go: DO NOT EDIT

package jwa

import (
	"github.com/pkg/errors"
)

// KeyType represents the key type ("kty") that are supported
type KeyType string

// Supported values for KeyType
const (
	EC             KeyType = "EC"  // Elliptic Curve
	InvalidKeyType KeyType = ""    // Invalid KeyType
	OctetSeq       KeyType = "oct" // Octet sequence (used to represent symmetric keys)
	RSA            KeyType = "RSA" // RSA
)

// Accept is used when conversion from values given by
// outside sources (such as JSON payloads) is required
func (v *KeyType) Accept(value interface{}) error {
	var tmp KeyType
	switch x := value.(type) {
	case string:
		tmp = KeyType(x)
	case KeyType:
		tmp = x
	default:
		return errors.Errorf(`invalid type for jwa.KeyType: %T`, value)
	}
	switch tmp {
	case EC, OctetSeq, RSA:
	default:
		return errors.Errorf(`invalid jwa.KeyType value`)
	}

	*v = tmp
	return nil
}

// String returns the string representation of a KeyType
func (v KeyType) String() string {
	return string(v)
}
