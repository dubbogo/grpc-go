package encoding

import (
	// perrors "github.com/pkg/errors"
)

import (
	// "github.com/dubbogo/grpc-go/encoding/proto_wrapper_api"
)

// PBWrapperTwoWayCodec is codec impl of pb
type PBWrapperTwoWayCodec struct {
	pbCodec Codec
	codec   Codec
	name    string
}

// NewPBWrapperTwoWayCodec new common.TwoWayCodec PBWrapperTwoWayCodec with @codecName defined Codec inside
func NewPBWrapperTwoWayCodec(name string, innerCodec, pbCodec Codec) TwoWayCodec {
	return &PBWrapperTwoWayCodec{
		codec:   innerCodec,
		name:    name,
		pbCodec: pbCodec,
	}
}

func (h *PBWrapperTwoWayCodec) Name() string {
	return h.name
}

// MarshalRequest marshal interface @v to []byte
func (h *PBWrapperTwoWayCodec) MarshalRequest(v interface{}) ([]byte, error) {
	// reqList := v.([]interface{})
	// argsBytes := make([][]byte, 0, len(reqList))
	// argsTypes := make([]string, 0, len(reqList))
	// for _, value := range reqList {
	// 	data, err := h.codec.Marshal(value)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	argsBytes = append(argsBytes, data)
	// 	argsTypes = append(argsTypes, GetArgType(value))
	// }

	// wrapperRequest := &proto_wrapper_api.TripleRequestWrapper{
	// 	SerializeType: h.Name(),
	// 	Args:          argsBytes,
	// 	ArgTypes:      argsTypes,
	// }
	return h.pbCodec.Marshal(v)
}

// UnmarshalRequest unmarshal bytes @data to interface
func (h *PBWrapperTwoWayCodec) UnmarshalRequest(data []byte, v interface{}) error {
	// wrapperRequest := proto_wrapper_api.TripleRequestWrapper{}
	err := h.pbCodec.Unmarshal(data, v)
	if err != nil {
		return err
	}

	// paramsInterfaces := v.([]interface{})
	// if len(paramsInterfaces) != len(wrapperRequest.Args) {
	// 	return perrors.Errorf("error ,request params len is %d, but exported method has %d", len(wrapperRequest.Args), len(paramsInterfaces))
	// }

	// for idx, value := range wrapperRequest.Args {
	// 	if err := h.codec.Unmarshal(value, paramsInterfaces[idx]); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// MarshalResponse marshal interface @v to []byte
func (h *PBWrapperTwoWayCodec) MarshalResponse(v interface{}) ([]byte, error) {
	// data, err := h.codec.Marshal(v)
	// if err != nil {
	// 	return nil, err
	// }

	// wrapperRequest := &proto_wrapper_api.TripleResponseWrapper{
	// 	SerializeType: h.name,
	// 	Data:          data,
	// 	Type:          GetArgType(v),
	// }
	return h.pbCodec.Marshal(v)
}

// UnmarshalResponse unmarshal bytes @data to interface
func (h *PBWrapperTwoWayCodec) UnmarshalResponse(data []byte, v interface{}) error {
	// wrapperResponse := proto_wrapper_api.TripleResponseWrapper{}
	return h.pbCodec.Unmarshal(data, v)
}

// PBTwoWayCodec is pb impl of TwoWayCodec
type PBTwoWayCodec struct {
	codec Codec
}
