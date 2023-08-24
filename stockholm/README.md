# Stockholm

This `README.md` file contains informations about the `Stockholm` project.

It is a ransomware / cryptolocker virus written in Go which encrypts files in a folder called `infection` placed in the user's home directory.

It's written for Linux specifically, but it should also work on MacOS / Windows.

## Compiling

To compile Stockholm, you must have Go installed (go1.18.1 was used to work on this project).

Then, simply run the `go build` command, or `make` (`make` will just run `go build`).

## Usage

Stockholm has three options

```
usage: stockholm [-h|--help] [-v|--version] [-r|--reverse "<value>"]
                 [-s|--silent]

                 A very friendly program (not really)

Arguments:

  -h  --help     Print help information
  -v  --version  Shows the version of the program
  -r  --reverse  Reverse the infection using the provided encryption key
  -s  --silent   Silent mode, no output
```

### Infection

To infect your computer, simply run the program without any arguments. The used encryption key will be displayed on the screen at the end.

### Reverse infection

To cure all your precious files, run `stockholm -r <key>`. This will search for `.ft` files in the `infection` directory and decrypt them using the provided hex key.

## How it works

Since this is a proof of concept, `stockholm` uses AES-256 symmetric encryption to encrypt files quickly.
