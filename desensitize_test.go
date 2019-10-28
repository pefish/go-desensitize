package go_desensitize

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDesensitizeClass_DesensitizeToString(t *testing.T) {
	type Test struct {
		A     string `json:"token65"`
		Token string `json:"41password"`
	}

	test := Test{
		A: `21`,
		Token: `sgshgj`,
	}
	a := Desensitize.DesensitizeToString(test)
	fmt.Println(a)
}

func TestDesensitizeClass_desensitizeToString(t *testing.T) {
	type Test struct {
		A     string `json:"a"`
		Token string `json:"token"`
	}

	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `test`,
			args: args{
				data: map[string]interface{}{
					`aa`:       `11`,
					`token`:    `wrts`,
					`password`: `625426`,
				},
			},
			want: `{"aa":"11","password":"****","token":"****"}`,
		},
		{
			name: `test1`,
			args: args{
				data: Test{
					A:     `11`,
					Token: `ywrths`,
				},
			},
			want: `{"a":"11","token":"****"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Desensitize.DesensitizeToString(tt.args.data); got != tt.want {
				t.Errorf("DesensitizeClass.desensitizeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesensitizeClass_desensitize(t *testing.T) {
	type fields struct {
		SensitiveStr string
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: `test`,
			fields: fields{
				SensitiveStr: `token`,
			},
			args: args{
				data: map[string]interface{}{
					`aa`:       `11`,
					`token`:    `wrts`,
					`password`: `625426`,
				},
			},
			want: `****`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Desensitize.Desensitize(tt.args.data); !reflect.DeepEqual(got.(map[string]interface{})[`token`], tt.want) {
				t.Errorf("DesensitizeClass.desensitize() = %v, want %v", got, tt.want)
			}
		})
	}
}
