package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"strings"
	"strconv"
	"regexp"
	"time"
	"crypto/md5"
	"io"
	"os"
	"github.com/wowzai/session"
	_ "github.com/wowzai/session/providers/memory" //会调用init函数
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println(r.Form)
	/*
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	*/
	/*
	for k,v := range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	*/
	fmt.Fprintf(w,"Hello astaxie!")
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curr := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curr,10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t,_ := template.ParseFiles("upload.gtpl.html")
		t.Execute(w,token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file,handler,err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f,err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method) //获取请求的方法
	if r.Method == "GET" {
		curr := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(curr,10))
		token := fmt.Sprintf("%x",h.Sum(nil))

		t,_ := template.ParseFiles("login.gtpl.html")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.FormValue("token1")
		if token != "" {
			//验证token的合法性
			fmt.Println("Right")
		} else {
			//不存在的token报错
			fmt.Println("报错！")
		}
		name := r.Form["username"][0] 
		if len(name) == 0 {
			fmt.Println("输入的用户名为空！")
			fmt.Fprintf(w,"输入的用户名为空！\n")
		}
		/*  也可以用FormValue的形式，用FromValue会自己
			调用Request的ParseForm方法，然后取出第一个参数的值
		if len(r.FormValue("username")) == 0 {
			fmt.Println("输入的用户名为空！FormValue")
			fmt.Fprintf(w,"输入的用户名为空！FormValue")
		}
		*/
		//判断是不是数字
		age,err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("年龄填写有问题！")
		} else {
			fmt.Fprintf(w,"age = " + strconv.Itoa(age) + "\n")
		}
		//判断是不是中文
		if m,_ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$",name); !m {
			fmt.Fprintf(w,name + "不是中文字符串" + "\n")
		} else {
			fmt.Fprintf(w,name + "是中文字符串" + "\n")
		}
		//判断是不是邮箱
		email := r.FormValue("email")
		if m,_ := regexp.MatchString("^[a-zA-Z0-9]+@([a-zA-Z1-9]{1,}).([a-z]{2,4})$",email); !m {
			fmt.Fprintf(w,email + "不是邮件地址" + "\n")
		} else {
			fmt.Fprintf(w,email + "是邮件地址" + "\n")
		}
		phone := r.FormValue("phoneNumber")
		if m,_ := regexp.MatchString("^1[3][5][8][0-9]\\d{4,8}$",phone); !m {
			fmt.Fprintf(w,phone + "不是手机号码" + "\n")
		} else {
			fmt.Fprintf(w,phone + "是手机号码" + "\n")
		}
		fruit := r.FormValue("fruit")
		fmt.Fprintf(w,"Fruit is " + fruit + "\n")
		gender := r.FormValue("gender")
		fmt.Fprintf(w,"Gender is " + gender + "\n")
		interest := r.FormValue("interest")
		fmt.Fprintf(w,"Interest is " + interest + "\n")

		//HTML转义
		fmt.Println("username:",template.HTMLEscapeString(name))
		fmt.Println("email:",template.HTMLEscapeString(email))
		template.HTMLEscape(w,[]byte(name))

		//请求的是登陆的数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}

}

func count(w http.ResponseWriter, r *http.Request) {
	//创建一个全局的session
	globalSessions,err := session.NewManager("memory","wowzaicookie",3600)
	if err != nil {
		fmt.Printf("error:%v\n",err)
	}

	sess := globalSessions.SessionStart(w,r)
	ct := sess.Get("countnum")
	if ct == nil {
		ct = sess.Set("countnum",1)
	} else {
		ct = sess.Set("countnum",(ct.(int)+1))
	}
	t,_ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type","text/html")
	t.Execute(w, sess.Get("countnum"))
}

func main() {
	http.HandleFunc("/", sayHelloName)  //设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}