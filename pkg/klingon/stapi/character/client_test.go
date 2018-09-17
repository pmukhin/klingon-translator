package character

import (
	"github.com/pmukhin/klingon-translator/internal/testing/httpmock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func getStubResource(filename string) string {
	bts, err := ioutil.ReadFile("stubs/" + filename)
	if err != nil {
		panic("can not read resource file: " + err.Error())
	}
	return string(bts)
}

func Test_defaultCharactersClient_Get(t *testing.T) {
	type args struct {
		uid      string
		response httpmock.MockClientResponse
	}
	tests := []struct {
		name    string
		args    args
		want    *Full
		wantErr bool
	}{
		{
			name: "valid response with at least one species",
			args: args{
				uid: "CHMA0000023576",
				response: httpmock.MockClientResponse{
					Status: http.StatusOK,
					Body:   getStubResource("get-uhura-with-species.json"),
				},
			},
			want: &Full{
				UID:    "CHMA0000023576",
				Name:   "Nyota Uhura",
				Gender: "F",
				CharacterSpecies: []Species{{
					UID:  "someUid",
					Name: "Human",
				}},
			},
			wantErr: false,
		},

		{
			name: "valid response with no species",
			args: args{
				uid: "CHMA0000023576",
				response: httpmock.MockClientResponse{
					Status: http.StatusOK,
					Body:   getStubResource("get-uhura-without-species.json"),
				},
			},
			want: &Full{
				UID:    "CHMA0000023576",
				Name:   "Nyota Uhura",
				Gender: "F",
				CharacterSpecies: []Species{{
					UID:  "",
					Name: "Human",
				}},
			},
			wantErr: false,
		},
		{
			name: "invalid response",
			args: args{
				uid: "CHMA0000023576",
				response: httpmock.MockClientResponse{
					Status: http.StatusBadGateway,
					Body:   "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultCharactersClient{
				client:  httpmock.New([]httpmock.MockClientResponse{tt.args.response}),
				baseUrl: "mockhttp://baseurl",
			}
			got, err := d.Get(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultCharactersClient.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultCharactersClient.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultCharactersClient_Search(t *testing.T) {
	type args struct {
		name     string
		response httpmock.MockClientResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []Short
		wantErr bool
	}{
		{
			name: "valid response",
			args: args{
				name: "Uhura",
				response: httpmock.MockClientResponse{
					Status: http.StatusOK,
					Body:   getStubResource("search-uhura.json"),
				},
			},
			want: []Short{
				{
					UID:    "CHMA0000023576",
					Name:   "Nyota Uhura",
					Gender: "F",
				},
			},
			wantErr: false,
		},
		{
			name: "empty response",
			args: args{
				name: "testCharacterWhichDoesNotExist(probably)",
				response: httpmock.MockClientResponse{
					Status: http.StatusOK,
					Body:   getStubResource("search-empty.json"),
				},
			},
			want:    []Short{},
			wantErr: false,
		},
		{
			name: "invalid response",
			args: args{
				name: "someSpecialName",
				response: httpmock.MockClientResponse{
					Status: http.StatusBadGateway,
					Body:   "",
				},
			},
			want:    []Short{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultCharactersClient{
				client:  httpmock.New([]httpmock.MockClientResponse{tt.args.response}),
				baseUrl: "mockhttp://baseurl",
			}
			got, err := d.Search(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultCharactersClient.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultCharactersClient.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
