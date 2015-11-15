// Copyright © 2015 Daniele Tricoli <eriol@mornie.org>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

mvshaker swaps randomly your files.

It was created to remember Warsaw's Second Law: "Never change anything after
3pm on a Friday."
See http://barry.warsaw.us/software/laws.html for more details.

Directories are ignored if you don't use `--recursive` flag (short version `-r`)
and files can be excluded using `--exclude` flag (short version `-e`).

For example:

	# mvshaker /bin/* --exclude bash

--exclude flag has a compact vesion (-e) useful when you want exclude multiple
files:

	# mvshaker /bin/* -e bash -e ls

Since version 0.2:

    # mvshaker -r /bin -e bash -e ls

*/
package main // import "eriol.xyz/mvshaker"
