/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package proto defines the protobuf codec. Importing this package will
// register the codec.
package proto

import (
	"github.com/dubbogo/grpc-go/encoding"
	"github.com/dubbogo/grpc-go/encoding/raw_proto"
)

// Name is the name registered for the proto compressor.
const Name = "proto"

func init() {
	encoding.RegisterCodec(NewPBTwoWayCodec())
}

// PBTwoWayCodec is pb impl of TwoWayCodec
type PBTwoWayCodec struct {
	codec encoding.Codec
}

func (h *PBTwoWayCodec) Name() string {
	return "proto"
}

// NewPBTwoWayCodec new PBTwoWayCodec instance
func NewPBTwoWayCodec() encoding.Codec {
	return &PBTwoWayCodec{
		codec: raw_proto.NewProtobufCodec(),
	}
}

// Marshal marshal interface @v to []byte
func (h *PBTwoWayCodec) Marshal(v interface{}) ([]byte, error) {
	return h.codec.Marshal(v)
}

// Unmarshal unmarshal bytes @data to interface
func (h *PBTwoWayCodec) Unmarshal(data []byte, v interface{}) error {
	return h.codec.Unmarshal(data, v)
}
