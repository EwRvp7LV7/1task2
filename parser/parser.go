package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Что мы хотим от вас получить
// Приложение, которое позволяет скачивать произвольную HTML-страницу
// посредством HTTP-запроса на жесткий диск компьютера и выдает статистику по
// количеству уникальных слов в консоль.
// Требования к приложению
// 1 В качестве входных данных в приложение принимает строку с адресом
// web-страницы. Пример входной строки: https://www.simbirsoft.com/
// 2 Приложение разбивает текст страницы на отдельные слова с помощью
// списка разделителей.
// Пример списка:
// {' ', ',', '.', '! ', '?','"', ';', ':', '[', ']', '(', ')', '\n', '\r', '\t'}
// 3 В качестве результата работы пользователь должен получить статистику по
// количеству уникальных слов в тексте.
// Пример:
// РАЗРАБОТКА -1
// ПРОГРАММНОГО - 2
// ОБЕСПЕЧЕНИЯ - 4
// 4 Приложение должно быть реализовано с помощью стандартных классов(не стоит добавлять
// собственные реализации списков, словарей и т.п.)
// 5 Приложение написано в соответствии с принципами ООП
// 6 Приложение написано на языке выбранного направления (Java, C#, Golang)
// Что нам понравится в приложении
// Для зачисления вполне достаточно, если ваше приложение работает и
// имеет более одного класса, но хорошим бонусом будет:
// 1 Хороший стиль кода, приближенный к общепринятым стандартам
// форматирования кода
// 2 Использование паттернов проектирования
// 3 Логирование ошибок в файл
// 4 Сохранение статистики в базу данных.
// 5 Возможность расширяемости проекта и многоуровневую архитектуру
// 6 Тесты

func Parser() {

	//get html
	ptext := GetPBody()

	if utf8.RuneCountInString(strings.Trim(ptext, " ")) < 5 {
		log.Fatalln("Short text")
	}

	//save page text to file 
  //rewrite!
	err := ioutil.WriteFile("./simbirsoft.txt", []byte(ptext), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//в задании uppercase
	ptext = strings.ToUpper(ptext)

	st := SplitText(ptext)

	//add unique values to map
	muw1 := make(map[string]int)

	for _, r := range st {
		_, ok := muw1[r]
		if ok {
			muw1[r]++
		} else {
			muw1[r] = 1
		}

	}
	//могут быть одинаковые int, поэтому работаем с struct
	//using structs slice for sorting
	type Pair struct {
		Key   int
		Value string
	}

	pl := []Pair{}

	for k, v := range muw1 {
		pl = append(pl, Pair{v, k})
	}

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key > pl[j].Key
	})

	//result!
	for _, p1 := range pl {
		fmt.Printf("%d, %v\n", p1.Key, p1.Value)
	}

		log.Println("Thats all!")
}

func GetPBody() string {
	resp, err := http.Get("https://www.simbirsoft.com/")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func SplitText(text string) []string {

	sp := func(c rune) bool {

		//в задании не указано нужно обрабатывать текст с html тегами или без.
		//если без, то доставать plain text "с помощью стандартных классов" -отдельная большая работа
		//если с тегами, то в сете не хватает как минимум <>
		//поэтому split by любыми знаками не \w

		// ar := [...]rune{' ', ',', '.', '!', '?', '"', ';', ':', '[', ']', '(', ')', '\n', '\r', '\t'}

		// for _, r := range ar {
		// 	if c == r {
		// 		return true
		// 	}
		// }

		// return false

		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	return strings.FieldsFunc(text, sp)

}
