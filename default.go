// BSD 2-Clause License
//
// Copyright (c) 2020 Don Owens <don@regexguy.com>.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice,
//   this list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package flagutil

import (
    "os"
)

var (
    CommandLine *FlagSet
    Usage func()
)

func init() {
    CommandLine = NewFlagSet(os.Args[0], ExitOnError)
    Usage = CommandLine.Usage
}

// Returns the non-flag arguments (command-line arguments left over after
// parsing the flags).
func Args() []string {
    return CommandLine.Args()
}

// Returns the i'th argument. Arg(0) is the first remaining argument after
// flags have been processed. Arg returns an empty string if the requested
// element does not exist.
func Arg(i int) string {
    return CommandLine.Arg(i)
}

// Defines a flag with the specified name, and usage string. The argument
// `store` points to a variable in which to store the value of the flag. The
// type of the flag is determined from the underlying type of the variable
// pointed to by `store`.
//
//
// If `store` is a pointer to a supported slice type, the same flag will be
// accepted multiple times on the command line, with each value stored in the
// provided slice. Supported slice types:
// - []int
// - []int64
// - []uint
// - []uint64
// - []float64
// - []string
//
// See `FlagSep()` for supporting multiple values specified in a single
// command line argument.
func Flag(store interface{}, name, usage string) error {
    return CommandLine.Flag(store, name, usage)
}

// Like `Flag()`, except specify a delimiter to use to split the argument into
// multiple when `store` is a slice. When a delimiter is specified, the
// arguments may be specified as both multiple command-line arguments and
// single command-line arguments with multiple values separated by the
// specified delimiter.
func FlagSep(store interface{}, name, usage, del string) error {
    return CommandLine.FlagSep(store, name, usage, del)
}

// Like `Flag()`, except that `store` must be a pointer to a struct. Exported
// fields (ones starting with capital letters) from the struct that have a tag
// `flagutil` are examined to determine the name of the flag and the usage
// message.
func FlagFromStruct(store interface{}) error {
    return CommandLine.FlagFromStruct(store)
}

// Parse parses the command-line flags from os.Args[1:]. Must be called after
// all flags are defined and before flags are accessed by the program.
func Parse() error {
    if len(os.Args) > 1 {
        return CommandLine.Parse(os.Args[1:])
    } else {
        return CommandLine.Parse([]string{})
    }
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
    return CommandLine.Parsed()
}
