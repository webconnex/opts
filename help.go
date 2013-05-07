// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opts

import (
	"fmt"
	"bytes"
	"text/tabwriter"
)

var printHelp *bool

// AddHelp adds the -h and --help options
func AddHelp(desc string) {
	printHelp = Flag("-h", "--help", desc)
}

// Help prints a generated help screen
func Help() {
	buf := &bytes.Buffer{}
	w := tabwriter.NewWriter(buf, 0, 2, 2, ' ', 0)
	for _, opt := range optionList {
		forms := opt.Forms()
		if len(forms) < 1 {
			continue
		}
		arg := opt.Arg()
		name := opt.ArgName()
		w.Write([]byte{'\t'})
		for i, f := range forms {
			if i > 0 {
				w.Write([]byte{',', ' '})
			} 
			w.Write([]byte(f))
			if arg > NOARG {
				if len(f) > 2 {
					w.Write([]byte{'='})
				} else {
					w.Write([]byte{' '})
				}
				if arg > OPTARG {
					w.Write([]byte(name))
				} else {
					w.Write([]byte{'['})
					w.Write([]byte(name))
					w.Write([]byte{']'})
				}
			}
		}
		w.Write([]byte{'\t'})
		w.Write([]byte(opt.Description()))
		w.Write([]byte{'\n'})
	}
	w.Flush()
	fmt.Printf("Usage: %s %s\n%s\n", Xname, Usage, Description)
	println(string(buf.Bytes()))
}