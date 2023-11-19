# CLI DevOps Tools

A Go-based Command Line Interface (CLI) featuring essential tools for DevOps operations. It encompasses functionalities such as Base64 encoding, port checking, timestamp manipulation, secure password generation, URL screenshot capture, and comprehensive WHOIS queries.

## Features

1. **Base64 Encoding and Decoding:**

   - `base64.go`: Handles data encoding and decoding in Base64.

2. **Port Availability Check:**

   - `checkport.go`: Verifies port availability on a specific host.

3. **Advanced Timestamp Manipulation:**

   - `epoch.go`: Converts timestamps into readable formats.

4. **Secure Password Generation:**

   - `password.go`: Generates robust and secure passwords.

5. **Screenshot Capture:**

   - `printscreen.go`: Captures screenshots of URLs.

6. **Comprehensive WHOIS Queries:**

   - `whois.go`: Provides detailed information about domains through WHOIS queries.

## Usage

1. **Compile the CLI:**
   ```
   go build -o cli-go
   ```
2. **Execute the CLI:**
   ```
   ./cli-go [command] [arguments]
   ```

## Contributing

Contributions are welcome! If you find a bug or have an enhancement, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
