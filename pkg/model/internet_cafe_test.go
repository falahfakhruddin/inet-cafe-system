package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllocatePc(t *testing.T) {
	tests := []struct {
		name           string
		setup          *InternetCafe
		parameter      *User
		expectedResult *PC
	}{
		{
			name: "PC Full Book",
			setup: &InternetCafe{
				Capacity: uint64(2),
				ListPC: []*PC{
					{
						ID: uint64(1),
						User: &User{
							Name: "falah",
						},
					},
					{
						ID: uint64(2),
						User: &User{
							Name: "fakhruddin",
						},
					},
				},
			},
			parameter: &User{
				Name: "budi",
				Age:  int64(12),
			},
			expectedResult: nil,
		},
		{
			name: "Slot 2 Available",
			setup: &InternetCafe{
				Capacity: uint64(2),
				ListPC: []*PC{
					{
						ID: uint64(1),
						User: &User{
							Name: "falah",
						},
					},
					{
						ID:   uint64(2),
						User: nil,
					},
				},
			},
			parameter: &User{
				Name: "budi",
				Age:  int64(12),
			},
			expectedResult: &PC{
				ID: uint64(2),
				User: &User{
					Name: "budi",
					Age:  int64(12),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inetCafe := tt.setup
			result := inetCafe.AllocatePC(tt.parameter)
			assert.Equal(t, tt.expectedResult, result, tt.name)
		})
	}
}
