# Troubleshooting - An Errors Comprehensive Guide
This guide is designed to help users identify and resolve common issues with the **GoFidential** package. Each error is 
assigned a unique code, making troubleshooting straightforward. With these codes, you can avoid diving into the source
code or scouring forums for solutions to simple problems.

Our goal is to document nearly all possible errors to enhance the user experience and streamline the development process 
for everyone using the package.

## Need Help?

If you’ve followed the troubleshooting steps and your issue is still unresolved, don’t worry — we’re here to help! Please 
follow these steps to make it easier for us to assist you:

### 1. **Check the Documentation**
   Before reaching out for help, make sure you've thoroughly read the documentation. It's possible that the solution to 
   your problem is already included, especially for common issues or specific configuration details.

### 2. **Check for Existing Issues**
   Take a moment to search the [GitHub repository's issues page](https://github.com/fueripe-desu/gofidential/issues) to see 
   if someone else has already encountered the same problem. If an existing issue is related to yours, you can add 
   additional context or follow the suggested solutions.

### 3. **Reproduce the Issue**
   Try to isolate the problem by creating a minimal reproduction of the issue. Simplifying your code or configuration can 
   help pinpoint what exactly is going wrong. This also makes it easier for others to diagnose the problem.

### 4. **Gather Information**
   When you’re ready to ask for help, please provide the following details to ensure we can assist you effectively:
   - **Error message**: Include the full error message you're seeing, including any stack traces or log output.
   - **Environment details**: Specify your operating system, Go version, and any other relevant software details.
   - **Steps to reproduce**: Describe the actions that lead to the issue, including any relevant code or configuration files.
   - **Relevant code**: If applicable, share the snippet of code or the `.env` file involved. Make sure to remove any 
   sensitive information (such as passwords or API keys) before sharing.

### 5. **Open an Issue**
   If your issue hasn’t been resolved after reviewing the documentation and searching existing issues, please open a new 
   issue on the [GitHub repository](https://github.com/fueripe-desu/gofidential/issues). Be sure to include all the gathered 
   information to help us understand and resolve your problem as quickly as possible.

### 6. **Be Patient**
   While we strive to respond to issues as quickly as possible, please remember that it may take some time for us to 
   investigate and provide a solution, especially if the issue is complex. We appreciate your patience and understanding!

---

## Contribute to the Troubleshooting Guide

If you'd like to contribute to improving the troubleshooting guide, check out the [CONTRIBUTING.md](./CONTRIBUTING.md) file 
for instructions on how to submit your changes.

---

## Loader Errors

### `MissingName (LOADER_MISSING_NAME)`

The `LOADER_MISSING_NAME` error occurs when the `Name` field of the `Environment` struct is either missing or an empty
string, as shown below:

```go
// The 'Name' field is empty.
env := gf.Environment{
  Name: "",
}

// The 'Name' field is missing entirely.
env := gf.Environment{}
```

#### How to Fix

To resolve this issue, specify a valid `Name` for the environment that follows these rules:

- Only contains lowercase letters and underscores.
- No leading or trailing underscores.
- Cannot consist entirely of underscores.

**Example of a valid `Name`:**

```go
env := gf.Environment{
  // The loader will look for the "example_name.env" file.
  Name: "example_name",
}
```
---

### `InvalidName (LOADER_INVALID_NAME)`

The `LOADER_INVALID_NAME` error occurs when the `Name` field of the `Environment` struct contains invalid characters, such as:

- Uppercase letters
- Numbers
- Special characters (other than underscores)

Here are some examples of invalid values for the `Name` field:

```go
// Contains an uppercase letter.
env := gf.Environment{
  Name: "Example",
}

// Contains numbers.
env := gf.Environment{
  Name: "example123",
}

// Contains special characters.
env := gf.Environment{
  Name: "ex@$ample",
}
```

---

#### How to Fix

To fix this issue, provide a valid `Name` that adheres to these rules:

- Only contains lowercase letters and underscores.
- No leading or trailing underscores.
- Cannot consist entirely of underscores.

**Example of a valid `Name`:**

```go
env := gf.Environment{
  // The loader will look for the "example_name.env" file.
  Name: "example_name",
}
```
---
### `UnderscoreOnlyName (LOADER_UNDERSCORE_ONLY_NAME)`

The `LOADER_UNDERSCORE_ONLY_NAME` error occurs when the `Name` field of the `Environment` struct is composed entirely 
of underscores, as shown below:

```go
env := gf.Environment{
  Name: "______",
}
```

---

#### How to Fix

Provide a valid `Name` that follows these rules:

- Only contains lowercase letters and underscores.
- No leading or trailing underscores.
- Cannot be composed entirely of underscores.

**Example of a valid `Name`:**

```go
env := gf.Environment{
  // The loader will look for the "example_name.env" file.
  Name: "example_name",
}
```
---
### `TrailingUnderscore (LOADER_TRAILING_UNDERSCORE)`

The `LOADER_TRAILING_UNDERSCORE` error occurs when the `Name` field of the `Environment` struct contains a trailing
underscore, as shown below:

```go
env := gf.Environment{
  Name: "example_",
}
```

---

#### How to Fix

Provide a valid `Name` that follows these rules:

- Only contains lowercase letters and underscores.
- No leading or trailing underscores.
- Cannot be composed entirely of underscores.

**Example of a valid `Name`:**

```go
env := gf.Environment{
  // The loader will look for the "example_name.env" file.
  Name: "example_name",
}
```
---

### `LeadingUnderscore (LOADER_LEADING_UNDERSCORE)`

The `LOADER_LEADING_UNDERSCORE` error occurs when the `Name` field of the `Environment` struct contains a leading
underscore, as shown below:

```go
env := gf.Environment{
  Name: "_example",
}
```

---

#### How to Fix

Provide a valid `Name` that follows these rules:

- Only contains lowercase letters and underscores.
- No leading or trailing underscores.
- Cannot be composed entirely of underscores.

**Example of a valid `Name`:**

```go
env := gf.Environment{
  // The loader will look for the "example_name.env" file.
  Name: "example_name",
}
```
---

### `InexistentDir (LOADER_INEXISTENT_DIR)`

The `LOADER_INEXISTENT_DIR` error occurs when the `OverridePath` field in the `Environment` struct points to a directory 
that does not exist.

---

#### A Simple Example

Imagine your file tree looks like this:

```
.
├── main.go
└── secrets
    └── dev.env
```

If you're executing from the root directory and want to load `dev.env` from `secrets`, your `OverridePath` should look 
like this:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets"
}
```

If you mistakenly use a non-existing directory:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "unknown"
}
```

This will cause the `LOADER_INEXISTENT_DIR` error because `./unknown/` does not exist.

---

#### A More Complex Example

Consider the following tree:

```
.
├── cmd
│   └── myapp
│       └── main.go
└── secrets
    └── dev.env
```

If you run the app from `cmd/myapp`, you need to adjust the `OverridePath` accordingly:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "../../secrets"
}
```

If you mistakenly set the path relative to `cmd/myapp`:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets"
}
```

The loader won't find the `.env` file because the path is incorrect.

---

#### How to Fix

To fix the issue, ensure that the `OverridePath` points to an existing directory, relative to the current working 
directory (cwd), or use an absolute path.

**Valid Example:**

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "../../secrets"
}
```

Make sure the directory you specify actually exists.

---

### `FailedToReadDir (LOADER_FAILED_TO_READ_DIR)`

The `LOADER_FAILED_TO_READ_DIR` error occurs when the application does not have permission to read the directory specified 
in the `OverridePath` field of the `Environment` struct.

---

#### Example

Suppose the directory permissions are as follows:

```
d--x--x--x  2 example-user example-user 4.0K Jan 20 14:48 secrets
```

And the environment is configured like this:

```go
env := gf.Environment{
  Name: "dev",
  // OverridePath points to a directory without read permissions.
  OverridePath: "secrets"
}
```

Since the application doesn't have permission to read the `secrets` directory, this will cause the 
`LOADER_FAILED_TO_READ_DIR` error.

---

#### How to Fix

To resolve the issue, change the permissions of the directory to allow read access, or update the `OverridePath` to point 
to a different directory with the appropriate permissions.

**Valid Example:**

After adjusting the directory permissions:

```
drwxr-xr-x  2 example-user example-user 4.0K Jan 20 14:48 secrets
```

The environment configuration should look like this:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets"
}
```
---

### `FailedToReadEnv (LOADER_FAILED_TO_READ_ENV)`

The `LOADER_FAILED_TO_READ_ENV` error occurs when the application does not have permission to read the target `.env` 
file. This issue can arise in three situations:

---

#### First Situation

When no directory path is specified in the `OverridePath` field in the `Environment` struct, the loader will first 
look for the `.env` file in the `./env/` directory. If the `.env` file exists there but does not have read permission, the 
loader will return the `LOADER_FAILED_TO_READ_ENV` error.

**Example:**

Consider the following file structure:

```
.
├── env
│   └── dev.env
└── main.go
```

The file permissions for `dev.env`:

```
---x--x--x  2 example-user example-user 4.0K Jan 20 14:48 dev.env
```

And the environment is configured like this:

```go
env := gf.Environment{
  Name: "dev",
}
```

Since no `OverridePath` is specified, the loader will look for the `dev.env` file in the `./env/` directory, but because the 
file does not have read permission, the `LOADER_FAILED_TO_READ_ENV` error will occur.

---

#### Second Situation

If no directory path is specified in the `OverridePath` field, the loader will first check the `./env/` directory for the 
`.env` file. If the `.env` file does not exist there, the loader will fallback to the current directory `./`. If the 
`.env` file exists in the current directory but the application does not have permission to read it, the 
`LOADER_FAILED_TO_READ_ENV` error will occur.

**Example:**

Consider the following file structure:

```
.
├── dev.env
└── main.go
```

The file permissions for `dev.env`:

```
---x--x--x  2 example-user example-user 4.0K Jan 20 14:48 dev.env
```

And the environment is configured like this:

```go
env := gf.Environment{
  Name: "dev",
}
```

Since no `OverridePath` is specified, the loader will first check `./env/`, and because the file doesn't exist there, it 
will fallback to the current directory `./`. Because the `dev.env` file in the current directory does not have read 
permission, the loader will return the `LOADER_FAILED_TO_READ_ENV` error.

---

#### Third Situation

When a directory path is specified in the `OverridePath` field, the loader will look for the `.env` file in the specified 
directory. If the file is present but lacks read permissions, the loader will return the `LOADER_FAILED_TO_READ_ENV` error.

**Example:**

Consider the following file structure:

```
.
├── main.go
└── secrets
    └── dev.env
```

The file permissions for `dev.env`:

```
---x--x--x  2 example-user example-user 4.0K Jan 20 14:48 dev.env
```

And the environment is configured like this:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets",
}
```

Since a path is specified in the `OverridePath` field, the loader will look for the `dev.env` file in 
the `./secrets/` directory. Since the `dev.env` file doesn't have read permissions, the loader will return 
the `LOADER_FAILED_TO_READ_ENV` error.

---

#### How to Fix

To resolve this issue, adjust the permissions of the target `.env` file to allow read access, or update your environment 
configuration to point to a `.env` file with appropriate permissions.

**Valid Example:**

Consider the following file structure:

```
.
├── dev.env
└── main.go
```

After adjusting the file permissions:

```
-rw-r--r--  2 example-user example-user 4.0K Jan 20 14:48 dev.env
```

The environment configuration should look like this:

```go
env := gf.Environment{
  Name: "dev",
}
```

Now, since the `dev.env` file has read permission, the loader will be able to read it without issues.

---

### `PathIsNotDir (LOADER_PATH_IS_NOT_DIR)`

The `LOADER_PATH_IS_NOT_DIR` error occurs when the `OverridePath` field defined in the `Environment` struct points to a 
file instead of a directory.

---

#### Example

Consider the following file structure:

```
.
├── secrets.env
└── main.go
```

And the environment is configured like this:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets.env"
}
```

In this case, the `OverridePath` field points to a file (`secrets.env`) rather than a directory. As a result, the loader 
will return the `LOADER_PATH_IS_NOT_DIR` error because it expects a directory but receives a file.

---

#### How to Fix

To resolve this issue, ensure that the path specified in the `OverridePath` field points to a valid directory.

**Valid Example:**

Consider the following file structure:

```
.
├── main.go
└── secrets
    └── dev.env
```

The environment configuration should be updated as follows:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets"
}
```

Now, since `./secrets` is a directory and not a file, the application should be able to read it without any issues.

---

### `PathIsNotFile (LOADER_PATH_IS_NOT_FILE)`

The `LOADER_PATH_IS_NOT_FILE` error is an internal error that occurs when the loader is not calculating the expected
filename correctly.

If you encounter this error, please open an issue [here](https://github.com/fueripe-desu/gofidential/issues) and provide the 
following contextual information to help us assist you:

- Your operating system
- Your Golang version
- Snippets of code that may be relevant to the issue

This will help us understand the problem better and work towards a solution faster.

---

### `EnvNotFound (LOADER_ENV_NOT_FOUND)`

The `LOADER_ENV_NOT_FOUND` error occurs when the loader cannot find the target `.env` file in the expected directory. This 
can happen in two situations:

---

#### First Situation

If no `OverridePath` is specified, the loader will first check the `./env/` directory for the `.env` file. If it's not
there, the loader will look in the current directory (`./`). If the file is not found in either location, the loader 
will return the `LOADER_ENV_NOT_FOUND` error.

**Example:**

Consider the following file structure:

```
.
├── prod.env
└── main.go
```

And the environment is configured like this:

```go
env := gf.Environment{
  Name: "dev",
}
```

Since no `OverridePath` is specified, the loader will check the `./env/` directory. If the file doesn't exist there, it 
will fall back to the current directory `./`. As the `dev.env` file is missing from both, the 
loader returns the `LOADER_ENV_NOT_FOUND` error.

---

#### Second Situation

When an `OverridePath` is specified, the loader will look for the `.env` file in that specific directory. If the file 
doesn't exist there, the loader will return the `LOADER_ENV_NOT_FOUND` error.

**Example:**

Consider the following file structure:

```
.
├── main.go
└── secrets
    └── prod.env
```

The environment is configured as:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets",
}
```

The loader will search for `dev.env` in the `./secrets/` directory. Since it does not exist, the loader will 
return the `LOADER_ENV_NOT_FOUND` error.

---

#### How to Fix

To resolve this issue, ensure the `.env` file is named according to the `Name` field of the `Environment` struct, and 
that it exists in the correct directory.

**Valid Example:**

Consider the following file structure:

```
.
├── main.go
└── secrets
    └── dev.env
```

The environment configuration should be updated as:

```go
env := gf.Environment{
  Name: "dev",
  OverridePath: "secrets"
}
```

Since `dev.env` exists under the `./secrets/` directory, the application will be able to locate it without issues.

## Parser Errors

### `MissingAssignment (PARSER_MISSING_ASSIGNMENT)`

The `PARSER_MISSING_ASSIGNMENT` error occurs when a key-value pair in the `.env` file is missing the assignment 
operator (`=`), which separates the key from its value. 

**Example of Invalid Syntax:**

```
EXAMPLE_KEY"example value"
```

In this case, the `=` is missing, causing the parser to return the `PARSER_MISSING_ASSIGNMENT` error.

---

#### How to Fix

To fix this issue, ensure that every key-value pair in the `.env` file includes an assignment operator (`=`) between 
the key and the value.

**Valid Example:**

```
EXAMPLE_KEY="example value"
```

By adding the `=`, the parser will be able to correctly interpret the key-value pair.

---

### `MissingKey (PARSER_MISSING_KEY)`

The `PARSER_MISSING_KEY` error occurs when a key-value pair in the `.env` file is missing a **key**—the identifier 
that appears before the assignment operator (`=`).

---

#### Example of Invalid Syntax

```env
="example value"
```

In this example, the parser cannot interpret the key-value pair because the **key** is missing.

---

#### How to Fix

Ensure that every key-value pair in the `.env` file includes a **valid key** before the assignment operator (`=`).

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By providing a proper key, the parser will be able to process the key-value pair correctly.

--- 

### `MissingValue (PARSER_MISSING_VALUE)`

The `PARSER_MISSING_VALUE` error occurs when a key-value pair in the `.env` file is missing the **value**—the part that 
is assigned to the key and follows the assignment operator (`=`).

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY=
```

In this case, the key `EXAMPLE_KEY` is defined, but no value is assigned to it, causing the parser to return an error.

---

#### How to Fix

To resolve this issue, ensure that every key-value pair in the `.env` file includes a valid **value** after the assignment 
operator (`=`). It is required to enclose the value in double quotes to comply with the parser rules.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By assigning a valid value, the parser will correctly interpret the key-value pair.

---

### `EmptyValue (PARSER_EMPTY_VALUE)`

The `PARSER_EMPTY_VALUE` error occurs when a key-value pair in the `.env` file has an empty value enclosed in double quotes.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY=""
```

In this example, the key `EXAMPLE_KEY` is defined, but its value is an empty string, causing the parser to return an error.

---

#### How to Fix

To resolve this issue, ensure that every key-value pair in the `.env` file includes a non-empty value after the 
assignment operator (`=`). If the value is intentionally unused, you can either assign a placeholder value or remove the 
key entirely.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```
--- 

### `SingleQuotedValue (PARSER_SINGLE_QUOTED_VALUE)`

The `PARSER_SINGLE_QUOTED_VALUE` error occurs when a key-value pair in the `.env` file has its value enclosed in single 
quotes, which is not allowed.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY='example value'
```

In this example, the key `EXAMPLE_KEY` is defined, but its value is enclosed in single quotes, violating the parser's 
strict rules.

---

#### How to Fix

To resolve this issue, replace the single quotes around the value with double quotes.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By using double quotes, the key-value pair will adhere to the parser's requirements.

---

### `UnquotedValue (PARSER_UNQUOTED_VALUE)`

The `PARSER_UNQUOTED_VALUE` error occurs when a key-value pair in the `.env` file does not have the value enclosed 
in double quotes, which is not allowed.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY=example value
```

In this example, the key `EXAMPLE_KEY` is defined, but the value is unquoted, violating the parser's strict rules.

---

#### How to Fix

To resolve this issue, enclose the value in double quotes.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By enclosing the value in double quotes, the key-value pair will comply with the parser's requirements.

---

### `UnterminatedQuotes (PARSER_UNTERMINATED_QUOTES)`

The `PARSER_UNTERMINATED_QUOTES` error occurs when a key-value pair in the `.env` file starts with an opening double quote 
but is missing the closing double quote for the value.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY="example value
```

In this example, the key `EXAMPLE_KEY` is defined, but the value is missing the closing double quote. Causing the parser
to return an error.

---

#### How to Fix

Ensure that the value is enclosed with both opening and closing double quotes.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By properly enclosing the value in double quotes, the key-value pair will meet the parser's requirements.

---  

### `SpacedSeparator (PARSER_SPACED_SEPARATOR)`

The `PARSER_SPACED_SEPARATOR` error occurs when a key-value pair in the `.env` file contains spaces around the assignment 
operator (`=`).

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY = "example value"
```

In this example, the key `EXAMPLE_KEY` is defined, but spaces are present around the assignment operator (`=`), which 
violates the parser's strict rules.

---

#### How to Fix

Ensure that the assignment operator is immediately after the key and before the value, without any spaces in between.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By removing spaces around the assignment operator, the key-value pair will comply with the parser's requirements.

---

### `LeadingSpace (PARSER_LEADING_SPACE)`

The `PARSER_LEADING_SPACE` error occurs when a key-value pair in the `.env` file contains leading spaces before the key.

---

#### Example of Invalid Syntax

```env
                 EXAMPLE_KEY = "example value"
```

In this example, the key `EXAMPLE_KEY` is defined, but leading spaces appear before the key, violating the parser's 
strict rules.

---

#### How to Fix

Ensure that there are no leading spaces before the key.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By removing the leading spaces before the key, the key-value pair will comply with the parser's requirements.

---

### `UnallowedEscape (PARSER_UNALLOWED_ESCAPE)`

The `PARSER_UNALLOWED_ESCAPE` error occurs when a key-value pair in the `.env` file contains an escape sequence in the 
value that is not permitted. The parser strictly allows only the following escape characters:

- `\n` for a newline,
- `\"` for a double quote,
- `\\` for a backslash.

All other escape sequences are considered invalid.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY="\texample value\n"
```

In this example, the key `EXAMPLE_KEY` is defined, but the value includes the unallowed tab escape character (`\t`), 
violating the parser's strict rules.

---

#### How to Fix

Remove any unallowed escape characters from the value.

**Valid Example:**

```env
EXAMPLE_KEY="example value\n"
```

By using only permitted escape characters, the key-value pair will comply with the parser's requirements.

---

### `UnescapedQuoteChar (PARSER_UNESCAPED_QUOTE_CHAR)`

The `PARSER_UNESCAPED_QUOTE_CHAR` error occurs when a key-value pair in the `.env` file contains multiple unescaped 
double quotes within the value. This typically happens when the user forgets to escape the double quote character (`\"`) that 
appears as part of the content.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY="Then he said: "I might not be able to come tomorrow."."
```

In this example, the key `EXAMPLE_KEY` is defined, but the value contains unescaped double quotes (`"`) within the 
string, causing the parser to return an error.

---

#### How to Fix

Escape any double quotes that are part of the literal value by using a backslash (`\"`).

**Valid Example:**

```env
EXAMPLE_KEY="Then he said: \"I might not be able to come tomorrow.\"."
```

By escaping the double quotes properly, the key-value pair will adhere to the parser's strict requirements.

--- 

### `LowercaseKey (PARSER_LOWERCASE_KEY)`

The `PARSER_LOWERCASE_KEY` error occurs when a key in the `.env` file contains one or more lowercase characters. The parser 
enforces strict rules requiring all keys to be written entirely in uppercase.

---

#### Example of Invalid Syntax

```env
example_key="example value"
```

In this example, the key `example_key` contains lowercase letters, which violates the parser's strict requirements.

---

#### How to Fix

Ensure that the key is written entirely in uppercase letters.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By using uppercase letters exclusively, the key-value pair will comply with the parser's rules.

--- 

### `NumericKeyChars (PARSER_NUMERIC_KEY_CHARS)`

The `PARSER_NUMERIC_KEY_CHARS` error occurs when a key in the `.env` file contains one or more numeric characters. The 
parser enforces strict rules requiring all keys to be written only using uppercase letters and underscores, without 
any numeric characters.

---

#### Example of Invalid Syntax

```env
EX4MPL3_K3Y="example value"
```

In this example, the key `EX4MPL3_K3Y` contains numeric characters, which violates the parser's strict requirements.

---

#### How to Fix

Ensure that the key is written using only uppercase letters and underscores, avoiding numeric characters entirely.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By following these rules, the key-value pair will comply with the parser's strict requirements.

--- 

### `LeadingUnderscore (PARSER_LEADING_UNDERSCORE)`

The `PARSER_LEADING_UNDERSCORE` error occurs when a key in the `.env` file starts with a leading underscore (`_`).

---

#### Example of Invalid Syntax

```env
_EXAMPLE_KEY="example value"
```

In this example, the key `_EXAMPLE_KEY` begins with a leading underscore, which violates the parser's strict requirements.

---

#### How to Fix

To resolve this issue, remove the leading underscore from the key.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By ensuring that keys do not start with leading underscores, the key-value pair will adhere to the parser's strict 
requirements.

--- 

### `TrailingUnderscore (PARSER_TRAILING_UNDERSCORE)`

The `PARSER_TRAILING_UNDERSCORE` error occurs when a key in the `.env` file ends with a trailing underscore (`_`).

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY_="example value"
```

In this example, the key `EXAMPLE_KEY_` ends with a trailing underscore, which violates the parser's strict requirements.

---

#### How to Fix

To resolve this issue, remove the trailing underscore from the key.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By ensuring that keys do not end with trailing underscores, the key-value pair will adhere to the parser's strict 
requirements.

--- 

### `InvalidKeyChars (PARSER_INVALID_KEY_CHARS)`

The `PARSER_INVALID_KEY_CHARS` error occurs when a key in the `.env` file contains characters other than uppercase 
letters and underscores. The parser strictly enforces these rules, rejecting keys with numbers, lowercase letters, or 
special characters.

---

#### Example of Invalid Syntax

```env
EX#@!MPLE_KEY="example value"
```

In this example, the key `EX#@!MPLE_KEY` contains special characters, violating the parser's requirements.

---

#### How to Fix

To resolve this issue, rewrite the key using only uppercase letters and underscores.

**Valid Example:**

```env
EXAMPLE_KEY="example value"
```

By adhering to these rules, the key-value pair will meet the parser's strict requirements.

--- 

### `InlineComment (PARSER_INLINE_COMMENT)`

The `PARSER_INLINE_COMMENT` error occurs when a key-value pair in the `.env` file includes an inline comment on the 
same line as the definition. The parser enforces a strict rule prohibiting inline comments.

---

#### Example of Invalid Syntax

```env
EXAMPLE_KEY="example value" # This key is merely an example
```

In this example, while the key-value pair `EXAMPLE_KEY="example value"` is valid, the inline comment at the end violates 
the parser's requirements.

---

#### How to Fix

To resolve this issue, place comments on their own line above the key-value pair.

**Valid Example:**

```env
# This key is merely an example
EXAMPLE_KEY="example value"
```

By separating comments from key-value pairs, the `.env` file will adhere to the parser's strict rules.

--- 

---

### `MultilineValue (PARSER_MULTILINE_VALUE)`

The `PARSER_MULTILINE_VALUE` error occurs when a key-value pair in the `.env` file includes a value that spans 
multiple lines. The parser strictly prohibits multiline values to maintain consistency and simplicity in 
the `.env` file format.

If you need to represent a line break within a value, use the `\n` escape character instead.

---

#### Example of Invalid Syntax

```env
MULTILINE_KEY="This is an example of \
a multiline \
value \
that is split \
across \
multiple lines"
```

In this example, the key `MULTILINE_KEY` contains a multiline value, which violates the parser's requirements.

---

#### How to Fix

Ensure all values are written as single-line strings.

**Valid Example:**

```env
EXAMPLE_KEY="This is an example of a single-line value"
```

If you need to include line breaks, use the `\n` escape character:

```env
EXAMPLE_KEY="This is an example of a value with a line break.\nThis is the next line."
```

By following these rules, the `.env` file will comply with the parser's strict requirements.

--- 

## Reflector Errors

### `SIsNil (REFLECTOR_S_IS_NIL)`

The `REFLECTOR_S_IS_NIL` error occurs when the `s` parameter passed to the `Load()` function is `nil`.

#### Example of Invalid Syntax

```go
err := Load(nil, env)
```

In this example, the `Load()` function receives `nil` instead of a valid struct pointer, which causes the error.

#### How to Fix

To resolve this issue, ensure that a valid, initialized struct pointer is passed to the `Load()` function.

**Valid Example:**

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // This will load "EXAMPLE_KEY" from the .env file.
  ExampleKey string
}

func main() {
  env := gf.Environment{
    Name: "dev",
  }

  // Create an empty instance of the Secrets struct.
  s := Secrets{}

  // Pass the struct pointer to the Load() function.
  err := Load(&s, env)

  if err != nil {
    log.Fatal("Could not read .env file.")
  }
}
```

In this example, we create an empty `Secrets` struct and pass its pointer to the `Load()` function. This resolves the 
issue and ensures the environment variables are correctly loaded into the struct.

---

### `SIsNotPtr (REFLECTOR_IS_S_NOT_PTR)`

The `REFLECTOR_S_IS_NOT_PTR` error occurs when the `s` parameter passed to the `Load()` function is not a pointer.

#### Example of Invalid Syntax

```go
err := Load(1, env)
```

In this example, the `Load()` function receives `1` instead of a valid struct pointer, which causes the error.

#### How to Fix

To resolve this issue, ensure that a valid, initialized struct pointer is passed to the `Load()` function.

**Valid Example:**

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // This will load "EXAMPLE_KEY" from the .env file.
  ExampleKey string
}

func main() {
  env := gf.Environment{
    Name: "dev",
  }

  // Create an empty instance of the Secrets struct.
  s := Secrets{}

  // Pass the struct pointer to the Load() function.
  err := Load(&s, env)

  if err != nil {
    log.Fatal("Could not read .env file.")
  }
}
```

In this example, we create an empty `Secrets` struct and pass its pointer to the `Load()` function. This resolves the 
issue and ensures the environment variables are correctly loaded into the struct.

---

### `NilStructPtr (REFLECTOR_NIL_STRUCT_PTR)`

The `REFLECTOR_NIL_STRUCT_PTR` error occurs when the `s` parameter passed to the `Load()` function is a `nil` pointer. The 
`Load()` function requires a valid, initialized struct pointer to populate the environment variables.

#### Example of Invalid Syntax

```go
var s *Secrets = nil
err := Load(s, env)
```

In this example, the `Load()` function receives a `nil` pointer (`s`) instead of a valid struct pointer, which causes 
the error.

#### How to Fix

To resolve this issue, ensure that the struct pointer is properly initialized before passing it to the `Load()` function.

**Valid Example:**

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // This will load "EXAMPLE_KEY" from the .env file.
  ExampleKey string
}

func main() {
  env := gf.Environment{
    Name: "dev",
  }

  // Create an empty instance of the Secrets struct.
  s := Secrets{}

  // Pass the struct pointer to the Load() function.
  err := Load(&s, env)

  if err != nil {
    log.Fatal("Could not read .env file.")
  }
}
```

In this example, we create an empty `Secrets` struct and pass its pointer (`&s`) to the `Load()` function. This ensures 
that the environment variables are correctly loaded into the struct, resolving the error.

---

### `InvalidStructPtr (REFLECTOR_INVALID_STRUCT_PTR)`

The `REFLECTOR_INVALID_STRUCT_PTR` error occurs when the `s` parameter passed to the `Load()` function is a pointer that 
does not point to a valid struct. The `Load()` function requires a pointer to a valid, initialized struct in order to 
correctly populate the environment variables.

#### Example of Invalid Syntax

```go
s := []string{}
err := Load(&s, env)
```

In this example, the `Load()` function receives a pointer to a slice (`[]string`) instead of a valid struct pointer, which 
causes the error.

#### How to Fix

To resolve this issue, ensure that a pointer to a valid struct is passed to the `Load()` function, not a pointer to a 
non-struct type.

**Valid Example:**

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // This will load "EXAMPLE_KEY" from the .env file.
  ExampleKey string
}

func main() {
  env := gf.Environment{
    Name: "dev",
  }

  // Create an empty instance of the Secrets struct.
  s := Secrets{}

  // Pass the struct pointer to the Load() function.
  err := Load(&s, env)

  if err != nil {
    log.Fatal("Could not read .env file.")
  }
}
```

In this example, we create an empty `Secrets` struct and pass its pointer (`&s`) to the `Load()` function. This ensures that 
the environment variables are correctly loaded into the struct, resolving the error.

By ensuring that the `s` parameter is a pointer to a struct, the error is avoided, and the `Load()` function can successfully 
load environment variables.

---

### `InvalidStructPtr (REFLECTOR_INVALID_STRUCT_PTR)`

The `REFLECTOR_INVALID_STRUCT_PTR` error occurs when the `s` parameter passed to the `Load()` function is a pointer that 
does not point to a valid struct. The `Load()` function requires a pointer to a valid, initialized struct in order to 
correctly populate the environment variables.

#### Example of Invalid Syntax

```go
s := []string{}
err := Load(&s, env)
```

In this example, the `Load()` function receives a pointer to a slice (`[]string`) instead of a valid struct pointer, which 
causes the error.

#### How to Fix

To resolve this issue, ensure that a pointer to a valid struct is passed to the `Load()` function, not a pointer to a 
non-struct type.

**Valid Example:**

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // This will load "EXAMPLE_KEY" from the .env file.
  ExampleKey string
}

func main() {
  env := gf.Environment{
    Name: "dev",
  }

  // Create an empty instance of the Secrets struct.
  s := Secrets{}

  // Pass the struct pointer to the Load() function.
  err := Load(&s, env)

  if err != nil {
    log.Fatal("Could not read .env file.")
  }
}
```

In this example, we create an empty `Secrets` struct and pass its pointer (`&s`) to the `Load()` function. This ensures that 
the environment variables are correctly loaded into the struct, resolving the error.

By ensuring that the `s` parameter is a pointer to a struct, the error is avoided, and the `Load()` function can successfully 
load environment variables.

---

### `InvalidEnvData (REFLECTOR_INVALID_ENV_DATA)`

The `REFLECTOR_INVALID_ENV_DATA` error is an internal error that occurs when the `data` parameter, which is a 
`map[string]string` containing all key-value pairs generated by the parser, is found to be `nil`. This issue typically 
happens when there is a bug in the loader or parser.

If you encounter this error, please open an issue [here](https://github.com/fueripe-desu/gofidential/issues) and provide the 
following contextual information to help us assist you:

- Your operating system
- Your Golang version
- Snippets of code that may be relevant to the issue

This will help us understand the problem better and work towards a solution faster.

---

### `DuplicateKeys (REFLECTOR_DUPLICATE_KEYS)`

The `REFLECTOR_DUPLICATE_KEYS` error is an internal error that occurs when the `data` parameter, a `map[string]string` 
containing all key-value pairs generated by the parser, contains conflicting keys. This issue typically arises when the 
parser fails to correctly block duplicate keys due to a bug or unexpected behavior.

If you encounter this error, please open an issue [here](https://github.com/fueripe-desu/gofidential/issues) and provide the 
following contextual information to help us assist you:

- Your operating system
- Your Golang version
- Snippets of code that may be relevant to the issue

This will help us understand the problem better and work towards a solution faster.

---

### `MissingField (REFLECTOR_MISSING_FIELD)`

The `REFLECTOR_MISSING_FIELD` error occurs when the struct intended to be populated is missing a field that corresponds 
to a key-value pair in the `.env` file. This can happen in two scenarios:

---

#### First Situation

Consider the following `.env` file:

```
APP_NAME="MyApp"
DEBUG="true"
BASE_URL="www.mybaseurl.com"
```

And your struct looks like this:

```go
type Secrets struct {
  AppName string
  Debug bool
}
```

In this case, the struct is missing a field for `BASE_URL`, so the reflector expects a `BaseUrl` field but doesn't 
find it, leading to the error.

---

#### Second Situation

Consider the following `.env` file:

```
APP_NAME="MyApp"
DEBUG="true"
BASE_URL="www.mybaseurl.com"
```

And your struct looks like this:

```go
type Secrets struct {
  AppName string
  Debug bool
  Baseurl string
}
```

Here, the field for `BASE_URL` is incorrectly named `Baseurl`, missing the capital `U`. The correct name should be `BaseUrl`, 
and because of this mismatch, the reflector won't find the expected field, causing the error.

---

#### How to Fix

To resolve this issue, ensure that all struct fields match the names of the `.env` keys, converted to Pascal case with the 
first letter capitalized (so the field is exported).

**Valid Example:**

```go
type Secrets struct {
  AppName string
  Debug bool
  BaseUrl string
}
```

In this example, the `Secrets` struct is defined correctly, with all fields exported and named properly in Pascal case. By 
ensuring that all fields match the expected names, the reflector will be able to populate them without any issues.

---

### `UnsupportedField (REFLECTOR_UNSUPPORTED_FIELD)`

The `REFLECTOR_UNSUPPORTED_FIELD` error occurs when the struct intended to be populated has a field of a type that is not 
supported by the reflector for mapping key-value pairs from the `.env` file. 

The reflector only supports native Go types, including:

- **Numeric types:** `int` (8, 16, 32, 64), `uint` (8, 16, 32, 64), `float` (32, 64), `complex` (64, 128)
- **String type:** `string`
- **Boolean type:** `bool`
- **Time type:** `time.Time`

---

#### Example of Invalid Syntax

```go
type PaymentInfo struct {
  CreditCard string
  ZipCode string
}

type Secrets struct {
  Name string
  Age int
  Payment *PaymentInfo
}
```

In this example, the `Secrets` struct contains a field of type `PaymentInfo`, which is an unsupported custom type. The 
reflector does not support complex types such as structs, slices, maps, or pointers. As a result, this causes the error.

---

#### How to Fix

To resolve this issue, avoid using nested structs or other complex types in the struct passed to the reflector. Instead, 
consolidate all fields into a single struct using only the supported primitive types.

Supported types include:
- **Numbers:** `int`, `uint`, `float`, `complex`
- **Text:** `string`
- **Booleans:** `bool`
- **Date/time:** `time.Time`

**Valid Example:**

```go
type Secrets struct {
  Name        string
  Age         int
  CreditCard  string
  ZipCode     string
}
```

In this example, the `Secrets` struct uses only supported types, ensuring the reflector can populate the struct without 
issues. By simplifying the structure and using primitive types, the error can be avoided.

---

### `UnsettableField (REFLECTOR_UNSETTABLE_FIELD)`

The `REFLECTOR_UNSETTABLE_FIELD` error is an internal error that occurs when the reflector (due to a bug) incorrectly 
allows either a pointer that does not properly reference a struct or attempts to set unexported (private) fields in the 
provided struct pointer.

If you encounter this error, please open an issue [here](https://github.com/fueripe-desu/gofidential/issues) and provide the 
following contextual information to help us assist you:

- Your operating system
- Your Golang version
- Snippets of code that may be relevant to the issue

This will help us understand the problem better and work towards a solution faster.

---

### `UnexportedField (REFLECTOR_UNEXPORTED_FIELD)`

The `REFLECTOR_UNEXPORTED_FIELD` error occurs when the struct intended to be populated contains unexported fields.

---

#### Example of Invalid Syntax

```go
type Secrets struct {
  name        string
  age         int
  nationality string
}
```

In this example, the `Secrets` struct contains unexported fields (fields with names starting with lowercase letters). These 
fields cannot be accessed or populated by the reflector, causing the error.

---

#### How to Fix

To resolve this issue, ensure that all fields in the struct are properly exported by starting their names with an uppercase 
letter.

**Valid Example:**

```go
type Secrets struct {
  Name        string
  Age         int
  Nationality string
}
```

In this example, the `Secrets` struct uses only exported fields (with names starting with uppercase letters). This allows the 
reflector to access and populate them without any issues.

---

### `InvalidInt (REFLECTOR_INVALID_INT)`

The `REFLECTOR_INVALID_INT` error occurs when the value from the `.env` file intended to populate an `int` (8, 16, 32, 64) 
field in the struct is not a valid integer string literal.

---

#### Example of Invalid Syntax

Consider the following `.env` file:

```
AGE="unknown"
```

And this struct:

```go
type Secrets struct {
  Age int
}
```

In this example, the `Secrets` struct contains an `int` field, but the value associated with the `AGE` key in the `.env` file 
cannot be converted to an integer because `"unknown"` is not a valid integer string literal. This mismatch causes the error.

---

#### How to Fix

To resolve this issue, ensure that the key-value pair in the `.env` file contains a valid integer string literal. A valid 
integer is any whole number, positive or negative, including zero (e.g., `-1`, `0`, `42`).

**Valid Example:**

Consider this corrected `.env` file:

```
AGE="18"
```

And this struct:

```go
type Secrets struct {
  Age int
}
```

In this example, the `AGE` key in the `.env` file contains a valid integer string literal (`"18"`). The reflector 
successfully parses this value, casts it to an integer, and populates the `Age` field in the `Secrets` struct without errors.

---

### `InvalidUint (REFLECTOR_INVALID_UINT)`

The `REFLECTOR_INVALID_UINT` error occurs when the value from the `.env` file intended to populate a `uint` (8, 16, 32, 64) 
field in the struct is not a valid unsigned integer string literal.

---

#### Example of Invalid Syntax

Consider the following `.env` file:

```
AGE="unknown"
```

And this struct:

```go
type Secrets struct {
  Age uint
}
```

In this example, the `Secrets` struct contains a `uint` field, but the value associated with the `AGE` key in 
the `.env` file cannot be converted to an unsigned integer because `"unknown"` is not a valid unsigned integer string 
literal. This mismatch causes the error.

---

#### How to Fix

To resolve this issue, ensure that the key-value pair in the `.env` file contains a valid unsigned integer string 
literal. Valid unsigned integers include positive whole numbers (e.g., `0`, `1`, `42`).

**Valid Example:**

Consider this corrected `.env` file:

```
AGE="18"
```

And this struct:

```go
type Secrets struct {
  Age uint
}
```

In this example, the `AGE` key in the `.env` file contains a valid unsigned integer string literal (`"18"`). The 
reflector can successfully parse this value, convert it to an unsigned integer, and populate the `Age` field 
in the `Secrets` struct without errors.

---

### `InvalidFloat (REFLECTOR_INVALID_FLOAT)`

The `REFLECTOR_INVALID_FLOAT` error occurs when the value from the `.env` file intended to populate 
a `float32` or `float64` field in the struct is not a valid float string literal.

---

#### Example of Invalid Syntax

Consider the following `.env` file:

```
BALANCE="unknown"
```

And this struct:

```go
type Secrets struct {
  Balance float32
}
```

In this example, the `Secrets` struct contains a `float32` field, but the value associated with the `BALANCE` key in 
the `.env` file cannot be converted to a float because `"unknown"` is not a valid float string literal. This mismatch 
causes the error.

---

#### How to Fix

To resolve this issue, ensure that the key-value pair in the `.env` file contains a valid float string literal, which 
can include integers, decimals, and scientific notation formats.

**Valid Example:**

Consider this corrected `.env` file:

```
BALANCE="1800.5"
```

And this struct:

```go
type Secrets struct {
  Balance float32
}
```

In this example, the `BALANCE` key in the `.env` file contains a valid float string literal (`"1800.5"`). The reflector 
can successfully parse this value, convert it to a float, and populate the `Balance` field in the `Secrets` struct 
without errors.

---

### `InvalidComplex (REFLECTOR_INVALID_COMPLEX)`

The `REFLECTOR_INVALID_COMPLEX` error occurs when the value from the `.env` file intended to 
populate a `complex64` or `complex128` field in the struct is not a valid complex number string literal.

---

#### Example of Invalid Syntax

Consider the following `.env` file:

```
COMPLEX="unknown"
```

And this struct:

```go
type Secrets struct {
  Complex complex64
}
```

In this example, the `Secrets` struct contains a `complex64` field, but the value associated with the `COMPLEX` key in 
the `.env` file cannot be converted to a complex number because `"unknown"` is not a valid complex number string 
literal. This mismatch causes the error.

---

#### How to Fix

To resolve this issue, ensure that the key-value pair in the `.env` file contains a valid complex number string literal 
in the format `<real>+<imaginary>i`.

**Valid Example:**

Consider this corrected `.env` file:

```
COMPLEX="3+4i"
```

And this struct:

```go
type Secrets struct {
  Complex complex64
}
```

In this example, the `COMPLEX` key in the `.env` file contains a valid complex number string literal (`"3+4i"`). The 
reflector can successfully parse this value, convert it to a complex number, and populate the `Complex` field 
in the `Secrets` struct without errors.

---

### `InvalidBool (REFLECTOR_INVALID_BOOL)`

The `REFLECTOR_INVALID_BOOL` error occurs when the value from the `.env` file intended to populate a `bool` field in 
the struct is not a valid boolean string literal.

---

#### Example of Invalid Syntax

Consider the following `.env` file:

```
DEBUG="unknown"
```

And this struct:

```go
type Secrets struct {
  Debug bool
}
```

In this example, the `Secrets` struct contains a `bool` field, but the value associated with the `DEBUG` key in 
the `.env` file cannot be converted to a boolean because `"unknown"` is not a valid boolean string literal. This 
mismatch causes the error.

---

#### How to Fix

To resolve this issue, ensure that the key-value pair in the `.env` file contains a valid boolean string literal. Boolean 
literals are represented as the strings `"true"` or `"false"`, case-insensitive.

**Valid Example:**

Consider this corrected `.env` file:

```
DEBUG="true"
```

And this struct:

```go
type Secrets struct {
  Debug bool
}
```

In this example, the `DEBUG` key in the `.env` file contains a valid boolean string literal (`"true"`). The reflector 
can successfully parse this value, convert it to a boolean, and populate the `Debug` field in the `Secrets` struct 
without errors.

---

### `InvalidTime (REFLECTOR_INVALID_TIME)`

The `REFLECTOR_INVALID_TIME` error occurs when the value from the `.env` file intended to populate a `time.Time` field in 
the struct is not a valid `RFC3339Nano` string literal.

---

#### Example of Invalid Syntax

Consider the following `.env` file:

```
CREATED_AT="unknown"
```

And this struct:

```go
type Secrets struct {
  CreatedAt time.Time
}
```

In this example, the `Secrets` struct contains a `time.Time` field, but the value associated with the `CREATED_AT` key in 
the `.env` file cannot be converted to a `time.Time` value because `"unknown"` is not a valid `RFC3339Nano` string 
literal. This mismatch causes the error.

---

#### How to Fix

To resolve this issue, ensure that the key-value pair in the `.env` file contains a valid `RFC3339Nano` string literal. The 
correct format for time literals is:

```
YYYY-MM-DDTHH:MM:SS.ssssssZ
```

Where:

- `YYYY`: The four-digit year (e.g., `2023`)
- `MM`: The two-digit month (01–12) (e.g., `01` for January)
- `DD`: The two-digit day of the month (01–31) (e.g., `18`)
- `T`: The separator between the date and time components
- `HH`: The two-digit hour in 24-hour format (00–23) (e.g., `15`)
- `MM`: The two-digit minute (00–59) (e.g., `04`)
- `SS`: The two-digit second (00–59) (e.g., `05`)
- `.ssssss`: (Optional) The fractional part of the second, expressed in microseconds or nanoseconds. This part is optional 
but can be included for higher precision (e.g., `123456` represents 123,456 microseconds, or 0.123456 seconds).
- `Z`: The UTC time zone indicator (indicating the time is in UTC). You could also use an offset like `+02:00` if the time 
zone isn't UTC.

**Valid Example:**

Consider this corrected `.env` file:

```
CREATED_AT="2023-01-18T15:04:05.123456Z"
```

And this struct:

```go
type Secrets struct {
  CreatedAt time.Time
}
```

In this example, the `CREATED_AT` key in the `.env` file contains a valid `RFC3339Nano` string literal 
(`"2023-01-18T15:04:05.123456Z"`). The reflector can successfully parse this value, convert it to a `time.Time`, and 
populate the `CreatedAt` field in the `Secrets` struct without errors.

If you don't need nanoseconds, you can omit the fractional part, so something like `"2023-01-18T15:04:05Z"` would also be 
valid.

