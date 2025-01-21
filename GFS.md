# GoFidential Standard v1.0.0 (GFSv1)

The **GoFidential Standard** is a set of rules created by the **GoFidential** package to enforce stricter but more 
reliable `.env` file practices.

The philosophy underpinning this standard is that stricter rules lead to more predictable behavior, reducing 
flexibility and thereby minimizing the likelihood of bugs.

This document defines all rules, guidelines, and examples necessary for creating `.env` files compliant with 
the **GoFidential Standard**.

---

## Table of Contents
1. [Keys](#keys)
    - [Rules for Keys](#rules-for-keys)
    - [Examples of Valid and Invalid Keys](#examples-of-valid-and-invalid-keys)
2. [Assignment Operator](#assignment-operator)
    - [Rules for Assignment Operator](#rules-for-assignment-operator)
    - [Examples of Correct and Incorrect Usage](#examples-of-correct-and-incorrect-usage)
3. [Values](#values)
    - [Rules for Values](#rules-for-values)
    - [Examples of Valid and Invalid Values](#examples-of-valid-and-invalid-values)
4. [Inline Comments](#inline-comments)
    - [Rules for Comments](#rules-for-comments)
    - [Examples of Correct and Incorrect Comments](#examples-of-correct-and-incorrect-comments)
5. [Multiline Values](#multiline-values)
    - [Rules for Multiline Values](#rules-for-multiline-values)
    - [Examples of Multiline Value Handling](#examples-of-multiline-value-handling)
6. [Why No Substitutions?](#why-no-substitutions)
7. [Examples of Complete Files](#examples-of-complete-files)

---

## Keys

### Rules for Keys
Keys are identifiers used to define environment variables. To ensure consistency and prevent errors, the following 
rules apply:

1. Keys **must**:
    - Use only uppercase letters (`A-Z`) and underscores (`_`).
    - Be unique within the file.
2. Keys **must not**:
    - Be omitted.
    - Begin or end with an underscore (`_`).
    - Consist entirely of underscores.
    - Contain leading or trailing spaces.
    - Be unused (every defined key must serve a purpose in the application).

### Examples of Valid and Invalid Keys

#### Valid Keys
- `API_KEY`
- `DATABASE_URL`
- `ENVIRONMENT`

#### Invalid Keys
- `_KEY` (leading underscore)
- `KEY_` (trailing underscore)
- `DEV2` (numeric characters)
- `key_name` (lowercase letters)
- `#EXAMPLE` (special characters)
- `__` (entirely underscores)
- ` DATABASE_URL` (leading space)
- `="missing key"` (omitted key)

---

## Assignment Operator

### Rules for Assignment Operator
The assignment operator connects a key to its value and must be used consistently to avoid ambiguity.

1. The `=` operator **must**:
    - Appear immediately after the key.
    - Have no spaces between the key, the operator, and the value.
2. The `=` operator **must not**:
    - Be omitted.
    - Be surrounded by spaces.

### Examples of Correct and Incorrect Usage

#### Correct
- `API_KEY="12345"`
- `DATABASE_URL="postgres://user:pass@host:port/db"`

#### Incorrect
- `API_KEY = "12345"` (space around `=`)
- `API_KEY"12345"` (missing `=`)
- `DATABASE_URL postgres://user:pass@host:port/db` (missing `=`)

---

## Values

### Rules for Values
Values represent the data assigned to each key. They must adhere to strict guidelines to ensure proper parsing and 
security.

1. Values **must**:
    - Be enclosed in double quotes (`"`).
    - Use only permitted escape sequences:
        - `\n` for newlines.
        - `\"` for double quotes.
        - `\\` for backslashes.
    - Be non-empty.
2. Values **must not**:
    - Use single quotes (`'`).
    - Contain unescaped double quotes.
    - Be missing.

### Examples of Valid and Invalid Values

#### Valid Values
- `DATABASE_URL="postgres://user:pass@host:port/db"`
- `GREETING="Hello, \"world\"!"`
- `INFO="Name: 'Felipe'\nAge: '18'\n"`

#### Invalid Values
- `MISSING=` (missing value)
- `DATABASE_URL=postgres://user:pass@host:port/db` (missing quotes)
- `GREETING='Hello, world!'` (single quotes not allowed)
- `SPEECH="He said: "I'll be back by tomorrow"."` (unescaped double quotes)
- `MESSAGE="Unterminated string` (unterminated quotes)
- `COST="\tR$10.0\n"` (unallowed `\t` escape character)

---

## Inline Comments

### Rules for Comments
Comments provide context for environment variables but must follow specific rules to maintain clarity.

1. Comments **must**:
    - Begin with a `#`.
    - Be placed on their own line, separate from key-value pairs.
2. Comments **must not**:
    - Appear inline with key-value pairs.

### Examples of Correct and Incorrect Comments

#### Correct
```env
# This is the API key for the service
API_KEY="12345"
```

#### Incorrect
```env
API_KEY="12345" # This is the API key
```

---

## Multiline Values

### Rules for Multiline Values
Multiline values are prohibited to ensure simplicity and portability. Instead, use escape sequences for line breaks.

1. Values **must**:
    - Be written as a single line.
    - Use `\n` for line breaks.
2. Values **must not**:
    - Span multiple lines in the file.

### Examples of Multiline Value Handling

#### Correct
```env
MESSAGE="Line 1\nLine 2\nLine 3"
```

#### Incorrect
```env
MESSAGE="Line 1 \
Line 2 \
Line 3"
```

---

## Why No Substitutions?

Substitutions, where one key references another, are prohibited in the **GoFidential Standard**. This restriction 
simplifies debugging and ensures `.env` files are predictable.

### Rationale
- Substitutions introduce complexity and ambiguity.
- Explicitly defining keys avoids issues caused by unexpected variable expansion.

### Alternative Approach
Instead of:
```env
BASE_URL="https://example.com"
API_ENDPOINT="${BASE_URL}/api"
```
Use:
```env
# Concatenate in application code
BASE_URL="https://example.com"
API_ENDPOINT="/api" 
```

---

## Examples of Complete Files

### Example 1: Valid `.env` File
```env
# Environment configuration for the application
API_KEY="12345"
DATABASE_URL="postgres://user:pass@host:port/db"
ENVIRONMENT="production"
GREETING="Hello, \"world\"!"
```

### Example 2: Invalid `.env` File
```env
# Invalid examples
API_KEY=12345                # Missing quotes
DATABASE_URL='postgres://...' # Single quotes not allowed
 ENVIRONMENT="staging"       # Leading space before key
```

By adhering to the **GoFidential Standard**, your `.env` files will be robust, predictable, and easy to maintain, 
reducing the risk of errors in your application.

