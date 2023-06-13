package nhashids

import (
	"errors"

	"github.com/speps/go-hashids/v2"
	"github.com/spf13/cast"
)

func Encode(data int64) (string, error) {

	hd := hashids.NewData()

	idx, salt := getSalt(data)

	hd.Salt = salt
	hd.MinLength = 12
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	id, err := h.EncodeInt64([]int64{data})
	if err != nil {
		return "", err
	}
	return idx + id, nil
}

func Decode(data string) (int64, error) {

	if len(data) < 1 {
		return 0, nil
	}

	idx := cast.ToInt64(data[0] - '0')
	_, salt := getSalt(idx)

	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 12
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return 0, err
	}
	id, err := h.DecodeInt64WithError(data[1:])
	if err != nil {
		return 0, err
	}

	if len(id) <= 0 {
		return 0, errors.New("not valid string")
	}

	return id[0], nil
}
