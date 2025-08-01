// Code generated by Kitex v0.13.1. DO NOT EDIT.

package view

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"
	kutils "github.com/cloudwego/kitex/pkg/utils"

	"github.com/coze-dev/coze-loop/backend/kitex_gen/coze/loop/observability/domain/common"
)

var (
	_ = common.KitexUnusedProtection
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *View) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetID bool = false
	var issetViewName bool = false
	var issetFilters bool = false
	var issetIsSystem bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetID = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetViewName = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetFilters = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetIsSystem = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetViewName {
		fieldId = 4
		goto RequiredFieldNotSetError
	}

	if !issetFilters {
		fieldId = 7
		goto RequiredFieldNotSetError
	}

	if !issetIsSystem {
		fieldId = 8
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_View[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_View[fieldId]))
}

func (p *View) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.ID = _field
	return offset, nil
}

func (p *View) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.EnterpriseID = _field
	return offset, nil
}

func (p *View) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.WorkspaceID = _field
	return offset, nil
}

func (p *View) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.ViewName = _field
	return offset, nil
}

func (p *View) FastReadField5(buf []byte) (int, error) {
	offset := 0

	var _field *common.PlatformType
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.PlatformType = _field
	return offset, nil
}

func (p *View) FastReadField6(buf []byte) (int, error) {
	offset := 0

	var _field *common.SpanListType
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.SpanListType = _field
	return offset, nil
}

func (p *View) FastReadField7(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Filters = _field
	return offset, nil
}

func (p *View) FastReadField8(buf []byte) (int, error) {
	offset := 0

	var _field bool
	if v, l, err := thrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.IsSystem = _field
	return offset, nil
}

func (p *View) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *View) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField8(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
		offset += p.fastWriteField5(buf[offset:], w)
		offset += p.fastWriteField6(buf[offset:], w)
		offset += p.fastWriteField7(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *View) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *View) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 1)
	offset += thrift.Binary.WriteI64(buf[offset:], p.ID)
	return offset
}

func (p *View) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetEnterpriseID() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.EnterpriseID)
	}
	return offset
}

func (p *View) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetWorkspaceID() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 3)
		offset += thrift.Binary.WriteI64(buf[offset:], *p.WorkspaceID)
	}
	return offset
}

func (p *View) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 4)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.ViewName)
	return offset
}

func (p *View) fastWriteField5(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetPlatformType() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 5)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.PlatformType)
	}
	return offset
}

func (p *View) fastWriteField6(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSpanListType() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 6)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.SpanListType)
	}
	return offset
}

func (p *View) fastWriteField7(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 7)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Filters)
	return offset
}

func (p *View) fastWriteField8(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.BOOL, 8)
	offset += thrift.Binary.WriteBool(buf[offset:], p.IsSystem)
	return offset
}

func (p *View) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *View) field2Length() int {
	l := 0
	if p.IsSetEnterpriseID() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.EnterpriseID)
	}
	return l
}

func (p *View) field3Length() int {
	l := 0
	if p.IsSetWorkspaceID() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I64Length()
	}
	return l
}

func (p *View) field4Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.ViewName)
	return l
}

func (p *View) field5Length() int {
	l := 0
	if p.IsSetPlatformType() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.PlatformType)
	}
	return l
}

func (p *View) field6Length() int {
	l := 0
	if p.IsSetSpanListType() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.SpanListType)
	}
	return l
}

func (p *View) field7Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Filters)
	return l
}

func (p *View) field8Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.BoolLength()
	return l
}

func (p *View) DeepCopy(s interface{}) error {
	src, ok := s.(*View)
	if !ok {
		return fmt.Errorf("%T's type not matched %T", s, p)
	}

	p.ID = src.ID

	if src.EnterpriseID != nil {
		var tmp string
		if *src.EnterpriseID != "" {
			tmp = kutils.StringDeepCopy(*src.EnterpriseID)
		}
		p.EnterpriseID = &tmp
	}

	if src.WorkspaceID != nil {
		tmp := *src.WorkspaceID
		p.WorkspaceID = &tmp
	}

	if src.ViewName != "" {
		p.ViewName = kutils.StringDeepCopy(src.ViewName)
	}

	if src.PlatformType != nil {
		tmp := *src.PlatformType
		p.PlatformType = &tmp
	}

	if src.SpanListType != nil {
		tmp := *src.SpanListType
		p.SpanListType = &tmp
	}

	if src.Filters != "" {
		p.Filters = kutils.StringDeepCopy(src.Filters)
	}

	p.IsSystem = src.IsSystem

	return nil
}
