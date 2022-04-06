# Another Web Assembly Example

A go wasm example that shows the basic bootstrap to run go code as a web assembly module
and also:

- Using functions to pass data from JS to go
- Using functions to pass data from go to JS
- Using functiosn to pass files from JS to go
- Unit testing web assembly (some flags required)

TL;DR: run `make build serve` and open http://localhost:8080/

## Design

Go stores a color in HSL format:
- Key events can (arrows up/down) allow changing the hue
- Javascript is used to
    - Initialize the application
    - Forward key events
    - Periodically display the current color on a rectangle

### Data input

This example shows two ways to get data into go wasm from JS:
- Pass simple variables and structs
- Pass arbirary file content (`[]byte`)

The keyboard & button events are forwarded to go, which uses arrow up and down to change the color hue.
On the javascript side there are handlers for both buttons on screen and keyboard events, these are sent to our
module via an exposed function. In the module we then parse the key events and change the hue for revelant inputs.

The initiallization allows setting the default color from a settings file. Since we don't have file system access,
we load the file on the javascript side, and pass the data in bytes to our module we can then parse the file and initialize
our engine.
This is a simplified example to show file handling, it's more interesting to pass binary files,
structured data would likely be easier to store in json and natively parse in javascritp for efficency.

### Data output

This part of the example shows how we can send structures (rather than just int or strings) back from wasm to js.

We naively fetch the data every 2.5 seconds and update the color of a rectangle. The interesting part is that we can
use `map[string]interface{}` to generate what will be converted to a JSON object once it gets to js.

### Tests

We need the `GOARCH=wasm` `GOOS=js` flags to compile with the `wasm/js` imports to work. So that means
our tests are compiled as wasm and cannot be executed the normal way. The `go_js_wasm_exec` wrapper is used to
execut them using node.js as runtime, so with the correct environment setup the tests run and output as usual.

Useful references:
- https://go.dev/misc/wasm/go_js_wasm_exec
- https://dev.to/nlepage/testing-go-webassembly-code-on-github-actions-j81

## Usage

Requirements:
- Go
- A web browser
- NodeJS (to run the tests)

Most of the commands are wrapped through the make file:
- Build the application `make build`
- Build and run a small dev server to run the application `make serve`
- List all options `make help`

## References and links:

[gptankit's go-wasm](https://github.com/gptankit/go-wasm#go-and-webassembly) repository has useful details
about the interaction between js and go wasm.

For a more complicated application check out [ganother-world](https://github.com/neophob/ganother-world),
it uses a similar setup to load data files, run a game engine sending keys as input and rendering
to a canvas as output.
