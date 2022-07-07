package encoding

// PBWrapperTwoWayCodec is codec impl of pb
type PBWrapperTwoWayCodec struct {
	pbCodec Codec
	codec   Codec
	name    string
}

// NewPBWrapperTwoWayCodec new common.TwoWayCodec PBWrapperTwoWayCodec with @codecName defined Codec inside
func NewPBWrapperTwoWayCodec(name string, innerCodec, pbCodec Codec) Codec {
	return &PBWrapperTwoWayCodec{
		codec:   innerCodec,
		name:    name,
		pbCodec: pbCodec,
	}
}

func (h *PBWrapperTwoWayCodec) Name() string {
	return h.name
}

// Marshal marshal interface @v to []byte
func (h *PBWrapperTwoWayCodec) Marshal(v interface{}) ([]byte, error) {
	return h.pbCodec.Marshal(v)
}

// Unmarshal unmarshal bytes @data to interface
func (h *PBWrapperTwoWayCodec) Unmarshal(data []byte, v interface{}) error {
	return h.pbCodec.Unmarshal(data, v)
}

// PBTwoWayCodec is pb impl of TwoWayCodec
type PBTwoWayCodec struct {
	codec Codec
}
