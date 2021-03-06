// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Chart struct {
	SeriesGid []*string `json:"seriesGID"`
	Labels    []*Label  `json:"labels"`
	X         *Series   `json:"x"`
	Y         []*Series `json:"y"`
}

type Label struct {
	SeriesGid []*string `json:"seriesGID"`
	Name      *string   `json:"name"`
	LabelType LabelType `json:"labelType"`
	Min       *float64  `json:"min"`
	Max       *float64  `json:"max"`
}

type Series struct {
	SeriesGid []*string  `json:"seriesGID"`
	Name      *string    `json:"name"`
	Values    []*float64 `json:"values"`
}

type LabelType string

const (
	LabelTypeMutable   LabelType = "MUTABLE"
	LabelTypeImmutable LabelType = "IMMUTABLE"
)

var AllLabelType = []LabelType{
	LabelTypeMutable,
	LabelTypeImmutable,
}

func (e LabelType) IsValid() bool {
	switch e {
	case LabelTypeMutable, LabelTypeImmutable:
		return true
	}
	return false
}

func (e LabelType) String() string {
	return string(e)
}

func (e *LabelType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LabelType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LABEL_TYPE", str)
	}
	return nil
}

func (e LabelType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
