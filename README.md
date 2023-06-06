# Network Port Scanner

This is a simple network port scanner written in Golang. It allows you to scan a range of ports on a given host or IP address and determine if they are open or closed.

## Installation

1. Make sure you have Go installed. You can download it from the official website: [https://go.dev/dl/](https://go.dev/dl/)

2. Clone the repository or download the source code:

   ```shell
   git clone https://github.com/Niteletsplay/GoPortScanner.git
   ```

3. Navigate to the directory where you cloned the repository or extracted the source code:

   ```shell
    cd GoPortScanner
    ```

4. Build the project:

    ```shell
    go build
    ```

## Usage

To use the port scanner, run the following command:

  ```shell
  ./GoPortScanner <hostname> <start_port-end_port>
  ```

  - `<hostname>` :The target host or IP address you want to scan.
  - `<start_port-end_port>`: The range of ports you want to scan, separated by a hyphen.

For example:

  ```shell
  ./GoPortScanner google.com 1-1000
  ```

This will scan ports 1 to 1000 on the example.com host and display the results.