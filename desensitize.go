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

func (d *DesensitizeClass) SetSensitiveStrs(str []string) {
	d.sensitiveStrArr = str
}

func (d *DesensitizeClass) MustDesensitizeToString(data interface{}) string {
	r, err := d.DesensitizeToString(data)
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DesensitizeClass) DesensitizeToString(data interface{}) (string, error) {
	if data == nil {
		return "Nil", nil
	}
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
			return "", err
		}
		str = string(marshalResult)
	}
	for _, v := range d.sensitiveStrArr {
		regStr := fmt.Sprintf(`("%s":").*?(")`, v)
		re := regexp.MustCompile(regStr)
		str = re.ReplaceAllString(str, "$1****$2")
	}
	return str, nil
}

func (d *DesensitizeClass) MustDesensitize(data interface{}) interface{} {
	r, err := d.Desensitize(data)
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DesensitizeClass) Desensitize(data interface{}) (interface{}, error) {
	var result interface{}
	str, err := d.DesensitizeToString(data)
	if err != nil {
		return nil, err
	}
	if str == "Nil" {
		return nil, nil
	}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
