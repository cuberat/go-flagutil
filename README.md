

# flagutil
`import "github.com/cuberat/go-flagutil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
flagutil is a wrapper around the flag module to implement command-line flag
parsing. Define flags using `Flag()` or `FlagFromStruct()`. Instead of
providing a different `*Var`, method for each flag type, flagutil provides
just one simlar method (`Flag()`) that determines the type based on the
type of variable used to bind the flag, in addition to supporting the use
of a struct to declare flags. In addition, slices are supported for options
that should have multiple values. Just pass a reference to a slice to
`Flag()`, or make the appropriate field a slice in a struct passed to
`FlagFromStruct`.

This declares an integer flag, `-flagname`, bound to the variable
`flagvar`.


	var flagvar int flagutil.Flag(&flagvar, "flagname", "The flag number.")

The default value for a flag is specified by assigning a value to the
variable before calling `flagutil.Parse()`.

`flagutil` allows the use of a struct to specify flags. The flags
parameters are specified using struct tags strings, using the tag name
`flagutil`. Elements in the tag should be comma-delimited. The first value
is the flag name. Remaining parameters must be of the form name='value',
where supported names are "del" (optional delimiter for multi-valued
flags), and "usage" (the usage string for the flag). In order for a struct
field to be used as a flag, the field name must start with an uppercase
letter (so that the field is exported), and the name parameter must be
specified. The delimiter and parameters are optional.

The following specifies a slice of IP addresses as strings. IP addresses
can be repeated as separate command-line arguments, each preceded by the
flag `ip`, or be specified as a single argument with IP addresses delimited
by commas, or a combination of the two. `Counts` is a slice of integers
that must be repeated as separate arguments (each preceded with the flag
`cnt`), since there is no delimiter specified. The `Quote` field
demonstrates how to put a single quote in the usage string.


	type MyFlags struct {
	    IPs []string `flagutil:"ip,del=',', usage='The IP address'"`
	    Counts []int `flagutil:"cnt,usage='The count'"`
	    Quote string `flagutil:"quote,usage='Field to show embedded (\\') chars'"`
	}
	
	data := new(MyFlags)
	flagutil.FlagFromStruct(data)

The default set of command-line flags is controlled by top-level functions.
The FlagSet type allows one to define independent sets of flags, such as to
implement subcommands in a command-line interface. The methods of FlagSet
are analogous to the top-level functions for the command-line flag set.




## <a name="pkg-index">Index</a>
* [func Arg(i int) string](#Arg)
* [func Args() []string](#Args)
* [func Flag(store interface{}, name, usage string) error](#Flag)
* [func FlagFromStruct(store interface{}) error](#FlagFromStruct)
* [func FlagSep(store interface{}, name, usage, del string) error](#FlagSep)
* [func Parse() error](#Parse)
* [func Parsed() bool](#Parsed)
* [type ErrorHandling](#ErrorHandling)
* [type FlagSet](#FlagSet)
  * [func NewFlagSet(name string, error_handling ErrorHandling) *FlagSet](#NewFlagSet)
  * [func (fs *FlagSet) Arg(i int) string](#FlagSet.Arg)
  * [func (fs *FlagSet) Args() []string](#FlagSet.Args)
  * [func (fs *FlagSet) Flag(store interface{}, name, usage string) error](#FlagSet.Flag)
  * [func (fs *FlagSet) FlagFromStruct(store interface{}) error](#FlagSet.FlagFromStruct)
  * [func (fs *FlagSet) FlagSep(store interface{}, name, usage, del string) error](#FlagSet.FlagSep)
  * [func (f *FlagSet) GetFlagSet() *flag.FlagSet](#FlagSet.GetFlagSet)
  * [func (fs *FlagSet) Output() io.Writer](#FlagSet.Output)
  * [func (fs *FlagSet) Parse(args []string) error](#FlagSet.Parse)
  * [func (f *FlagSet) Parsed() bool](#FlagSet.Parsed)
  * [func (fs *FlagSet) PrintDefaults()](#FlagSet.PrintDefaults)
  * [func (fs *FlagSet) SetOutput(w io.Writer)](#FlagSet.SetOutput)
  * [func (f *FlagSet) Var(value flag.Value, name string, usage string)](#FlagSet.Var)
  * [func (f *FlagSet) Visit(fn func(*flag.Flag))](#FlagSet.Visit)
  * [func (f *FlagSet) VisitAll(fn func(*flag.Flag))](#FlagSet.VisitAll)
* [type MultiArgFloat64](#MultiArgFloat64)
  * [func NewMultiArgFloat64(delimiter string) *MultiArgFloat64](#NewMultiArgFloat64)
  * [func (ma *MultiArgFloat64) Get() interface{}](#MultiArgFloat64.Get)
  * [func (ma *MultiArgFloat64) GetFloat64s() []float64](#MultiArgFloat64.GetFloat64s)
  * [func (ma *MultiArgFloat64) Set(val string) error](#MultiArgFloat64.Set)
  * [func (ma *MultiArgFloat64) String() string](#MultiArgFloat64.String)
* [type MultiArgInt](#MultiArgInt)
  * [func NewMultiArgInt(delimiter string) *MultiArgInt](#NewMultiArgInt)
  * [func (ma *MultiArgInt) Get() interface{}](#MultiArgInt.Get)
  * [func (ma *MultiArgInt) GetInts() []int](#MultiArgInt.GetInts)
  * [func (ma *MultiArgInt) Set(val string) error](#MultiArgInt.Set)
  * [func (ma *MultiArgInt) String() string](#MultiArgInt.String)
* [type MultiArgInt64](#MultiArgInt64)
  * [func NewMultiArgInt64(delimiter string) *MultiArgInt64](#NewMultiArgInt64)
  * [func (ma *MultiArgInt64) Get() interface{}](#MultiArgInt64.Get)
  * [func (ma *MultiArgInt64) GetInt64s() []int64](#MultiArgInt64.GetInt64s)
  * [func (ma *MultiArgInt64) Set(val string) error](#MultiArgInt64.Set)
  * [func (ma *MultiArgInt64) String() string](#MultiArgInt64.String)
* [type MultiArgString](#MultiArgString)
  * [func NewMultiArgString(delimiter string) *MultiArgString](#NewMultiArgString)
  * [func (mas *MultiArgString) Get() interface{}](#MultiArgString.Get)
  * [func (mas *MultiArgString) GetStrings() []string](#MultiArgString.GetStrings)
  * [func (mas *MultiArgString) Set(val string) error](#MultiArgString.Set)
  * [func (mas *MultiArgString) String() string](#MultiArgString.String)
* [type MultiArgUint](#MultiArgUint)
  * [func NewMultiArgUint(delimiter string) *MultiArgUint](#NewMultiArgUint)
  * [func (ma *MultiArgUint) Get() interface{}](#MultiArgUint.Get)
  * [func (ma *MultiArgUint) GetUints() []uint](#MultiArgUint.GetUints)
  * [func (ma *MultiArgUint) Set(val string) error](#MultiArgUint.Set)
  * [func (ma *MultiArgUint) String() string](#MultiArgUint.String)
* [type MultiArgUint64](#MultiArgUint64)
  * [func NewMultiArgUint64(delimiter string) *MultiArgUint64](#NewMultiArgUint64)
  * [func (ma *MultiArgUint64) Get() interface{}](#MultiArgUint64.Get)
  * [func (ma *MultiArgUint64) GetUint64s() []uint64](#MultiArgUint64.GetUint64s)
  * [func (ma *MultiArgUint64) Set(val string) error](#MultiArgUint64.Set)
  * [func (ma *MultiArgUint64) String() string](#MultiArgUint64.String)

#### <a name="pkg-examples">Examples</a>
* [FlagSet](#example_FlagSet)

#### <a name="pkg-files">Package files</a>
[default.go](/src/github.com/cuberat/go-flagutil/default.go) [flagutil.go](/src/github.com/cuberat/go-flagutil/flagutil.go) [values.go](/src/github.com/cuberat/go-flagutil/values.go) 





## <a name="Arg">func</a> [Arg](/src/target/default.go?s=1926:1948#L42)
``` go
func Arg(i int) string
```
Returns the i'th argument. Arg(0) is the first remaining argument after
flags have been processed. Arg returns an empty string if the requested
element does not exist.



## <a name="Args">func</a> [Args](/src/target/default.go?s=1693:1713#L35)
``` go
func Args() []string
```
Returns the non-flag arguments (command-line arguments left over after
parsing the flags).



## <a name="Flag">func</a> [Flag](/src/target/default.go?s=2618:2672#L64)
``` go
func Flag(store interface{}, name, usage string) error
```
Defines a flag with the specified name, and usage string. The argument
`store` points to a variable in which to store the value of the flag. The
type of the flag is determined from the underlying type of the variable
pointed to by `store`.

If `store` is a pointer to a supported slice type, the same flag will be
accepted multiple times on the command line, with each value stored in the
provided slice. Supported slice types:
- []int
- []int64
- []uint
- []uint64
- []float64
- []string

See `FlagSep()` for supporting multiple values specified in a single
command line argument.



## <a name="FlagFromStruct">func</a> [FlagFromStruct](/src/target/default.go?s=3414:3458#L81)
``` go
func FlagFromStruct(store interface{}) error
```
Like `Flag()`, except that `store` must be a pointer to a struct. Exported
fields (ones starting with capital letters) from the struct that have a tag
`flagutil` are examined to determine the name of the flag and the usage
message.



## <a name="FlagSep">func</a> [FlagSep](/src/target/default.go?s=3046:3108#L73)
``` go
func FlagSep(store interface{}, name, usage, del string) error
```
Like `Flag()`, except specify a delimiter to use to split the argument into
multiple when `store` is a slice. When a delimiter is specified, the
arguments may be specified as both multiple command-line arguments and
single command-line arguments with multiple values separated by the
specified delimiter.



## <a name="Parse">func</a> [Parse](/src/target/default.go?s=3658:3676#L87)
``` go
func Parse() error
```
Parse parses the command-line flags from os.Args[1:]. Must be called after
all flags are defined and before flags are accessed by the program.



## <a name="Parsed">func</a> [Parsed](/src/target/default.go?s=3885:3903#L96)
``` go
func Parsed() bool
```
Parsed reports whether the command-line flags have been parsed.




## <a name="ErrorHandling">type</a> [ErrorHandling](/src/target/flagutil.go?s=4339:4361#L81)
``` go
type ErrorHandling int
```

``` go
const (
    ContinueOnError ErrorHandling = ErrorHandling(flag.ContinueOnError)
    ExitOnError     ErrorHandling = ErrorHandling(flag.ExitOnError)
    PanicOnError    ErrorHandling = ErrorHandling(flag.PanicOnError)
)
```









## <a name="FlagSet">type</a> [FlagSet](/src/target/flagutil.go?s=4989:5162#L103)
``` go
type FlagSet struct {
    Usage func()
    // contains filtered or unexported fields
}
```
A FlagSet represents a set of defined flags.

Flag names must be unique within a FlagSet. An attempt to define a flag
whose name is already in use will cause a panic.

This FlagSet is a wrapper around flag.FlagSet, providing convenience
methods to make configuring flags simpler.


``` go
var (
    CommandLine *FlagSet
    Usage       func()
)
```






### <a name="NewFlagSet">func</a> [NewFlagSet](/src/target/flagutil.go?s=5353:5420#L116)
``` go
func NewFlagSet(name string, error_handling ErrorHandling) *FlagSet
```
Returns a new, empty flag set with the specified name and error handling
property. If the name is not empty, it will be printed in the default usage
message and in error messages.





### <a name="FlagSet.Arg">func</a> (\*FlagSet) [Arg](/src/target/flagutil.go?s=6430:6466#L148)
``` go
func (fs *FlagSet) Arg(i int) string
```
Returns the i'th argument. Arg(0) is the first remaining argument after
flags have been processed. Arg returns an empty string if the requested
element does not exist.




### <a name="FlagSet.Args">func</a> (\*FlagSet) [Args](/src/target/flagutil.go?s=6179:6213#L141)
``` go
func (fs *FlagSet) Args() []string
```
Returns the non-flag arguments (command-line arguments left over after
parsing the flags).




### <a name="FlagSet.Flag">func</a> (\*FlagSet) [Flag](/src/target/flagutil.go?s=12808:12876#L366)
``` go
func (fs *FlagSet) Flag(store interface{}, name, usage string) error
```
Defines a flag with the specified name, and usage string. The argument
`store` points to a variable in which to store the value of the flag. The
type of the flag is determined from the underlying type of the variable
pointed to by `store`.

If `store` is a pointer to a supported slice type, the same flag will be
accepted multiple times on the command line, with each value stored in the
provided slice. Supported slice types:
- []int
- []int64
- []uint
- []uint64
- []float64
- []string

See `FlagSep()` for supporting multiple values specified in a single
command line argument.




### <a name="FlagSet.FlagFromStruct">func</a> (\*FlagSet) [FlagFromStruct](/src/target/flagutil.go?s=10637:10695#L291)
``` go
func (fs *FlagSet) FlagFromStruct(store interface{}) error
```
Like `Flag()`, except that `store` must be a pointer to a struct. Exported
fields (ones starting with capital letters) from the struct that have a tag
`flagutil` are examined to determine the name of the flag and the usage
message. See `Flag()` for a list of supported slice types.

Tag format:


	type FlagData struct {
	    Val []string `flagutil:"val,del=',',usage='Usage string'"`
	    //                       ^     ^         ^
	    //                      name delimiter  usage_string
	}

The name is required for the field to be used as a flag (this is the flag
name as appears on the command line). The delimiter ("del" field) is
optional and is used to split arguments into a slice, where appropriate.
The usage string ("usage"), if provided, will be used as in the usage
message.




### <a name="FlagSet.FlagSep">func</a> (\*FlagSet) [FlagSep](/src/target/flagutil.go?s=13248:13324#L375)
``` go
func (fs *FlagSet) FlagSep(store interface{}, name, usage, del string) error
```
Like `Flag()`, except specify a delimiter to use to split the argument into
multiple when `store` is a slice. When a delimiter is specified, the
arguments may be specified as both multiple command-line arguments and
single command-line arguments with multiple values separated by the
specified delimiter.




### <a name="FlagSet.GetFlagSet">func</a> (\*FlagSet) [GetFlagSet](/src/target/flagutil.go?s=15887:15933#L454)
``` go
func (f *FlagSet) GetFlagSet() *flag.FlagSet
```
Returns the underlying `*flag.FlagSet`.




### <a name="FlagSet.Output">func</a> (\*FlagSet) [Output](/src/target/flagutil.go?s=6629:6666#L154)
``` go
func (fs *FlagSet) Output() io.Writer
```
Returns the destination for usage and error messages. os.Stderr is returned
if output was not set or was set to nil.




### <a name="FlagSet.Parse">func</a> (\*FlagSet) [Parse](/src/target/flagutil.go?s=7166:7211#L169)
``` go
func (fs *FlagSet) Parse(args []string) error
```
Parses flag definitions from the argument list, which should not include
the command name. Must be called after all flags in the FlagSet are defined
and before flags are accessed by the program. The return value will be
ErrHelp if -help or -h were set but not defined.




### <a name="FlagSet.Parsed">func</a> (\*FlagSet) [Parsed](/src/target/flagutil.go?s=15151:15182#L431)
``` go
func (f *FlagSet) Parsed() bool
```
Pass-through to the underlying `flag` object.

Parsed reports whether `f.Parse` has been called.




### <a name="FlagSet.PrintDefaults">func</a> (\*FlagSet) [PrintDefaults](/src/target/flagutil.go?s=6006:6040#L135)
``` go
func (fs *FlagSet) PrintDefaults()
```
PrintDefaults prints, to standard error unless configured otherwise, the
default values of all defined command-line flags in the set. See the
documentation for the global function PrintDefaults for more information.




### <a name="FlagSet.SetOutput">func</a> (\*FlagSet) [SetOutput](/src/target/flagutil.go?s=6787:6828#L160)
``` go
func (fs *FlagSet) SetOutput(w io.Writer)
```
Sets the destination for usage and error messages. If output is nil,
os.Stderr is used.




### <a name="FlagSet.Var">func</a> (\*FlagSet) [Var](/src/target/flagutil.go?s=14931:14997#L424)
``` go
func (f *FlagSet) Var(value flag.Value, name string, usage string)
```
Pass-through to the underlying `flag` object.

Var defines a flag with the specified name and usage string. The type and
value of the flag are represented by the first argument, of type `Value`,
which typically holds a user-defined implementation of Value. For instance,
the caller could create a flag that turns a comma-separated string into a
slice of strings by giving the slice the methods of `Value`; in particular,
Set would decompose the comma-separated string into the slice.




### <a name="FlagSet.Visit">func</a> (\*FlagSet) [Visit](/src/target/flagutil.go?s=15453:15497#L440)
``` go
func (f *FlagSet) Visit(fn func(*flag.Flag))
```
Pass-through to the underlying `flag` object. Note that the function passed
takes a `*flag.Flag`.

Visit visits the flags in lexicographical order, calling fn for each. It
visits only those flags that have been set.




### <a name="FlagSet.VisitAll">func</a> (\*FlagSet) [VisitAll](/src/target/flagutil.go?s=15759:15806#L449)
``` go
func (f *FlagSet) VisitAll(fn func(*flag.Flag))
```
Pass-through to the underlying `flag` object. Note that the function passed
takes a `*flag.Flag`.

VisitAll visits the flags in lexicographical order, calling fn for each. It
visits all flags, even those not set.




## <a name="MultiArgFloat64">type</a> [MultiArgFloat64](/src/target/values.go?s=5745:5810#L176)
``` go
type MultiArgFloat64 struct {
    Args []float64
    Del  string
}
```
Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
implement flags as slices.







### <a name="NewMultiArgFloat64">func</a> [NewMultiArgFloat64](/src/target/values.go?s=5989:6049#L184)
``` go
func NewMultiArgFloat64(delimiter string) *MultiArgFloat64
```
Returns a new object initialized with the specified delimiter. If a
delimiter is specified, it is used to split individual command line
arguments into multiple values.





### <a name="MultiArgFloat64.Get">func</a> (\*MultiArgFloat64) [Get](/src/target/values.go?s=6153:6199#L189)
``` go
func (ma *MultiArgFloat64) Get() interface{}
```
Returns the resulting []float64 as an interface{}.




### <a name="MultiArgFloat64.GetFloat64s">func</a> (\*MultiArgFloat64) [GetFloat64s](/src/target/values.go?s=6260:6312#L194)
``` go
func (ma *MultiArgFloat64) GetFloat64s() []float64
```
Returns the resulting []float64.




### <a name="MultiArgFloat64.Set">func</a> (\*MultiArgFloat64) [Set](/src/target/values.go?s=6597:6645#L205)
``` go
func (ma *MultiArgFloat64) Set(val string) error
```
Appends `val` to the underlying []float64, after splitting on the specified
delimiter (if not "").




### <a name="MultiArgFloat64.String">func</a> (\*MultiArgFloat64) [String](/src/target/values.go?s=6405:6447#L199)
``` go
func (ma *MultiArgFloat64) String() string
```
Returns the resulting []float64 as a string formatted with "+v".




## <a name="MultiArgInt">type</a> [MultiArgInt](/src/target/values.go?s=1666:1723#L29)
``` go
type MultiArgInt struct {
    Args []int
    Del  string
}
```
Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
implement flags as slices.







### <a name="NewMultiArgInt">func</a> [NewMultiArgInt](/src/target/values.go?s=1902:1954#L37)
``` go
func NewMultiArgInt(delimiter string) *MultiArgInt
```
Returns a new object initialized with the specified delimiter. If a
delimiter is specified, it is used to split individual command line
arguments into multiple values.





### <a name="MultiArgInt.Get">func</a> (\*MultiArgInt) [Get](/src/target/values.go?s=2050:2092#L42)
``` go
func (ma *MultiArgInt) Get() interface{}
```
Returns the resulting []int as an interface{}.




### <a name="MultiArgInt.GetInts">func</a> (\*MultiArgInt) [GetInts](/src/target/values.go?s=2149:2189#L47)
``` go
func (ma *MultiArgInt) GetInts() []int
```
Returns the resulting []int.




### <a name="MultiArgInt.Set">func</a> (\*MultiArgInt) [Set](/src/target/values.go?s=2473:2517#L58)
``` go
func (ma *MultiArgInt) Set(val string) error
```
Parses and appends `val` to the underlying []int, after splitting on the
specified delimiter (if not "").




### <a name="MultiArgInt.String">func</a> (\*MultiArgInt) [String](/src/target/values.go?s=2278:2316#L52)
``` go
func (ma *MultiArgInt) String() string
```
Returns the resulting []int as a string formatted with "+v".




## <a name="MultiArgInt64">type</a> [MultiArgInt64](/src/target/values.go?s=3057:3118#L80)
``` go
type MultiArgInt64 struct {
    Args []int64
    Del  string
}
```
Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
implement flags as slices.







### <a name="NewMultiArgInt64">func</a> [NewMultiArgInt64](/src/target/values.go?s=3297:3353#L88)
``` go
func NewMultiArgInt64(delimiter string) *MultiArgInt64
```
Returns a new object initialized with the specified delimiter. If a
delimiter is specified, it is used to split individual command line
arguments into multiple values.





### <a name="MultiArgInt64.Get">func</a> (\*MultiArgInt64) [Get](/src/target/values.go?s=3453:3497#L93)
``` go
func (ma *MultiArgInt64) Get() interface{}
```
Returns the resulting []int64 as an interface{}.




### <a name="MultiArgInt64.GetInt64s">func</a> (\*MultiArgInt64) [GetInt64s](/src/target/values.go?s=3556:3602#L98)
``` go
func (ma *MultiArgInt64) GetInt64s() []int64
```
Returns the resulting []int64.




### <a name="MultiArgInt64.Set">func</a> (\*MultiArgInt64) [Set](/src/target/values.go?s=3892:3938#L109)
``` go
func (ma *MultiArgInt64) Set(val string) error
```
Parses and appends `val` to the underlying []int64, after splitting on the
specified delimiter (if not "").




### <a name="MultiArgInt64.String">func</a> (\*MultiArgInt64) [String](/src/target/values.go?s=3693:3733#L103)
``` go
func (ma *MultiArgInt64) String() string
```
Returns the resulting []int64 as a string formatted with "+v".




## <a name="MultiArgString">type</a> [MultiArgString](/src/target/values.go?s=4473:4536#L131)
``` go
type MultiArgString struct {
    Args []string
    Del  string
}
```
Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
implement flags as slices.







### <a name="NewMultiArgString">func</a> [NewMultiArgString](/src/target/values.go?s=4715:4773#L139)
``` go
func NewMultiArgString(delimiter string) *MultiArgString
```
Returns a new object initialized with the specified delimiter. If a
delimiter is specified, it is used to split individual command line
arguments into multiple values.





### <a name="MultiArgString.Get">func</a> (\*MultiArgString) [Get](/src/target/values.go?s=4875:4921#L144)
``` go
func (mas *MultiArgString) Get() interface{}
```
Returns the resulting []string as an interface{}.




### <a name="MultiArgString.GetStrings">func</a> (\*MultiArgString) [GetStrings](/src/target/values.go?s=4982:5032#L149)
``` go
func (mas *MultiArgString) GetStrings() []string
```
Returns the resulting []string.




### <a name="MultiArgString.Set">func</a> (\*MultiArgString) [Set](/src/target/values.go?s=5318:5366#L160)
``` go
func (mas *MultiArgString) Set(val string) error
```
Appends `val` to the underlying []string, after splitting on the specified
delimiter (if not "").




### <a name="MultiArgString.String">func</a> (\*MultiArgString) [String](/src/target/values.go?s=5126:5168#L154)
``` go
func (mas *MultiArgString) String() string
```
Returns the resulting []string as a string formatted with "%+v".




## <a name="MultiArgUint">type</a> [MultiArgUint](/src/target/values.go?s=8603:8662#L278)
``` go
type MultiArgUint struct {
    Args []uint
    Del  string
}
```
Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
implement flags as slices.







### <a name="NewMultiArgUint">func</a> [NewMultiArgUint](/src/target/values.go?s=8841:8895#L286)
``` go
func NewMultiArgUint(delimiter string) *MultiArgUint
```
Returns a new object initialized with the specified delimiter. If a
delimiter is specified, it is used to split individual command line
arguments into multiple values.





### <a name="MultiArgUint.Get">func</a> (\*MultiArgUint) [Get](/src/target/values.go?s=8993:9036#L291)
``` go
func (ma *MultiArgUint) Get() interface{}
```
Returns the resulting []uint as an interface{}.




### <a name="MultiArgUint.GetUints">func</a> (\*MultiArgUint) [GetUints](/src/target/values.go?s=9094:9137#L296)
``` go
func (ma *MultiArgUint) GetUints() []uint
```
Returns the resulting []uint.




### <a name="MultiArgUint.Set">func</a> (\*MultiArgUint) [Set](/src/target/values.go?s=9413:9458#L307)
``` go
func (ma *MultiArgUint) Set(val string) error
```
Appends `val` to the underlying []uint, after splitting on the specified
delimiter (if not "").




### <a name="MultiArgUint.String">func</a> (\*MultiArgUint) [String](/src/target/values.go?s=9227:9266#L301)
``` go
func (ma *MultiArgUint) String() string
```
Returns the resulting []uint as a string formatted with "+v".




## <a name="MultiArgUint64">type</a> [MultiArgUint64](/src/target/values.go?s=7182:7245#L227)
``` go
type MultiArgUint64 struct {
    Args []uint64
    Del  string
}
```
Implements the `flag.Value` and `flag.Getter` interfaces. Useful for
passing to `flag.Var()` or `flagutil.Var()`. Used by `flagutil.Flag()` to
implement flags as slices.







### <a name="NewMultiArgUint64">func</a> [NewMultiArgUint64](/src/target/values.go?s=7424:7482#L235)
``` go
func NewMultiArgUint64(delimiter string) *MultiArgUint64
```
Returns a new object initialized with the specified delimiter. If a
delimiter is specified, it is used to split individual command line
arguments into multiple values.





### <a name="MultiArgUint64.Get">func</a> (\*MultiArgUint64) [Get](/src/target/values.go?s=7584:7629#L240)
``` go
func (ma *MultiArgUint64) Get() interface{}
```
Returns the resulting []uint64 as an interface{}.




### <a name="MultiArgUint64.GetUint64s">func</a> (\*MultiArgUint64) [GetUint64s](/src/target/values.go?s=7689:7738#L245)
``` go
func (ma *MultiArgUint64) GetUint64s() []uint64
```
Returns the resulting []uint64.




### <a name="MultiArgUint64.Set">func</a> (\*MultiArgUint64) [Set](/src/target/values.go?s=8020:8067#L256)
``` go
func (ma *MultiArgUint64) Set(val string) error
```
Appends `val` to the underlying []uint64, after splitting on the specified
delimiter (if not "").




### <a name="MultiArgUint64.String">func</a> (\*MultiArgUint64) [String](/src/target/values.go?s=7830:7871#L250)
``` go
func (ma *MultiArgUint64) String() string
```
Returns the resulting []uint64 as a string formatted with "+v".








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
