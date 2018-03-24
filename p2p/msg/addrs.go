package msg

import (
	"bytes"
	"encoding/binary"
)

type Addrs struct {
	Count     uint64
	PeerAddrs []PeerAddr
}

func NewAddrs(addrs []PeerAddr) *Addrs {
	msg := new(Addrs)
	msg.Count = uint64(len(addrs))
	msg.PeerAddrs = addrs
	return msg
}

func (addrs *Addrs) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, addrs.Count)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, addrs.PeerAddrs)
	if err != nil {
		return nil, err
	}

	return BuildMessage("addr", buf.Bytes())
}

func (addrs *Addrs) Deserialize(msg []byte) error {
	buf := bytes.NewReader(msg)
	err := binary.Read(buf, binary.LittleEndian, &addrs.Count)
	if err != nil {
		return err
	}

	addrs.PeerAddrs = make([]PeerAddr, addrs.Count)
	err = binary.Read(buf, binary.LittleEndian, &addrs.PeerAddrs)
	if err != nil {
		return err
	}

	return nil
}
