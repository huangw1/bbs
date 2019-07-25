/**
 * @Author: huangw1
 * @Date: 2019/5/27 22:15
 */

package crypto

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashByte), err
}

func Compare(hashByte, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashByte), []byte(password))
	if err != nil {
		return false
	}
	return true
}
