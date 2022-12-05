# mvshaker #

`mvshaker` swaps randomly your files.

[![asciicast](https://asciinema.org/a/9gf89grw31j8z8jvymoyfqvhl.png)](https://asciinema.org/a/9gf89grw31j8z8jvymoyfqvhl)

It was created to remember Warsaw's Second Law: "Never change anything after
3pm on a Friday."
See http://barry.warsaw.us/software/laws.html for more details.

Directories are ignored if you don't use `--recursive` flag (short version `-r`)
and files can be excluded using `--exclude` flag (short version `-e`).

## Installation ##

To build `mvshaker` and install it to `$GOPATH/bin/mvshaker`you need a working
Go compiler:

    % go install noa.mornie.org/eriol/mvshaker@latest

## Examples ##

    # mvshaker /bin/* --exclude bash

`--exclude` flag has a compact vesion (`-e`) useful when you want exclude
multiple files:

    # mvshaker /bin/* -e bash -e ls

Since version 0.2:

    # mvshaker -r /bin -e bash -e ls
