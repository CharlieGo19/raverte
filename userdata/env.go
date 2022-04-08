package userdata

import (
	"errors"
	"fmt"
	"os"
	"raverte/appdata"
	"runtime"
	"strings"
)

// TODO: Errors into errors

var windowsRaverteFolderPath string = "\\Raverte\\"
var macosRaverteFolderPath string = "/Raverte/"
var linuxRaverteFolderPath string = "/.raverte/"

// Returns path to directory that should be holding the data for Raverte to function effectively.
func GetRaverteAsset(asset string) (string, error) {
	var filePath string

	if asset != appdata.KEYSTORE && asset != appdata.PROFILE {
		return "", fmt.Errorf(appdata.INVALID_RAVERTE_ASSET)
	}

	if runtime.GOOS == "windows" {
		appData, err := os.UserCacheDir()
		if err != nil {
			return "", fmt.Errorf(appdata.APPDATA_NOT_FOUND_WIN)
		}
		filePath = appData + windowsRaverteFolderPath
	} else if runtime.GOOS == "darwin" {
		appData, err := os.UserConfigDir()
		if err != nil {
			return "", fmt.Errorf(appdata.APPSUPP_NOT_FOUND_DARWIN)
		}
		filePath = appData + macosRaverteFolderPath
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf(appdata.HOMEDIR_NOT_FOUND_UNIX)
		}
		filePath = homeDir + linuxRaverteFolderPath

	}

	return filePath + asset, nil
}

// Checks if Raverte asset exists and has permissions defined on creation.
//
// Returns error if folder/file does not exists and permissions of afore mentioned are not 0600.
func checkRaverteAsset(path string) error {
	if info, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New(appdata.FILE_DOES_NOT_EXIST)
	} else if info.Mode().Perm() != 0600 {
		return errors.New(appdata.INCORRECT_PERMISSIONS)
	}
	return nil
}

// Creates and configures Raverte assets, permissions allow only user to read and write.
func configureRaverteAsset(filePath, asset string) error {
	var raverteFolder string = strings.ReplaceAll(filePath, asset, "")
	if _, err := os.Stat(raverteFolder); os.IsNotExist(err) {
		err = os.Mkdir(filePath, 0600)
		if err != nil {
			return errors.New(appdata.CAN_NOT_CREATE_RAVERTE_DIR)
		}
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if _, err := os.Create(filePath); err != nil {
			return errors.New(fmt.Sprintln(appdata.CAN_NOT_CREATE_FILE, filePath))
		}
		fmt.Println("Created profile at ", filePath)
	}
	if err := os.Chmod(filePath, 0600); err != nil {
		fmt.Println("Permissions problems?")
		return errors.New(fmt.Sprintln(appdata.CAN_NOT_SET_PERMISSIONS, filePath))
	}
	return nil
}

// Deletes Raverte asset.
func destroyRaverteAsset(asset string) {
	os.Remove(asset)
}
