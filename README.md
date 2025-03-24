# HTPass CLI

**HTPass** is a user password handler for `.htpasswd` files that uses `bcrypt` under the hood, providing useful functionality to manage htpasswd files effectively.

## Features

- **Create and Manage .htpasswd Files**: Easily create and manage `.htpasswd` files with a user-friendly command-line interface.
- **Bcrypt Encryption**: Securely handle passwords using `bcrypt`, a robust and secure hashing algorithm.
- **Comprehensive User Management**: Add, update, delete, and list users with ease.
- **Password Verification**: Match passwords securely against stored hashes.

## Installation

To install the `htpass` CLI, build from source using:

```bash
go install github.com/tilliboy/htpass@latest
```
