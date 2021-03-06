// generated by jwg -type Sample -output misc/fixture/a/model_json.go misc/fixture/a; DO NOT EDIT

package a

import (
	"encoding/json"
)

// SampleJSON is jsonized struct for Sample.
type SampleJSON struct {
	Foo string `json:"foo,omitempty"`
}

// SampleJSONList is synonym about []*SampleJSON.
type SampleJSONList []*SampleJSON

// SamplePropertyEncoder is property encoder for [1]sJSON.
type SamplePropertyEncoder func(src *Sample, dest *SampleJSON) error

// SamplePropertyDecoder is property decoder for [1]sJSON.
type SamplePropertyDecoder func(src *SampleJSON, dest *Sample) error

// SamplePropertyInfo stores property information.
type SamplePropertyInfo struct {
	name    string
	Encoder SamplePropertyEncoder
	Decoder SamplePropertyDecoder
}

// SampleJSONBuilder convert between Sample to SampleJSON mutually.
type SampleJSONBuilder struct {
	_properties map[string]*SamplePropertyInfo
	Foo         *SamplePropertyInfo
}

// NewSampleJSONBuilder make new SampleJSONBuilder.
func NewSampleJSONBuilder() *SampleJSONBuilder {
	return &SampleJSONBuilder{
		_properties: map[string]*SamplePropertyInfo{},
		Foo: &SamplePropertyInfo{
			name: "Foo",
			Encoder: func(src *Sample, dest *SampleJSON) error {
				if src == nil {
					return nil
				}
				dest.Foo = src.Foo
				return nil
			},
			Decoder: func(src *SampleJSON, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.Foo = src.Foo
				return nil
			},
		},
	}
}

// AddAll adds all property to SampleJSONBuilder.
func (b *SampleJSONBuilder) AddAll() *SampleJSONBuilder {
	b._properties["Foo"] = b.Foo
	return b
}

// Add specified property to SampleJSONBuilder.
func (b *SampleJSONBuilder) Add(info *SamplePropertyInfo) *SampleJSONBuilder {
	b._properties[info.name] = info
	return b
}

// Remove specified property to SampleJSONBuilder.
func (b *SampleJSONBuilder) Remove(info *SamplePropertyInfo) *SampleJSONBuilder {
	delete(b._properties, info.name)
	return b
}

// Convert specified non-JSON object to JSON object.
func (b *SampleJSONBuilder) Convert(orig *Sample) (*SampleJSON, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &SampleJSON{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ConvertList specified non-JSON slice to JSONList.
func (b *SampleJSONBuilder) ConvertList(orig []*Sample) (SampleJSONList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(SampleJSONList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

// Convert specified JSON object to non-JSON object.
func (orig *SampleJSON) Convert() (*Sample, error) {
	ret := &Sample{}

	b := NewSampleJSONBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// Convert specified JSONList to non-JSON slice.
func (jsonList SampleJSONList) Convert() ([]*Sample, error) {
	orig := ([]*SampleJSON)(jsonList)

	list := make([]*Sample, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

// Marshal non-JSON object to JSON string.
func (b *SampleJSONBuilder) Marshal(orig *Sample) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}
