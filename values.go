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
    // "flag"
    "fmt"
    "strconv"
    "strings"
)

// Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
// passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
// implement flags as slices.
type MultiArgInt struct {
    Args []int
    Del string
}

// Returns a new object initialized with the specified delimiter. If a
// delimiter is specified, it is used to split individual command line
// arguments into multiple values.
func NewMultiArgInt(delimiter string) (*MultiArgInt) {
    return &MultiArgInt{Del: delimiter}
}

// Returns the resulting []int as an interface{}.
func (ma *MultiArgInt) Get() (interface{}) {
    return ma.Args
}

// Returns the resulting []int.
func (ma *MultiArgInt) GetInts() ([]int) {
    return ma.Args
}

// Returns the resulting []int as a string formatted with "+v".
func (ma *MultiArgInt) String() string {
    return fmt.Sprintf("%+v", ma.Args)
}

// Parses and appends `val` to the underlying []int, after splitting on the
// specified delimiter (if not "").
func (ma *MultiArgInt) Set(val string) error {
    var args []string
    if ma.Del == "" {
        args = []string{val}
    } else {
        args = strings.Split(val, ma.Del)
    }

    for _, str_val := range args {
        int_val, err := strconv.ParseInt(str_val, 10, 64)
        if err != nil {
            return err
        }
        ma.Args = append(ma.Args, int(int_val))
    }

    return nil
}

// Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
// passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
// implement flags as slices.
type MultiArgInt64 struct {
    Args []int64
    Del string
}

// Returns a new object initialized with the specified delimiter. If a
// delimiter is specified, it is used to split individual command line
// arguments into multiple values.
func NewMultiArgInt64(delimiter string) (*MultiArgInt64) {
    return &MultiArgInt64{Del: delimiter}
}

// Returns the resulting []int64 as an interface{}.
func (ma *MultiArgInt64) Get() (interface{}) {
    return ma.Args
}

// Returns the resulting []int64.
func (ma *MultiArgInt64) GetInt64s() ([]int64) {
    return ma.Args
}

// Returns the resulting []int64 as a string formatted with "+v".
func (ma *MultiArgInt64) String() string {
    return fmt.Sprintf("%+v", ma.Args)
}

// Parses and appends `val` to the underlying []int64, after splitting on the
// specified delimiter (if not "").
func (ma *MultiArgInt64) Set(val string) error {
    var args []string
    if ma.Del == "" {
        args = []string{val}
    } else {
        args = strings.Split(val, ma.Del)
    }

    for _, str_val := range args {
        int_val, err := strconv.ParseInt(str_val, 10, 64)
        if err != nil {
            return err
        }
        ma.Args = append(ma.Args, int_val)
    }

    return nil
}

// Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
// passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
// implement flags as slices.
type MultiArgString struct {
    Args []string
    Del string
}

// Returns a new object initialized with the specified delimiter. If a
// delimiter is specified, it is used to split individual command line
// arguments into multiple values.
func NewMultiArgString(delimiter string) (*MultiArgString) {
    return &MultiArgString{Del: delimiter}
}

// Returns the resulting []string as an interface{}.
func (mas *MultiArgString) Get() (interface{}) {
    return mas.Args
}

// Returns the resulting []string.
func (mas *MultiArgString) GetStrings() ([]string) {
    return mas.Args
}

// Returns the resulting []string as a string formatted with "%+v".
func (mas *MultiArgString) String() string {
    return fmt.Sprintf("%+v", mas.Args)
}

// Appends `val` to the underlying []string, after splitting on the specified
// delimiter (if not "").
func (mas *MultiArgString) Set(val string) error {
    var args []string
    if mas.Del == "" {
        args = []string{val}
    } else {
        args = strings.Split(val, mas.Del)
    }

    mas.Args = append(mas.Args, args...)

    return nil
}

// Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
// passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
// implement flags as slices.
type MultiArgFloat64 struct {
    Args []float64
    Del string
}

// Returns a new object initialized with the specified delimiter. If a
// delimiter is specified, it is used to split individual command line
// arguments into multiple values.
func NewMultiArgFloat64(delimiter string) (*MultiArgFloat64) {
    return &MultiArgFloat64{Del: delimiter}
}

// Returns the resulting []float64 as an interface{}.
func (ma *MultiArgFloat64) Get() (interface{}) {
    return ma.Args
}

// Returns the resulting []float64.
func (ma *MultiArgFloat64) GetFloat64s() ([]float64) {
    return ma.Args
}

// Returns the resulting []float64 as a string formatted with "+v".
func (ma *MultiArgFloat64) String() string {
    return fmt.Sprintf("%+v", ma.Args)
}

// Appends `val` to the underlying []float64, after splitting on the specified
// delimiter (if not "").
func (ma *MultiArgFloat64) Set(val string) error {
    var args []string
    if ma.Del == "" {
        args = []string{val}
    } else {
        args = strings.Split(val, ma.Del)
    }

    for _, str_val := range args {
        float_val, err := strconv.ParseFloat(str_val, 64)
        if err != nil {
            return err
        }
        ma.Args = append(ma.Args, float_val)
    }

    return nil
}

// Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
// passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
// implement flags as slices.
type MultiArgUint64 struct {
    Args []uint64
    Del string
}

// Returns a new object initialized with the specified delimiter. If a
// delimiter is specified, it is used to split individual command line
// arguments into multiple values.
func NewMultiArgUint64(delimiter string) (*MultiArgUint64) {
    return &MultiArgUint64{Del: delimiter}
}

// Returns the resulting []uint64 as an interface{}.
func (ma *MultiArgUint64) Get() (interface{}) {
    return ma.Args
}

// Returns the resulting []uint64.
func (ma *MultiArgUint64) GetUint64s() ([]uint64) {
    return ma.Args
}

// Returns the resulting []uint64 as a string formatted with "+v".
func (ma *MultiArgUint64) String() string {
    return fmt.Sprintf("%+v", ma.Args)
}

// Appends `val` to the underlying []uint64, after splitting on the specified
// delimiter (if not "").
func (ma *MultiArgUint64) Set(val string) error {
    var args []string
    if ma.Del == "" {
        args = []string{val}
    } else {
        args = strings.Split(val, ma.Del)
    }

    for _, str_val := range args {
        int_val, err := strconv.ParseUint(str_val, 10, 64)
        if err != nil {
            return err
        }
        ma.Args = append(ma.Args, int_val)
    }

    return nil
}

// Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
// passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
// implement flags as slices.
type MultiArgUint struct {
    Args []uint
    Del string
}

// Returns a new object initialized with the specified delimiter. If a
// delimiter is specified, it is used to split individual command line
// arguments into multiple values.
func NewMultiArgUint(delimiter string) (*MultiArgUint) {
    return &MultiArgUint{Del: delimiter}
}

// Returns the resulting []uint as an interface{}.
func (ma *MultiArgUint) Get() (interface{}) {
    return ma.Args
}

// Returns the resulting []uint.
func (ma *MultiArgUint) GetUints() ([]uint) {
    return ma.Args
}

// Returns the resulting []uint as a string formatted with "+v".
func (ma *MultiArgUint) String() string {
    return fmt.Sprintf("%+v", ma.Args)
}

// Appends `val` to the underlying []uint, after splitting on the specified
// delimiter (if not "").
func (ma *MultiArgUint) Set(val string) error {
    var args []string
    if ma.Del == "" {
        args = []string{val}
    } else {
        args = strings.Split(val, ma.Del)
    }

    for _, str_val := range args {
        int_val, err := strconv.ParseUint(str_val, 10, 64)
        if err != nil {
            return err
        }
        ma.Args = append(ma.Args, uint(int_val))
    }

    return nil
}
