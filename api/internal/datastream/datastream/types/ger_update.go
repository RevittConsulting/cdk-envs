package types

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ledgerwatch/erigon-lib/common"
)

const (
	gerUpdateDataLength         = 106
	gerUpdateDataLengthPreEtrog = 102

	// EntryTypeL2Block represents a L2 block
	EntryTypeGerUpdate EntryType = 4
)

type GerUpdate struct {
	BatchNumber    uint64         // 8 bytes
	Timestamp      uint64         // 8 bytes
	GlobalExitRoot common.Hash    // 32 bytes
	Coinbase       common.Address // 20 bytes
	ForkId         uint16         // 2 bytes
	ChainId        uint32         // 4 bytes
	StateRoot      common.Hash    // 32 bytes
}

func (g *GerUpdate) EncodeToBytes() []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, g.BatchNumber); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.LittleEndian, g.Timestamp); err != nil {
		return nil
	}
	buf.Write(g.GlobalExitRoot.Bytes())
	buf.Write(g.Coinbase.Bytes())
	if err := binary.Write(buf, binary.LittleEndian, g.ForkId); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.LittleEndian, g.ChainId); err != nil {
		return nil
	}
	buf.Write(g.StateRoot.Bytes())

	return buf.Bytes()
}

func (g *GerUpdate) EncodeToBytesBigEndian() []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, g.BatchNumber); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, g.Timestamp); err != nil {
		return nil
	}
	buf.Write(g.GlobalExitRoot.Bytes())
	buf.Write(g.Coinbase.Bytes())
	if err := binary.Write(buf, binary.BigEndian, g.ForkId); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, g.ChainId); err != nil {
		return nil
	}
	buf.Write(g.StateRoot.Bytes())

	return buf.Bytes()
}

// decodes a StartL2Block from a byte array
func DecodeGerUpdate(data []byte) (*GerUpdate, error) {
	if len(data) != gerUpdateDataLength {
		if len(data) == gerUpdateDataLengthPreEtrog {
			return decodeGerUpdatePreEtrog(data)
		}
		return &GerUpdate{}, fmt.Errorf("expected data length: %d, got: %d", gerUpdateDataLength, len(data))
	}

	var ts uint64
	buf := bytes.NewBuffer(data[8:16])
	if err := binary.Read(buf, binary.LittleEndian, &ts); err != nil {
		return &GerUpdate{}, err
	}

	return &GerUpdate{
		BatchNumber:    binary.LittleEndian.Uint64(data[:8]),
		Timestamp:      ts,
		GlobalExitRoot: common.BytesToHash(data[16:48]),
		Coinbase:       common.BytesToAddress(data[48:68]),
		ForkId:         binary.LittleEndian.Uint16(data[68:70]),
		ChainId:        binary.LittleEndian.Uint32(data[70:74]),
		StateRoot:      common.BytesToHash(data[74:106]),
	}, nil
}

// decodes a StartL2Block from a byte array
func decodeGerUpdatePreEtrog(data []byte) (*GerUpdate, error) {
	var ts uint64
	buf := bytes.NewBuffer(data[8:16])
	if err := binary.Read(buf, binary.LittleEndian, &ts); err != nil {
		return &GerUpdate{}, err
	}

	return &GerUpdate{
		BatchNumber:    binary.LittleEndian.Uint64(data[:8]),
		Timestamp:      ts,
		GlobalExitRoot: common.BytesToHash(data[16:48]),
		Coinbase:       common.BytesToAddress(data[48:68]),
		ForkId:         binary.LittleEndian.Uint16(data[68:70]),
		StateRoot:      common.BytesToHash(data[70:102]),
	}, nil
}

// decodes a StartL2Block from a byte array
func DecodeGerUpdateBigEndian(data []byte) (*GerUpdate, error) {
	if len(data) != gerUpdateDataLength {
		return &GerUpdate{}, fmt.Errorf("expected data length: %d, got: %d", gerUpdateDataLength, len(data))
	}

	var ts uint64
	buf := bytes.NewBuffer(data[8:16])
	if err := binary.Read(buf, binary.BigEndian, &ts); err != nil {
		return &GerUpdate{}, err
	}

	return &GerUpdate{
		BatchNumber:    binary.BigEndian.Uint64(data[:8]),
		Timestamp:      ts,
		GlobalExitRoot: common.BytesToHash(data[16:48]),
		Coinbase:       common.BytesToAddress(data[48:68]),
		ForkId:         binary.BigEndian.Uint16(data[68:70]),
		ChainId:        binary.LittleEndian.Uint32(data[70:74]),
		StateRoot:      common.BytesToHash(data[74:106]),
	}, nil
}
