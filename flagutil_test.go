package flagutil_test

import (
    "flag"
    flagutil "github.com/cuberat/go-flagutil"
    "fmt"
    "os"
    "reflect"
    "sort"
    "strings"
    "testing"
)

func TestMultiArgString(t *testing.T) {
    options := flag.NewFlagSet("test", flag.ContinueOnError)

    args := []string{"-ip", "127.0.0.1", "-ip", "127.0.0.2,127.0.0.3"}
    expected := []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"}

    ip_var := flagutil.NewMultiArgString(",")
    options.Var(ip_var, "ip", "IP Address")

    err := options.Parse(args)
    if err != nil {
        t.Errorf("error parsing options: %s", err)
        return
    }

    ips := ip_var.GetStrings()
    if !reflect.DeepEqual(ips, expected) {
        t.Errorf("Slices not equal. Got %s, expected %s", ips, expected)
    }
}

func TestMultiArgInt(t *testing.T) {
    options := flag.NewFlagSet("test", flag.ContinueOnError)

    args := []string{"-num", "1", "-num", "2,3"}
    expected := []int{1, 2, 3}

    num_var := flagutil.NewMultiArgInt(",")
    options.Var(num_var, "num", "Number")

    err := options.Parse(args)
    if err != nil {
        t.Errorf("error parsing options: %s", err)
        return
    }

    nums := num_var.GetInts()
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Slices not equal. Got %+v, expected %+v", nums, expected)
    }
}

func TestMultiArgInt64(t *testing.T) {
    options := flag.NewFlagSet("test", flag.ContinueOnError)

    args := []string{"-num", "1", "-num", "2,3"}
    expected := []int64{1, 2, 3}

    num_var := flagutil.NewMultiArgInt64(",")
    options.Var(num_var, "num", "Number")

    err := options.Parse(args)
    if err != nil {
        t.Errorf("error parsing options: %s", err)
        return
    }

    nums := num_var.GetInt64s()
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Slices not equal. Got %+v, expected %+v", nums, expected)
    }
}

func TestMultiArgFloat64(t *testing.T) {
    options := flag.NewFlagSet("test", flag.ContinueOnError)

    args := []string{"-num", "1.5", "-num", "2.6,3.7"}
    expected := []float64{1.5, 2.6, 3.7}

    num_var := flagutil.NewMultiArgFloat64(",")
    options.Var(num_var, "num", "Number")

    err := options.Parse(args)
    if err != nil {
        t.Errorf("error parsing options: %s", err)
        return
    }

    nums := num_var.GetFloat64s()
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Slices not equal. Got %+v, expected %+v", nums, expected)
    }
}

func TestMultiArgUint64(t *testing.T) {
    options := flag.NewFlagSet("test", flag.ContinueOnError)

    args := []string{"-num", "1", "-num", "2,3"}
    expected := []uint64{1, 2, 3}

    num_var := flagutil.NewMultiArgUint64(",")
    options.Var(num_var, "num", "Number")

    err := options.Parse(args)
    if err != nil {
        t.Errorf("error parsing options: %s", err)
        return
    }

    nums := num_var.GetUint64s()
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Slices not equal. Got %+v, expected %+v", nums, expected)
    }
}

func TestMultiArgUint(t *testing.T) {
    options := flag.NewFlagSet("test", flag.ContinueOnError)

    args := []string{"-num", "1", "-num", "2,3"}
    expected := []uint{1, 2, 3}

    num_var := flagutil.NewMultiArgUint(",")
    options.Var(num_var, "num", "Number")

    err := options.Parse(args)
    if err != nil {
        t.Errorf("error parsing options: %s", err)
        return
    }

    nums := num_var.GetUints()
    if !reflect.DeepEqual(nums, expected) {
        t.Errorf("Slices not equal. Got %+v, expected %+v", nums, expected)
    }
}

type TstTypesData struct {
    Args []string
    Expected interface{}
    Got interface{}
}

func TestFlagSetTypes(t *testing.T) {
    test_data := map[string]*TstTypesData {
        "int": &TstTypesData{
            Args: []string{"-1"}, Expected: int(-1), Got: ptr_to(int(0)),
        },
        "int64": &TstTypesData{
            Args: []string{"-4"}, Expected: int64(-4), Got: ptr_to(int64(0)),
        },
        "uint": &TstTypesData{
            Args: []string{"2"}, Expected: uint(2), Got: ptr_to(uint(0)),
        },
        "uint64": &TstTypesData{
            Args: []string{"3"}, Expected: uint64(3), Got: ptr_to(uint64(0)),
        },
        "float64": &TstTypesData{
            Args: []string{"1.2"},
            Expected: float64(1.2),
            Got: ptr_to(float64(0)),
        },
        "string": &TstTypesData{
            Args: []string{"d3d"}, Expected: "d3d", Got: ptr_to(""),
        },
        "bool": &TstTypesData{
            Args: []string{}, Expected: true, Got: ptr_to(false),
        },
    }

    param_names := make([]string, 0, len(test_data))
    for param, _ := range test_data {
        param_names = append(param_names, param)
    }
    sort.Strings(param_names)

    for _, param := range param_names {
        this_test := test_data[param]
        t.Run(param, func(st *testing.T) {
            flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)

            val := this_test.Got
            err := flags.Flag(val, param, fmt.Sprintf("use of param %s", param))
            if err != nil {
                st.Errorf("error in Flag() call for param %q: %s", param, err)
                return
            }

            param_args := this_test.Args
            these_args := []string{"-" + param}
            these_args = append(these_args, param_args...)
            if err := flags.Parse(these_args); err != nil {
                st.Errorf("error parsing options: %s", err)
                return
            }

            value := reflect.ValueOf(this_test.Got)
            if value.Kind() != reflect.Ptr {
                st.Errorf("Not a pointer! (%+v)", this_test.Got)
                return
            }
            got_param := value.Elem().Interface()

            if !reflect.DeepEqual(got_param, this_test.Expected) {
                st.Errorf("%s vals not equal for param %q. Got %+v, expected %+v",
                    value.Elem().Kind().String(), param, got_param,
                    this_test.Expected)
            }
        })
    }
}

func TestFlagSetSlices(t *testing.T) {
    test_data := map[string]*TstTypesData {
        "int": &TstTypesData{
            Args: []string{"-1", "-2"},
            Expected: []int{-1, -2},
            Got: ptr_to([]int{}),
        },
        "int64": &TstTypesData{
            Args: []string{"-4", "-5"},
            Expected: []int64{-4, -5},
            Got: ptr_to([]int64{}),
        },
        "uint": &TstTypesData{
            Args: []string{"2", "3"},
            Expected: []uint{2, 3},
            Got: ptr_to([]uint{}),
        },
        "uint64": &TstTypesData{
            Args: []string{"3", "4"},
            Expected: []uint64{3, 4},
            Got: ptr_to([]uint64{}),
        },
        "float64": &TstTypesData{
            Args: []string{"1.2", "3.4"},
            Expected: []float64{1.2, 3.4},
            Got: ptr_to([]float64{}),
        },
        "string": &TstTypesData{
            Args: []string{"d3d", "b3f"},
            Expected: []string{"d3d", "b3f"},
            Got: ptr_to([]string{}),
        },
    }

    param_names := make([]string, 0, len(test_data))
    for param, _ := range test_data {
        param_names = append(param_names, param)
    }
    sort.Strings(param_names)

    for _, param := range param_names {
        this_test := test_data[param]
        t.Run(param, func(st *testing.T) {
            flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)

            val := this_test.Got
            err := flags.Flag(val, param, fmt.Sprintf("use of param %s", param))
            if err != nil {
                st.Errorf("error in Flag() call for param %q: %s", param, err)
                return
            }

            these_args := make([]string, 0, 2 * len(this_test.Args))
            for _, arg := range this_test.Args {
                these_args = append(these_args, "-" + param)
                these_args = append(these_args, arg)
            }

            if err := flags.Parse(these_args); err != nil {
                st.Errorf("error parsing options: %s", err)
                return
            }

            value := reflect.ValueOf(this_test.Got)
            if value.Kind() != reflect.Ptr {
                st.Errorf("Not a pointer! (%+v)", this_test.Got)
                return
            }
            got_param := value.Elem().Interface()

            if !reflect.DeepEqual(got_param, this_test.Expected) {
                st.Errorf("%s vals not equal for param %q. Got %+v, expected %+v",
                    value.Elem().Kind().String(), param, got_param,
                    this_test.Expected)
            }
        })
    }
}

type MyFlagStruct struct {
    IP string `flagutil:"ip, usage='The IP address'"`
    Count int `flagutil:"cnt,usage='The count'"`
    Default string `flagutil:"default,usage='Field to show defaults'"`
    Quote string `flagutil:"quote,usage='Field to show embedded (\\') chars'"`
    Spaces string `flagutil:"spacing , usage = 'Spec with spaces'"`
}

type MyFlagStructWithSlices struct {
    IPs []string `flagutil:"ip,del=',', usage='The IP address'"`
    Counts []int `flagutil:"cnt,usage='The count'"`
}

func TestFromStruct(t *testing.T) {
    flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)
    data := new(MyFlagStruct)
    data.Default = "foo"
    err := flags.FlagFromStruct(data)
    if err != nil {
        t.Errorf("error adding flags: %s", err)
    }

    args := []string{"-ip", "127.0.0.1", "-cnt", "5"}
    err = flags.Parse(args)
    if err != nil {
        t.Errorf("error parsing flags: %s", err)
    }

    if data.IP != "127.0.0.1" {
        t.Errorf("IP incorrect. Got %q, expected 127.0.0.1", data.IP)
    }
    if data.Count != 5 {
        t.Errorf("Count incorrect. Got %d, expected 5", data.Count)
    }
    if data.Default != "foo" {
        t.Errorf("Default incorrect. Got %s, expected %q", data.Default,
            "foo")
    }
}

func ExampleFlagSet() {
    ip := ""
    default_ip := "127.0.0.1"
    count := int(0)

    args := []string{"-ip", "127.0.0.2", "-count", "2"}

    flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)
    flags.Flag(&ip, "ip", "IP address (string)")
    flags.Flag(&default_ip, "default_ip", "IP address with default value")
    flags.Flag(&count, "count", "Count (int)")

    flags.Parse(args)

    fmt.Printf("ip: %s\n", ip)
    fmt.Printf("default_ip: %s\n", default_ip)
    fmt.Printf("count: %d\n", count)

    // Output:
    // ip: 127.0.0.2
    // default_ip: 127.0.0.1
    // count: 2
}

func ExampleMultivalue() {
    ips := []string{}
    count := int(0)

    args := []string{"-ip", "127.0.0.2", "-ip", "127.0.0.1", "-count", "5"}

    flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)
    flags.Flag(&ips, "ip", "IP address (string)")
    flags.Flag(&count, "count", "Count (int)")

    flags.Parse(args)

    fmt.Printf("ips: %s\n", ips)
    fmt.Printf("count: %d\n", count)

    // Output:
    // ips: [127.0.0.2 127.0.0.1]
    // count: 5
}

func ExampleFromStruct() {
    flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)
    flags.SetOutput(os.Stdout)
    data := new(MyFlagStruct)
    data.Default = "foo"
    err := flags.FlagFromStruct(data)
    if err != nil {
        fmt.Printf("error adding flags: %s", err)
    }

    args := []string{"-ip", "127.0.0.1", "-cnt", "5"}
    err = flags.Parse(args)
    if err != nil {
        fmt.Printf("error parsing flags: %s", err)
    }

    fmt.Printf("IP: %s\n", data.IP)
    fmt.Printf("Count: %d\n", data.Count)
    fmt.Printf("Default: %s\n", data.Default)

    writer := new(strings.Builder)
    flags.SetOutput(writer)

    flags.PrintDefaults()

    // For `go test` to properly recognize the output match.
    out_str := writer.String()
    out_str = strings.ReplaceAll(out_str, "\t", "    ")
    fmt.Printf("%s", out_str)

    // Output:
    // IP: 127.0.0.1
    // Count: 5
    // Default: foo
    //   -cnt int
    //         The count
    //   -default string
    //         Field to show defaults (default "foo")
    //   -ip string
    //         The IP address
    //   -quote string
    //         Field to show embedded (') chars
    //   -spacing string
    //         Spec with spaces
}

func ExampleFromStructWithSlices() {
    flags := flagutil.NewFlagSet("test", flagutil.ContinueOnError)
    flags.SetOutput(os.Stdout)
    data := new(MyFlagStructWithSlices)
    err := flags.FlagFromStruct(data)
    if err != nil {
        fmt.Printf("error adding flags: %s", err)
    }

    args := []string{"-ip", "127.0.0.1,127.0.0.2", "-cnt", "5",
        "-ip", "127.0.0.3", "-cnt", "6"}
    err = flags.Parse(args)
    if err != nil {
        fmt.Printf("error parsing flags: %s", err)
    }

    writer := new(strings.Builder)
    flags.SetOutput(writer)

    fmt.Fprintf(writer, "IP: %#v\n", data.IPs)
    fmt.Fprintf(writer, "Count: %#v\n", data.Counts)
    fmt.Fprintf(writer, "Usage:\n")
    flags.PrintDefaults()

    // For `go test` to properly recognize the output match.
    out_str := writer.String()
    out_str = strings.ReplaceAll(out_str, "\t", "    ")
    fmt.Printf("%s", out_str)

    // Output:
    // IP: []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"}
    // Count: []int{5, 6}
    // Usage:
    //   -cnt value
    //         The count
    //   -ip value
    //         The IP address
}

func ptr_to(intfc_val interface{}) interface{} {
    value := reflect.ValueOf(intfc_val)
    ptr_value := reflect.New(value.Type())

    reflect.Indirect(ptr_value).Set(value)

    return ptr_value.Interface()
}
