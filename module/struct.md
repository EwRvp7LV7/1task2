## [Challenge](https://gist.github.com/slntopp/000435b4443a9da511876d6fbe51b757) and my solution

### At the first stage, 
hash benchmark was checked. As expected, shortest 16-bit md5 was faster.

Please, pay attention here: BenchmarkHash_**md5**-4           	 2345684	       **512.6 ns/op**[^1]
[^1]: Nanoseconds per one operation
```
6356d331afbb33be325cabba0e49a332
BenchmarkHash_md5-4           	 2345684	       512.6 ns/op
5b560e8aaca68ad7aa71e850904884217d6683e3
BenchmarkHash_sha1-4          	 1771825	       686.7 ns/op
7acea5f33bb35aa008e32a871a5a226fd3e901f335af70d21f679c4e46b412aa
BenchmarkHash_sha256-4        	 1000000	      1035 ns/op
dad37009c927060de3523a9d7e5d86695b9cc75ebd579f09eef1bdc5
BenchmarkHash_sha512_224-4    	 1283527	       921.0 ns/op
d4dc71526dd29ba4b445c5097950c61c7bbc02dd0f7383201f27d98566fb987e
BenchmarkHash_sha512_256-4    	 1342377	       889.1 ns/op
```
I think, 16-bit hash enough for this purpose, because it's not critical information and it isn't big data. Also, there is some effect of comparation shortest strings.

### At second, 
compared two methods converts struct to binary: `encoding/gob` and `encoding/json`. Noted that `json.Marshal(struct)` was about 4 times faster and more effective in common. Also, I think `encoding/gob` is convenient to work by his buffer with big structs or slow connections.
```
BenchmarkHash_Marshal-4   	  694574	      1573 ns/op	     224 B/op	       2 allocs/op
BenchmarkHash_gob-4       	  136437	      8780 ns/op	    2216 B/op	      31 allocs/op
```
### At third,
developed function for convert whole struct into hash. 
```golang
//get hash of whole struct
func GetWholeStructHash(st interface{}) string {
	bt, _ := json.Marshal(st)
	// fmt.Printf("Marshal:'%s'\n", bt)
	byteSl := md5.Sum(bt)
	return hex.EncodeToString(byteSl[:])
}
``` 
By the way, two lines can be written as one, but they looks more readable and don't worry, the compiler will remove unnecessary variables. \
:exclamation:Note, that `json.Marshal(struct)` convert only public fields i.e. capitalized fields. And if you don't want to consider any fields in hash you can simply make those private. 

### At fourth,
we have to consider tags `hasher:"+"` which looks like
```golang
type PersonalData struct {
	Address string `json:"address" hasher:"+"`
}
```
*Plus* indicates inclusion in a hash.

The naive way to do it is simply to get tags from the fields as `field.Tag.Get("hasher")`. 

:exclamation:But! It works only for a non-nested structs. Also, we should consider that any struct has its own type. That is why we should work with common type like `var interface{}`. And attention... Wellcome `reflect`!

We have to:
- [x] Explore a nested structure as a tree. It's convenient to do it by recursion.
- [x] Check two types of nesting: single nested struct and array of structs, and as well as pointers to it.
- [x] Get string values from string fields with the `hasher:"+"` tag. Of course, it may be there other types of fields, but we will consider only string type for simplifity.
- [x] Concatenate, and return string values as one string. And convert this to hash.

Let's go!

```golang
func readStructStr(val reflect.Value) (out string) {
	//for pointers, converts it to objects
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ { //loop around fields

		f := val.Field(i)

		switch f.Kind() {
		case reflect.Struct://for single nested struct
			out += readStructStr(f)
		case reflect.Slice://for nested slice
			for j := 0; j < f.Len(); j++ {
				out += readStructStr(f.Index(j))
			}
		case reflect.String://for string field
			if val.Type().Field(i).Tag.Get("hasher") == "+" {//with right tag
				out += f.String()
			}
		}
	}

	return out //return concatenated string values
}
//get hash of struct fields taged as hasher:"+"
func GetStructHash(st interface{}) string {
	str := readStructStr(reflect.ValueOf(st)) //convert interface{} to reflect value and run recursion
	//fmt.Printf("str:'%s'\n", str) //check string involving field values 
	byteSl := md5.Sum([]byte(str)) //get hash
	return hex.EncodeToString(byteSl[:]) //return hash as hex string
}
```
Also, this recursion may be more fastly if we put "out strings" on common string by pointer, but this isn't commited for simplifity.

And now, you can explore it in [The Go Playground](https://play.golang.org/p/lsVEetRNg4T).


