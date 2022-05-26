package main

import (
	"bufio"
	"fmt"
	"io"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func main() {
	_,err := ReadConfig()
	if err != nil {
		//org err: *os.PathError open /Users/js/.setting.xml: no such file or directory
		fmt.Printf("org err: %T %v\n",errors.Cause(err),errors.Cause(err))
		// wrap err: could not read config: open fail: open /Users/js/.setting.xml: no such file or directory
		fmt.Printf("wrap err: %v\n",err)
		// 详细堆栈信息
		fmt.Printf("stack trace: \n %+v\n",err)
	}
}

func ReadFile(path string)([]byte,error)  {
	f,err := os.Open(path)
	if err != nil {
		// 只对err进行包装，不破坏原错误，携带了附加的错误信息&堆栈信息
		return nil, errors.Wrap(err,"open fail")
	}
	defer f.Close()
	return nil, err
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config,err:=ReadFile(filepath.Join(home,".setting.xml"))
	// 这里也只直接
	return config,errors.WithMessage(err,"could not read config")
}

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines = 0
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, nil
	}
	return lines, nil
}

func CountLines2(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "http/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}
	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s:%s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "\r\n"); err != nil {
		return err
	}
	_, err = io.Copy(w, body)
	return err
}

func WriteResponse2(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWrite{Writer:w}
	fmt.Fprintf(ew, "http/1.1 %d %s\r\n", st.Code, st.Reason)
	for _, h := range headers {
		fmt.Fprintf(ew, "%s:%s\r\n", h.Key, h.Value)
	}
	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}

type errWrite struct {
	io.Writer
	err error
}

func (e *errWrite) Write(buf []byte)(int,error)  {
	if e.err != nil {
		return 0,e.err
	}
	n := 0
	n,e.err = e.Writer.Write(buf)
	return n,nil
}

//输入0,返回1
//输入1,返回1
//输入项，返回值

// GetFeiValueByIndex 获取XXXX第i项的值
// index 表示第几项， 若小于0将返回err和0，否则返回对应项的值
func GetFeiValueByIndex(index int) (int, error) {
	if index < 0 {
		return 0, errors.New("wrong input")
	}

	if index == 0 {
		return 0, nil
	}
	if index < 2 {
		return 1, nil
	}
	l0 := 0
	l1 := 1
	result := 0
	for index > 1 {
		result = l0 + l1
		l0 = l1
		l1 = result
		index--
	}
	return result, nil
}
