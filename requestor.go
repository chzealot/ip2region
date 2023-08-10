package ip2region

import (
	"encoding/binary"
	"errors"
	"github.com/chzealot/ip2region/assets"
	"github.com/zu1k/nali/pkg/qqwry"
	"github.com/zu1k/nali/pkg/wry"
	"net"
)

type Requestor struct {
	dataFile *qqwry.QQwry
}

func NewRequestor() (*Requestor, error) {
	if !qqwry.CheckFile(assets.QQWryContent) {
		return nil, errors.New("illegal qqwry.dat")
	}
	header := assets.QQWryContent[0:8]
	start := binary.LittleEndian.Uint32(header[:4])
	end := binary.LittleEndian.Uint32(header[4:])
	dataFile := &qqwry.QQwry{
		IPDB: wry.IPDB[uint32]{
			Data: assets.QQWryContent,

			OffLen:   3,
			IPLen:    4,
			IPCnt:    (end-start)/7 + 1,
			IdxStart: start,
			IdxEnd:   end,
		},
	}
	return &Requestor{
		dataFile: dataFile,
	}, nil
}

func (r *Requestor) Query(ip string) (*Location, error) {
	ipObj := net.ParseIP(ip)
	if ipObj == nil {
		return nil, errors.New("query should be IPv4")
	}
	ip4 := ipObj.To4()
	if ip4 == nil {
		return nil, errors.New("query should be IPv4")
	}
	ip4uint := binary.BigEndian.Uint32(ip4)

	offset := (*r.dataFile).SearchIndexV4(ip4uint)
	if offset <= 0 {
		return nil, errors.New("query not valid")
	}

	reader := wry.NewReader((*r.dataFile).Data)
	reader.Parse(offset + 4)

	loc := NewLocationFromCnFull(reader.Result.DecodeGBK().String())
	loc.City = reader.Result.Country
	loc.Area = reader.Result.Area
	return loc, nil
}
