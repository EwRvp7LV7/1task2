package module

import (

	// "fmt"
	// "crypto/md5"
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"log"
	"os"
	"path/filepath"
)

func MainIO() {
	fname := "ttt.txt"
	// fmt.Scanf("%s", &fname)

	abs, _ := filepath.Abs("testdir")
	inputFile, err := os.Open(abs + "/" + fname)
	if err != nil {
		log.Println(err.Error())
		// conn.Write([]byte(err.Error()))
		return
	}
	defer inputFile.Close()

	stats, _ := inputFile.Stat()
	//send file Size
	//по результатам тестов при минимальной паузе Read расклеивает пакеты на раздельные
	log.Println("stats.Size()", stats.Size())
	// time.Sleep(time.Millisecond * time.Duration(100))
	// conn.Write([]byte(strconv.FormatInt(stats.Size(), 10)))
	// time.Sleep(time.Millisecond * time.Duration(100))

	// buffer := make([]byte, 1024)
	// for {
	// 	_, err := inputFile.Read(buffer)
	// 	if err == io.EOF {
	// 		// если есть буфер он дописывает его нулями, и только в следующий раз объявляет EOF
	// 		// поэтому здесь не требуется дочитывание
	// 		break
	// 	}
	// 	os.Stdout.Write(buffer)
	// }

	io.Copy(os.Stdout, inputFile)

	// 		buf := bufio.NewScanner(conn)
	// 	for buf.Scan() {
	// buf.Text()

	// 	}

}

// реализация io.Writer
//https://medium.com/@as27/a-simple-beginners-tutorial-to-io-writer-in-golang-2a13bfefea02
// The io package specifies the io.Reader interface, which represents the read end of a stream of data.
// https://tour.golang.org/methods/21
type Person struct {
	Id   int
	Name string
	Age  int
}

func (p *Person) Read(w io.Writer) {
	b, _ := json.Marshal(*p)
	// Inside our function we just write into the io.Writer
	// We don't care about which writer we use
	w.Write(b)
}

func InterfaceIOW() {
	// var out bytes.Buffer
	// json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	// out.WriteTo(os.Stdout)

	me := Person{
		Id:   1,
		Name: "Me",
		Age:  64}

	// For testing we use a bytes.Buffer
	// That type has an implementaion of the io.Writer
	var b bytes.Buffer

	// We call the write Method with our io.Writer
	// For testing we don't write our data into a
	// file. Just into a buffer.
	me.Read(&b)

	// Let's look what the Write method wrote into our
	// buffer
	fmt.Printf("%s\n", b.String())

}
