package module

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/EwRvp7LV7/1test/cipherst"
)

// var ciphertext1 []byte

// https://pkg.go.dev/crypto/cipher
// https://pkg.go.dev/crypto/aes

func ExampleNewCBCEncrypterSt() []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	password := []byte("MyDarkSecret")
	key := md5.Sum(password)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	// fmt.Println(len(key)) //16 bytes
	// key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("Вице-премьер РФ Татьяна Голикова сообщила об осложнении ситуации с распространением коронавируса в десяти регионах России, сообщает ТАСС")

	buf := bytes.NewReader(plaintext)
	bufOut := new(bytes.Buffer)
	// io.CopyN(bufOut, rand.Reader, aes.BlockSize)
	bufOut.Write([]byte{109, 111, 28, 0, 190, 156, 201, 84, 60, 137, 125, 147, 240, 80, 79, 18})
	modeSt := cipherst.NewCBCEncrypter(block, bufOut.Bytes()) //return BlockMode interface cbc

	for {
		mySlice := make([]byte, aes.BlockSize)
		//copy(arr, make([]int, len(arr)))
		//TODO провести бенчмарк экономней каждый раз создавать новый массив или наполнять его пустыми значениями

		n, err := buf.Read(mySlice)

		if err != nil {
			if err == io.EOF {
				//TODO if file end x00 or x01 it will cut, add one more block
				break
			}
			fmt.Println(n, string(mySlice), mySlice, err)
			panic(err)
		}
		if n < aes.BlockSize {
			if n > 0 {
				mySlice[n] = 1
			}
			bufOut.Write(modeSt.CryptBlocksSt(mySlice))
			//bufOut.Write(mySlice)
			break
		}

		bufOut.Write(modeSt.CryptBlocksSt(mySlice))
		//bufOut.Write(mySlice)
	}

	// fmt.Println("bufOut", bufOut.Len()%aes.BlockSize, bufOut.String(), bufOut.Bytes())
	fmt.Printf("bufOut\n%x\n\n", bufOut.Bytes())
	return bufOut.Bytes()

}

func ExampleNewCBCEncrypter() []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	password := []byte("MyDarkSecret")
	key := md5.Sum(password)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	// fmt.Println(len(key)) //16 bytes
	// key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("Вице-премьер РФ Татьяна Голикова сообщила об осложнении ситуации с распространением коронавируса в десяти регионах России, сообщает ТАСС")
	// fmt.Println("plaintext", plaintext)//[208 146 208 184 209 134 208 181 45 208 ...162 208 144 208 161 208 161]
	// var b bytes.Buffer
	// arr := make([]byte, aes.BlockSize)
	// bufOut := new(bytes.Buffer)
	// bufOut := make([]byte, aes.BlockSize)

	// for {
	// 	// bytes, err := io.CopyN(os.Stdout, buf, aes.BlockSize)
	// 	bytes, err := io.CopyN(bufOut, buf, aes.BlockSize)
	// 	// If error is not nil then panics
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println(bytes, bufOut.String(), err)
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("hello", bytes, bufOut.String())
	// }

	buffer := make([]byte, int(math.Ceil(float64(len(plaintext))/aes.BlockSize)*aes.BlockSize))

	copy(buffer, plaintext)
	//  fmt.Println(buffer, len(buffer)) //16 bytes

	if len(plaintext)%aes.BlockSize != 0 {
		buffer[len(plaintext)] = 1

	}

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(buffer)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	/* 	// The IV needs to be unique, but not secure. Therefore it's common to
	   	// include it at the beginning of the ciphertext.
	   	ciphertext := make([]byte, aes.BlockSize+len(buffer))

	   	iv := ciphertext[:aes.BlockSize]

	   	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	   		panic(err)
	   	}

	   	mode := cipherst.NewCBCEncrypter(block, iv)
	   	mode.CryptBlocks(ciphertext[aes.BlockSize:], buffer) */

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(buffer))

	// iv := make([]byte, aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	// // iv=[]byte{109, 111, 28, 0, 190, 156, 201, 84, 60, 137, 125, 147, 240, 80, 79, 18}
	// //
	// for i, b := range []byte{109, 111, 28, 0, 190, 156, 201, 84, 60, 137, 125, 147, 240, 80, 79, 18} {
	// 	ciphertext[i] = b
	// }

	// fmt.Println(iv)
	// fmt.Println(ciphertext[:aes.BlockSize])

	// fmt.Println("hello1", len(ciphertext))
	mode := cipherst.NewCBCEncrypter(block, iv) //return BlockMode interface cbc
	mode.CryptBlocks(ciphertext[aes.BlockSize:], buffer)
	// fmt.Println("hello2", len(ciphertext))

	// 		fmt.Println( string(plaintext))
	// return  []byte{0}

	// fmt.Println("hello")
	// ciphertext1 := make([]byte, len(buffer))
	// mode.CryptBlocks(ciphertext1, buffer)
	// fmt.Printf("encoded\n%x\n\n", ciphertext1)
	// return ciphertext1

	// 	// It's important to remember that ciphertexts must be authenticated
	// 	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// 	// be secure.
	// 	// ciphertext1 = ciphertext

	fmt.Printf("encoded\n%x\n\n", ciphertext)
	return ciphertext

}

// func ExampleNewCBCDecrypterSt(ciphertext []byte) {
//Не получается потоком, потому что расщифровывается с конца к началу!
// }

func ExampleNewCBCDecrypter(ciphertext []byte) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	password := []byte("MyDarkSecret")

	key := md5.Sum(password)

	// key, _ := hex.DecodeString("6368616e676520746869732070617373")
	// ciphertext, _ := hex.DecodeString("73c86d43a9d700a253a96c85b0f6b03ac9792e0e757f869cca306bd3cba1c62b")

	// ciphertext, _ := hex.DecodeString("201a65c253e418fcc0f0a0754449bf0fa68a3d1fd86eb07209147262ad3accae705525b952d09915d39025d581c7ea0a4e7aab5f0a8831af9dfe424a4e56e1bd969941676e34cddb6ee2d1fa4fa4bff23bedfd956828f9be8fb32a8c7602edb25de4e24026ec910749b02e048255de9d86e8b52e73d051c74a58ff529eb00e66460ca7a90b88d05bf654c2c885d1c1c123a7ff44d90418613120b30f17cf13be6bc68b1b32fa1f7c702db93a8210abbdda60bcc5c900e619e73791b15c03fa0b25a9754a2f59985d4c07f7360dedbc43473c252c8fe345abf41801d1b4453dc7980e0e799b3e752559af34ad258701d1cd4b655c0de3f3273690632aee3f0eef05342d82e083432f987cca9750fc530a")
	// ciphertext := ciphertext1

	// fmt.Println(len(key))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipherst.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	ciphertext = bytes.TrimRight(ciphertext, "\x00")
	ciphertext = bytes.TrimRight(ciphertext, "\x01")
	// ciphertext = ciphertext[:len(ciphertext)-1] //именно 1!
	fmt.Printf("decoded\n%s!\n", ciphertext)
	// Output: exampleplaintext

}

// // fmt.Println(ComputeHmac256("Message", "secret"))
// // Пакет hmac реализует код аутентификации ключевых хеш-сообщений (HMAC),как определено в Федеральных стандартах США по обработке информации
// func ComputeHmac256(message string, secret string) string {
// 	key := []byte(secret)
// 	h := hmac.New(sha256.New, key)
// 	h.Write([]byte(message))
// 	return base64.StdEncoding.EncodeToString(h.Sum(nil))
// }

// func ExampleStreamWriter() {
// 	// Load your secret key from a safe place and reuse it across multiple
// 	// NewCipher calls. (Obviously don't use this example key for anything
// 	// real.) If you want to convert a passphrase to a key, use a suitable
// 	// package like bcrypt or scrypt.
// 	key, _ := hex.DecodeString("6368616e676520746869732070617373")

// 	bReader := bytes.NewReader([]byte("some secret text"))

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// If the key is unique for each ciphertext, then it's ok to use a zero
// 	// IV.
// 	var iv [aes.BlockSize]byte
// 	stream := cipher.NewOFB(block, iv[:])

// 	var out bytes.Buffer

// 	writer := &cipher.StreamWriter{S: stream, W: &out}
// 	// Copy the input to the output buffer, encrypting as we go.
// 	if _, err := io.Copy(writer, bReader); err != nil {
// 		panic(err)
// 	}

// 	// Note that this example is simplistic in that it omits any
// 	// authentication of the encrypted data. If you were actually to use
// 	// StreamReader in this manner, an attacker could flip arbitrary bits in
// 	// the decrypted result.

// 	fmt.Printf("%x\n", out.Bytes())
// 	// Output: cf0495cc6f75dafc23948538e79904a9
// }

// func ExampleStreamReader() {
// 	// Load your secret key from a safe place and reuse it across multiple
// 	// NewCipher calls. (Obviously don't use this example key for anything
// 	// real.) If you want to convert a passphrase to a key, use a suitable
// 	// package like bcrypt or scrypt.
// 	key, _ := hex.DecodeString("6368616e676520746869732070617373")

// 	encrypted, _ := hex.DecodeString("cf0495cc6f75dafc23948538e79904a9")
// 	bReader := bytes.NewReader(encrypted)

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// If the key is unique for each ciphertext, then it's ok to use a zero
// 	// IV.
// 	var iv [aes.BlockSize]byte
// 	stream := cipher.NewOFB(block, iv[:])

// 	reader := &cipher.StreamReader{S: stream, R: bReader}
// 	// Copy the input to the output stream, decrypting as we go.
// 	if _, err := io.Copy(os.Stdout, reader); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println()
// 	// Note that this example is simplistic in that it omits any
// 	// authentication of the encrypted data. If you were actually to use
// 	// StreamReader in this manner, an attacker could flip arbitrary bits in
// 	// the output.

// 	// Output: some secret text
// }

func CBCEncrypter(password string, sl []byte) ([]byte, error) {

	key := md5.Sum([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	//требует чтобы размер файла был кратным 16 байтам поэтому в конце дописываем единицу и нули,
	//которые обрежем в Decrypter
	sl16 := make([]byte, int(math.Ceil(float64(len(sl))/aes.BlockSize)*aes.BlockSize)) //%16 bytes
	copy(sl16, sl)
	sl16[len(sl)] = 1

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(sl16)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	ciphertext := make([]byte, aes.BlockSize+len(sl16))
	iv := ciphertext[:aes.BlockSize]

	// iv :=  make([]byte, aes.BlockSize)

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, (err)
	}

	mode := cipherst.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext[aes.BlockSize:], sl16)

	// ciphertext := make([]byte, len(sl16))
	// mode.CryptBlocks(ciphertext, sl16)

	return ciphertext, nil
}

func CBCDecrypter(password string, ciphertext []byte) ([]byte, error) {
	key := md5.Sum([]byte(password))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, (err)
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipherst.NewCBCDecrypter(block, iv)

	// without return
	// // CryptBlocks can work in-place if the two arguments are the same.
	// mode.CryptBlocks(ciphertext, ciphertext)
	// //обрезаем см начало CBCEncrypter
	// ciphertext = bytes.TrimRight(ciphertext, "\x00")
	// ciphertext = ciphertext[:len(ciphertext)-1] //именно 1!

	ciphertextOut := make([]byte, len(ciphertext))
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertextOut, ciphertext)
	//обрезаем см начало CBCEncrypter
	ciphertextOut = bytes.TrimRight(ciphertextOut, "\x00")
	ciphertextOut = ciphertextOut[:len(ciphertextOut)-1] //именно 1!
	return ciphertextOut, nil
}

// func checkFileMD5Hash(path string) {

// 	hashFile, _ := os.Open(path)
// 	defer hashFile.Close()
// 	h := md5.New()
// 	if _, err := io.Copy(h, hashFile); err != nil {
// 		log.Println(err)
// 	}
// 	statsh, _ := hashFile.Stat()
// 	log.Printf("File %s\nmd5 hash: %x\n", statsh.Name(), h.Sum(nil))
// }
