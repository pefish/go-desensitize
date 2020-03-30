
## Description

Clear sensitive information

## Quick Start

```golang
func ExampleDesensitizeClass_DesensitizeToString() {
	a := Desensitize.DesensitizeToString(`{"a": "57", "token": "uejdsh"}`)
	fmt.Println(a)

	type Test struct {
		A     string `json:"token65"`
		Token string `json:"41password"`
	}

	test := Test{
		A: `21`,
		Token: `sgshgj`,
	}
	a1 := Desensitize.DesensitizeToString(test)
	fmt.Println(a1)
	// Output:
	// {"a":"57","token":"****"}
	// {"token65":"****","41password":"****"}
}
```

