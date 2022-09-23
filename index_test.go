package merr

import (
	"fmt"
	"log"
	"testing"
)

var test1 = RegisterModule("test1")
var test2 = RegisterModule("test2")

func init() {
	RegisterErrorHandler(func(idx Index, err error) {
		log.Println("new error", idx.ModuleName())
		log.Println("new error", err.Error())
	})
	for i := 0; i < 100; i++ {
		test1.New(fmt.Sprintf("test1.error: %d", i))
	}
	for i := 0; i < 100; i++ {
		test2.New(fmt.Sprintf("test2.error: %d", i))
	}

}

func TestMakeErrIndex(t *testing.T) {
	type args struct {
		prefix uint32
		index  uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				prefix: 0x01,
				index:  1,
			},
			want: "Module[test1]: test1.error: 0",
		},
		{
			name: "",
			args: args{
				prefix: 0x02,
				index:  2,
			},
			want: "Module[test2]: test2.error: 1",
		},
		{
			name: "",
			args: args{
				prefix: 0,
				index:  0,
			},
			want: "unknown error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeErrIndex(tt.args.prefix, tt.args.index)
			if got.String() != tt.want {
				t.Errorf("makeErrIndex() = %v, want %v", got, tt.want)
			}
			e := IndexError(got)
			if e.Error() != tt.want {
				t.Errorf("IndexError() = %v, want %v", got, tt.want)
			}
			//if err := test1.New(got); err != nil {
			//	t.Errorf("IndexError() = %v, want %v", got, tt.want)
			//	return
			//}
		})
	}
}
