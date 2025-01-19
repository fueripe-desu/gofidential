package loader

const (
	// Occurs when the 'name' parameter is not specified in the provided Environment.
	MissingName string = "LOADER_MISSING_NAME"

	// Occurs when the 'name' parameter is not lowercase and underscores only.
	InvalidName string = "LOADER_INVALID_NAME"

	// Occurs when the 'name' parameter is composed by only underscores.
	UnderscoreOnlyName string = "LOADER_UNDERSCORE_ONLY_NAME"

	// Occurs when the 'name' parameter contains a trailing underscore.
	TrailingUnderscore string = "LOADER_TRAILING_UNDERSCORE"

	// Occurs when the 'name' parameter contains a leading underscore.
	LeadingUnderscore string = "LOADER_LEADING_UNDERSCORE"

	// Occurs when the 'path' parameter specified in the Environment is not an existent dir.
	InexistentDir string = "LOADER_INEXISTENT_DIR"

	// Occurs when the 'path' parameter specified in the Environment could not be read. E.g. Permission denied.
	FailedToReadDir string = "LOADER_FAILED_TO_READ_DIR"

	// Occurs when a file could not be read due to an unexpected error. E.g. Permission denied.
	FailedToReadEnv string = "LOADER_FAILED_TO_READ_ENV"

	// Occurs when the 'path' parameter specified in the Environment does not point to a dir.
	PathIsNotDir string = "LOADER_PATH_IS_NOT_DIR"

	// Occurs when the 'path' parameter specified in the Environment does not point to a file.
	PathIsNotFile string = "LOADER_PATH_IS_NOT_FILE"

	// Occurs when the file loader could not find the target .env file.
	EnvNotFound string = "LOADER_ENV_NOT_FOUND"
)
