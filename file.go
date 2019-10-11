package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//文件的操作
func main() {

	//读取文件 通过ioutil
	body, err := readFileByIoutil("./data/file.txt")

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(body))

	//逐渐行读取
	body, _ = readFileByPearLine("./data/file.txt")
	fmt.Println(string(body))

	//根据buffer读取
	body, _ = readFileByBuff("./data/file.txt", 200)
	fmt.Println(string(body))

	//判断文件是否存在
	ret := FileIsExist("./data/file.txt")

	fmt.Println(ret)

	ret = FileIsExist("./file1.txt")

	fmt.Println(ret)

	//新建文件
	ret = NewFile("./data/debug.txt")

	fmt.Println(ret)

	//写文件
	ret, _ = WriteFile("./data/debug.txt", "测试\n")
	fmt.Println(ret)
	//删除文件
	ret = RemoveFile("./data/debug.txt")

	fmt.Println(ret)

	ret = RemoveFile("./data/debug1.txt")
	fmt.Println()

}

/**
 * 通过os.Open打开文件 再通过 ioutil读取
 * 读取文件内容 通过ioutil读取
 * 优点 不需要估算文件大小 自动申请大小
 * 缺点 性能不是特别好 容易造成内存浪费
 */
func readFileByIoutil(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ioutil.ReadAll(f)
}

/**
 * 逐行读取文件问内容
 * 优点 内存省
 * 缺点 性能差
 */
func readFileByPearLine(path string) ([]byte, error) {

	var body []byte

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	//关闭文件
	defer f.Close()

	handler := bufio.NewReader(f)
	for {
		//换行地方结束
		line, err := handler.ReadBytes('\n')

		body = ByteCombine(body, line)
		if err != nil {
			//文件末行
			if err == io.EOF {
				return body, nil
			}

			return nil, err
		}
	}
}

// []byte 整合到一起
func ByteCombine(b ...[]byte) []byte {
	return bytes.Join(b, []byte(""))
}

/**
 * 根据buffer 分块读取文件内容
 * 优点 不会内存泄漏 性能 > 逐行
 * 缺点 需要自己衡量buffer大小 容易翻车
 */
func readFileByBuff(path string, bufferSize int) ([]byte, error) {

	var body []byte

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf := make([]byte, bufferSize)

	handler := bufio.NewReader(f)

	for {
		n, err := handler.Read(buf)

		body = ByteCombine(body, buf[:n])
		if err != nil {
			if err == io.EOF {
				return body, nil
			}

			return nil, err
		}

	}
}

func RemoveFile(path string) bool {
	err := os.Remove(path)
	if err != nil {
		return false
	}

	return true
}

func WriteFile(path string, content string) (bool, error) {
	//os.Open 只能只读
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return false, err
	}

	_, err = io.WriteString(f, content)

	fmt.Println(f)
	if err != nil {
		return false, err
	}

	defer f.Close()
	defer f.Sync()

	return true, nil
}

func NewFile(path string) (bool) {
	ret := FileIsExist(path)

	if ret {
		return true
	}

	//os.Create直接创建文件 NewFile 创建文件 不保存 返回*File对象
	f, err := os.Create(path)

	if err != nil {
		return false
	}

	defer f.Close()

	return true
}

/**
 * 文件是否存在
 */
func FileIsExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func writeFileByIoutil(path string, body string) {

}
