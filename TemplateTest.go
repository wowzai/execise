package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

func EmailDeal(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace by @ by "at"
	return (substrs[0] + " at " + substrs[1])
}

func RangeAndWith() {
	f1 := Friend{Fname: "dayuan"}
	f2 := Friend{Fname: "xichao"}
	t := template.New("Persons")
	t, _ = t.Parse(
		`Hello, {{.UserName}}!
{{range .Emails}}
an email {{.}}
{{end}}
{{with .Friends}}
	{{range .}}
	my friend name is {{.Fname}}
	{{end}}
{{end}}`)
	p := Person{UserName: "Jack Ding",
		Emails:  []string{"123@123.com", "345@123.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func Pipeline() {
	f1 := Friend{Fname: "dayuan"}
	f2 := Friend{Fname: "xichao"}
	t := template.New("Persons")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDeal})
	t, _ = t.Parse(
		`Hello, {{.UserName}}!
{{range .Emails}}
an email {{.|emailDeal}}
{{end}}
{{with .Friends}}
	{{range .}}
	my friend name is {{.Fname}}
	{{end}}
{{end}}`)
	p := Person{UserName: "Jack Ding",
		Emails:  []string{"123@123.com", "345@123.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func PipelineAndIf() {
	tEmpty := template.New("template")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo:{{if ``}} 不会输出.{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("withvalue")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出.{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("ifelse")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if 部分 {{else}} else 部分. {{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

func Must() {
	tOK := template.New("first")
	template.Must(tOK.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parse OK")

	fmt.Println("The next one ought to fail")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }")) //少了个}
}

func NestTemplate() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println("--before content--")
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println("--before footer--")
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println("--before nil--")
	s1.Execute(os.Stdout, nil)
}

func main() {
	RangeAndWith()
	fmt.Println("----------------")
	PipelineAndIf()
	fmt.Println("----------------")
	Pipeline()
	fmt.Println("----------------")
	t := template.New("fieldname")
	t, _ = t.Parse(`Hello, {{.UserName}}! {{.email}}\n`)
	p := Person{UserName: "Jack Ding"}
	t.Execute(os.Stdout, p)
	fmt.Println("----------------")
	//Must()
	fmt.Println("----------------")
	NestTemplate()
	fmt.Println("----------------")
}
