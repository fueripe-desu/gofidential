# Contributing to GoFidential

Thank you for considering contributing to **GoFidential**! 🎉 Your contributions help improve this project and make it more useful for developers worldwide. Whether it's a bug fix, a new feature, or just a small improvement, we appreciate your time and effort!

Please follow these steps to contribute.

## Getting Started

### 1. **Fork the Repository**

   To get started, **fork** the repository to your own GitHub account. Forking allows you to create your own copy of the project where you can make changes without affecting the original project.

   To fork the repository:
   - Navigate to the [GoFidential repository](https://github.com/fueripe-desu/gofidential).
   - Click the **Fork** button (top-right of the page).
   - Select your GitHub account to fork the project into.

   Once the repository is forked, you will have your own copy of **GoFidential** to make changes to.

### 2. **Clone the Forked Repository**

   After forking, you need to clone the repository to your local machine. This allows you to work on the code locally.

   ```bash
   git clone https://github.com/yourusername/gofidential.git
   cd gofidential
   ```

   Replace `yourusername` with your actual GitHub username.

### 3. **Create a New Branch**

   We use **clean commit prefixes** to keep things organized. Each branch should correspond to a single task, such as adding a feature or fixing a bug. This makes it easier to track changes and maintain code quality.

   **Branch Naming Rules**:

   - **Prefix the branch name** with a clear, standardized type of task using **hyphens** to separate words.
   - Make sure your branch name **represents one task** (e.g., one feature, one bug fix). Don't try to implement more than one feature in the same branch.

   **Valid Prefixes**:

   - `feat/`: Use this for new features.
   - `fix/`: Use this for bug fixes.
   - `docs/`: Use this for documentation updates.
   - `chore/`: Use this for chores like dependencies updates or build-related tasks.
   - `test/`: Use this for adding or improving tests.
   - `refactor/`: Use this for code refactoring (no functional changes).

   **Examples of valid branch names**:

   - `feat/add-missing-env-check` — New feature: adding a check for missing environment variables.
   - `fix/resolve-env-variable-issue` — Bug fix: fixing an issue where environment variables were not being loaded correctly.
   - `docs/update-readme-with-contributing-guide` — Documentation update: adding a contributing guide to the README.
   - `chore/update-dependencies` — Chore: updating the dependencies in `go.mod`.
   - `test/improve-env-variable-test-coverage` — Tests: improving test coverage for environment variable handling.
   - `refactor/simplify-env-loading-code` — Refactor: simplifying the code for loading environment variables.

   The branch name should be **short, descriptive, and specific** to the task.

   **General guidelines**:
   - Separate words with **hyphens** (`-`), not underscores or camelCase.
   - Be concise but clear about the branch's purpose.
   - Avoid generic names like `fix1`, `feature1`, or `bugfix`.

   For more information on clean commit naming conventions, check out the [Clean Commit Conventions Guide](https://www.conventionalcommits.org/).

### 4. **Make Your Changes**

   Now you're ready to make changes to the code! Whether you’re fixing a bug, adding a new feature, or improving documentation, make sure to:

   - Follow the **coding guidelines** and project conventions.
   - Include relevant tests (if applicable) to verify that your changes work.
   - Update the **documentation** to reflect the new changes you have just made.

   Take your time to ensure the changes are correct and well-structured.

### 5. **Test Your Changes**

   Before committing your changes, ensure that all tests pass locally. This is important because your changes may unintentionally break code written by other maintainers. Consider your changes ready for commit only once all tests (including those you've added for your new functionality) pass.

To run the tests locally, use the following command:

   ```bash
   go test ./...
   ```

Once all tests pass, you are good to go ❤️

### 6. **Commit Your Changes**

Once you're happy with the changes, commit them with a clear and concise message that explains what you did and why.

**Important**: Your commit messages **must follow the clean commit message format**. This means:
- The commit message should **start with a prefix** indicating the type of change (e.g., `feat:` for new features, `fix:` for bug fixes).
- If your changes are specific to a **sub-package** of the main package, include the sub-package name in parentheses after the prefix (e.g., `feat(parser):`).
- Use **present tense** (e.g., "fix bug with environment variables").
- Be concise but specific about what was changed.

#### Common Commit Prefixes
Here are some common prefixes to follow when naming your commits:

- **`feat/`**: New feature implementation.
  - Example: `feat(auth): add login functionality`
- **`fix/`**: Bug fixes.
  - Example: `fix(parser): handle missing environment variables`
- **`refactor/`**: Code refactoring, improving the structure or readability of the code without changing its functionality.
  - Example: `refactor(api): restructure user routes`
- **`chore/`**: Routine tasks or changes that don’t directly affect the app’s functionality.
  - Example: `chore(build): update dependencies`
- **`docs/`**: Changes related to documentation.
  - Example: `docs: update README with usage examples`
- **`style/`**: Code style changes that don’t affect the functionality (e.g., formatting, white-space adjustments).
  - Example: `style: fix indentation in the auth module`
- **`test/`**: Changes related to adding or modifying tests.
  - Example: `test(auth): add unit tests for login feature`

If your changes fall under one of these categories, use the corresponding prefix. If you are unsure, consult the [Clean Commit](https://www.conventionalcommits.org/en/v1.0.0/) website for more information and guidelines on commit conventions.

**Example**:
```bash
git add .
git commit -m "feat(parser): add support for single line .env comments"
```

### 7. **Push to Your Fork**

   After committing, push your changes back to your forked repository on GitHub. You can do this by running the following command:

   ```bash
   git push origin feature/your-feature-name
   ```

### 8. **Create a Pull Request**

Now that all of your changes are committed and pushed to your fork, it's time to create a **Pull Request** (PR) to propose your changes to the original repository. Please make sure **all** the changes you want to include are committed and pushed before opening the PR. Opening a PR with incomplete or uncommitted changes can cause confusion and slow down the review process.

#### Steps to open a Pull Request:
1. **Ensure All Changes Are Committed and Pushed**: Before creating the PR, confirm that you've committed all your changes and pushed them to your fork. Do not open a PR unless you are ready for it to be reviewed. A PR should contain **only** the changes you want to propose.
   
2. Go to the **Pull Requests** tab on the [GoFidential GitHub page](https://github.com/fueripe-desu/gofidential).
3. Click the **New Pull Request** button.
4. Select the branch you just pushed from your fork and compare it with the `main` branch of the original repository.
   
#### Writing a Good PR Description:

A well-written description will help the maintainers understand your changes better and speed up the review process. Here's what to include:

- **Provide a Summary of the Changes**: Start with a brief but clear summary of what your PR does. Be specific about the feature added, bug fixed, or task completed.

- **Explain Why You Made These Changes**: Mention the problem you were trying to solve, the feature you wanted to add, or the issue you addressed with your changes. This gives context to your PR.

- **Describe the Solution**: Explain how you solved the problem. If applicable, link any issues this PR addresses (e.g., `Fixes #123`).

- **List Major Changes**: Highlight important changes made in the codebase. For example, if you added a new function or refactored a major part of the code, mention it here.

- **Test Coverage**: If you wrote new tests or modified existing ones, mention this in your PR description. 

- **Known Issues**: If there are any known issues or limitations with your PR, make sure to mention them clearly. If there are things you plan to improve later, explain that too.

- **Link to the Documentation**: If the PR requires or has associated documentation, link to it in your PR description.

Here’s an example of a well-structured PR description:

```markdown
## What does this PR do?
This PR adds the JWT authentication system to the project, replacing the previous cookie-based session management. This enhances security and scalability.

## Why are these changes necessary?
The current authentication system is not scalable for large applications and has some potential security issues. JWT tokens allow for stateless authentication and better scalability.

## Changes made:
- Implemented JWT authentication with login and logout functionality.
- Updated authentication middleware to validate tokens.
- Refactored login page to integrate the JWT-based authentication flow.

## Test coverage:
- Added unit tests for `login` and `logout` functions.
- Ran all existing tests; they all passed.

## Known issues:
- The token expiration time is hardcoded and needs to be configured in the next iteration.
```
There's no need to manually copy this example. When you open a new Pull Request, an automatically generated template will be provided, which you can simply fill in with the relevant details.

### 9. **Review and Feedback**

   Once you submit your PR, the project maintainers will review your changes. They might provide feedback, ask for modifications, or approve your changes.

   - If feedback is given, make the necessary changes and push the updated code to your branch.
   - If everything is good, your PR will be merged into the `main` branch!

---

## Code Style Guidelines

- **Use descriptive and meaningful names** for variables, functions, and classes to enhance readability and maintainability.
- Adhere to **Go coding standards**:
  - **Indentation**: Use 2 spaces for indentation, **do not use tabs**.
  - **Comments**: Always add comments to explain complex logic and non-obvious code behavior.
  - **Function Documentation**: Document any new functions you introduce, providing context on their purpose, parameters, and return values.
  - **Error Handling**: Utilize **Go’s built-in error handling** pattern, and leverage custom error payloads where applicable to provide more context and clarity.
  - **Write Tests**: Ensure that new functionality is covered by automated tests, and that existing tests are maintained or improved with any changes.
- **Keep your code well-tested**. Tests should cover all edge cases, and you should aim for high code coverage to prevent regressions.

---

## Reporting Issues

If you encounter any bugs or issues, please follow these steps to report them:

1. **Search for existing issues**: Before creating a new issue, please check if the issue has already been reported. You can search for existing issues using the search bar in the [Issues](https://github.com/fueripe-desu/gofidential/issues) tab.
   
2. **Create a new issue**: If you don't find an existing issue, create a new one by following these guidelines and providing as much detail as possible:
   - **A clear description of the issue**: What exactly is happening? Be as descriptive as possible.
   - **Steps to reproduce the issue**: Provide a step-by-step guide on how to replicate the issue. Include any specific setup or configurations.
   - **Expected behavior**: What did you expect the program to do?
   - **Actual behavior**: What happened instead? Include any error messages, logs, or unexpected behavior observed.
   - **Environment details**: Information about your environment (e.g., operating system, Go version, relevant libraries, etc.).
   - **Additional context**: Any other details that might be relevant, such as related configuration settings, recent changes, or other observations.
   
By providing these details, you help us understand the issue and resolve it faster!

---

## Code of Conduct

By contributing to **GoFidential**, you agree to follow our [Code of Conduct](./CODE_OF_CONDUCT.md). We expect all contributors to engage respectfully, kindly, and professionally with each other.

## Thank You!

Thank you again for contributing to **GoFidential**! Your time and effort mean a lot, and help make the project better for everyone. 🙌

If you have any questions or need help, feel free to reach out by opening an issue or sending a message. We’re happy to assist!
