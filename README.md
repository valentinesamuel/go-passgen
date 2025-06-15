# passgen

A simple, flexible command-line password generator written in Go.

Generates strong random passwords using crypto-grade randomness, with customizable length and character sets.

---

## 🚀 Features

- ✅ Cryptographically secure random generation
- ✅ Custom password length
- ✅ Toggle inclusion of:
    - Uppercase letters (A–Z)
    - Lowercase letters (a–z)
    - Digits (0–9)
    - Symbols (e.g., `!@#$%^&*`)
- ✅ Basic input validation
- ✅ Simple help flag

---

## 📦 Installation

**Build from source:**

1. Make sure you have [Go installed](https://go.dev/dl/).
2. Clone this repo:
    ```bash
   git clone https://github.com/valentinesamuel/go-passgen
   cd passgen
    ```


3. Build it:

   ```bash
   go build -o passgen main.go
   ```

---

## 🔑 Usage

Run it from your terminal:

```bash
./passgen [flags]
```

### Available Flags

| Flag              | Description                        | Default |
| ----------------- | ---------------------------------- | ------- |
| `-passwordLength` | Length of the password to generate | `12`    |
| `-upper`          | Include uppercase letters          | `true`  |
| `-lower`          | Include lowercase letters          | `true`  |
| `-digits`         | Include digits                     | `true`  |
| `-symbols`        | Include symbols                    | `false` |
| `-help`           | Show usage instructions            | `false` |

---

### Examples

**Generate a 16-character password with symbols:**

```bash
./passgen -passwordLength 16 -symbols true
```

**Generate a 20-character password with only lowercase letters and digits:**

```bash
./passgen -passwordLength 20 -upper false -symbols false
```

**Show help:**

```bash
./passgen -help
```

---

## ✅ **Tests & Benchmarks**

This project uses Go’s standard testing tools for correctness and performance.

### 📌 **Run all tests**

```bash
go test -v
```

### 📌 **Run table-driven unit tests**

> (Included in `go test`; see `password_test.go`)

```bash
go test -v
```

### 📌 **Run benchmark tests**

```bash
go test -bench=.
```

Example output:

```
goos: darwin
goarch: arm64
pkg: github.com/yourusername/passgen
BenchmarkBuildPasswordPool-8   	1000000	      1234 ns/op
PASS
ok  	github.com/yourusername/passgen	2.345s
```

### 📌 **Check test coverage**

```bash
go test -cover
```

### 📌 **Generate and view coverage report**

```bash
# Generate coverage profile
go test -coverprofile=coverage.out

# Open an HTML report in your browser
go tool cover -html=coverage.out
```

---

## 📌 **Run tests in CI**

To automate tests in GitHub Actions or any CI:

```bash
go test ./...
```

---

## ⚙️ Notes

* The generator uses the `crypto/rand` package for secure random bytes.
* If no character sets are selected (all flags disabled), the program will print an error.
* If you set a password length below 10, it will warn you — adjust the threshold in code if you want stricter or looser limits.

---

## ✨ Author

Built with ❤️ by Me (and ChatGPT as my technical pairing buddy who helped me create this README😄).
 
 