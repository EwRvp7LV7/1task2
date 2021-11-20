package module

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	// "strings"
)

func MainConsole() {

	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	//  Bold = "\x1b[1m"
	//     Dim = "\x1b[2m"
	//     Italic = "\x1b[3m"
	//     Underlined = "\x1b[4m"
	//     Blink = "\x1b[5m"
	//     Reverse = "\x1b[7m"
	//     Hidden = "\x1b[8m"
	//     # Reset part
	//     Reset = "\x1b[0m"
	//     Reset_Bold = "\x1b[21m"
	//     Reset_Dim = "\x1b[22m"
	//     Reset_Italic = "\x1b[23m"
	//     Reset_Underlined = "\x1b[24"
	//     Reset_Blink = "\x1b[25m"
	//     Reset_Reverse = "\x1b[27m"
	//     Reset_Hidden = "\x1b[28m"
	//  # Foreground
	//     F_Default = "\x1b[39m"
	//     F_Black = "\x1b[30m"
	//     F_Red = "\x1b[31m"
	//     F_Green = "\x1b[32m"
	//     F_Yellow = "\x1b[33m"
	//     F_Blue = "\x1b[34m"
	//     F_Magenta = "\x1b[35m"
	//     F_Cyan = "\x1b[36m"
	//     F_LightGray = "\x1b[37m"
	//     F_DarkGray = "\x1b[90m"
	//     F_LightRed = "\x1b[91m"
	//     F_LightGreen = "\x1b[92m"
	//     F_LightYellow = "\x1b[93m"
	//     F_LightBlue = "\x1b[94m"
	//     F_LightMagenta = "\x1b[95m"
	//     F_LightCyan = "\x1b[96m"
	//     F_White = "\x1b[97m"
	//     # Background
	//     B_Default = "\x1b[49m"
	//     B_Black = "\x1b[40m"
	//     B_Red = "\x1b[41m"
	//     B_Green = "\x1b[42m"
	//     B_Yellow = "\x1b[43m"
	//     B_Blue = "\x1b[44m"
	//     B_Magenta = "\x1b[45m"
	//     B_Cyan = "\x1b[46m"
	//     B_LightGray = "\x1b[47m"
	//     B_DarkGray = "\x1b[100m"
	//     B_LightRed = "\x1b[101m"
	//     B_LightGreen = "\x1b[102m"
	//     B_LightYellow = "\x1b[103m"
	//     B_LightBlue = "\x1b[104m"
	//     B_LightMagenta = "\x1b[105m"
	//     B_LightCyan = "\x1b[106m"
	//     B_White = "\x1b[107m"

	//  var myname string
	// fmt.Scanf("%s", &myname)
	// // fmt.Print("\033[32m") // FgGreen
	// fmt.Print("\x1b[32m") // FgGreen
	//     fmt.Print("password: ")
	// fmt.Print("\033[0m") // Reset
	// fmt.Print("\033[8m") // Hidden
	// fmt.Scan(&myname)
	// fmt.Print("\033[28m") // Reset_Hidden
	//     fmt.Println("Hello", myname)

	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("-> ")
	// 	text, _ := reader.ReadString('\n')
	// 	// convert CRLF to LF
	// 	text = strings.Replace(text, "\n", "", -1)

	// 	if strings.Compare("hi", text) == 0 {
	// 		fmt.Println("hello, Yourself")
	// 	}

	// }

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// // To create dynamic array
	// arr := make([]string, 0)
	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	//     fmt.Print("Enter Text: ")
	//     // Scans a line from Stdin(Console)
	//     scanner.Scan()
	//     // Holds the string that scanned
	//     text := scanner.Text()
	//     if len(text) != 0 {
	//         fmt.Println(text)
	//         arr = append(arr, text)
	//     } else {
	//         break
	//     }

	// }
	// // Use collected inputs
	// fmt.Println(arr)

}

func MainFlagsConsole() {

	// Базовые определения флагов доступны для типов
	// string, integer и boolean. Тут мы определяем
	// флаг-строку `word` со значением по-умолчанию
	// `"foo"` и коротким описанием. Функция `flag.String`
	// возвращает указатель на строку (а не саму строку);
	// мы посмотрим, как использовать этот указатель ниже.
	wordPtr := flag.String("word", "foo", "a string")

	// Тут мы декларируем флаги `numb` и `fork`, используя
	// аналогичный подход, как и с флагом `word`.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// Также флаги можно определять, используя
	// существующую переменную, объявленную ранее в программе.
	// Обратите внимание, что мы передаем в функцию указатель.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Как только все флаги определены, вызываем
	// `flag.Parse()`, чтобы выполнить непосредственно
	// разбор командной строки.
	flag.Parse()

	// Тут мы выводим на экран считанные значения флагов,
	// и оставшиеся аргументы командной строки. Обратите
	// внимание. что мы должны разыменовывать указатель,
	// `*wordPtr`, например, чтобы получить реальные значения.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
