package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/sha3"
	"hash"
	"io"
	"log"
	"os"
	"strings"
)

// getHashFunction 返回对应的哈希函数
func getHashFunction(algorithm string) (func() hash.Hash, error) {
	switch strings.ToLower(algorithm) {
	case "md5":
		return md5.New, nil
	case "sha-1":
		return sha1.New, nil
	case "sha-256":
		return sha256.New, nil
	case "sha-512":
		return sha512.New, nil
	case "sha3-256":
		return sha3.New256, nil
	case "sha3-512":
		return sha3.New512, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s. Valid algorithms are: md5, sha-1, sha-256, sha-512, sha3-256, sha3-512", algorithm)
	}
}

// calculateHash 计算文件的哈希值
func calculateHash(filePath, algorithm string) (string, error) {
	hashFunc, err := getHashFunction(algorithm)
	if err != nil {
		return "", err
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}
	// 延迟执行。在当前函数结束前执行这一句。
	defer file.Close()

	// 创建哈希实例
	hash := hashFunc()

	// 分块读取文件内容并计算哈希
	buf := make([]byte, 1024*1024) // 1MB 缓冲区
	for {
		n, err := file.Read(buf)
		if n > 0 {
			hash.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("error reading file: %v", err)
		}
	}

	// 获取计算出的哈希值
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// printHashResults 打印每个文件的哈希结果
func printHashResults(filePaths []string, algorithm string) {
	for _, filePath := range filePaths {
		hash, err := calculateHash(filePath, algorithm)
		if err != nil {
			log.Printf("Error calculating hash for %s: %v\n\n", filePath, err)
			continue
		}
		fmt.Printf("File: %s\nAlgo: %s\nHash: %s\n\n", filePath, algorithm, hash)
	}
}

func main() {
	// 检查命令行参数
	if len(os.Args) < 3 {
		fmt.Println("Usage: hash-it [algorithm] [file_path]...")
		fmt.Println("Supported algorithms: md5, sha-1, sha-256, sha-512, sha3-256, sha3-512")
		os.Exit(1)
	}

	// 获取算法和文件路径
	algorithm := os.Args[1]
	filePaths := os.Args[2:]

	// 打印哈希结果
	printHashResults(filePaths, algorithm)
}
