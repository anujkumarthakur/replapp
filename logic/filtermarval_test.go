package logic

import (
	"reflect"
	"replapp/model"
	"testing"
)

func TestGetAllMarvalCharacter(t *testing.T) {
	type args struct {
		name   string
		offset string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.SearchCharacterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "thor",
			args: args{
				name:   "thor",
				offset: "0",
			},
			want:    &model.SearchCharacterResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllMarvalCharacter(tt.args.name, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMarvalCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMarvalCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
