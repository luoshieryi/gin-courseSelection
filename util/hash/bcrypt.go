package hash

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	cost int
}

func NewHash() (b *Bcrypt) {
	return new(Bcrypt)
}

//Make 加密方法
func (b *Bcrypt) Make(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.cost)
}

//Check 检查方法
func (b *Bcrypt) Check(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
