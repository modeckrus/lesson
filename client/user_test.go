package client

import "testing"

func TestUser_isExist(t *testing.T) {
	type fields struct {
		ID   int64
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			"isExist",
			fields{
				ID:   int64(0),
				Name: "test",
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			got, err := u.isExist()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.isExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.isExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
