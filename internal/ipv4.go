package internal

import (
	"errors"
	"regexp"
	"strconv"
)

type IPv4 struct {
	Address uint32
	Mask    uint32
}

func ParseString(s string) (*IPv4, error) {
	reg := regexp.MustCompile(`(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})\/(\d{1,2})`)
	parts := reg.FindStringSubmatch(s)

	if parts == nil {
		return nil, errors.New("Invalid ip format")
	}

	address, err := parseIpAddress(parts[1:5])

	if err != nil {
		return nil, err
	}

	mask, err := parseIpMask(parts[5])

	if err != nil {
		return nil, err
	}

	return &IPv4{address, mask}, nil
}

func parseIpAddress(values []string) (uint32, error) {
	var r uint32 = 0

	for i := 0; i < len(values); i++ {
		u, err := strconv.ParseUint(values[i], 10, 8)
		if err != nil {
			return 0, errors.New("Invalid ip number format")
		}
		r = (r | uint32(u))

		if i < 3 {
			r <<= 8
		}
	}

	return r, nil
}

func parseIpMask(mask string) (uint32, error) {
	n, err := strconv.ParseUint(mask, 10, 32)
	if err != nil || n > 32 {
		return 0, errors.New("Invalid ip mask format")
	}
	shift := 32 - n
	return (^uint32(0) >> shift) << shift, nil
}

func (ip *IPv4) IsSubnetOf(other *IPv4) bool {
	if ip.Mask <= other.Mask {
		return false
	}

	ipNetwork := ip.Address & other.Mask
	otherNetwork := other.Address & other.Mask

	return ipNetwork == otherNetwork
}

func (ip *IPv4) IsOnSameNetworkAs(other *IPv4) bool {
	if ip.Mask != other.Mask {
		return false
	}

	return (ip.Mask & ip.Address) == (other.Mask & other.Address)
}
