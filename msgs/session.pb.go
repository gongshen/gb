// Code generated by protoc-gen-go from "session.proto"
// DO NOT EDIT!

package msgs

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error


type Session struct {
	Id			*int64	"PB(varint,1,req,name=id)"
	Timeout			*int64	"PB(varint,2,req,name=timeout)"
	XXX_unrecognized	[]byte
}

func (this *Session) Reset()		{ *this = Session{} }
func (this *Session) String() string	{ return proto.CompactTextString(this) }

func init() {
}