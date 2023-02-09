package u_file

import (
	"SamgeWxApi/cmd/iface"
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// Read 读取文件
func Read(filePath string) (string, error) {
	exits, _ := PathExists(filePath)
	if !exits {
		return "", nil
	}
	txt, err := os.ReadFile(filePath)
	return string(txt), err
}

// Open 打开文件
func Open(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("文件打开失败：%v", err)
		return nil, err
	}
	return file, nil
}

// ReadXml 读取xml文件
// Go读文件：https://blog.csdn.net/slphahaha/article/details/122995124
// Go中[]byte与string转换：https://blog.csdn.net/slphahaha/article/details/109405685
// xml文件写入与读取：http://c.biancheng.net/view/4551.html
// xml在线格式化：https://c.runoob.com/front-end/710/
// xml转go的struct：https://tool.hiofd.com/xml-to-go/
func ReadXml(filePath string, infoStruct iface.BaseInterface) error {
	file, err := Open(filePath)
	defer file.Close()
	decoder := xml.NewDecoder(file) //创建 xml 解码器
	err = decoder.Decode(&infoStruct)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return err
	}
	return nil
}

// ReadXml2 读取xml文件，方式2
func ReadXml2(filePath string, infoStruct iface.BaseInterface) error {
	xmlStr, err := Read(filePath)
	if err != nil {
		return err
	}
	return ReadXmlByStr(xmlStr, infoStruct)
}

// ReadXmlByStr 读取文本格式的xml并解析为结构体
func ReadXmlByStr(xmlStr string, infoStruct iface.BaseInterface) error {
	xmlSrc := []byte(xmlStr)
	if err := xml.Unmarshal(xmlSrc, infoStruct); err != nil {
		return err
	}
	return nil
}

// SaveCover 覆盖保存
func SaveCover(filePath string, txt string) error {
	return Save(filePath, txt, os.O_WRONLY|os.O_CREATE)
}

// SaveAppend 追加保存
func SaveAppend(filePath string, txt string) error {
	return Save(filePath, txt, os.O_WRONLY|os.O_APPEND)
}

// Save 保存文件
func Save(filePath string, txt string, flag int) error {

	var file *os.File
	var err error
	file, err = os.OpenFile(filePath, flag, 0666)
	if err != nil {
		log.Printf("该文本不存在，自动新建该文件")
		file, err = os.Create(filePath)
		if err != nil {
			log.Printf("新建文件错误= %v \n", err)
			return err
		}
	}

	//及时关闭
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("关闭文件时出错：%v\n", err.Error())
		}
	}(file)

	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(txt)
	if err != nil {
		return err
	}

	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	status, _ := PathExists(path)
	return status
}

// CreateMultiDir 调用os.MkdirAll递归创建文件夹
func CreateMultiDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}
