# trie-autocompletion

Word autocompletion algorithm that uses tries (prefix trees) for storing dictionaries

## Algorithm

The specification of algorithm is presented in the [ALGORITHM.md](ALGORITHM.md) file.
It also contains design principles, clarifications and description of possible soft spots and improvements.

## React webapp demo

There is a demo of autocompletion system that uses the algorithm.

### Running the react demo

1. Make sure you have installed [Node.js](https://nodejs.org) (>=19.8.1) and [Yarn](https://yarnpkg.com/).

2. Go to the `react-demo` directory:
   ```bash
    $ cd react-demo
   ```
3. Install required dependencies:

   ```bash
    $ yarn install
   ```

4. Run the app
   ```bash
   yarn start
   ```

Demo app should open automatically in your browser. If not, type [localhost:3000](http://localhost:3000) in the URL bar and hit the enter button.

### Building the react demo

1. Make sure you have installed [Node.js](https://nodejs.org) (>=19.8.1) and [Yarn](https://yarnpkg.com/).

2. Go to the `react-demo` directory:
   ```bash
    $ cd react-demo
   ```
3. Install required dependencies:
   ```bash
    $ yarn install
   ```
4. Build the project:
   ```bash
    $ yarn build
   ```

## CLI implementation in Go

In the `go-implementation` directory, there is a simple example application that showcases a possible CLI program autocompletion feature.

### Running the CLI implementation

1. Make sure that you have installed [Go](https://go.dev) (>=1.20).

2. Go to the `go-implementation` directory:
   ```bash
   $ cd go-implementation
   ```
3. Run the application:
   ```bash
   $ go run .
   ```

All dependencies should be installed automatically.
