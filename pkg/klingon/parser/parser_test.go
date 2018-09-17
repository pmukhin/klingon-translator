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
		// breaking alphabet into groups for better readability
		{
			name:    "Uhura",
			want:    []rune{0xF8E5, 0xF8D6, 0xF8E5, 0xF8E1, 0xF8D0},
			wantErr: false,
		},
		{
			name:    "abchDegh",
			want:    []rune{0xF8D0, 0xF8D1, 0xF8D2, 0xF8D3, 0xF8D4, 0xF8D5},
			wantErr: false,
		},
		{
			name:    "HIjlmnng",
			want:    []rune{0xF8D6, 0xF8D7, 0xF8D8, 0xF8D9, 0xF8DA, 0xF8DB, 0xF8DC},
			wantErr: false,
		},
		{
			name:    "opqQr",
			want:    []rune{0xF8DD, 0xF8DE, 0xF8DF, 0xF8E0, 0xF8E1},
			wantErr: false,
		},
		{
			name:    "sttlhu",
			want:    []rune{0xF8E2, 0xF8E3, 0xF8E4, 0xF8E5},
			wantErr: false,
		},
		{
			name:    "vwy'",
			want:    []rune{0xF8E6, 0xF8E7, 0xF8E8, 0xF8E9},
			wantErr: false,
		},
		{
			name:    "vwy' v",
			want:    []rune{0xF8E6, 0xF8E7, 0xF8E8, 0xF8E9, 0x20, 0xF8E6},
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
