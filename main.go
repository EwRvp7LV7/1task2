package main

import (
	"log"

	"github.com/EwRvp7LV7/1test/module"
)

func run() (err error) {

	// module.Recursive()
	// module.Recursive3()
	// module.ArrayReverse()
	// module.SliceReverse()
	// module.StringReverse()
	// module.RunesReverse()
	// module.Closures()
	//module.MainSeekInFile()
	// module.MainStruct()
	// module.MainStructTags()
	// module.MainBcrypt()
	// module.MainMD5crypt()
	// module.MainConsole()
	// module.MainFlagsConsole()
	// module.MainIO()
	// module.ExampleStreamReader()
	// module.ExampleStreamWriter()
	// module.CheckArrSize()
	// module.CheckStruct()
	module.MainStructTags()



	// module.ExampleNewCBCEncrypter()
	// module.ExampleNewCBCDecrypter(module.ExampleNewCBCEncrypterSt())

	// module.InterfaceIOW()

	// module.Channels()
	// module.WorkersPool()

	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
