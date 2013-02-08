# Thru

Worlds simplest web server.

  * Servers static files via a URL

## Usage

  * Navigate to the directory of files you want to serve.
  * Type `thru -o`

## Examples

### Output to a specific port

    thru :8080

### Automagically open browser for me

    thru -o

### Bind to a specific I.P. address

    thru 127.0.0.1

### Bind to a specific I.P. address, with specific port, and open browser

    thru 127.0.0.1:3000 -o

## Get thru

Download the latest version of thru by navigating the `build` directory.  Or see if the version you want is in this list:

  * [thru - v1 - Mac OS X](https://github.com/stretchrcom/thru/blob/master/build/Mac%20OS%20X/v1/thru.zip?raw=true)

Download the file somewhere nice (like `/usr/bin`), and ensure it is added to your `PATH` so you can use it in the command line.

### Build your own

If you want to build your own:

  * Get the repo
  * Make sure you have Go installed
  * Do `go build`
  * Put the `thru` output file in some nice bin folder (like `/usr/bin`) and add it to your `PATH`
