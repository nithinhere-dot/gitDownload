# GitHub Repository Downloader

This Go program allows you to download a GitHub repository as a ZIP file and extract its contents to a specified destination folder.

## Features

- Downloads a GitHub repository in ZIP format.
- Extracts the downloaded ZIP file to a specified destination folder.
- Allows users to specify the repository and the destination via command-line flags.

## Prerequisites

- Go (version 1.x) should be installed on your machine.

## Installation

1. Clone or download the repository.

    ```bash
    git clone https://github.com/yourusername/your-repository.git
    cd your-repository
    ```

2. Build the Go program.

    ```bash
    go build -o github-downloader main.go
    ```

## Usage

You can run the program from the command line and provide the following flags:

- `-Repo` (required): The GitHub repository in the format `owner/repository`.
- `-Dest` (optional): The destination folder where the repository will be extracted. By default, it extracts to the current directory.

### Example

```bash
./github-downloader -Repo "nithinhere-dot/E-Cleanse" -Dest "/path/to/destination"
