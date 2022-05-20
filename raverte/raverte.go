package raverte

import (
	"fmt"
	"raverte/appdata"
	"raverte/userdata"
	"strings"
)

type Raverte struct {
	ChartOnlyMode bool
	Profile       *userdata.Profile
	KeyRing       *userdata.ApiKeyRing
}

func RaverteInit() *Raverte {
	rav := Raverte{}
	rav.ChartOnlyMode = true // incase something goes wrong in LoadProfile.
	return &rav
}

func (r *Raverte) GetChartOnlyMode() bool {
	return r.ChartOnlyMode
}

// Loads the users profile defined in userdata/profile.go
//
// Attempts to open user profile, if it doesn't exist it will create it. If a keystore is NOT configured, it will flag the application for chart only mode.
func (r *Raverte) LoadProfile() error {
	r.Profile = &userdata.Profile{}
	var err error = r.Profile.LoadProfile()
	if err != nil {
		if strings.EqualFold(err.Error(), appdata.FILE_DOES_NOT_EXIST) {
			err = r.Profile.InitialiseProfile()
			if err != nil {
				return fmt.Errorf("%s; %s; %s", appdata.INIT_PROFILE_FAILED, err.Error(), appdata.WIKI_PROFILE_SETUP_URL)
			}
		} else {
			return fmt.Errorf("%s; %s; %s", appdata.LOAD_PROFILE_FAILED, err.Error(), appdata.WIKI_PROFILE_TROUBLESHOOTING_URL)
		}
	}
	r.ChartOnlyMode = !r.Profile.Keystore

	return nil
}

func (r *Raverte) UnlockKeys(password string) error {
	r.KeyRing = &userdata.ApiKeyRing{}

	err := r.KeyRing.UnlockKeys(password, *r.Profile)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
