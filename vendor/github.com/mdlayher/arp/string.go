// Code generated by "stringer -output=string.go -type=Operation"; DO NOT EDIT.

package arp

import "strconv"

const _Operation_name = "OperationRequestOperationReply"

var _Operation_index = [...]uint8{0, 16, 30}

func (i Operation) String() string {
	i -= 1
	if i >= Operation(len(_Operation_index)-1) {
		return "Operation(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Operation_name[_Operation_index[i]:_Operation_index[i+1]]
}
