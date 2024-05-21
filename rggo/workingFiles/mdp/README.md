## Markdown Preview Tool

This tool converts Markdown files to HTML and previews them in the default web browser. 

### Installation

#### Install Go

To install Go on a Debian-based system, follow these steps:

1. **Update your package list**:
    ```sh
    sudo apt update
    ```

2. **Install Go**:
    ```sh
    sudo apt install -y golang
    ```

3. **Verify Go installation**:
    ```sh
    go version
    ```
4. For detailed installation instructions on other systems, refer to the official [Go installation guide](https://go.dev/doc/install)

You should see output indicating the installed version of Go.

### Usage

1. **Clone the repository**:
    ```sh
    git clone <repository_url>
    cd <repository_directory>
    ```

2. **Build the application**:
    ```sh
    go build -o mdpreview main.go
    ```

3. **Run the application**:

    ```sh
    ./mdpreview -file=<path_to_markdown_file>
    ```

    - `-file`: Specifies the path to the Markdown file to preview.
    - `-s`: Optional flag to skip auto-preview.
    - `-t`: Optional flag to specify an alternate template file.

### Example

To preview a Markdown file named `example.md`:

```sh
./mdpreview -file=example.md
```

If you have a custom template named `custom_template.html`, you can use it like this:

```sh
./mdpreview -file example.md -t custom_template.html.tmpl
```

### Notes

- The tool generates a temporary HTML file and opens it in the default web browser.
- The temporary HTML file will be deleted after the preview unless the `-s` flag is used to skip the preview.
