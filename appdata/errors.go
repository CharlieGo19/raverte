package appdata

const (
	FILE_DOES_NOT_EXIST        string = "file does not exist"
	CAN_NOT_SET_PERMISSIONS    string = "couldn't set permissions for "
	CAN_NOT_CREATE_FILE        string = "couldn't create raverte asset "
	CAN_NOT_CREATE_RAVERTE_DIR string = "couldn't create raverte directory"
	INCORRECT_PERMISSIONS      string = "incorrect permissions, should be rw for your user only"

	INVALID_RAVERTE_ASSET    string = "invalid raverte asset"
	HOMEDIR_NOT_FOUND_UNIX   string = "couldn't find home folder"
	APPSUPP_NOT_FOUND_DARWIN string = "couldn't find application support folder"
	APPDATA_NOT_FOUND_WIN    string = "couldn't find appdata folder"

	INIT_PROFILE_FAILED    string = "could not initialise profile"
	LOAD_PROFILE_FAILED    string = "could not load profile"
	PROFILE_ALREADY_EXISTS string = "profile already exists"
)
