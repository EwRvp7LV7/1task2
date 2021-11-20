package module

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Id    int     `validate:"-"`
	Name  string  `validate:"presence,min=2,max=32"`
	Email string  `validate:"email,required"`
	Tmp   float32 `json:"followers,omitempty,string"`
	//The “omitempty” option specifies that the field should be omitted from the encoding if the field has an empty value
	//The “string” option signals that a field is stored as JSON inside a JSON-encoded string.
	FirstName string `json:"fname"`           // `fname` as field name
	LastName  string `json:"lname,omitempty"` // discard if value is empty
	Email1    string `json:"-"`               // always discard
	Age       int    `json:"-,"`              // `-` as field name
	IsMale    bool   `json:",string"`         // keep original field name, coerce to a string
	Profile   bool   `json:""`                // no effect
}

// https://medium.com/rungo/structures-in-go-76377cc106a2
// Nested struct
// особенность Go в том, что когда мы используем анонимная вложенная структура, все поля вложенной структуры автоматически доступны в родительской структуре.
// Это называетсяпродвижение на местах.
// when we use an anonymous nested struct, all the nested struct fields are automatically available on parent struct. This is called field promotion.
// Вложенный интерфейс
// interface type is a declaration of method signatures
// Using this polymorphism principle, we can have a struct field of an interface type
// Экспортируемые поля FirstName, LastName string
// Function fields
// function as a type and function as a value
// Struct comparison
// Two structs are comparable if they belong to the same type and have the same field values.

// Name of the struct tag used in examples
const tagName = "validate"

func MainStruct() {
	// user := User{
	// 	Id:    1,
	// 	Name:  "John Doe",
	// 	Email: "john@example",
	// }

	// user := User{1,"John Doe","john@example", big.NewInt(), }

	//Anonymous struct
	user1 := struct {
		id    int    `validate:"-"`
		name  string `validate:"presence,min=2,max=32"`
		email string `validate:"email,required"`
	}{
		id:    1,
		name:  "John Doe",
		email: "john@example",
	}

	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	// t := reflect.TypeOf(user)
	t := reflect.TypeOf(user1)

	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		//{Email  string validate:"email,required" 24 [2] false}
		// type StructField struct {
		// 	// Name is the field name.
		// 	Name string

		// 	// PkgPath is the package path that qualifies a lower case (unexported)
		// 	// field name. It is empty for upper case (exported) field names.
		// 	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
		// 	PkgPath string

		// 	Type      Type      // field type
		// 	Tag       StructTag // field tag string
		// 	Offset    uintptr   // offset within struct, in bytes //смещение в структуре в байтах
		// 	Index     []int     // index sequence for Type.FieldByIndex
		// 	Anonymous bool      // is an embedded field
		// }

		fmt.Println("Field:", field)

		// Get the field tag value
		tag := field.Tag.Get(tagName)

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}

// https://gist.github.com/slntopp/000435b4443a9da511876d6fbe51b757
type PersonalData struct {
	Title   string `json:"title"`
	Address string `json:"address" hasher:"+"`
}
type Whatever struct {
	Uuid string       `json:"uuid"`
	Data PersonalData `json:"data"`
}

// // https://stackoverflow.com/questions/68299983/reading-nested-structure-using-reflection
// func ReadStruct(st interface{}) {
// 	readStruct(reflect.ValueOf(st))
// }

// func readStruct(val reflect.Value) {

// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem()
// 	}

// 	for i := 0; i < val.NumField(); i++ {
// 		// fmt.Println(val.Type().Field(i).Type.Kind())
// 		f := val.Field(i)
// 		fmt.Println("Type.Kind:", f.Kind())
// 		switch f.Kind() {
// 		case reflect.Struct:
// 			readStruct(f)
// 		case reflect.Slice:
// 			for j := 0; j < f.Len(); j++ {
// 				readStruct(f.Index(i))
// 			}
// 		case reflect.String:
// 			fmt.Printf("String:%v=%v\n", val.Type().Field(i).Name, val.Field(i))
// 		}
// 	}
// }

func GetWholeStructHash(st interface{}) string {
	// var network bytes.Buffer // Stand-in for a network connection
	// enc := gob.NewEncoder(&network)
	// enc.Encode(st)
	// byteSl := md5.Sum(network.Bytes())
	//1655 vs 5443 ns/op
	bt, _ := json.Marshal(st)
	fmt.Printf("Marshal:'%s'\n", bt)
	byteSl := md5.Sum(bt)
	return hex.EncodeToString(byteSl[:])
}

func GetStructHash(st interface{}) string {
	str := readStructStr(reflect.ValueOf(st))
	fmt.Printf("str:'%s'\n", str)
	byteSl := md5.Sum([]byte(str))
	return hex.EncodeToString(byteSl[:])
}

func readStructStr(val reflect.Value) (out string) {
	//if pointer is sended
	if val.Kind() == reflect.Ptr {
		// fmt.Println("Type.Kind:", val.Type())
		val = val.Elem()
	}
	// if val.Kind() != reflect.Struct { //slice!
	// 	return ""
	// }

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		// fmt.Println("val.Field:", f.Kind())
		switch f.Kind() {
		case reflect.Struct:
			out += readStructStr(f)
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				out += readStructStr(f.Index(j))
			}
		case reflect.String:
			// fmt.Printf("String:%v=%v\n", val.Type().Field(i).Name, val.Field(i).Interface())
			if val.Type().Field(i).Tag.Get("hasher") == "+" {
				// fmt.Printf("Tag:%v=%v\n", val.Type().Field(i).Name, val.Type().Field(i).Tag.Get("hasher"))
				out += f.String()

			}
		}
	}

	return out
}

func MainStructTags() {

	// fieldBySlice := func(strct interface{}) interface{} {
	// 	t := reflect.TypeOf(strct)
	// 	// val := reflect.ValueOf(strct)
	// 	for i := 0; i < t.NumField(); i++ {
	// 		field := t.Field(i)
	// 		// val = val.FieldByName(field.Name)
	// 		fmt.Println("Field:", field)
	// 	}
	// 	// return val.Interface()
	// 	return t
	// }
	// https://gist.github.com/slntopp/000435b4443a9da511876d6fbe51b757
	obj := Whatever{"a-b-c-d", PersonalData{"i'm whatever", "my address"}}
	fmt.Println(obj)
	//hash of whole struct
	fmt.Println(GetWholeStructHash(obj))
	//hash ignore fields with tag hasher:"+"
	fmt.Println(GetStructHash(obj))
	// ReadStruct(obj)
	// hash := hasher.Marshal(obj)
	// fmt.Println(hash) // => "f714d0221fef51f2c0d26fa1ac1e3513"
	// a few moments later
	// obj.Data.Title = "i'm whoever"
	// fmt.Println(obj)
	// ReadStruct(obj)
	// // hash = hasher.Marshal(obj)
	// // fmt.Println(hash) // => "605671cbae106e973f9b0f50dbbcf416"

	user1 := struct { //&struct { no difference
		uuid    string `validate:"-"` //This field is excluded in Marshal!
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

	fmt.Println(user1)
	//hash of whole struct
	fmt.Println(GetWholeStructHash(user1))
	//hash ignore fields with tag hasher:"+"
	fmt.Println(GetStructHash(user1))

}
