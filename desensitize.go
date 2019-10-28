package go_desensitize

import (
	"encoding/json"
	"fmt"
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
	marshalResult, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	str := string(marshalResult)
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
