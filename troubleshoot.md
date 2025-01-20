# Troubleshooting - An Errors Comprehensive Guide
This guide is designed to help users identify and resolve common issues with the **GoFidential** package. Each error is assigned a unique code, making troubleshooting straightforward. With these codes, you can avoid diving into the source code or scouring forums for solutions to simple problems.

Our goal is to document nearly all possible errors to enhance the user experience and streamline the development process for everyone using the package.

## Loader Errors

### `MissingName (LOADER_MISSING_NAME)`

The `LOADER_MISSING_NAME` error occurs when using the `Load()` function without specifying a valid name for the environment. Below is the definition of the `Load()` function:

```go
func Load(s any, env Environment) error
```

As shown, this function requires an `Environment` struct, which is used to configure the environment. The `Environment` struct allows you to define the target environment (e.g., `dev`, `test`, `prod`) and other optional parameters. Here's its definition:

```go
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
```

The `Name` field in this struct is **required**. If it is missing or set to an empty string, the loader will return the `LOADER_MISSING_NAME` error. To fix this, provide a valid name that adheres to the following rules:

- Only contains lowercase letters and underscores.
- No leading or trailing underscores.

Here’s an example of how to resolve this error:

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // A key to be loaded from the .env file.
  ExampleField string
}

func main() {
  // Define the environment with a valid name.
  env := gf.Environment{
    // This name ensures the loader will look for "example.env".
    Name: "example",
  }

  s := Secrets{}

  // Load the .env file into the struct.
  err := gf.Load(&s, env)

  if err != nil {
    fmt.Println("Failed to load .env file:", err)
  }
}
```

#### Key Details in the Example

1. Correctly defining the `Name` field in the `Environment` struct:
   ```go
   env := gf.Environment{
     Name: "example",
   }
   ```

2. Passing the `Environment` struct to the `Load()` function:
   ```go
   err := gf.Load(&s, env)
   ```
---

### `InvalidName (LOADER_INVALID_NAME)`

The `LOADER_INVALID_NAME` error occurs when using the `Load()` function with an invalid value for the `Name` field in the `Environment` struct. Below is the definition of the `Load()` function:

```go
func Load(s any, env Environment) error
```

This function requires an `Environment` struct, which is used to configure the environment (e.g., specify `dev`, `test`, or `prod`). Here's the definition of the `Environment` struct:

```go
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
```

#### Rules for the `Name` Field

The `Name` field is **required** and must follow these rules:
- **Lowercase letters** and **underscores** are the only allowed characters.
- No **leading** or **trailing underscores**.
- **Numbers** and **special characters** are not permitted.

If the `Name` field violates any of these rules, the loader will return the `LOADER_INVALID_NAME` error. 

#### Fixing the Error

To resolve this error, ensure the `Name` field adheres to the rules. Here's an example of correct usage:

```go
package main

import gf "github.com/fueripe-desu/gofidential"

type Secrets struct {
  // A key to be loaded from the .env file.
  ExampleField string
}

func main() {
  // Define the environment with a valid name.
  env := gf.Environment{
    // Valid name composed of lowercase letters, and optinally, underscores.
    Name: "example",
  }

  s := Secrets{}

  // Load the .env file into the struct.
  err := gf.Load(&s, env)

  if err != nil {
    fmt.Println("Failed to load .env file:", err)
  }
}
```

#### Common Mistakes

Here are some invalid examples of the `Name` field and why they fail:

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

// Contains a trailing underscore.
env := gf.Environment{
  Name: "example_",
}

// Contains a leading underscore.
env := gf.Environment{
  Name: "_example",
}
```

To avoid errors, ensure the `Name` field only includes valid characters. Here’s an example of a valid `Name`:

```go
env := gf.Environment{
  // Valid name for the environment.
  Name: "example_name",
}
```

In this case, the loader will look for a `.env` file named `example_name.env`.

#### Key Details in the Example

1. Correctly defining the `Name` field in the `Environment` struct:
   ```go
   env := gf.Environment{
     // Contains only lowercase letters, so it's a valid name.
     Name: "example",
   }
   ```

2. Passing the `Environment` struct to the `Load()` function:
   ```go
   err := gf.Load(&s, env)
   ```
---
## Troubleshooting Steps

If the error still occurs after troubleshooting, please open an Issue on the [GitHub repository](https://github.com/fueripe-desu/gofidential) with the relevant details.
