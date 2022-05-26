package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"os"
)

func main() {
	var a interface{}
	i := 1
	a = i
	k := a.(int)
	println(k)
}

func case1(ch, ch2 chan int) {
	select {
	case ch <- 1:
		println(1)
	case <-ch2:
		println(2)
	}
}
func case2() bool {
	println("case2")
	return false
}
func case3() bool {
	println("case3")
	return false
}

type i interface {
}

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s S) Set(age int) {
	s.Age = age
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func testGoSort(x int) {
	chanList := make([]chan struct{}, 0)
	for i := 0; i <= x; i++ {
		//chanList = append(chanList, make(chan struct{}))
	}
	stopChan := make(chan struct{})
	for i := 0; i <= x; i++ {
		go func(chList []chan struct{}, index int) {
			<-chanList[index]
			fmt.Println(index, "done")
			if index == x {
				stopChan <- struct{}{}
			} else {
				chanList[index+1] <- struct{}{}
			}
		}(chanList, i)
	}
	chanList[0] <- struct{}{}
	select {
	case <-stopChan:
		println("done")
	}
}

func ttttt(x int) {
	chanList := make([]chan struct{}, 0)
	for i := 0; i <= x; i++ {
		chanList = append(chanList, make(chan struct{}))
	}
	stopChan := make(chan struct{})
	for i := 0; i <= x; i++ {
		go func(chList []chan struct{}, index int) {
			if index == 0 {
				fmt.Println(index, "done")
				chanList[index+1] <- struct{}{}
				return
			}
			if index == x {
				<-chanList[index]
				fmt.Println(index, "done")
				stopChan <- struct{}{}
				return
			} else {
				<-chanList[index]
				fmt.Println(index, "done")
				chanList[index+1] <- struct{}{}
			}
		}(chanList, i)
	}
	select {
	case <-stopChan:
		println("done")
	}
}

func checkGoroutineErr(errCtx context.Context) error {
	select {
	case <-errCtx.Done():
		return errCtx.Err()
	default:
		return nil
	}
}

func TestErrGroup() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	for i := 0; i < 3; i++ {
		index := i
		group.Go(func() error {
			fmt.Println("index=", index)
			if index == 0 {
				fmt.Println("index == 0, end!")
			} else if index == 1 {
				fmt.Println("index == 1, start...")
				//time.Sleep(time.Second * 3)
				cancel()
				fmt.Println("inde == 1, has error!")
			} else if index == 2 {
				fmt.Println("index == 2, start...")
				if err := checkGoroutineErr(errCtx); err != nil {
					return err
				}
				fmt.Println("index == 2, has done!")
			}
			return nil
		})
	}

	err := group.Wait()
	if err != nil {
		fmt.Println("Get error: ", err)
	} else {
		fmt.Println("All Done!")
	}
}

type orgList struct {
	List []org `json:"list"`
}

type org struct {
	OrgCode string `json:"orgcode"`
}

func readFile(path string) ([]byte, error) {
	// 打开文件
	jsonFile, err := os.Open(path)

	// 最好要处理以下错误
	if err != nil {
		panic(fmt.Sprintf("打开配置文件错误, err: %v", err))
	}

	// 要记得关闭
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}

type appInfo struct {
	Appid string `json:"orgcode"`
}

type response struct {
	AppInfo []appInfo `json:"list"`
}

type JsonResult struct {
	Resp response `json:"resp"`
}

func testMain() {
	//jsonstr := `{"respCode": "000000","respMsg": "成功","list": [{"orgcode": "d12abd3da59d47e6bf13893ec43730b8"},{"orgcode": "d12abd3da59d47e6bf13893ec43730b8"}]}`

	var JsonRes response
	value, _ := readFile("/Users/js/orglist.test.txt")
	json.Unmarshal(value, &JsonRes)
	fmt.Println("after parse", JsonRes.AppInfo[0].Appid)
	fmt.Println("after parse", JsonRes.AppInfo[1].Appid)
}
