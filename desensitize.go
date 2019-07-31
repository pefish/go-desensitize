package go_desensitize

import (
	"encoding/json"
	"regexp"
)

type DesensitizeClass struct {
	SensitiveStr string
}

var DEFAULT_DESENSITIVESTR = `pass|token|password|key|pkey`

var desensitize = DesensitizeClass{
	SensitiveStr: DEFAULT_DESENSITIVESTR,
}

func (this *DesensitizeClass) SetSensitiveStrs(str string) {
	this.SensitiveStr = str
}

func (this *DesensitizeClass) desensitizeToString(data interface{}) string {
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	str := string(result)
	re := regexp.MustCompile(`("(`+ this.SensitiveStr  +`)":").*?(")`)
	rep := re.ReplaceAllString(str, "$1****$3")
	return rep
}

func (this *DesensitizeClass) desensitize(data interface{}) interface{} {
	var result interface{}
	str := this.desensitizeToString(data)
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}
	return result
}
