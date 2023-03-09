package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
)

//国密加解密

type SM2 struct {
	pubKey string
	priKey string
	mode   int

	pubMen *sm2.PublicKey
	priMen *sm2.PrivateKey
}

// 创建SM2对象.
// mode 加密模式 0:C1C3C2  1:C1C2C3
func NewSM2(pubKey, priKey string, mode int) (sm *SM2, err error) {

	var pubMen *sm2.PublicKey
	if pubKey != "" {
		pubMen, err = x509.ReadPublicKeyFromHex(pubKey)
		if err != nil {
			return
		}
	}
	var priMen *sm2.PrivateKey

	if priKey != "" {
		priMen, err = x509.ReadPrivateKeyFromHex(priKey)
		if err != nil {
			return
		}

	}

	return &SM2{
		pubKey: pubKey,
		priKey: priKey,
		mode:   mode,
		pubMen: pubMen,
		priMen: priMen,
	}, nil
}

// 公钥加密
func (o *SM2) Encrypt(en []byte) (rst []byte, err error) {

	if o.pubMen == nil {
		return nil, errors.New("no public key")
	}

	rst, err = sm2.Encrypt(o.pubMen, en, nil, o.mode)
	return
}

// 私钥解密
func (o *SM2) Decrypt(de []byte) (rst []byte, err error) {

	if o.priMen == nil {
		return nil, errors.New("no private key")
	}

	rst, err = sm2.Decrypt(o.priMen, de, o.mode)

	return
}

/*
// 国密 公钥加密
// mode 加密模式 0:C1C3C2  1:C1C2C3
func Sm2Encode(endstr string, publickKey string, mode int) (result string, err error) {

	var smmode int = sm2.C1C3C2
	if mode == 1 {
		smmode = sm2.C1C2C3
	} else if mode == 0 {
		smmode = sm2.C1C3C2
	} else {
		return "", errors.New("mode error")
	}

	pubMen, err := x509.ReadPublicKeyFromHex(publickKey)

	if err != nil {
		return
	}

	msg := []byte(endstr)
	ciphertxt, err := sm2.Encrypt(pubMen, msg, nil, smmode)
	if err != nil {
		return
	}
	result = hex.EncodeToString(ciphertxt)

	return

}

// 国密 私解密
// mode 加密模式 0:C1C3C2  1:C1C2C3
func Sm2Decode(ciphertxt string, privateKey string, mode int) (result string, err error) {

	var smmode = sm2.C1C3C2
	if mode == 1 {
		smmode = sm2.C1C2C3
	} else if mode == 0 {
		smmode = sm2.C1C3C2
	} else {
		return "", errors.New("mode error")
	}

	privateKeys, err := x509.ReadPrivateKeyFromHex(privateKey)
	if err != nil {
		return
	}

	ciphertxtbyte, err := hex.DecodeString(ciphertxt)
	if err != nil {
		return
	}

	ciphertxt3, err := sm2.Decrypt(privateKeys, ciphertxtbyte, smmode)
	if err != nil {
		return
	}

	result = string(ciphertxt3)

	return
}
*/
// 生成sm2密钥对
func GenerateSm2Key() (priKey, pubKey string, err error) {
	var priv *sm2.PrivateKey
	priv, err = sm2.GenerateKey(rand.Reader)

	if err != nil {
		return
	}

	priKey = x509.WritePrivateKeyToHex(priv)
	pubKey = x509.WritePublicKeyToHex(&priv.PublicKey)

	return
}

// sm3生成签名
func Sm3Encode(enStr string) []byte {
	return sm3.Sm3Sum([]byte(enStr))
}

type SM4 struct {
	pwd          string //128位，16字节
	iv           string
	block        cipher.Block
	blockmode_en cipher.BlockMode
	blockmode_de cipher.BlockMode
}

// 创建对称加密对象
//
// pwd: 密钥128位，16字节;
// iv: 初始化向量,位数与密钥相同
func NewSM4(pwd, iv string) (*SM4, error) {

	block, err := sm4.NewCipher([]byte(pwd))
	if err != nil {
		return nil, err
	}

	ivBytes := []byte(iv)
	blockmode_en := cipher.NewCBCEncrypter(block, ivBytes)
	blockmode_de := cipher.NewCBCDecrypter(block, ivBytes)

	return &SM4{
		pwd:          pwd,
		iv:           iv,
		block:        block,
		blockmode_en: blockmode_en,
		blockmode_de: blockmode_de,
	}, nil
}

// 加密
func (o *SM4) Encrypt(src []byte) []byte {
	src = paddingText(src, o.block.BlockSize())
	// blockmode := cipher.NewCBCEncrypter(o.block, o.ivBytes)
	dst := make([]byte, len(src))
	o.blockmode_en.CryptBlocks(dst, src)
	return dst
}

// 解密
func (o *SM4) Decrypt(src []byte) []byte {
	//创建解密模式
	dst := make([]byte, len(src))
	//解密
	o.blockmode_de.CryptBlocks(dst, src)
	//去除填充
	dst = unPaddingText(dst)
	return dst

}

// 给最后一组数据填充至blockSize字节
func paddingText(src []byte, blockSize int) []byte {

	//求出最后一个分组需要填充的字节数
	padding := blockSize - len(src)%blockSize
	//创建新的切片，切片字节数为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//将新创建的切片和带填充的数据进行拼接
	nextText := append(src, padText...)
	return nextText

}

// 取出数据尾部填充的赘余字符
func unPaddingText(src []byte) []byte {
	//获取待处理数据长度
	len := len(src)
	//取出最后一个字符
	num := int(src[len-1])
	newText := src[:len-num]
	return newText
}
