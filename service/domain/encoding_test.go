package domain

import "testing"

func Test_base62URLIDEncoder_Decode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint64
		wantErr bool
	}{
		{
			name:    "invalid input",
			input:   "@",
			wantErr: true,
		},
		{
			name:  "input=b",
			input: "b",
			want:  1,
		},
		{
			name:  "input=b",
			input: "qQ3Pfb",
			want:  1000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := base62URLIDEncoder{}
			got, err := b.Decode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base62URLIDEncoder_Encode(t *testing.T) {
	tests := []struct {
		name    string
		input   uint64
		want    string
		wantErr bool
	}{
		{
			name:    "invalid ID",
			input:   0,
			wantErr: true,
		},
		{
			name:  "id=1",
			input: 1,
			want:  "b",
		},
		{
			name:  "id=1 000 000 000",
			input: 1000000000,
			want:  "qQ3Pfb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := base62URLIDEncoder{}
			got, err := b.Encode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
