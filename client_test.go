package main

import "testing"

func Test_checkProtocolSchema(t *testing.T) {
	type args struct {
		urlPath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "error on empty url path",
			args:    args{""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "url path without scheme",
			args:    args{"google.com"},
			want:    "http://google.com",
			wantErr: false,
		},
		{
			name:    "url path with http scheme",
			args:    args{"http://google.com"},
			want:    "http://google.com",
			wantErr: false,
		},
		{
			name:    "url path with https scheme",
			args:    args{"https://google.com"},
			want:    "https://google.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkProtocolSchema(tt.args.urlPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkProtocolSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkProtocolSchema() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMD5Hash(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test string",
			args: args{[]byte("test string for md5")},
			want: "1db6b460538f9e7afe09d25663c468c9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMD5Hash(tt.args.str); got != tt.want {
				t.Errorf("getMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
