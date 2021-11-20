package module

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"
	// "github.com/goccy/go-json"
)

// const passwordStr = `Package sha3 implements the SHA-3 fixed-output-length hash functions and the SHAKE variable-output-length hash functions defined by FIPS-202.`

// //go:generate go test -v -test.bench=. hash_test.go
// func BenchmarkHash_md5(b *testing.B) {
// 	var out [16]byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = md5.Sum(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }

// func BenchmarkHash_sha1(b *testing.B) {
// 	var out [20]byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = sha1.Sum(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }
// func BenchmarkHash_sha256(b *testing.B) {
// 	var out [32]byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = sha256.Sum256(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }
// func BenchmarkHash_sha512_224(b *testing.B) {
// 	var out [28]byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = sha512.Sum512_224(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }
// func BenchmarkHash_sha512_256(b *testing.B) {
// 	var out [32]byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = sha512.Sum512_256(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }
// func BenchmarkHash_fnv_New32a(b *testing.B) {
// 	var out []byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = fnv.New32a().Sum(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }
// func BenchmarkHash_fnv_New64(b *testing.B) {
// 	var out []byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = fnv.New64().Sum(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }
// func BenchmarkHash_fnv_New128a(b *testing.B) {
// 	var out []byte
// 	password := []byte(passwordStr)
// 	for i := 0; i < b.N; i++ {
// 		out = fnv.New128a().Sum(password)
// 	}
// 	fmt.Printf("%x\n", out)
// }

func BenchmarkHash_gob(b *testing.B) {
	user1 := struct {
		uuid    string `validate:"-"`
		Name    string `validate:"presence,min=2,max=32" hasher:"+"`
		Email   string `validate:"email,required" hasher:"+"`
		Address []struct {
			City   string `validate:"-"`
			Street string `validate:"-" hasher:"+"`
		}
	}{
		uuid:  "a-b-c-d",
		Name:  "John Doe",
		Email: "john@example",
		Address: []struct {
			City   string `validate:"-"`
			Street string `validate:"-" hasher:"+"`
		}{{"Minsk", "Main Road"}, {"Novosibirsk", "Red Road"}},
	}

	for i := 0; i < b.N; i++ {
		var network bytes.Buffer // Stand-in for a network connection
		enc := gob.NewEncoder(&network)
		enc.Encode(user1)
		byteSl := md5.Sum(network.Bytes())
		//plug
		if i == 0 {
			fmt.Printf("gob:%s\n", network.String())
			fmt.Printf("md5:%x\n", byteSl)
		}

	}
}

func BenchmarkHash_Marshal(b *testing.B) {
	user1 := struct {
		uuid    string `validate:"-"`
		Name    string `validate:"presence,min=2,max=32" hasher:"+"`
		Email   string `validate:"email,required" hasher:"+"`
		Address []struct {
			City   string `validate:"-"`
			Street string `validate:"-" hasher:"+"`
		}
	}{
		uuid:  "a-b-c-d",
		Name:  "John Doe",
		Email: "john@example",
		Address: []struct {
			City   string `validate:"-"`
			Street string `validate:"-" hasher:"+"`
		}{{"Minsk", "Main Road"}, {"Novosibirsk", "Red Road"}},
	}

	for i := 0; i < b.N; i++ {
		bt, _ := json.Marshal(user1)
		byteSl := md5.Sum(bt)
		//plug
		if i == 0 {
			fmt.Printf("Marshal:%s\n", bt)
			fmt.Printf("md5:%x\n", byteSl)
		}

	}
}
