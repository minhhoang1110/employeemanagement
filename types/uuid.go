package types

import (
	"encoding/base64"

	"github.com/google/uuid"
)

type UUIDBase64 uuid.UUID

// NewUUIDBase64 creates a new UUIDBase64 and return the pointer
func NewUUIDBase64() UUIDBase64 {
	u, _ := uuid.NewRandom()
	return *(*UUIDBase64)(&u)
}

// String returns the base64 encoded UUID
func (u UUIDBase64) String() string {
	if !u.IsValid() {
		return ""
	}
	s := base64.RawURLEncoding.EncodeToString(u[:])
	return s
}

// UUID returns the uuid.UUID representation
func (u UUIDBase64) UUID() uuid.UUID {
	return *(*uuid.UUID)(&u)
}

// MarshalText implements encoding.TextMarshaler.
func (u UUIDBase64) MarshalText() ([]byte, error) {
	if !u.IsValid() {
		return nil, nil
	}
	var result []byte
	result = make([]byte, 22)
	base64.RawURLEncoding.Encode(result, u[:])
	return result, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UUIDBase64) UnmarshalText(data []byte) error {
	defer func() {
		if r := recover(); r != nil {
			// Prevent panic from decoding error
		}
	}()
	var result []byte
	result = make([]byte, 16)
	base64.RawURLEncoding.Decode(result, data)
	copy(u[:], result)
	return nil
}

// IsValid validates the uuidBase64
func (u UUIDBase64) IsValid() bool {
	return u.IsValidUUID() && !u.IsNilUUID()
}

// IsValidUUID validates the uuid
func (u UUIDBase64) IsValidUUID() bool {
	return u.UUID().Variant() != uuid.Invalid
}

// IsNilUUID validates if the uuid is nil
func (u UUIDBase64) IsNilUUID() bool {
	return u.UUID() == uuid.Nil
}
