package structtag_interface

import (
	"encoding/json"
	"github.com/tomoyamachi/logmask-patterns/util"
	"testing"
)

type MaskResponse struct {
	Str               string            `json:"str"`
	MaskStr           string            `json:"mstr" log:"-"`
	Int               int               `json:"int"`
	MaskInt           int               `json:"mint" log:"-"`
	Slice             []string          `json:"slice""`
	MaskSlice         []string          `json:"mslice" log:"-"`
	Map               map[string]string `json:"map"`
	MaskMap           map[string]string `json:"mmap" log:"-"`
	Struct            SubMask           `json:"struct"`
	MaskStruct        SubMask           `json:"mstruct" log:"-"`
	PointerStruct     *SubMask          `json:"pstruct"`
	MaskPointerStruct *SubMask          `json:"mpstruct" log:"-"`
}

type SubMask struct {
	Str     string `json:"str"`
	MaskStr string `json:"mstr" log:"-"`
}

func initMask() MaskResponse {
	return MaskResponse{
		Str:       "a",
		MaskStr:   "a",
		Int:       100,
		MaskInt:   100,
		Slice:     []string{"a"},
		MaskSlice: []string{"a"},
		Map:       map[string]string{"a": "b"},
		MaskMap:   map[string]string{"a": "b"},
		Struct: SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		MaskStruct: SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		PointerStruct: &SubMask{
			Str:     "a",
			MaskStr: "a",
		},
		MaskPointerStruct: &SubMask{
			Str:     "a",
			MaskStr: "a",
		},
	}
}

func initMaskPtr() *MaskResponse {
	m := initMask()
	return &m
}

func TestMask(t *testing.T) {
	tests := []struct {
		in         interface{}
		ok         bool
		expectLog  string
		expectJson string
	}{
		{
			in: initMask(),
			ok: true,
			expectLog: `{
  "str": "a",
  "mstr": "*",
  "int": 100,
  "mint": "*",
  "slice": [
    "a"
  ],
  "mslice": "*",
  "map": {
    "a": "b"
  },
  "mmap": "*",
  "struct": {
    "mstr": "*",
    "str": "a"
  },
  "mstruct": "*",
  "pstruct": {
    "mstr": "*",
    "str": "a"
  },
  "mpstruct": "*"
}`,
			expectJson: `{
  "str": "a",
  "mstr": "a",
  "int": 100,
  "mint": 100,
  "slice": ["a"],
  "mslice": ["a"],
  "map": {"a": "b"},
  "mmap": {"a": "b"},
  "struct": {
    "mstr": "a",
    "str": "a"
  },
  "mstruct": {
    "mstr": "a",
    "str": "a"
  },
  "pstruct": {
    "mstr": "a",
    "str": "a"
  },
  "mpstruct": {
    "mstr": "a",
    "str": "a"
  }
}`,
		},
		{
			in: initMaskPtr(),
			ok: true,
			expectLog: `{
  "str": "a",
  "mstr": "*",
  "int": 100,
  "mint": "*",
  "slice": [
    "a"
  ],
  "mslice": "*",
  "map": {
    "a": "b"
  },
  "mmap": "*",
  "struct": {
    "mstr": "*",
    "str": "a"
  },
  "mstruct": "*",
  "pstruct": {
    "mstr": "*",
    "str": "a"
  },
  "mpstruct": "*"
}`,
			expectJson: `{
  "str": "a",
  "mstr": "a",
  "int": 100,
  "mint": 100,
  "slice": ["a"],
  "mslice": ["a"],
  "map": {"a": "b"},
  "mmap": {"a": "b"},
  "struct": {
    "mstr": "a",
    "str": "a"
  },
  "mstruct": {
    "mstr": "a",
    "str": "a"
  },
  "pstruct": {
    "mstr": "a",
    "str": "a"
  },
  "mpstruct": {
    "mstr": "a",
    "str": "a"
  }
}`,
		},
	}
	for i, tt := range tests {
		b, err := Log(tt.in)
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				t.Errorf("test %d, unexpected failure: %v", i, err)
			} else {
				t.Errorf("test %d, unexpected success", i)
			}
			continue
		}

		// check log output
		{
			ok, err := util.CompareJsonBytes(b, []byte(tt.expectLog))
			if err != nil {
				t.Errorf("test %d, unexpected error with compare log output", i)
			}
			if !ok {
				t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(b), tt.expectLog)
			}
		}

		// check json marshal
		got, err := json.Marshal(tt.in)
		if err != nil {
			t.Errorf("test %d, failed json.Marshal", i)
		}
		{
			ok, err := util.CompareJsonBytes(got, []byte(tt.expectJson))
			if err != nil {
				t.Errorf("test %d, unexpected error with compare log output", i)
			}
			if !ok {
				t.Errorf("test %d, Marshal(%#v) = %s, want %s", i, tt.in, string(b), tt.expectLog)
			}
		}

	}
}