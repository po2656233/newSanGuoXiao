package base

import (
	"context"
	rand2 "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/po2656233/goleaf/log"
	"gorm.io/gorm/logger"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unicode/utf8"
	"unsafe"
)

// 生成随机字符串
const (
	character = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789;:'[]{}~!@#$%^&*()<>?/\\`.,|-_+="
	letter    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandomStr 随机生成字符
func RandomStr(count int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, count)
	for i := range b {
		b[i] = character[rand.Intn(len(character))]
	}
	return string(b)
}

// RandomStrLetter 随机生成字母
func RandomStrLetter(count int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, count)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// CheckError 异常处理
func CheckError(err error) bool {
	if err != nil {
		//fmt.Println(err.Error())
		log.Error("数据异常:%v", err) // 正式开服的时候，这里一定是错误类型 错误内容 等异常信息输出到日志上面的
		return false
	}
	return true
}

// 将任意切片类型转换为 []interface{}
// 示例
// tiles := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
// arraySlice:=ToAnyTypeSlice(tiles)
//
//	if newArray,ok:=arraySlice.([]int32);ok{
//	   .....
//	}

func ToAnyTypeSlice(slice interface{}) []interface{} {
	//判断是否是切片类型
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}
	sliceLen := v.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = v.Index(i).Interface()
	}
	return out
}

// slice转为数组结构体

// CopyInsert 插入某元素
func CopyInsert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface()
}

// DeleteKey 删除key 慎用
func DeleteKey(slice interface{}, index int) interface{} {
	//判断是否是切片类型
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}
	//参数检查
	if v.Len() == 0 || index < 0 || index > v.Len()-1 {
		return nil
	}
	return reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())).Interface()
}

// DeleteValue 删除值
func DeleteValue(slice interface{}, value interface{}) interface{} {
	//判断是否是切片类型
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}
	for i := 0; i < v.Len(); i++ {
		if reflect.ValueOf(value).IsValid() {
			if v.Index(i).Kind() == reflect.ValueOf(value).Kind() {
				if reflect.DeepEqual(v.Index(i).Interface(), value) {
					return reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())).Interface()
				}
			}
		}
	}
	return slice
}

// SliceRemoveDuplicate 删除重复的数值
func SliceRemoveDuplicate(data interface{}) interface{} {
	inArr := reflect.ValueOf(data)
	if inArr.Kind() != reflect.Slice && inArr.Kind() != reflect.Array {
		return data
	}

	existMap := make(map[interface{}]bool)
	outArr := reflect.MakeSlice(inArr.Type(), 0, inArr.Len())

	for i := 0; i < inArr.Len(); i++ {
		iVal := inArr.Index(i)

		if _, ok := existMap[iVal.Interface()]; !ok {
			outArr = reflect.Append(outArr, inArr.Index(i))
			existMap[iVal.Interface()] = true
		}
	}

	return outArr.Interface()
}

// ClearSlice 清空slice
func ClearSlice(s *[]interface{}) {
	*s = (*s)[0:0]
}

//  ------------------------------------

// IntAbs 取整数绝对值
func IntAbs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// Float2Digit 保留位数
func Float2Digit(f float64, digit int) float64 {
	format := "%1." + strconv.Itoa(digit) + "f"
	str := fmt.Sprintf(format, f)
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(str, err.Error())
	}
	return i
}

// Wrap 将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// Unwrap 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// WrapToFloat64 精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

// UnwrapToInt64 精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}

// Int64ToBytes 整型转bytes数组
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Int16ToBytes(i int16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// BytesToPB proto的转化
func BytesToPB(data []byte, toPb proto.Message) error {
	return proto.Unmarshal(data, toPb)
}
func PBToBytes(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

func JSON2PB(formJsonStr string, toPb proto.Message) error {
	// json字符串转pb
	return json.Unmarshal([]byte(formJsonStr), &toPb)
}

func PB2JSON(fromPb proto.Message, toJsonStr string) error {
	// pb转json字符串
	jsonStr, err := json.Marshal(fromPb)
	if err == nil {
		toJsonStr = string(jsonStr)
	}

	return err
}

func BytesToRune(byteArray []byte) []rune {
	runeArray := make([]rune, 0)
	for _, b := range byteArray {
		if b != 0 {
			r, _ := utf8.DecodeRune([]byte{b})
			runeArray = append(runeArray, r)
		}
	}
	return runeArray
}

// IsSatisfy 是否命中 百分比
func IsSatisfy(rate int) bool {
	r2, _ := rand2.Int(rand2.Reader, big.NewInt(100))
	return 0 == r2.Int64()%int64(rate)
}

// GetOneDay 获取当天0点和24点时间戳
// beginTimeNum  0点
// endTimeNum  24点
func GetOneDay() (beginTimeNum, endTimeNum int64) {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	beginTimeNum = t.Unix()
	endTimeNum = beginTimeNum + 86400
	return beginTimeNum, endTimeNum
}

//写成宏，方便移植
//#define setbit(x,y) x|=(1<<y) //将X的右起 第Y-1位置1
//#define clrbit(x,y) x&=～(1<<y) //将X的右起 第Y-1位清0
//指定位数上置一
//func BitSetOne(num int64, bit int) int64 {
//	if bit < 1 {
//		return num
//	}
//	return num | 1<<(bit-1)
//}

//指定位数上置零
//func BitSetZero(num int, bit int) int32 {
//	if bit < 1 {
//		return num
//	}
//	subtrahend := [32]int{Exponent1, Exponent2, Exponent3, Exponent4, Exponent5, Exponent6, Exponent7, Exponent8, Exponent9, Exponent10, Exponent11, Exponent12, Exponent13, Exponent14, Exponent15, Exponent16, Exponent17, Exponent18, Exponent19, Exponent20, Exponent21, Exponent22, Exponent23, Exponent24, Exponent25, Exponent26, Exponent27, Exponent28, Exponent29, Exponent30, Exponent31, Exponent32}
//
//	return num &^ subtrahend[bit-1]
//}

// RandomNum 生成固定随机 elementMax元素的最大值(注:最后一个元素不受限制,其他元素必须等于或小于elementMax) remainder余数 count元素个数
func RandomNum(elementMax, remainder, count int) ([]int64, bool) {
	var data []int64

	if count < 1 || elementMax < 1 || remainder == 0 {
		return data, false
	}

	if remainder/2 < elementMax {
		elementMax = remainder / 2
	}

	//fmt.Printf("RandomNum-> remain:%v count:%v ", remainder, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count-1; i++ {
		data = append(data, int64(rand.Int()%elementMax)+1)
	}
	sum := int64(0)
	for _, v := range data {
		sum += v
	}
	data = append(data, int64(remainder)-sum%int64(remainder))
	return data, true
}

// 群组聊天id
var gChatID int64

func GetGChatID() int64 {
	return atomic.AddInt64(&gChatID, 1)
}

// 个人聊天id
var pChatID int64

func GetPChatID() int64 {
	return atomic.AddInt64(&pChatID, 1)
}

func MergeID(a, b int64) string {
	sum := a + b
	if b < a {
		a, b = b, a
	}
	return "*" + strconv.FormatInt(sum, 10) + strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10)
}

//func KeFuID(uid int64) string {
//	return "**_" + strconv.FormatUint(uid, 10)
//}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////

// PrintFuncInfo 打印函数名和行号
func PrintFuncInfo(layer int) {
	//参数 layer 函数所在的层数
	file, fileName, line, ok := runtime.Caller(layer + 1)
	if ok {
		funcName := runtime.FuncForPC(file).Name()
		log.Release(funcName, " -> ", line, " -> ", fileName)
		//fmt.Printf("%s:%d -> -> ->%s\n", funcName, line, fileName)
	}
}

// GoID 获取goroutine的id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	log.Release("id:", id)
	return id
}

// DumpStacks 堆栈信息输出
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

// id 生成器
////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	maxStack  = 20
	separator = "---------------------------------------\n"
)

var panicHandler func(string)

func OnPanic(h func(string)) {
	panicHandler = func(str string) {
		defer func() {
			recover()
		}()
		h(str)
	}
}

func HandlePanic() {
	if err := recover(); err != nil {
		errstr := fmt.Sprintf("\n%sruntime error: %v\ntraceback:\n", separator, err)

		i := 2
		for {
			pc, file, line, ok := runtime.Caller(i)
			if !ok || i > maxStack {
				break
			}
			errstr += fmt.Sprintf("\tstack: %d %v [file: %s] [func: %s] [line: %d]\n", i-1, ok, file, runtime.FuncForPC(pc).Name(), line)
			i++
		}
		errstr += separator

		if panicHandler != nil {
			panicHandler(errstr)
		} else {
			log.Error(errstr)
		}
	}
}
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func IsName(name string) bool {
	if name != "" {
		if isOk, _ := regexp.MatchString("^[\u4e00-\u9fa5_a-zA-Z0-9]+$", name); isOk {
			return true
		}
	}
	return false
}

func Safe(cb func()) {
	defer HandlePanic()
	cb()
}

// //////////////////////////////////与游戏强相关//////////////////////////////////////////////////////////

// GenerateGameNum 生成游戏牌局号
func GenerateGameNum(gameName string, level, tableID int32) string {
	return gameName + strconv.Itoa(int(level)) + strconv.Itoa(int(tableID)) + strconv.FormatInt(time.Now().Unix(), 10)
}

// GenerateRobotNum 生成机器人牌局号
func GenerateRobotNum(tableID int64) string {
	return strconv.Itoa(int(tableID)) + strconv.FormatInt(time.Now().Unix(), 10)
}

/////////////////////////////////////////////////////////////

// ToArray slice转为数组结构体
func ToArray(slice interface{}, dataType interface{}) interface{} {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		return nil
	}
	sliceLen := value.Len()
	typeValue := reflect.ValueOf(dataType)

	code := 0
	switch value.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		code = 1
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		code = 2
	case reflect.Float64, reflect.Float32:
		code = 3
	case reflect.String:
		code = 4
	case reflect.Interface:
		code = 5
	default:
		return value.Interface()
	}

	switch typeValue.Kind() {
	case reflect.Int:
		r := make([]int, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int(value.Index(i).Int())
			case 2:
				r[i] = int(value.Index(i).Uint())
			case 3:
				r[i] = int(value.Index(i).Float())
			case 4:
				if v, err := strconv.Atoi(value.Index(i).String()); err == nil {
					r[i] = v
				}
			case 5:
				r[i] = value.Index(i).Interface().(int)
			}
		}
		return r
	case reflect.Int8:
		r := make([]int8, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int8(value.Index(i).Int())
			case 2:
				r[i] = int8(value.Index(i).Uint())
			case 3:
				r[i] = int8(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int8(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int8)
			}
		}
		return r
	case reflect.Int16:
		r := make([]int16, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int16(value.Index(i).Int())
			case 2:
				r[i] = int16(value.Index(i).Uint())
			case 3:
				r[i] = int16(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int16(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int16)
			}
		}
		return r
	case reflect.Int32:
		r := make([]int32, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int32(value.Index(i).Int())
			case 2:
				r[i] = int32(value.Index(i).Uint())
			case 3:
				r[i] = int32(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int32(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int32)
			}
		}
		return r
	case reflect.Int64:
		r := make([]int64, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = value.Index(i).Int()
			case 2:
				r[i] = int64(value.Index(i).Uint())
			case 3:
				r[i] = int64(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = v
				}
			case 5:
				r[i] = value.Index(i).Interface().(int64)
			}
		}
		return r
	case reflect.Uint:
		r := make([]uint, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint(value.Index(i).Int())
			case 2:
				r[i] = uint(value.Index(i).Uint())
			case 3:
				r[i] = uint(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint)
			}
		}
		return r
	case reflect.Uint8:
		r := make([]uint8, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint8(value.Index(i).Int())
			case 2:
				r[i] = uint8(value.Index(i).Uint())
			case 3:
				r[i] = uint8(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint8(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint8)
			}
		}
		return r
	case reflect.Uint16:
		r := make([]uint16, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint16(value.Index(i).Int())
			case 2:
				r[i] = uint16(value.Index(i).Uint())
			case 3:
				r[i] = uint16(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint16(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint16)
			}
		}
		return r
	case reflect.Uint32:
		r := make([]int32, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int32(value.Index(i).Int())
			case 2:
				r[i] = int32(value.Index(i).Uint())
			case 3:
				r[i] = int32(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int32(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int32)
			}
		}
		return r
	case reflect.Uint64:
		r := make([]int64, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = value.Index(i).Int()
			case 2:
				r[i] = int64(value.Index(i).Uint())
			case 3:
				r[i] = int64(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int64(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int64)
			}
		}
		return r
	case reflect.Float32:
		r := make([]float32, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = float32(value.Index(i).Int())
			case 2:
				r[i] = float32(value.Index(i).Uint())
			case 3:
				r[i] = float32(value.Index(i).Float())
			case 4:
				if f, err := strconv.ParseFloat(value.Index(i).String(), 32); err == nil {
					r[i] = float32(f)
				}
			case 5:
				r[i] = value.Index(i).Interface().(float32)
			}
		}
		return r
	case reflect.Float64:
		r := make([]float64, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = float64(value.Index(i).Int())
			case 2:
				r[i] = float64(value.Index(i).Uint())
			case 3:
				r[i] = value.Index(i).Float()
			case 4:
				r[i], _ = strconv.ParseFloat(value.Index(i).String(), 64)

			case 5:
				r[i] = value.Index(i).Interface().(float64)
			}
		}
		return r
	case reflect.String:
		r := make([]string, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = strconv.FormatInt(value.Index(i).Int(), 10)
			case 2:
				r[i] = strconv.FormatUint(value.Index(i).Uint(), 10)
			case 3:
				r[i] = strconv.FormatFloat(value.Index(i).Float(), 'f', -1, 64)
			case 4:
				r[i] = value.Index(i).String()
			case 5:
				r[i] = value.Index(i).Interface().(string)
			}
		}
		return r
	}
	return value.Interface()
}

////////////////////////////

type DiscardLogger struct{}

func (d DiscardLogger) LogMode(logger.LogLevel) logger.Interface {
	return &d
}
func (d DiscardLogger) Info(context.Context, string, ...interface{}) {
}
func (d DiscardLogger) Warn(context.Context, string, ...interface{}) {
}
func (d DiscardLogger) Error(context.Context, string, ...interface{}) {
}
func (d DiscardLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}
