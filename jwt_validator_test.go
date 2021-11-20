package shield

import (
	"testing"
)

const es256pubKey = "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXFaTbtCtNYsU5/x1ogVlX0USVHRW\ng3wyaoKInlMdvAWpFWga4Di2zdzFOIdj6P7VSmW9Zsf3Dl3WrQEFmgzY0w==\n-----END PUBLIC KEY-----"

func TestValidator_Valid(t *testing.T) {
	var tests = []struct {
		name    string
		token   string
		pubKey  []byte
		keyType string
		err     error
	}{
		{
			name:    "1. ES256 - token and key mismatch, expect invalid token error.",
			token:   "eyJhbGciOiJFUzI1NiJ9.eyJ0ZXN0IjoidGVzdCJ9.w-rMmS-endA4lG0yb68xvIlrZfLDU8oC5Vt6DUsOGmoxpCmhqRnC_fCUSnd1Rdcc999yMd0MtiKlObAWdfGuNw",
			pubKey:  []byte("-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEu0tFkmRlStZkoXEmCbyHshMQwo5M\nd9L0nn1dkNNkXkpVYK9UqUH1a8rTTeslh8Na7xypvUQ8PZKzF62YC778ug==\n-----END PUBLIC KEY-----"),
			keyType: "ES256",
			err:     ErrSignatureInvalid,
		},
		{
			name:    "2. ES256 - token and key is valid, expect error equal nil",
			token:   "eyJhbGciOiJFUzI1NiJ9.eyJ0ZXN0IjoidGVzdCJ9.w-rMmS-endA4lG0yb68xvIlrZfLDU8oC5Vt6DUsOGmoxpCmhqRnC_fCUSnd1Rdcc999yMd0MtiKlObAWdfGuNw",
			pubKey:  []byte("-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE71mdiBfDaCKeKtPCjjjobqTtOh4D\nGFm/xIaElV6EfU1J4nbqplH+1i9qlORxDWihLrUz3szswIlGbBUbOsfpUA==\n-----END PUBLIC KEY-----"),
			keyType: "ES256",
			err:     nil,
		},
		{
			name:    "3. ES256 - token and key is valid, check exp claim",
			token:   "eyJhbGciOiJFUzI1NiJ9.eyJ0ZXN0IjoidGVzdCIsImV4cCI6MTYzMjE0MTE3NH0.thv_nAd5O-VkSLg-YGC7bzBcHrYHGwphe9kHF1bYieHznFQRPvu1qORxfmZtdkDjJfRN5XCMac5meRPtF8oqwA",
			pubKey:  []byte("-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE71mdiBfDaCKeKtPCjjjobqTtOh4D\nGFm/xIaElV6EfU1J4nbqplH+1i9qlORxDWihLrUz3szswIlGbBUbOsfpUA==\n-----END PUBLIC KEY-----"),
			keyType: "ES256",
			err:     ErrTokenExpired,
		},
		{
			name:    "4. ES256 - AM test token and key",
			token:   "eyJhbGciOiJFUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJJX3NsQlhJMWI4NzcwQ3B1c1RIRHVJTUpYYS00dkI4R3pnYzFPQTFCZjdZIn0.eyJleHAiOjE2MzI5OTE1ODYsImlhdCI6MTYzMjk4Nzk4NiwianRpIjoiYjY2OGFkZTgtOTlkNi00NDdkLWEyMjktNmE5NWRlZTZjY2RmIiwiaXNzIjoiaHR0cHM6Ly9rZXljbG9hay10ZXN0Mi5hbGZhLWJhbmsua3o6ODQ0NS9hdXRoL3JlYWxtcy9tb2JpbGUiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZjoyYjM4NWE3OS1iY2I2LTQ3NTQtODUxYS0wNDk1OTRhZjg5MTA6MTA2NzU4IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibW9iaWxlX2FwcCIsInNlc3Npb25fc3RhdGUiOiJiNzUzZTEwZC05NWQwLTQ2ODAtOGRkNC0wNGQxMmY2Y2E5OWQiLCJhY3IiOiIxIiwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiZW1haWwgcGhvbmUiLCJzaWQiOiJiNzUzZTEwZC05NWQwLTQ2ODAtOGRkNC0wNGQxMmY2Y2E5OWQiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6IjcwMjcxMTcyNzMiLCJlbWFpbCI6ImFzZGFhMUBhc2RhZC5jb20ifQ.f7n7j8JdjU3_s2ogT-uoOw1UzCAOS1MXNOwmy0ryDLI8V0HSv91nX9tdztiZQMZLAemfNgkLHXKBNWRb27aT7w",
			pubKey:  []byte(es256pubKey),
			keyType: "ES256",
			err:     ErrTokenExpired,
		},
		{
			name:    "5. ES256 - AM token is good",
			token:   "eyJhbGciOiJFUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJJX3NsQlhJMWI4NzcwQ3B1c1RIRHVJTUpYYS00dkI4R3pnYzFPQTFCZjdZIn0.eyJleHAiOjE2MzI5OTE1ODYsImlhdCI6MTYzMjk4Nzk4NiwianRpIjoiYjY2OGFkZTgtOTlkNi00NDdkLWEyMjktNmE5NWRlZTZjY2RmIiwiaXNzIjoiaHR0cHM6Ly9rZXljbG9hay10ZXN0Mi5hbGZhLWJhbmsua3o6ODQ0NS9hdXRoL3JlYWxtcy9tb2JpbGUiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZjoyYjM4NWE3OS1iY2I2LTQ3NTQtODUxYS0wNDk1OTRhZjg5MTA6MTA2NzU4IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibW9iaWxlX2FwcCIsInNlc3Npb25fc3RhdGUiOiJiNzUzZTEwZC05NWQwLTQ2ODAtOGRkNC0wNGQxMmY2Y2E5OWQiLCJhY3IiOiIxIiwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiZW1haWwgcGhvbmUiLCJzaWQiOiJiNzUzZTEwZC05NWQwLTQ2ODAtOGRkNC0wNGQxMmY2Y2E5OWQiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6IjcwMjcxMTcyNzMiLCJlbWFpbCI6ImFzZGFhMUBhc2RhZC5jb20ifQ.f7n7j8JdjU3_s2ogT-uoOw1UzCAOS1MXNOwmy0ryDLI8V0HSv91nX9tdztiZQMZLAemfNgkLHXKBNWRb27aT7w",
			pubKey:  []byte(es256pubKey),
			keyType: "ES256",
			err:     ErrTokenExpired,
		},
	}
	for _, tc := range tests {
		v, err := NewValidator(tc.pubKey, tc.keyType)
		if err != nil {
			t.Errorf("%s: failed to parse key: %s", tc.name, err.Error())
			continue
		}
		_, err = v.Valid(tc.token)
		if err == tc.err {
		} else if err != nil {
			t.Errorf("%s: %s", tc.name, err.Error())
		} else {
			t.Errorf("%s", tc.name)
		}
	}
}
