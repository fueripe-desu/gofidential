package loader

const (
	// The LOADER_MISSING_NAME error occurs when the Name field of the Environment
	// struct is either missing or an empty string.
	MissingName string = "LOADER_MISSING_NAME"

	// The LOADER_INVALID_NAME error occurs when the Name field of the Environment
	// struct contains invalid characters, such as:
	//
	//	- Uppercase letters
	//	- Numbers
	//	- Special characters (other than underscores)
	InvalidName string = "LOADER_INVALID_NAME"

	// The LOADER_UNDERSCORE_ONLY_NAME error occurs when the Name field of the Environment
	// struct is composed entirely of underscores.
	UnderscoreOnlyName string = "LOADER_UNDERSCORE_ONLY_NAME"

	// The LOADER_TRAILING_UNDERSCORE error occurs when the Name field of the Environment
	// struct contains a trailing underscore.
	TrailingUnderscore string = "LOADER_TRAILING_UNDERSCORE"

	// The LOADER_LEADING_UNDERSCORE error occurs when the Name field of the Environment
	// struct contains a leading underscore.
	LeadingUnderscore string = "LOADER_LEADING_UNDERSCORE"

	// The LOADER_INEXISTENT_DIR error occurs when the OverridePath field in the Environment
	// struct points to a directory that does not exist.
	InexistentDir string = "LOADER_INEXISTENT_DIR"

	// The LOADER_FAILED_TO_READ_DIR error occurs when the application does not have permission
	// to read the directory specified in the OverridePath field of the Environment struct.
	FailedToReadDir string = "LOADER_FAILED_TO_READ_DIR"

	// The LOADER_FAILED_TO_READ_ENV error occurs when the application does not have permission
	// to read the target .env file.
	FailedToReadEnv string = "LOADER_FAILED_TO_READ_ENV"

	// The LOADER_PATH_IS_NOT_DIR error occurs when the OverridePath field defined in the Environment
	// struct points to a file instead of a directory.
	PathIsNotDir string = "LOADER_PATH_IS_NOT_DIR"

	// The LOADER_PATH_IS_NOT_FILE error is an internal error that occurs when the loader is not
	// calculating the expected filename correctly.
	PathIsNotFile string = "LOADER_PATH_IS_NOT_FILE"

	// The LOADER_ENV_NOT_FOUND error occurs when the loader cannot find the target .env file in the
	// expected directory.
	EnvNotFound string = "LOADER_ENV_NOT_FOUND"
)
