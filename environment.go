package gofidential

// The [gofidential.Environment] struct is used to configure an environment.
//
// Constraints:
//   - "Name" must not be empty.
//   - "Name" must only contain lowercase letters and underscores.
//   - "Name" must not have leading or trailing underscores.
//   - "Name" must not consist solely of underscores.
//   - If specified, "OverridePath" must be a valid directory path (not a filepath).
//
// Notes:
//   - Both "OverridePath" and "IgnoreFilename" are optional.
type Environment struct {
	// The name of the environment. This field is required.
	Name string

	// An optional custom directory path to look for the .env file.
	OverridePath string

	// Determines the naming convention for the .env file.
	// If true, the filename will always be ".env". If false, the filename will include
	// the environment name (e.g., "dev.env", "test.env", "prod.env").
	IgnoreFilename bool
}
