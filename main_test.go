package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		args   []string
		limit  int
		client Requester
	}
	tests := []struct {
		name string
		args args
		want []Result
	}{
		{
			name: "error on invalid url",
			args: args{
				args:  []string{""},
				limit: 1,
				client: &TestRequester{
					m: map[string]returns{},
				},
			},
			want: []Result{
				{
					Resp: Response{},
					Err:  errors.New("url is empty"),
				},
			},
		},
		{
			name: "test url",
			args: args{
				args:  []string{"testurl.com"},
				limit: 1,
				client: &TestRequester{
					m: map[string]returns{
						"http://testurl.com": {
							Bytes: []byte("test string for md5"),
							Err:   nil,
						},
					},
				},
			},
			want: []Result{
				{
					Resp: Response{
						Url:  "http://testurl.com",
						Hash: "1db6b460538f9e7afe09d25663c468c9",
					},
					Err: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collect(Run(tt.args.args, tt.args.limit, tt.args.client)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

type returns struct {
	Bytes []byte
	Err   error
}

type TestRequester struct {
	m map[string]returns
}

func (t *TestRequester) SendRequest(urlPath string) ([]byte, error) {
	r, ok := t.m[urlPath]
	if !ok {
		panic("unexpected mock path: " + urlPath)
	}

	return r.Bytes, r.Err
}

func collect(ch chan Result) []Result {
	var slice []Result
	for elem := range ch {
		slice = append(slice, elem)
	}
	return slice
}
