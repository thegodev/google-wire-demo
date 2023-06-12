# Google Wire Demo

This is a demo repository for using [Google Wire](https://github.com/google/wire) to generate dependency injection code in Go.

## Getting Started

To use this demo, you will need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

Once you have Go installed, you can clone this repository by running the following command:

```
$ git clone https://github.com/thegodev/google-wire-demo.git
```

Then, navigate to the cloned repository:

```
$ cd google-wire-demo
```

## Usage

To generate the DI code, run the following command:

```
$ go generate ./...
```

This will generate the `wire_gen.go` file, which contains the generated DI code.

To build and run the program, run the following commands:

```
$ go build
$ ./google-wire-demo
```

You should see the following output:

```
hi there!
```