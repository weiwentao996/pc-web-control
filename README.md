# Mouse and Keyboard Controller via Web Interface

This Go application creates a simple web interface to control the mouse and keyboard on a local machine. It uses the [Robotgo](https://github.com/go-vgo/robotgo) library to send mouse and keyboard events to the operating system. The web interface is served over HTTP and allows basic interactions like clicking the mouse, moving the cursor, and pressing arrow keys.

## Features

- Mouse clicking (left and right)
- Arrow key pressing (up, down, left, right)
- Cursor movement within the screen boundaries
- Static file serving for the frontend (HTML, CSS, JavaScript)

## Requirements

- Go 1.16+ (for `embed` functionality)
- Robotgo library (install via `go get github.com/go-vgo/robotgo`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/weiwentao996/pc-web-control.git
2. Build the program using Go.
