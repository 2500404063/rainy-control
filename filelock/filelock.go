package filelock

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"os"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

//AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func File_0x01(file string, pwd string, extentKB int) error {
	if len(pwd) != 16 {
		return errors.New("the length of pwd must be 16")
	}
	fd_reader, err_reader := os.OpenFile(file, os.O_RDONLY, 0666)
	fd_writer, err_writer := os.OpenFile(file, os.O_WRONLY, 0666)
	if err_reader != nil {
		return err_reader
	}
	if err_writer != nil {
		return err_writer
	}
	var buf = make([]byte, 1024*extentKB-1)
	len, err := fd_reader.Read(buf)
	if err != nil {
		return err
	}
	encrypted, err2 := AesEncrypt(buf[:len], []byte(pwd))
	if err2 != nil {
		return err2
	}
	fd_writer.Write(encrypted)
	fd_writer.Close()
	fd_reader.Close()
	return nil
}

func File_0x02(file string, pwd string, extentKB int) error {
	if len(pwd) != 16 {
		return errors.New("the length of pwd must be 16")
	}
	fd_reader, err_reader := os.OpenFile(file, os.O_RDONLY, 0666)
	fd_writer, err_writer := os.OpenFile(file, os.O_WRONLY, 0666)
	if err_reader != nil {
		return err_reader
	}
	if err_writer != nil {
		return err_writer
	}
	var buf = make([]byte, 1024*extentKB)
	len, err := fd_reader.Read(buf)
	if err != nil {
		return err
	}
	decrypted, err2 := AesDecrypt(buf[:len], []byte(pwd))
	if err2 != nil {
		return err2
	}
	fd_writer.Write(decrypted)
	fd_writer.Close()
	fd_reader.Close()
	return nil
}
