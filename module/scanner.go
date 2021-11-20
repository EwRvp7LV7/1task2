package module

import (
    "bufio"
    "fmt"
    "os"
)

//Есть три способа читать байты из потока
// https://pkg.go.dev/io
// https://pkg.go.dev/bufio#Scanner.Scan
// https://pkg.go.dev/bytes#IndexByte
func MainScan() {
    fd, err := os.Open("big.txt")
    if err != nil {
        panic(err)
    }
    defer fd.Close()

    n, err := fd.Seek(0, 0)
    if err != nil {
        panic(err)
    }
    fmt.Println("n =", n) // 0

    scanner := bufio.NewScanner(fd)
    //Вы можете увеличить размер буфера вашего сканера
    //buf := make([]byte, 0, 64*1024)
    //scanner.Buffer(buf, 1024*1024) //1024*1024 => 1mb max (you can change value here to read larger files
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        break
    }

    offset, err := fd.Seek(0, 1)
    if err != nil {
        panic(err)
    }
    fmt.Println("offset =", offset) //4096

    offsetreset, err := fd.Seek(offset, 0)
    if err != nil {
        panic(err)
    }
    fmt.Println("offsetreset =", offsetreset) //4096

    offset, err = fd.Seek(0, 1)
    if err != nil {
        panic(err)
    }
    fmt.Println("offset =", offset) //4096

}






