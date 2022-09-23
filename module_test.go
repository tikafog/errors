package merr

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewModule(t *testing.T) {
	type args struct {
		str    string
		err    error
		errStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				str:    "test1",
				err:    errors.New("test1 error"),
				errStr: "test1 errorStr",
			},
			want: "Module[test1]: test1 errorStr: test1 error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RegisterModule(tt.args.str)
			err := got.Wrap(tt.args.err, tt.args.errStr)
			fmt.Printf("%x", err.Index())
			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("NewModule() = %v, want %v", err.Error(), tt.want)
			}

		})
	}
}
