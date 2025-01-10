# Maintaining the Project

This document provides guidelines for maintaining the project. It is meant for core contributors or those taking on significant maintenance roles.

## General Guidelines

- **Decide by cost-benefit**: When making decisions, always weigh the benefits against the costs of implementation and maintenance.
- **Document decisions**: Write down what was decided for future reference, especially when it concerns architectural changes.
- **Constraints are good**: Embrace constraints as they help streamline decisions and prioritize work.
- **Automation is key**: Use automation to solve repetitive tasks and prevent human error.
- **API stability**: Never break the public API but be open to UI changes where necessary.

## Issue Triage

Managing issues efficiently is crucial to the project’s long-term success. Follow these steps:

- **Prioritize based on completion stage**: 
    - Issues with an existing PR are a higher priority.
    - Label issues that are closer to completion (e.g., `has:plan`).
    - Ignore issues that do not yet have a plan or PR.
- **Forecast milestones**: Use a simple forecasting system with the following milestones:
    - Next bugfix release
    - Next feature release

Issues that aren't aligned with upcoming milestones should be considered lower priority.

## Release Policy

Releases should happen frequently but only when stable.

- **Master branch**: It should contain the latest changes but should not be released until stability is ensured.
- **Release branches**: 
    - For maintenance releases, create a `release-x.y` branch.
    - Cherry-pick critical fixes to the release branch and trigger a release using the automation scripts.

### Release Automation

Utilize automation tools like [GitHub Actions](https://github.com/features/actions) to streamline release processes.

## Deprecating and Removing Features

When deprecating features, follow this process:

1. **Soft deprecation** (next release):
   - Document deprecation and annotate code (`@deprecated`).
   - Use a deprecation message in the release notes.
2. **Hard deprecation** (following release):
   - Warn users when using the deprecated feature.
3. **Removal** (next release or later):
   - Features marked for removal should be eliminated in a subsequent release.

Deprecation should be communicated clearly to users, ideally through release notes and documentation.

## Third-party Dependencies

For managing third-party Go dependencies, we rely on Go modules to handle versioning and updates. 

- **Managing Go dependencies**: 
  - Dependencies are tracked using Go modules (via `go.mod` and `go.sum`).
  - Use `go get` to update dependencies and ensure compatibility with your project.
  
- **Vendoring Dependencies**:
  - If you need to vendor dependencies (i.e., include them directly in your repository), you can run the following command:
    ```bash
    go mod vendor
    ```
    This will copy the dependencies into the `vendor/` directory.
  - Ensure to commit the `vendor/` directory to your repository if you're using vendoring.
  
- **Updating Dependencies**:
  - To update dependencies, you can run:
    ```bash
    go get -u
    ```
    This will update all dependencies to their latest versions.
  
- **Handling Forks or Patches**:
  - If you maintain forks of any dependencies (for example, to incorporate patches), you can replace the original dependency with your fork in the `go.mod` file:
    ```go
    replace github.com/original/dependency => github.com/your-fork/dependency v0.0.0-<commit-hash>
    ```
    - After replacing, run `go mod tidy` to clean up the module.

## Refactoring

- Refactoring in Go packages should be approached carefully. Some legacy code should be preserved as-is until someone takes ownership. Avoid significant changes unless absolutely necessary to maintain stability or improve performance.

## CI and Automation

- **GitHub Actions** should be used for continuous integration tasks, ensuring the Go build and tests are automatically run on every pull request.
  - Example action for Go:
    ```yaml
    name: Go CI

    on:
      push:
        branches:
          - main
      pull_request:
        branches:
          - main

    jobs:
      build:
        runs-on: ubuntu-latest
        steps:
          - name: Checkout code
            uses: actions/checkout@v2
          - name: Set up Go
            uses: actions/setup-go@v2
            with:
              go-version: '1.18'  # Use the appropriate Go version
          - name: Install dependencies
            run: go mod tidy
          - name: Run tests
            run: go test ./...
    ```
  
- **Special CI labels**: 
  - `ci:skip-news` to skip unnecessary actions related to updating release notes.
  - `needs:response` to flag PRs that require feedback, ensuring they are closed after a period of inactivity.

## Maintenance Tasks

- **Regular audits**: Review open issues and PRs regularly. Make sure any Go-related issues, such as dependency updates or version conflicts, are addressed. Clean up inactive or outdated issues and PRs to maintain a healthy repository.
  
- **Community contributions**: Encourage contributors to follow the Go coding standards and the PR process outlined in the contributing guide. This ensures consistency and quality across the project.

## Special Guidelines

- **Handling security**: If vulnerabilities or security issues are discovered in dependencies or your code, they must be fixed immediately. A new patch release should be created to address the security concern and tagged accordingly in Go.
  
- **Handling major changes**: Major changes, such as refactors or breaking changes (e.g., Go API changes), should be discussed with the community and reviewed by maintainers before merging. Use semantic versioning to indicate breaking changes and update `go.mod` as necessary.

## Conclusion

Maintaining this Go package requires ongoing attention to stability, usability, and community interaction. These guidelines provide a structure that ensures the project remains healthy and vibrant for all contributors, with a focus on maintaining a smooth development process through Go-specific tools and practices.
