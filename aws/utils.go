package aws

import (
	"encoding/json"
	"reflect"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func looksLikeJsonString(s interface{}) bool {
	return regexp.MustCompile(`^\s*{`).MatchString(s.(string))
}

func jsonBytesEqual(b1, b2 []byte) bool {
	var o1 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}

	var o2 interface{}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}

	return reflect.DeepEqual(o1, o2)
}

func isResourceTimeoutError(err error) bool {
	timeoutErr, ok := err.(*resource.TimeoutError)
	return ok && timeoutErr.LastError == nil
}
func isResourceNotFoundError(err error) bool {
	_, ok := err.(*resource.NotFoundError)
	return ok
}
