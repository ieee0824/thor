package util

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var AwsConfig *aws.Config

func init() {
	var cred *credentials.Credentials
	cred = credentials.NewSharedCredentials("", "default")
	AwsConfig = &aws.Config{
		Credentials: cred,
		Region:      aws.String("ap-northeast-1"),
	}
}

type LANG string

func IsMultiByteChar(r rune) bool {
	c := string([]rune{r})
	return 1 != len(c)
}

func ConvertRune(s string) (rune, error) {
	if len([]rune(s)) != 1 {
		var e rune
		return e, errors.New("not char")
	}
	return []rune(s)[0], nil
}

func MultiString(s string) string {
	array := strings.Split(s, "")
	var ret string

	for _, s := range array {
		r, _ := ConvertRune(s)
		if IsMultiByteChar(r) {
			ret = ret + s + " "
		} else {
			ret += s
		}
	}
	return ret
}
