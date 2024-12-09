# peekr

This is a small CLI tool to detect bot mitigation services used by any given website.

## Detectors

- [x] Akamai
- [ ] Akamai Pixel
- [ ] Cloudflare
- [ ] Imperva
- [x] PerimeterX
- [ ] Reblaze
- [ ] Radware
- [x] DataDome
- [x] Kasada

## Prerequisites

- Go 1.18 or later

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sameelarif/peekr.git
   cd peekr
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1. Change the URL in `main.go` to the URL you want to fetch and parse.

2. Run the application:

   ```bash
   go run main.go
   ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
