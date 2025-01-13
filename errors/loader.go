package errors

const (
	// Occurs when the 'name' parameter is not specified in the provided Environment.
	MissingEnvNameCode string = "MISSING_ENV_NAME"

	// Occurs when the 'path' parameter specified in the Environment is not an existent folder.
	EnvFolderNotExistCode string = "ENV_FOLDER_NOT_EXIST"

	// Occurs when the 'path' parameter specified in the Environment could not be read. E.g. Permission denied.
	FailedToReadFolderCode string = "FAILED_TO_READ_FOLDER"

	// Occurs when a file could not be read due to an unexpected error. E.g. Permission denied.
	FailedToReadFileCode string = "FAILED_TO_READ_FILE"

	// Occurs when the 'path' parameter specified in the Environment does not point to a folder.
	PathIsNotFolderCode string = "PATH_IS_NOT_FOLDER"

	// Occurs when the 'path' parameter specified in the Environment does not point to a file.
	PathIsNotFileCode string = "PATH_IS_NOT_FILE"

	// Occurs when the file loader could not find the root folder.
	RootNotFoundCode string = "ROOT_NOT_FOUND"

	// Occurs when the file loader could not find the target .env file.
	EnvNotFoundCode string = "ENV_NOT_FOUND"
)
