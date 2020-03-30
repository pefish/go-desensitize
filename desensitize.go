package go_desensitize

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
)

type DesensitizeClass struct {
	sensitiveStrArr []string
}

var DEFAULT_DESENSITIVESTR = []string{
	`.*?pass.*?`,
	`.*?token.*?`,
	`.*?key.*?`,
	`.*?secret.*?`,
}

var Desensitize = DesensitizeClass{
	sensitiveStrArr: DEFAULT_DESENSITIVESTR,
}

func (this *DesensitizeClass) SetSensitiveStrs(str []string) {
	this.sensitiveStrArr = str
}

func (this *DesensitizeClass) DesensitizeToString(data interface{}) string {
	type_ := reflect.TypeOf(data)
	str := ``
	if type_.Kind() == reflect.String {
		str = data.(string)
		// 去除所有空格
		re := regexp.MustCompile(` `)
		str = re.ReplaceAllString(str, "")
	} else {
		marshalResult, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		str = string(marshalResult)
	}
	for _, v := range this.sensitiveStrArr {
		regStr := fmt.Sprintf(`("%s":").*?(")`, v)
		re := regexp.MustCompile(regStr)
		str = re.ReplaceAllString(str, "$1****$2")
	}
	return str
}

func (this *DesensitizeClass) Desensitize(data interface{}) interface{} {
	var result interface{}
	str := this.DesensitizeToString(data)
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}
	return result
}
