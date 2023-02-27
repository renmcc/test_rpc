package pb_test

import (
	"fmt"
	"test-grpc/protobuf/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestMarsha1(t *testing.T) {
	should := assert.New(t)
	str := &pb.String{Value: "hello"}
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Printf("pbBytes: %v\n", pbBytes)
	}

	// 反序列化 []byte --protobuf --> object
	obj := new(pb.String)
	err = proto.Unmarshal(pbBytes, obj)
	if should.NoError(err) {
		fmt.Printf("obj: %+v\n", *obj)
	}
}
