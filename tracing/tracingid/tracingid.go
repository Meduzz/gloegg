package tracingid

import (
	"encoding/base64"
	"encoding/json"

	"github.com/Meduzz/helper/hashing"
)

/*
	Package tracingid contains formating of what data a tracingid is made up of
	and logic on how to build & manipulate it.

	Tracingids are transported in a jwt inspired fashion, ie. serialized to
	json and then base64 encoded.

	An ex. of a tracingid for method4 that depends on method3 -> method2 -> method1
	would look like this:

	names [method1, method2, method3]
	ids: [sha1, sha1, sha1]
	id: unique sha1-hash
	name: method4
*/

type (
	TraceId struct {
		Names []string `json:"names"`
		IDS   []string `json:"ids"`
		ID    string   `json:"id"`
		Name  string   `json:"name"`
	}
)

func NewTracingID(name string) *TraceId {
	id := hashing.Token()
	empty := make([]string, 0)

	return &TraceId{empty, empty, id, name}
}

func NewTracingIDFromParent(parent, name string) (*TraceId, error) {
	p, err := FromString(parent)

	if err != nil {
		return nil, err
	}

	id := hashing.Token()

	return &TraceId{p.BuildNames(), p.BuildIDS(), id, name}, nil
}

func FromString(it string) (*TraceId, error) {
	bs, err := base64.RawStdEncoding.DecodeString(it)

	if err != nil {
		return nil, err
	}

	trace := &TraceId{}
	err = json.Unmarshal(bs, trace)

	if err != nil {
		return nil, err
	}

	return trace, nil
}

func ToString(trace *TraceId) (string, error) {
	bs, err := json.Marshal(trace)

	if err != nil {
		return "", err
	}

	it := base64.RawStdEncoding.EncodeToString(bs)

	return it, nil
}

func (t *TraceId) BuildIDS() []string {
	return append(t.IDS, t.ID)
}

func (t *TraceId) BuildNames() []string {
	return append(t.Names, t.Name)
}
