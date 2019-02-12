package apps

import (
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/models"
)

func Test_appsPostgreSQLRepository_GetApps(t *testing.T) {
	tests := []struct {
		name    string
		want    *[]models.App
		wantErr bool
	}{
		{
			name: "appsPostgreSQLRepository.GetApps() should return a list of apps",
			want: &[]models.App{
				models.App{
					ID:                    1,
					AppID:                 "com.aerogear.app1",
					AppName:               "app1",
					NumOfDeployedVersions: 1,
					NumOfAppLaunches:      1,
					NumOfClients:          1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &appsPostgreSQLRepository{}
			got, err := a.GetApps()
			if (err != nil) != tt.wantErr {
				t.Errorf("appsPostgreSQLRepository.GetApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appsPostgreSQLRepository.GetApps() = %v, want %v", got, tt.want)
			}
		})
	}
}
