/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 03:18:44
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 03:20:01
 * @Description:
 * @FilePath: /spike_system/utils/shr256.go
 * @GitHub: https://github.com/liuyuchen777
 */
package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSHA256HashCode(msg string) string {
	message := []byte(msg)
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)

	return hashCode
}
