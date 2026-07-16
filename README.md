<div id="top"></div>
<div align="center">
  <h1>ForgeKey</h1>
  <h3>
    Cryptographically Secure Password Generator in
    <a href="https://go.dev">Go</a> :closed_lock_with_key:
  </h3>
</div>

<details>
  <summary>
    <strong>Table of Contents</strong>
  </summary>
  <ol>
    <li><a href="#about">About</a></li>
    <li><a href="#how-it-works">How It Works</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#requirements">Requirements</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a>
      <ul>
        <li><a href="#1-passing-the-length-directly-recommended">1. Passing the Length Directly</a></li>
        <li><a href="#2-interactive-mode">2. Interactive Mode</a></li>
      </ul>
    </li>
    <li><a href="#security-architecture">Security Architecture</a></li>
    <li><a href="#sample-output">Sample Output</a></li>
  </ol>
</details>

---

## About

**ForgeKey** is a minimalist, high-performance command-line interface (CLI) tool built in Go. It is designed to generate highly secure, mathematically balanced passwords on demand. By leveraging the operating system's native cryptographic reader, it ensures that every generated sequence is unpredictable and strictly avoids the common probability vulnerabilities found in standard random number generators.

---

## How It Works

- **Input Validation:** Processes CLI arguments or interactive input, verifying that the requested length is a positive integer.
- **Buffer Initialization:** Allocates the memory buffer and character boundaries upfront to optimize execution efficiency.
- **Secure Generation:** Iterates through the requested length, retrieving entropy from `crypto/rand` for each character position.
- **Output:** Maps the entropy values against the character set and returns the final high-entropy string to the standard output.

---

## Features

- **Mathematical Integrity:** Uses `math/big` to ensure every character has an equal probability of selection, avoiding bias.
- **High-Performance:** Optimized memory management ensures rapid, efficient execution even when generating complex, high-entropy passwords.
- **Flexible Input:** Supports both automated command-line arguments and user-friendly interactive prompts.
- **Fail-Safe:** Gracefully handles invalid inputs and operational errors, providing clean error messages instead of system stack traces.

---

## Requirements

- **Go Compiler:** Version 1.18 or higher.
- **Operating System:** Cross-platform support (Windows, macOS, and Linux distributions).

---

## Installation

Clone the repository and compile the highly optimized static binary using standard Go toolchains:

```sh
git clone https://github.com/LeonardoCG12/ForgeKey.git
cd ForgeKey
go build -o forgekey main.go
```

## Usage

ForgeKey offers two seamless execution paradigms to fit into manual generation or continuous delivery pipelines:

### 1. Passing the Length Directly (Recommended)

Execute the utility pass-through workflow by specifying the desired password length as an absolute argument directly in your shell:

```sh
./forgekey 32
```

### 2. Interactive Mode

Run the built application binary blindly without passing tracking arguments to enter an interactive console wizard:

```bash
./forgekey
Enter password length: 64
```

## Security Architecture

ForgeKey avoids the predictability of standard random modules by using `crypto/rand` to draw from a 94-character set:

```go
const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%&*()-_=+-,./|\\{}[]^~;:?<>\"'"
```

By leveraging `math/big` with `crypto/rand`, the tool eliminates Modulo Bias, ensuring each character has an equal mathematical probability of selection, regardless of the password length.

## Sample Output

When executing the generation command successfully, a terminal diagnostic receipt is elegantly logged onto your environment:

```text

----------------------------------------
 Length   : 32
 Password : zT8#kLp9(vM$2xQ!aB7&nR1@yC5*jW4^
----------------------------------------

```

If an invalid length parameter is matched:

```text
2026/07/16 12:49:05 Error generating password: password length must be greater than zero.
```

<p align="right">[<a href="#top">Back to top</a>]</p>
