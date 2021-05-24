package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

type User struct {
	Browsers []string `json:"browsers"`
	Company  string   `json:"-"`
	Email    string   `json:"email"`
	Job      string   `json:"-"`
	Name     string   `json:"name"`
	Phone    string   `json:"-"`
}

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers := make(map[string]bool)
	var foundUsers strings.Builder

	sc := bufio.NewScanner(file)
	for i := 0; sc.Scan(); i++ {
		user := User{}
		err := user.UnmarshalJSON([]byte(sc.Text()))
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				_, notSeenBefore := seenBrowsers[browser]
				if !notSeenBefore {
					seenBrowsers[browser] = true
				}
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				_, notSeenBefore := seenBrowsers[browser]
				if !notSeenBefore {
					seenBrowsers[browser] = true
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		foundUsers.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email))
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCcaf38dDecodeGolangCourseraHw3BenchQwer(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCcaf38dEncodeGolangCourseraHw3BenchQwer(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCcaf38dEncodeGolangCourseraHw3BenchQwer(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCcaf38dEncodeGolangCourseraHw3BenchQwer(w, v)
}

func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCcaf38dDecodeGolangCourseraHw3BenchQwer(&r, v)
	return r.Error()
}

func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCcaf38dDecodeGolangCourseraHw3BenchQwer(l, v)
}