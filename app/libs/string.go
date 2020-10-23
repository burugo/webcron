package libs

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/axgle/mahonia"
)

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

//ConvertUtf8 gbk 解码
func ConvertUtf8(str string) string {
	if !utf8.ValidString(str) {
		utf8Decoder := mahonia.NewDecoder("gbk")
		return utf8Decoder.ConvertString(str)
	}
	return str
}
