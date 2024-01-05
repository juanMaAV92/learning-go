package main

import "fmt"

type PasswordProtector struct {
	user          string
	password      string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(user, password string, hashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		password:      password,
		hashAlgorithm: hashAlgorithm,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type Sha256 struct{}

func (s *Sha256) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing with SHA256 for %s\n", p.password)
}

type MD5 struct{}

func (s *MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing with MD5 for %s\n", p.password)
}

func main() {
	sha := &Sha256{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("user", "password", sha)
	passwordProtector.Hash()

	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
