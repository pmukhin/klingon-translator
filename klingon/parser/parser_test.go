package parser

import (
	"reflect"
	"testing"
)

func Test_defaultParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		want    []rune
		wantErr bool
	}{
		{
			name:    "Uhura",
			want:    []rune{0xF8E5, 0xF8D6, 0xF8E5, 0xF8E1, 0xF8D0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.name)

			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
