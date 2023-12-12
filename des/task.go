package des

import (
	"os"
	"strings"

	"github.com/samuelsih/chat-encrypt/des/table"
)

type Encryption int

const (
	EncryptionDebug Encryption = iota
	EncryptionHex
	EncryptionBase64
)

type Decryption int

const (
	DecryptionDebug Decryption = iota
	DecryptionHex
	DecryptionBase64
)

func blockKeys() []string {
	envKeys := os.Getenv("DES_KEY")
	return generateKeys(envKeys)
}

func split(s string, size int) []string {
	if size >= len(s) {
		return []string{s}
	}

	var blocks []string
	block := make([]rune, size)
	len := 0

	for _, r := range s {
		block[len] = r
		len++
		if len == size {
			blocks = append(blocks, string(block))
			len = 0
		}
	}

	if len > 0 {
		blocks = append(blocks, string(block[:len]))
	}

	return blocks
}

func Encrypt(message string, msgType Encryption) string {
	messageBlock := split(message, 8)
	keys := blockKeys()

	var res string

	for _, msg := range messageBlock {
		for len(msg) < 8 {
			msg += "#"
		}

		binaryMessage := StrToBinary(msg)
		binaryBlock := strings.Split(binaryMessage, "")

		ip := table.InitialPermutation(binaryBlock)
		left, right := enciphering(ip, keys, false)
		leftRightCombined := append(right, left...)
		inverseIP := table.InverseInitialPermutation(leftRightCombined)

		if msgType != EncryptionDebug {
			res += ToString(strings.Join(inverseIP, ""))
		} else {
			res += strings.Join(inverseIP, "")
		}
	}

	if msgType == EncryptionHex {
		return Hex(res)
	}

	if msgType == EncryptionBase64 {
		return Base64(res)
	}

	return res
}

func Decrypt(message string, msgType Decryption) string {
	if msgType == DecryptionHex {
		message = HexToBinary(message)
	} else if msgType == DecryptionBase64 {
		message = Base64ToBinary(message)
	}

	messageBlock := split(message, 64)
	keys := blockKeys()

	var res string

	for _, msg := range messageBlock {
		binaryBlock := strings.Split(msg, "")

		ip := table.InitialPermutation(binaryBlock)
		left, right := enciphering(ip, keys, true)
		leftRightCombined := append(right, left...)
		inverseIP := table.InverseInitialPermutation(leftRightCombined)

		res += ToString(strings.Join(inverseIP, ""))
	}

	hasExtraPadding, index := indexStartExtraPadding(res)
	if hasExtraPadding {
		return cutStringFrom(res, index)
	}

	return res
}

func indexStartExtraPadding(str string) (bool, int) {
	index := -1

	for i := len(str) - 1; i >= 0; i-- {
		s := str[i]

		if s != '#' {
			break
		}

		index = i
	}

	if index == -1 {
		return false, index
	}

	return true, index
}

func cutStringFrom(str string, index int) string {
	sb := strings.Builder{}

	for i := 0; i < index; i++ {
		s := str[i]

		sb.WriteRune(rune(s))
	}

	return sb.String()
}
