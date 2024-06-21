package service

import (
	"github.com/NathanaelAma/url-shortener/model"
	"reflect"
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	type args struct {
		longUrl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal URL",
			args: args{"www.google.com"},
			want: "GRNHv-Vd",
		},
		{
			name: "Very Long URL",
			args: args{"www.testsite.co.za/this/is/a/test/url/with/many/params?and=many&more=params"},
			want: "4YDVI7Th",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateShortUrl(tt.args.longUrl); got != tt.want {
				t.Errorf("GenerateShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLongUrl(t *testing.T) {
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Normal URL",
			args:    args{"GRNHv-Vd"},
			want:    "www.google.com",
			wantErr: false,
		},
		{
			name:    "Very Long URL",
			args:    args{"4YDVI7Th"},
			want:    "www.testsite.co.za/this/is/a/test/url/with/many/params?and=many&more=params",
			wantErr: false,
		},
		{
			name:    "Invalid URL",
			args:    args{"invalid"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var url model.Url
			var err error
			if tt.args.shortUrl != "" {
				url, err = ShortenUrl(model.Url{
					Long: tt.want,
				})
			}
			if err != nil {
				return
			}
			got, errorMessage := GetLongUrl(url.Short)
			if (errorMessage != nil) != tt.wantErr {
				t.Errorf("GetLongUrl() error = %v, wantErr %v", errorMessage, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLongUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenUrl(t *testing.T) {
	type args struct {
		url model.Url
	}
	tests := []struct {
		name    string
		args    args
		want    model.Url
		wantErr bool
	}{
		{
			name: "Normal URL",
			args: args{model.Url{
				Long: "www.google.com",
			},
			},
			want: model.Url{
				Long:  "www.google.com",
				Short: "GRNHv-Vd",
			},
		},
		{
			name: "Very Long URL",
			args: args{model.Url{
				Long: "www.testsite.co.za/this/is/a/test/url/with/many/params?and=many&more=params",
			},
			},
			want: model.Url{
				Long:  "www.testsite.co.za/this/is/a/test/url/with/many/params?and=many&more=params",
				Short: "4YDVI7Th",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShortenUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortenUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortenUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
