# GoFidential - A simple secret loader for Golang.
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Q5Q018RJ15)

GoFidential is designed to simplify how your applications handle environment variables and secrets across different development environments. Unlike traditional secret managers, our tool focuses on flexibility and seamless integration, making it a breeze to load ```.env``` files or fetch secrets from the cloud based on your environment configuration.

## Table of Contents  

- [Getting Started](#getting-started)  
- [Features](#features)  
- [Contributing](#contributing)
- [Maintaining](#how-to-maintain)  
- [Code of Conduct](#code-of-conduct)  
- [Issues](#issues)  
- [Show Your Support](#show-your-support)

## Features
- ⚙️ **Environment-Aware:** Automatically adapts to different environments (development, staging, production).
- 🛠 **Highly Customizable:** Offers extensive configuration options to fit your workflow.
- 🧩 **Seamless Integration:** Works effortlessly with your Go projects, solving pain points often encountered with other libraries.
- ☁️ **Cloud-Ready:** Fetch secrets dynamically from cloud providers when needed.
- 🔄 **Lightweight and Reliable**: Designed for performance and simplicity without unnecessary complexity
- ⚡ **Fail Fast:** Provides immediate error feedback during runtime initialization. If a required environment variable is missing, the application will fail early, ensuring you catch issues before they become runtime errors.
- 🔍 **Custom Error Codes:** Easily distinguish errors with custom error codes, making debugging simpler and more organized.
- 📚 **Context-Rich Errors:** Out-of-the-box support for detailed and context-rich error messages, helping you quickly identify the source of issues without unnecessary guesswork.

## Why GoFidential?
Managing environment variables and secrets shouldn't be a hassle. That is precisely why GoFidential was created, it addresses the shortcomings of existing solutions, providing a more intuitive and flexible way to load secrets while maintaining minimal overhead.

---

## Getting Started

### Installation

To install **GoFidential**, ensure you have Go 1.18 or later installed.

#### Step 1: Initialize a Go Module (if needed)

If your project doesn't use Go modules yet, initialize it:

```bash
go mod init <your-module-name>
```

Replace `<your-module-name>` with your project or module name, e.g.,:

```bash
go mod init github.com/yourusername/myproject
```

#### Step 2: Install GoFidential

Add **GoFidential** to your project:

```bash
go get github.com/fueripe-desu/gofidential
```

This command downloads the package and updates your `go.mod` file.

### Basic Usage

1. **Import GoFidential**:

```go
import gf "github.com/fueripe-desu/gofidential"
```

2. **Create a `.env` file**:

Use a file named according to your environment, e.g., `prod.env`:

```
APP_NAME="MyApp"
VERSION="1"
DEBUG="true"
ENVIRONMENT="production"
```

3. **Define a struct**:

The struct must use **PascalCase** field names matching the `.env` keys:

```go
type Secrets struct {
  AppName     string
  Version     int
  Debug       bool
  Environment string
}
```

4. **Set up the environment**:

Create an instance of the `Environment` struct:

```go
// Name makes the application look for the `prod.env` file.
env := gf.Environment{Name: "prod"}
```

5. **Create a struct instance**:

Initialize an empty struct to hold the `.env` data:

```go
s := Secrets{}
```

6. **Load the environment file**:

Populate the struct with `.env` data:

```go
err := gf.Load(&s, env)
if err != nil {
  log.Fatal(err)
}
```

### Complete Example

```go
package main

import (
  "fmt"
  "log"

  gf "github.com/fueripe-desu/gofidential"
)

type Secrets struct {
  AppName     string
  Version     int
  Debug       bool
  Environment string
}

func main() {
  env := gf.Environment{Name: "prod"}
  s := Secrets{}

  err := gf.Load(&s, env)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("App name: %v\n", s.AppName)
  fmt.Printf("Version: %v\n", s.Version)
  fmt.Printf("Debug: %v\n", s.Debug)
  fmt.Printf("Environment: %v\n", s.Environment)
}
```

### Need More Examples?

Check the `examples/` folder in the project repository for additional usage scenarios.

### Having Trouble?

Refer to the [Troubleshooting Guide](./TROUBLESHOOT.md) for detailed solutions to common errors.

### GoFidential Standard v1.0.0 (GFSv1)

Learn about the recommended `.env` file conventions in the [GFS documentation](./GFS.md).

---

## Contributing
We welcome contributions to GoFidential! Whether you're fixing a bug, adding a feature, or improving documentation, your help is greatly appreciated.

## How to Contribute
Please refer to our [CONTRIBUTING.md](./CONTRIBUTING.md) for detailed instructions on how to contribute to the project.

## How to Maintain
Please refer to our [MAINTAIN.md](./MAINTAIN.md) for detailed instructions on how to act as a core project mantainer.

## Issues
If you encounter any bugs or have suggestions for improvement, feel free to [open an issue](https://github.com/fueripe-desu/gofidential/issues) in the repository.

## Code of Conduct
We are committed to providing a welcoming and inclusive environment for all contributors. By participating in this project, you agree to follow our [Code of Conduct](./CODE_OF_CONDUCT.md).  

Please make sure to read it and adhere to the guidelines for respectful and collaborative interaction.

## Show Your Support  

If you find GoFidential helpful, please consider supporting the project in the following ways:  

- ⭐ **Star the repo**: Give the project a star to show your support! [Star the repo here](https://github.com/fueripe-desu/gofidential) 🌟  
- 🍴 **Fork the repo**: Fork the repository to contribute, improve, and make your own customizations.  
- 🐛 **Report issues**: If you find a bug, please report it by opening an [issue](https://github.com/yourusername/gofidential/issues).  
- 💬 **Spread the word**: Tell others about GoFidential! Share it on social media, blogs, or anywhere developers gather.  
- 🤝 **Contribute**: Help improve the project by contributing bug fixes, features, or improvements. Please check out the [Contributing Guide](./CONTRIBUTING.md).
- ☕ **Support me on Ko-Fi:** If you love the project and want to show your support, consider buying me a coffee! Every coffee fuels my creativity to keep improving the project and building awesome things. Support me on [Ko-Fi](https://ko-fi.com/fueripedesu) here 💙  

Your support helps keep the project alive and improves its reach. Thanks for being awesome! 🙌  
