# GM - A Text Censorship Tool

GM is a high-performance text censorship tool written in Go. It processes text files to filter out banned words based on customizable rules. This tool is ideal for content moderation, text analysis, and ensuring compliance with content guidelines.

## Features
- **High Performance**: Built with Go for speed and efficiency.
- **Customizable Censorship**: Define your own list of banned words in a separate file.
- **Confidence Scoring**: Optionally display a confidence score for the censorship process.
- **Flexible Input**: Supports custom input and banned word files.

## Project Structure
- **main.go**: The main entry point of the application. Handles file input, argument parsing, and program execution.
- **greetings.go**: Contains helper functions for text processing, including the censorship logic and confidence scoring.
- **go.mod**: Manages dependencies for the Go project.
- **source.txt**: Example input file containing text to be processed.
- **ban.txt**: Example file containing a list of banned words.

## Getting Started

### Prerequisites
- Go installed on your system. You can download it from [Go's official website](https://golang.org/).

### Running the Project
1. Clone the repository or download the source code.
2. Navigate to the project directory.
3. Prepare your input files:
   - **source.txt**: Add the text content you want to process.
   - **ban.txt**: Add the list of banned words, one word per line.
4. Run the following command to execute the program:
   ```
   go run main.go [!source.txt] [%ban.txt] [true|false]
   ```
   - Replace `!source.txt` with the path to your input text file.
   - Replace `%ban.txt` with the path to your banned words file.
   - Use `true` to enable confidence scoring or `false` to disable it.

### Example
Given the following files:
- **source.txt**:
  ```
  This is a sample text with some badword and inappropriate content.
  ```
- **ban.txt**:
  ```
  badword
  inappropriate
  ```
Running the command:
```
go run main.go !source.txt %ban.txt true
```
Will output:
```
This is a sample text with some b****** (6.3) and i*********** (7.1) content.
```

## Contributing
Contributions are welcome! Feel free to fork this repository and submit pull requests to improve the tool or add new features.

## License
This project is licensed under the MIT License.