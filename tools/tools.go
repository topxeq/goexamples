package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/topxeq/txtk"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// StartsWith 检查字符串strA开始是否是subStrA
func StartsWith(strA string, subStrA string) bool {

	return strings.HasPrefix(strA, subStrA)
}

// EndsWith 检查字符串strA结尾是否是subStrA
func EndsWith(strA string, subStrA string) bool {

	return strings.HasSuffix(strA, subStrA)
}

// Trim 去除字符串首尾的空白字符
func Trim(strA string) string {
	return strings.TrimSpace(strA)
}

// GetFlag 检查命令行切片中是否存在某标志参数，如果存在则返回该标志参数的值，否则返回空字符串
// 例：如果对命令行 “clix calbmi -w=70.0 -h=1.75”调用GetFlag(args, "-w=")，则结果为字符串“70.0”
func GetFlag(argsA []string, flagA string) string {
	for _, argT := range argsA {
		if StartsWith(argT, flagA) {
			argLen := len(flagA)
			tmpStr := argT[argLen:]

			return tmpStr
		}

	}

	return ""
}

// FlagExists 判断命令行参数中是否存在开关，用法：flag := FlagExists(args, "-value")
func FlagExists(argsA []string, flagA string) bool {
	for _, argT := range argsA {
		if StartsWith(argT, flagA) {
			return true
		}
	}

	return false
}

// 用于标志是否初始化过随机数种子的变量
var ifRandomizedG = false

// Randomize 初始化随机数种子，不会重复操作
func Randomize() {
	if !ifRandomizedG {
		rand.Seed(time.Now().Unix())
		ifRandomizedG = true
	}
}

// GenerateRandomString 生成一个可定制的随机字符串
func GenerateRandomString(minCharCountA, maxCharCountA int, hasUpperA, hasLowerA, hasDigitA, hasSpecialCharA, hasSpaceA bool, hasOtherChars bool) string {
	Randomize()

	if minCharCountA <= 0 {
		return ""
	}

	if maxCharCountA <= 0 {
		return ""
	}

	if minCharCountA > maxCharCountA {
		return ""
	}

	countT := minCharCountA + rand.Intn(maxCharCountA+1-minCharCountA)

	baseT := ""
	if hasUpperA {
		baseT += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if hasLowerA {
		baseT += "abcdefghijklmnopqrstuvwxyz"
	}

	if hasDigitA {
		baseT += "0123456789"
	}

	if hasSpecialCharA {
		baseT += "!@#$%^&*-=[]{}."
	}

	if hasSpaceA {
		baseT += " "
	}

	if hasOtherChars {
		baseT += "/\\:*\"<>|(),+?;"
	}

	rStrT := ""
	var idxT int

	for i := 0; i < countT; i++ {
		idxT = rand.Intn(len(baseT))
		rStrT += baseT[idxT:(idxT + 1)]
	}

	return rStrT
}

// GenerateRandomStringX 生成一个可定制的随机字符串，使用strings.Builder效率更高
func GenerateRandomStringX(minCharCountA, maxCharCountA int, hasUpperA, hasLowerA, hasDigitA, hasSpecialCharA, hasSpaceA bool, hasOtherChars bool) string {
	Randomize()

	if minCharCountA <= 0 {
		return ""
	}

	if maxCharCountA <= 0 {
		return ""
	}

	if minCharCountA > maxCharCountA {
		return ""
	}

	countT := minCharCountA + rand.Intn(maxCharCountA+1-minCharCountA)

	baseT := ""
	if hasUpperA {
		baseT += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if hasLowerA {
		baseT += "abcdefghijklmnopqrstuvwxyz"
	}

	if hasDigitA {
		baseT += "0123456789"
	}

	if hasSpecialCharA {
		baseT += "!@#$%^&*-=[]{}."
	}

	if hasSpaceA {
		baseT += " "
	}

	if hasOtherChars {
		baseT += "/\\:*\"<>|(),+?;"
	}

	var builderT strings.Builder
	var idxT int

	for i := 0; i < countT; i++ {
		idxT = rand.Intn(len(baseT))
		builderT.WriteByte(baseT[idxT])
	}

	return builderT.String()
}

// Printf 仅仅封装了fmt.Printf函数，与其完全一致
func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// Printfln 仅仅封装了fmt.Printf函数，但结尾会多输出一个换行符
func Printfln(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

// IntToString int类型转换为string类型
func IntToString(valueA int) string {
	return strconv.Itoa(valueA)
}

// ByteToString byte类型转换为string类型
func ByteToString(valueA byte) string {
	return strconv.FormatInt(int64(valueA), 10)
}

// IntegerToString 所有整数类型转换为string类型
func IntegerToString(valueA interface{}) string {
	switch valueA.(type) {
	case byte:
		return strconv.FormatInt(int64(valueA.(byte)), 10)
	case rune:
		return strconv.FormatInt(int64(valueA.(rune)), 10)
	case int64:
		return strconv.FormatInt(valueA.(int64), 10)
	default:
		return ""
	}
}

// NumberToString 所有主要数字类型转换为string类型
func NumberToString(valueA interface{}) string {
	switch valueA.(type) {
	case bool:
		return strconv.FormatBool(valueA.(bool))
	case byte:
		return strconv.FormatInt(int64(valueA.(byte)), 10)
	case rune:
		return strconv.FormatInt(int64(valueA.(rune)), 10)
	case int:
		return strconv.FormatInt(int64(valueA.(int)), 10)
	case int64:
		return strconv.FormatInt(valueA.(int64), 10)
	case uint32:
		return strconv.FormatUint(uint64(valueA.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(valueA.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(valueA.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(valueA.(float64), 'f', -1, 64)

	// 对于其他不能处理的类型返回空字符串
	default:
		return ""
	}
}

// LoadStringListFromFile 从文件中读取所有内容并返回为字符串切片，文件中每行为字符串切片中的一项
func LoadStringListFromFile(fileNameA string) []string {
	fileContentT, err := ioutil.ReadFile(fileNameA)
	if err != nil {
		return nil
	}

	strT := string(fileContentT)

	strT = strings.Replace(strT, "\r", "", -1)

	listT := strings.Split(strT, "\n")

	return listT
}

// LoadStringFromFile 从文件中读取所有内容并返回为字符串，如果出错则返回defaultA参数指定的字符串
func LoadStringFromFile(fileNameA string, defaultA string) string {
	fileContentT, err := ioutil.ReadFile(fileNameA)
	if err != nil {
		return defaultA
	}

	return string(fileContentT)
}

// LoadLinesFromFile 从文件中读取指定数量的行
func LoadLinesFromFile(fileNameA string, limitA int) string {

	fileT, err := os.Open(fileNameA)
	if err != nil {
		return "\u0001\u0001\u0001" + err.Error()
	}

	defer fileT.Close()

	var buf strings.Builder

	reader := bufio.NewReader(fileT)

	limitT := 0

	for true {
		strT, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		buf.WriteString(strT)

		limitT++

		if (limitA > 0) && (limitT >= limitA) {
			break
		}
	}

	return buf.String()
}

// SaveStringToFile 将字符串存入文件，如有原来有同名文件则其内容将被冲掉
func SaveStringToFile(strA string, fileNameA string) string {
	fileT, errT := os.Create(fileNameA)

	if errT != nil {
		return errT.Error()
	}

	defer fileT.Close()

	writerT := bufio.NewWriter(fileT)

	writerT.WriteString(strA)

	writerT.Flush()

	return ""
}

// AppendStringToFile 向文件中追加字符串，如果文件不存在则新建该文件后再追加
func AppendStringToFile(strA string, fileNameA string) string {

	fileT, errT := os.OpenFile(fileNameA, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errT != nil {
		return errT.Error()
	}

	defer fileT.Close()

	writerT := bufio.NewWriter(fileT)

	writerT.WriteString(strA)

	writerT.Flush()

	return ""
}

// FileExists 判断文件或目录是否存在
func FileExists(fileNameA string) bool {
	_, errT := os.Stat(fileNameA)
	return errT == nil || os.IsExist(errT)
}

// IsFile 判断路径名是否是文件
func IsFile(fileNameA string) bool {
	f, errT := os.Open(fileNameA)
	if errT != nil {
		return false
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false
	}

	if mode := fi.Mode(); mode.IsRegular() {
		return true
	} else {
		return false
	}
}

// IsDirectory 判断路径名是否是目录
func IsDirectory(dirNameA string) bool {
	f, err := os.Open(dirNameA)
	if err != nil {
		return false
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false
	}

	if mode := fi.Mode(); mode.IsDir() {
		return true
	} else {
		return false
	}
}

// ConvertBytesFromGB18030ToUTF8 转换GB18030编码的字节切片为UTF-8编码
func ConvertBytesFromGB18030ToUTF8(srcA []byte) []byte {

	bufT := make([]byte, len(srcA)*4)

	transformer := simplifiedchinese.GB18030.NewDecoder()

	countT, _, errT := transformer.Transform(bufT, srcA, true)

	if errT != nil {
		return nil
	}

	return bufT[:countT]
}

// ConvertBytesFromUTF8ToGB18030 转换UTF-8编码的字节切片为GB18030编码
func ConvertBytesFromUTF8ToGB18030(srcA []byte) []byte {

	bufT := make([]byte, len(srcA)*4)

	transformer := simplifiedchinese.GB18030.NewEncoder()

	countT, _, errT := transformer.Transform(bufT, srcA, true)

	if errT != nil {
		return nil
	}

	return bufT[:countT]
}

// ConvertBytesFromISO8859_1ToUTF8 转换ISO-8859-1编码的字节切片为UTF-8编码
func ConvertBytesFromISO8859_1ToUTF8(srcA []byte) []byte {

	bufT := make([]byte, len(srcA)*4)

	transformer := charmap.ISO8859_1.NewDecoder()

	countT, _, errT := transformer.Transform(bufT, srcA, true)

	if errT != nil {
		return nil
	}

	return bufT[:countT]
}

// GetFileTypeByHead 根据文件头的特殊字节判断常见文件类型
func GetFileTypeByHead(fileNameA string) (string, error) {

	fileT, errT := os.Open(fileNameA)

	if errT != nil {
		return "", errT
	}

	bufT := make([]byte, 10)

	_, errT = fileT.Read(bufT)

	if errT != nil {
		return "", errT
	}

	if bytes.HasPrefix(bufT, []byte{0xFF, 0xD8, 0xFF}) {
		return "jpg", nil
	}

	if bytes.HasPrefix(bufT, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png", nil
	}

	if bytes.HasPrefix(bufT, []byte{0x49, 0x44, 0x33, 03, 00}) || bytes.HasPrefix(bufT, []byte{0x49, 0x44, 0x33, 04, 00}) {
		return "mp3", nil
	}

	if bytes.HasPrefix(bufT, []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70}) {
		return "mp4", nil
	}

	if bytes.HasPrefix(bufT, []byte{0x4D, 0x5A}) {
		return "exe", nil
	}

	if bytes.HasPrefix(bufT, []byte{0x50, 0x4B, 0x03, 0x04}) {
		return "zip", nil
	}

	return "未知文件类型", nil

}

// AbsInt 获得整数（int类型）的绝对值
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// StringToInt 转换字符串为整数
func StringToInt(strA string) (int, error) {
	nT, errT := strconv.ParseInt(strA, 10, 0)
	if errT != nil {
		return 0, errT
	}

	return int(nT), nil
}

func DrawLine(imageA *image.NRGBA, x1, y1, x2, y2 int, colorA color.Color) {
	dx := AbsInt(x2 - x1)
	dy := AbsInt(y2 - y1)

	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}

	if y1 >= y2 {
		sy = -1
	}

	errT := dx - dy

	for {
		imageA.Set(x1, y1, colorA)

		if x1 == x2 && y1 == y2 {
			return
		}

		e2 := errT * 2

		if e2 > -dy {
			errT -= dy
			x1 += sx
		}

		if e2 < dx {
			errT += dx
			y1 += sy
		}
	}

}

// DownloadPageUTF8 用于下载UTF-8或兼容UTF-8编码的网页
func DownloadPageUTF8(urlA string, postDataA url.Values, timeoutSecsA time.Duration) (string, error) {
	client := &http.Client{
		Timeout: time.Second * timeoutSecsA,
	}

	var errT error

	var respT *http.Response

	if postDataA == nil {
		respT, errT = client.Get(urlA)
	} else {
		respT, errT = client.PostForm(urlA, postDataA)
	}

	if errT != nil {
		return "", errT
	}

	defer respT.Body.Close()

	if respT.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response status: %v", respT.StatusCode)
	}

	body, errT := ioutil.ReadAll(respT.Body)

	if errT != nil {
		return "", errT
	}

	return string(body), nil
}

// Fibonacci 计算斐波那契数列
func Fibonacci(c int64) int64 {
	if c < 2 {
		return c
	}

	return Fibonacci(c-2) + Fibonacci(c-1)
}

// CalPi 是使用随机落点法计算圆周率Π值的函数
// 一般来说，输入参数pointCountA的值越大，计算结果越准，但耗费时间也越多
func CalPi(pointCountA int) float64 {
	inCircleCount := 0

	var x, y float64
	var Pi float64

	for i := 0; i < pointCountA; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if x*x+y*y < 1 {
			inCircleCount++
		}
	}

	Pi = (4.0 * float64(inCircleCount)) / float64(pointCountA)

	return Pi
}

// CalPiX 是使用随机落点法计算圆周率Π值的函数
// 与CalPi唯一的不同是使用了更快的随机数发生器
// 但有可能不是并发安全的，建议仅在单线程中使用
func CalPiX(pointCountA int) float64 {
	inCircleCount := 0

	var x, y float64
	var Pi float64

	r := txtk.NewRandomGenerator()

	for i := 0; i < pointCountA; i++ {
		x = r.Float64()
		y = r.Float64()

		if x*x+y*y < 1 {
			inCircleCount++
		}
	}

	Pi = (4.0 * float64(inCircleCount)) / float64(pointCountA)

	return Pi
}

// CalCosSim 计算两个向量的余弦相似度
func CalCosSim(f1, f2 []float64) float64 {
	if f1 == nil || f2 == nil {
		Printfln("某个向量是空值nil")
		return -1
	}

	l1 := len(f1)
	l2 := len(f2)

	if l1 != l2 {
		Printfln("两个向量长度不一致，f1的长度是：%v，f2的长度是：%v", l1, l2)
		return -1
	}

	var rr float64 = 0.0
	var f1r float64 = 0.0
	var f2r float64 = 0.0

	for i := 0; i < l1; i++ {
		rr += f1[i] * f2[i]
		f1r += f1[i] * f1[i]
		f2r += f2[i] * f2[i]
	}

	var rs float64 = rr / (math.Sqrt(f1r) * math.Sqrt(f2r))

	return rs
}

// CalCosSimBig 计算两个向量的余弦相似度，使用big包避免计算溢出
func CalCosSimBig(f1, f2 []float64) float64 {
	if f1 == nil || f2 == nil {
		Printfln("某个向量是空值nil")
		return -1
	}

	l1 := len(f1)
	l2 := len(f2)

	if l1 != l2 {
		Printfln("两个向量长度不一致，f1的长度是：%v，f2的长度是：%v", l1, l2)
		return -1
	}

	rr := new(big.Float).SetFloat64(0.0)
	f1r := new(big.Float).SetFloat64(0.0)
	f2r := new(big.Float).SetFloat64(0.0)

	for i := 0; i < l1; i++ {
		f1Sub := new(big.Float).SetFloat64(f1[i])
		f2Sub := new(big.Float).SetFloat64(f2[i])

		rr.Add(rr, new(big.Float).Mul(f1Sub, f2Sub))
		f1r.Add(f1r, new(big.Float).Mul(f1Sub, f1Sub))
		f2r.Add(f2r, new(big.Float).Mul(f2Sub, f2Sub))
	}

	tmp1 := new(big.Float).Mul(new(big.Float).Sqrt(f1r), new(big.Float).Sqrt(f2r))

	tmp2 := new(big.Float).Quo(rr, tmp1)

	tmp3, _ := tmp2.Float64()

	return tmp3
}
