# go-purge

![GO_PURGE_LOGO](assets/logo.png)

![GitHub](https://img.shields.io/github/license/nathan-mittelette/go-purge)
![Build Status](https://img.shields.io/github/actions/workflow/status/nathan-mittelette/go-purge/ci.yml)
![GitHub contributors](https://img.shields.io/github/contributors/nathan-mittelette/go-purge)
![GitHub repo size](https://img.shields.io/github/repo-size/nathan-mittelette/go-purge)
![GitHub issues](https://img.shields.io/github/issues/nathan-mittelette/go-purge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/nathan-mittelette/go-purge)
![Buy Me a Coffee](https://img.shields.io/badge/Buy%20me%20a%20coffee-donate-%23FFDD00)


`go-purge` is a CLI (Command Line Interface) designed to help developers clean their systems or projects by removing caches, unnecessary files, and specific directories that accumulate over time.

---

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
  - [Available Commands](#available-commands)
  - [Usage Examples](#usage-examples)
- [Features](#features)
  - [Global Cleaning](#global-cleaning)
  - [Directory Cleaning](#directory-cleaning)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

If you already have Go installed on your system, you can install `go-purge` very easily by running the following command:

```bash
go install github.com/nathan-mittelette/go-purge@latest
```

This command downloads and installs the latest version of `go-purge` on your machine in your `$GOPATH/bin`.

**Verify the installation:**
```bash
go-purge --help
```

You should then see the CLI help appear in your terminal.

---

## Usage

Here's an overview of how to use `go-purge`:

```bash
go-purge [global options] command [command options]
```

For example:
- To clean directories:
  ```bash
  go-purge directory
  ```

- To perform a global system cleanup:
  ```bash
  go-purge global
  ```

---

## Commands

### Available Commands

`go-purge` provides the following commands:

1. **`global` (alias `g`)**:
    - Cleans your entire environment (Docker, Podman, Maven, Go cache, NPM caches, PNPM, Yarn, etc.).

2. **`directory` (alias `d`)**:
    - Cleans specific folders in the current directory or its subdirectories (e.g., `node_modules`, `.terraform`, etc.).

3. **`help` (alias `h`)**:
    - Displays the list of commands or help for a specific command.

### Usage Examples

1. Global cleanup with confirmation at each step:
   ```bash
   go-purge global
   ```

2. Global cleanup without confirmation (force mode activated):
   ```bash
   go-purge global --force
   ```

3. Cleaning directories in the current folder:
   ```bash
   go-purge directory
   ```

4. Cleaning directories in force mode (without confirmation):
   ```bash
   go-purge directory --force
   ```

---

## Features

### Global Cleaning

The `global` command cleans up various caches and resources from tools and technologies that developers often encounter. Here's what is currently supported:

- **Podman**:
    - Checks if a Podman machine is running.
    - Starts the Podman machine if necessary.
    - Cleans Podman with the `podman system prune --all --volumes` command.

- **Docker**:
    - Cleans Docker using `docker system prune --all --volumes`.

- **Maven**:
    - Deletes the `~/.m2/repository` folder.

- **Gradle**:
    - Deletes the cache directory located at `~/.gradle/caches`.

- **Composer**:
    - Deletes the Composer cache directory located at `~/.composer/cache`.

- **Dart**:
    - Deletes the Dart cache directory located at `~/.dart_tool`.

- **Pip**:
    - Deletes the Python cache directory.

- **Cargo**:
    - Deletes the Cargo cache directory located at `~/.cargo`.

- **Dotnet**:
    - Deletes the Dotnet cache directory.

- **SDKMAN**:
    - Deletes the SDKMAN cache directory.

- **Go**:
    - Cleans the Go cache using the command `go clean -cache -modcache -testcache`.

- **Brew** (macOS only):
    - Removes unnecessary files for Homebrew.

- **NPM / Yarn / PNPM**:
    - Cleans the cache of JavaScript package managers (`npm`, `yarn`, `pnpm`).

You can execute this command with or without confirmation:
- **With confirmation (default)**: prompts you before each action.
- **Force mode (`--force`)**: skips confirmations and performs all cleanups automatically.

---

### Directory Cleaning

The `directory` command focuses on cleaning specific directories in the current folder and its subfolders. The following directories are targeted for removal:

- `.terraform`
- `.dart_tool`
- `node_modules`
- `target` (used in Java or Maven projects)

This command scans your working folder for these directories and prompts you before deleting them, unless the force mode is activated.

Commands executed in the background:
```bash
find . -type d -name ".terraform" -exec rm -rf {} +
find . -type d -name ".dart_tool" -exec rm -rf {} +
find . -type d -name "node_modules" -exec rm -rf {} +
find . -type d -name "target" -exec rm -rf {} +
```

---

## Contributing

Contributions are welcome! If you'd like to add features, fix bugs, or report issues, please don't hesitate to open an issue or a pull request on the repository.

1. Fork this repository.
2. Clone your fork:
   ```bash
   git clone https://github.com/nathan-mittelette/go-purge.git
   ```
3. Make your changes locally.
4. Push your changes and create a pull request.

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more information.
