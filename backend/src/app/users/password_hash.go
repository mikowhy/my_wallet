package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = fmt.Errorf("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = fmt.Errorf("incompatible version of argon2")
)

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultParams = &Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

// GetParamsFromEnv attempts to load Argon2 parameters from environment variables.
// If an environment variable is not set or is invalid, it falls back to DefaultParams.
func GetParamsFromEnv() *Params {
	params := &Params{
		Memory:      DefaultParams.Memory,
		Iterations:  DefaultParams.Iterations,
		Parallelism: DefaultParams.Parallelism,
		SaltLength:  DefaultParams.SaltLength,
		KeyLength:   DefaultParams.KeyLength,
	}

	if memStr := os.Getenv("ARGON2_MEMORY"); memStr != "" {
		if mem, err := strconv.ParseUint(memStr, 10, 32); err == nil {
			params.Memory = uint32(mem)
		}
	}

	if iterStr := os.Getenv("ARGON2_ITERATIONS"); iterStr != "" {
		if iter, err := strconv.ParseUint(iterStr, 10, 32); err == nil {
			params.Iterations = uint32(iter)
		}
	}

	if parStr := os.Getenv("ARGON2_PARALLELISM"); parStr != "" {
		if par, err := strconv.ParseUint(parStr, 10, 8); err == nil {
			params.Parallelism = uint8(par)
		}
	}

	if saltLenStr := os.Getenv("ARGON2_SALT_LENGTH"); saltLenStr != "" {
		if saltLen, err := strconv.ParseUint(saltLenStr, 10, 32); err == nil {
			params.SaltLength = uint32(saltLen)
		}
	}

	if keyLenStr := os.Getenv("ARGON2_KEY_LENGTH"); keyLenStr != "" {
		if keyLen, err := strconv.ParseUint(keyLenStr, 10, 32); err == nil {
			params.KeyLength = uint32(keyLen)
		}
	}

	return params
}

func CreateHash(password string, params *Params) (string, error) {
	salt := make([]byte, params.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
	params, salt, hash, err := DecodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func DecodeHash(encodedHash string) (*Params, []byte, []byte, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	if vals[1] != "argon2id" {
		return nil, nil, nil, ErrInvalidHash
	}

	v := argon2.Version
	version, err := fmt.Sscanf(vals[2], "v=%d", &v)
	if err != nil || version != 1 {
		return nil, nil, nil, err
	}

	params := &Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}
