package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Test struct {
	Name  string
	Array []byte
}

// func (t *Test) MarshalJSON() ([]byte, error) {
//     var array string
//     if t.Array == nil {
//         array = "null"
//     } else {
//         array = strings.Join(strings.Fields(fmt.Sprintf("%d", t.Array)), ",")
//     }
//     jsonResult := fmt.Sprintf(`{"Name":%q,"Array":%s}`, t.Name, array)
//     return []byte(jsonResult), nil
// }

//Kson косячит с преобразоввнием []byte
// As per the docs: https://golang.org/pkg/encoding/json/#Marshal
// Array and slice values encode as JSON arrays,
//except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON object.
func JsonBytesArr() {
	for i := 0; i < 1000; i++ {
		token := make([]byte, 32)
		rand.Seed(time.Now().UnixNano())
		rand.Read(token)
		t := &Test{"Go", token}

		m, err := json.Marshal(t)
		if err != nil {
			fmt.Println(err)
		}
		var t2 Test
		err = json.Unmarshal(m, &t2)
		if err != nil {
			fmt.Println(err)
		}

		res := bytes.Compare(t.Array, t2.Array)
		if res != 0 {
			fmt.Printf("%v\n", t.Array)
			fmt.Printf("%v\n", t2.Array)
		}
	}
}
