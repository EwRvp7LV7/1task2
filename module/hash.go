package module

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash/fnv"

	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/ripemd160"
)

// https://betterprogramming.pub/a-short-guide-to-hashing-in-go-e8bb0173e97e

func MainBcrypt() {
	password := []byte("MyDarkSecret")

	// Hashing the password with the default cost of 10
	//хеш каждый раз рзный
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err == nil) // nil means it is a match
}

func MainMD5crypt() {
	password := []byte("MyDarkSecret")

	// fmt.Printf("%x\n", md4.New().Sum(password)) // import golang.org/x/crypto/md4
	fmt.Printf("md5 %x\n", md5.Sum(password))   // import crypto/md5
	fmt.Printf("sha1 %x\n", sha1.Sum(password)) // import crypto/sha1
	// fmt.Printf("%x\n", sha224.Sum(password)) // import crypto/sha224
	fmt.Printf("sha256 %x\n", sha256.Sum256(password)) // import crypto/sha256
	// fmt.Printf("%x\n", sha384.Sum(password)) // import crypto/sha512
	fmt.Printf("sha512 %x\n", sha512.Sum512(password))          // import crypto/sha512
	// fmt.Printf("ripemd160 %x\n", ripemd160.New().Sum(password)) // import golang.org/x/crypto/ripemd160 not reccomended
	//SHA-3 (Keccak — произносится как «кечак») — алгоритм хеширования переменной разрядности
	// fmt.Printf("%x\n", sha3_224.Sum(password))        // import golang.org/x/crypto/sha3
	// fmt.Printf("%x\n", sha3_256.Sum(password))        // import golang.org/x/crypto/sha3
	// fmt.Printf("%x\n", sha3_384.Sum(password))        // import golang.org/x/crypto/sha3
	// fmt.Printf("%x\n", sha3_512.Sum(password))        // import golang.org/x/crypto/sha3
	fmt.Printf("sha512 %x\n", sha512.Sum512_224(password)) // import crypto/sha512
	fmt.Printf("sha512 %x\n", sha512.Sum512_256(password)) // import crypto/sha512
	// fmt.Printf("%x\n", blake2s_256.Sum(password)) // import golang.org/x/crypto/blake2s
	// fmt.Printf("%x\n", blake2b_256.Sum(password)) // import golang.org/x/crypto/blake2b
	// fmt.Printf("%x\n", blake2b_384.Sum(password)) // import golang.org/x/crypto/blake2b
	// fmt.Printf("%x\n", blake2b_512.Sum(password)) // import golang.org/x/crypto/blake2b

	fmt.Printf("sha512 %x\n", fnv.New32a().Sum(password))          // import crypto/sha512

	h := fnv.New32a()
	h.Write(password)
	fmt.Println(h.Sum32())

}
