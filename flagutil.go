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

// flagutil is a wrapper around the flag module to implement command-line flag
// parsing. Define flags using `Flag()` or `FlagFromStruct()`. Instead of
// providing a different `*Var`, method for each flag type, flagutil provides
// just one simlar method (`Flag()`) that determines the type based on the
// type of variable used to bind the flag, in addition to supporting the use
// of a struct to declare flags. In addition, slices are supported for options
// that should have multiple values. Just pass a reference to a slice to
// `Flag()`, or make the appropriate field a slice in a struct passed to
// `FlagFromStruct`.
//
// This declares an integer flag, `-flagname`, bound to the variable
// `flagvar`.
//
//     var flagvar int flagutil.Flag(&flagvar, "flagname", "The flag number.")
//
// The default value for a flag is specified by assigning a value to the
// variable before calling `flagutil.Parse()`.
//
// `flagutil` allows the use of a struct to specify flags. The flags
// parameters are specified using struct tags strings, using the tag name
// `flagutil`. Elements in the tag should be comma-delimited. The first value
// is the flag name. Remaining parameters must be of the form name='value',
// where supported names are "del" (optional delimiter for multi-valued
// flags), and "usage" (the usage string for the flag). In order for a struct
// field to be used as a flag, the field name must start with an uppercase
// letter (so that the field is exported), and the name parameter must be
// specified. The delimiter and parameters are optional.
//
// The following specifies a slice of IP addresses as strings. IP addresses
// can be repeated as separate command-line arguments, each preceded by the
// flag `ip`, or be specified as a single argument with IP addresses delimited
// by commas, or a combination of the two. `Counts` is a slice of integers
// that must be repeated as separate arguments (each preceded with the flag
// `cnt`), since there is no delimiter specified. The `Quote` field
// demonstrates how to put a single quote in the usage string.
//
//   type MyFlags struct {
//       IPs []string `flagutil:"ip,del=',', usage='The IP address'"`
//       Counts []int `flagutil:"cnt,usage='The count'"`
//       Quote string `flagutil:"quote,usage='Field to show embedded (\\') chars'"`
//   }
//
//   data := new(MyFlags)
//   flagutil.FlagFromStruct(data)
//
// The default set of command-line flags is controlled by top-level functions.
// The FlagSet type allows one to define independent sets of flags, such as to
// implement subcommands in a command-line interface. The methods of FlagSet
// are analogous to the top-level functions for the command-line flag set.
package flagutil

import (
    "flag"
    "fmt"
    "io"
    // log "github.com/cuberat/go-log"
    "os"
    "reflect"
    // "strconv"
    "strings"
    textparser "github.com/cuberat/go-textparser"
    "unicode"
)

type ErrorHandling int

const (
    ContinueOnError ErrorHandling = ErrorHandling(flag.ContinueOnError)
    ExitOnError ErrorHandling = ErrorHandling(flag.ExitOnError)
    PanicOnError ErrorHandling = ErrorHandling(flag.PanicOnError)
)

// Specs for each flag
type flag_spec struct {
    val_ptr interface{}
    arg interface{}
    set_func func()
}

// A FlagSet represents a set of defined flags.
//
// Flag names must be unique within a FlagSet. An attempt to define a flag
// whose name is already in use will cause a panic.
//
// This FlagSet is a wrapper around flag.FlagSet, providing convenience
// methods to make configuring flags simpler.
type FlagSet struct {
    Usage func()

    special_flags []*flag_spec
    name string
    error_handling ErrorHandling
    flag_flagset *flag.FlagSet
    output io.Writer
}

// Returns a new, empty flag set with the specified name and error handling
// property. If the name is not empty, it will be printed in the default usage
// message and in error messages.
func NewFlagSet(name string, error_handling ErrorHandling) *FlagSet {
    flagset := &FlagSet{
        name: name,
        error_handling: error_handling,
        flag_flagset: flag.NewFlagSet(name, flag.ErrorHandling(error_handling)),
        output: os.Stderr,
    }

    flagset.Usage = func() {
        fmt.Fprintf(flagset.Output(), "Usage of %s:\n", os.Args[0])
        flagset.PrintDefaults()
    }

    return flagset
}

// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined command-line flags in the set. See the
// documentation for the global function PrintDefaults for more information.
func (fs *FlagSet) PrintDefaults() {
    fs.flag_flagset.PrintDefaults()
}

// Returns the non-flag arguments (command-line arguments left over after
// parsing the flags).
func (fs *FlagSet) Args() []string {
    return fs.flag_flagset.Args()
}

// Returns the i'th argument. Arg(0) is the first remaining argument after
// flags have been processed. Arg returns an empty string if the requested
// element does not exist.
func (fs *FlagSet) Arg(i int) string {
    return fs.flag_flagset.Arg(i)
}

// Returns the destination for usage and error messages. os.Stderr is returned
// if output was not set or was set to nil.
func (fs *FlagSet) Output() io.Writer {
    return fs.output
}

// Sets the destination for usage and error messages. If output is nil,
// os.Stderr is used.
func (fs *FlagSet) SetOutput(w io.Writer) {
    fs.output = w
    fs.flag_flagset.SetOutput(w)
}

// Parses flag definitions from the argument list, which should not include
// the command name. Must be called after all flags in the FlagSet are defined
// and before flags are accessed by the program. The return value will be
// ErrHelp if -help or -h were set but not defined.
func (fs *FlagSet) Parse(args []string) error {
    err := fs.flag_flagset.Parse(args)
    if err != nil {
        return nil
    }

    for _, f := range fs.special_flags {
        if f.set_func != nil {
            f.set_func()
        }
    }

    return nil
}

type tag_data struct {
    flag_name string
    delimiter string
    usage_string string
}

func parse_tag(tag_str string) *tag_data {
    if tag_str == "" {
        return nil
    }

    fields := make(map[string]string, 2)

    p := textparser.NewScannerString(tag_str)
    field_idx := 0
    field_name := ""
    expecting_name := true
    expecting_value := false

    for p.Scan() {
        if err := p.Err(); err != nil {
            return nil
        }
        token := p.Token()
        token_type := token.Type
        token_text := token.Text

        if field_idx == 0 {
            // The first token should be the field name, by itself, as either
            // an identifier or a quoted string.
            if token_type != textparser.TokenTypeIdent &&
                token_type != textparser.TokenTypeString {
                // FIXME: return an error
                return nil
            }

            fields["name"] = token_text
            field_idx++

            continue
        }

        switch token_type {
        case textparser.TokenTypeSymbol:
            switch token_text {
            case ",":
                field_idx++
                expecting_name = true
            case "=":
                if field_name == "" {
                    // FIXME: return an error
                    return nil
                }
                expecting_value = true
            default:
                // ignore for now
            }
        case textparser.TokenTypeIdent:
            fallthrough
        case textparser.TokenTypeString:
            if expecting_name {
                field_name = token_text
                expecting_name = false
                continue
            }
            if expecting_value {
                if field_name == "" {
                    // FIXME: return an error
                    return nil
                }
                if token_type == textparser.TokenTypeString {
                    token_text = token_text[1:len(token_text) - 1]
                }
                fields[field_name] = token_text
                expecting_value = false
                field_name = ""
                continue
            }
        }
    }

    tag_info := &tag_data{
        flag_name: fields["name"],
        delimiter: fields["del"],
        usage_string: fields["usage"],
    }

    return tag_info
}

// Like `Flag()`, except that `store` must be a pointer to a struct. Exported
// fields (ones starting with capital letters) from the struct that have a tag
// `flagutil` are examined to determine the name of the flag and the usage
// message. See `Flag()` for a list of supported slice types.
//
// Tag format:
//  type FlagData struct {
//      Val []string `flagutil:"val,del=',',usage='Usage string'"`
//      //                       ^     ^         ^
//      //                      name delimiter  usage_string
//  }
//
// The name is required for the field to be used as a flag (this is the flag
// name as appears on the command line). The delimiter ("del" field) is
// optional and is used to split arguments into a slice, where appropriate.
// The usage string ("usage"), if provided, will be used as in the usage
// message.
func (fs *FlagSet) FlagFromStruct(store interface{}) error {
    ptr_value := reflect.ValueOf(store)
    if ptr_value.Kind() != reflect.Ptr {
        return fmt.Errorf("`store` value must be a pointer")
    }

    data := reflect.Indirect(ptr_value)
    if data.Kind() != reflect.Struct {
        return fmt.Errorf("`store` must point to a struct, not a %s",
            data.Kind().String())
    }

    data_type := data.Type()
    num_fields := data_type.NumField()
    for i := 0; i < num_fields; i++ {
        type_field := data_type.Field(i)
        field_name := type_field.Name
        if !first_char_is_upper(field_name) {
            continue
        }

        tag := type_field.Tag
        my_tag_str := tag.Get("flagutil")
        tag_data := parse_tag(my_tag_str)
        if tag_data == nil || tag_data.flag_name == "" {
            continue
        }


        param_name := tag_data.flag_name
        usage_str := tag_data.usage_string
        delimiter := tag_data.delimiter

        data_field := data.Field(i)
        data_field_ptr := data_field.Addr()
        err := fs.FlagSep(data_field_ptr.Interface(), param_name, usage_str,
            delimiter)
        if err != nil {
            return fmt.Errorf("couldn't set up field %q for %s: %s",
                field_name, data_type.Name(), err)
        }

    }

    return nil
}

func first_char_is_upper(s string) bool {
    r := strings.NewReader(s)
    ch, _, err := r.ReadRune()
    if err != nil {
        return false
    }

    return unicode.IsUpper(ch)
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
func (fs *FlagSet) Flag(store interface{}, name, usage string) error {
    return fs.FlagSep(store, name, usage, "")
}

// Like `Flag()`, except specify a delimiter to use to split the argument into
// multiple when `store` is a slice. When a delimiter is specified, the
// arguments may be specified as both multiple command-line arguments and
// single command-line arguments with multiple values separated by the
// specified delimiter.
func (fs *FlagSet) FlagSep(store interface{}, name, usage, del string) error {
    if store == nil {
        return fmt.Errorf("`store` must be a non-nil pointer")
    }

    ptr_value := reflect.ValueOf(store)
    kind := ptr_value.Kind()
    if kind != reflect.Ptr {
        return fmt.Errorf("`store` must be a pointer to a supported type")
    }

    elem := ptr_value.Elem()
    elem_kind := elem.Kind()

    if elem_kind == reflect.Slice {
        return fs.set_slice(ptr_value, name, usage, elem, del)
    }

    switch v := store.(type) {
    case *bool:
        fs.flag_flagset.BoolVar(v, name, *v, usage)
    case *int:
        fs.flag_flagset.IntVar(v, name, *v, usage)
    case *int64:
        fs.flag_flagset.Int64Var(v, name, *v, usage)
    case *uint:
        fs.flag_flagset.UintVar(v, name, *v, usage)
    case *uint64:
        fs.flag_flagset.Uint64Var(v, name, *v, usage)
    case *float64:
        fs.flag_flagset.Float64Var(v, name, *v, usage)
    case *string:
        fs.flag_flagset.StringVar(v, name, *v, usage)
    default:
        return fmt.Errorf("Unsuported type %s (%s) for flag %q",
            kind.String(), v, name)
    }

    return nil
}

// Pass-through to the underlying `flag` object.
//
// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type `Value`,
// which typically holds a user-defined implementation of Value. For instance,
// the caller could create a flag that turns a comma-separated string into a
// slice of strings by giving the slice the methods of `Value`; in particular,
// Set would decompose the comma-separated string into the slice.
func (f *FlagSet) Var(value flag.Value, name string, usage string) {
    f.flag_flagset.Var(value, name, usage)
}

// Pass-through to the underlying `flag` object.
//
// Parsed reports whether `f.Parse` has been called.
func (f *FlagSet) Parsed() bool {
    return f.flag_flagset.Parsed()
}

// Pass-through to the underlying `flag` object. Note that the function passed
// takes a `*flag.Flag`.
//
// Visit visits the flags in lexicographical order, calling fn for each. It
// visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(*flag.Flag)) {
    f.flag_flagset.Visit(fn)
}

// Pass-through to the underlying `flag` object. Note that the function passed
// takes a `*flag.Flag`.
//
// VisitAll visits the flags in lexicographical order, calling fn for each. It
// visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(*flag.Flag)) {
    f.flag_flagset.VisitAll(fn)
}

// Returns the underlying `*flag.FlagSet`.
func (f *FlagSet) GetFlagSet() (*flag.FlagSet) {
    return f.flag_flagset
}

func (fs *FlagSet) set_slice(
    ptr_value reflect.Value,
    name, usage string,
    the_slice reflect.Value,
    sep string,
) error {
    slice_type := the_slice.Type().Elem()
    elem_kind := slice_type.Kind()

    spec := new(flag_spec)
    intfc_ptr_value := ptr_value.Interface()

    switch elem_kind {
    case reflect.Int:
        int_arg := NewMultiArgInt(sep)
        fs.flag_flagset.Var(int_arg, name, usage)
        spec.arg = int_arg
        int_ptr, ok := intfc_ptr_value.(*[]int)
        if !ok {
            return fmt.Errorf("not a *[]int")
        }
        spec.val_ptr = int_ptr
        spec.set_func = func() {
            if vals := int_arg.GetInts(); len(vals) > 0 {
                *int_ptr = vals
            }
        }
        fs.special_flags = append(fs.special_flags, spec)
    case reflect.Int64:
        int64_arg := NewMultiArgInt64(sep)
        fs.flag_flagset.Var(int64_arg, name, usage)
        spec.arg = int64_arg
        int64_ptr, ok := intfc_ptr_value.(*[]int64)
        if !ok {
            return fmt.Errorf("not a *[]int64")
        }
        spec.val_ptr = int64_ptr
        spec.set_func = func() {
            if vals := int64_arg.GetInt64s(); len(vals) > 0 {
                *int64_ptr = vals
            }
        }
        fs.special_flags = append(fs.special_flags, spec)

    case reflect.Uint:
        uint_arg := NewMultiArgUint(sep)
        fs.flag_flagset.Var(uint_arg, name, usage)
        spec.arg = uint_arg
        uint_ptr, ok := intfc_ptr_value.(*[]uint)
        if !ok {
            return fmt.Errorf("not a *[]uint")
        }
        spec.val_ptr = uint_ptr
        spec.set_func = func() {
            if vals := uint_arg.GetUints(); len(vals) > 0 {
                *uint_ptr = vals
            }
        }
        fs.special_flags = append(fs.special_flags, spec)

    case reflect.Uint64:
        uint64_arg := NewMultiArgUint64(sep)
        fs.flag_flagset.Var(uint64_arg, name, usage)
        spec.arg = uint64_arg
        uint64_ptr, ok := intfc_ptr_value.(*[]uint64)
        if !ok {
            return fmt.Errorf("not a *[]uint64")
        }
        spec.val_ptr = uint64_ptr
        spec.set_func = func() {
            if vals := uint64_arg.GetUint64s(); len(vals) > 0 {
                *uint64_ptr = vals
            }
        }
        fs.special_flags = append(fs.special_flags, spec)

    case reflect.Float64:
        float64_arg := NewMultiArgFloat64(sep)
        fs.flag_flagset.Var(float64_arg, name, usage)
        spec.arg = float64_arg
        float64_ptr, ok := intfc_ptr_value.(*[]float64)
        if !ok {
            return fmt.Errorf("not a *[]float64")
        }
        spec.val_ptr = float64_ptr
        spec.set_func = func() {
            if vals := float64_arg.GetFloat64s(); len(vals) > 0 {
                *float64_ptr = vals
            }
        }
        fs.special_flags = append(fs.special_flags, spec)

    case reflect.String:
        string_arg := NewMultiArgString(sep)
        fs.flag_flagset.Var(string_arg, name, usage)
        spec.arg = string_arg
        string_ptr, ok := intfc_ptr_value.(*[]string)
        if !ok {
            return fmt.Errorf("not a *[]string")
        }
        spec.val_ptr = string_ptr
        spec.set_func = func() {
            if vals := string_arg.GetStrings(); len(vals) > 0 {
                *string_ptr = vals
            }
        }
        fs.special_flags = append(fs.special_flags, spec)

    default:
        return fmt.Errorf("Unsuported slice type %s for flag %q",
            elem_kind.String(), name)
    }

    return nil
}
