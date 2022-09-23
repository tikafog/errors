package merr

import "testing"

func TestStackError(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				err: "test error",
			},
			wantErr: true,
			want:    "test error",
		},
		{
			name: "",
			args: args{
				err: "test error",
			},
			wantErr: true,
			want:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StackError(tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("StackError() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !MessageIs(tt.want, err) {
				t.Errorf("StackError() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestErrorMessageIs(t *testing.T) {
	type args struct {
		str string
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				str: "test1",
				err: WrapString(nil, "test1"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: "test1",
				err: WrapString(nil, "test2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MessageIs(tt.args.str, tt.args.err); got != tt.want {
				t.Errorf("ErrorMessageIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorIndexIs(t *testing.T) {
	type args struct {
		i   Index
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				i:   1,
				err: WrapIndex(nil, 1),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				i:   1,
				err: WrapIndex(nil, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexIs(tt.args.i, tt.args.err); got != tt.want {
				t.Errorf("ErrorIndexIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorIs(t *testing.T) {
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				err:    WrapIndex(nil, 0),
				target: WrapString(nil, "core is nil"),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				err:    WrapIndex(nil, 1),
				target: WrapString(nil, "core is nil"),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				err:    StackError("core is nil"),
				target: WrapString(nil, "core is nil"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				err:    StackError("core is nil1"),
				target: WrapString(nil, "core is nil"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("ErrorIs() = %v, want %v", got, tt.want)
			}
		})
	}
}
