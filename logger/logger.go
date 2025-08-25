package logger

import (
	"encoding/base64"
	"fmt"
	"iox/option"
	"os"
)

const (
	pWARN    = "[!] "
	pINFO    = "[+] "
	pSUCCESS = "[*] "
	pRAW     = ""
	pBASE64  = "[base64] "
)

func Info(format string, args ...interface{}) {
	if option.VERBOSE {
		fmt.Fprintf(os.Stdout, pINFO+format+"\n", args...)
	}
}

func Warn(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pWARN+format+"\n", args...)
}

func Success(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, pSUCCESS+format+"\n", args...)
}

func to_raw(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, pRAW+format+"\n", args...)
}

func to_base64(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, pBASE64+format+"\n", args...)
}

func File(data []byte) {
	if option.BASE64 {
		encoded := base64.StdEncoding.EncodeToString(data)
		to_base64("%s", encoded)
	} else {
		to_raw("%s", string(data))
	}
}
