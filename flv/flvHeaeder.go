package flv

import (
	"errors"
	"flvParser/bitbuffer"
	"fmt"
	"encoding/hex"
)

type FlvHeader struct {
	SignatureF         uint8  //1byte	//F
	SignatureL         uint8  //1byte	//L
	SignatureV         uint8  //1byte	//V
	Version            uint8  //1byte	//版本号
	TypeFlagsReserved5 uint8  //5bit		//保留的，一般为0
	TypeFlagsAudio     uint8  //1bit		//音频标签标识
	TypeFlagsReserved1 uint8  //1bit		//保留的，一般为0
	TypeFlagsVideo     uint8  //1bit		//视频标签表示
	DataOffset         uint32 //4byte	//flv头的长度 version1 一般为9
}

func (header *FlvHeader) Parse(buf []byte) (err error) {
	fmt.Println(hex.Dump(buf))
	bitBuf := new(bitbuffer.BitBuffer)
	bitBuf.Set(buf)

	header.SignatureF, err = bitBuf.PeekUint8(8)
	if err != nil {
		return err
	}

	header.SignatureL, err = bitBuf.PeekUint8(8)
	if err != nil {
		return err
	}

	header.SignatureV, err = bitBuf.PeekUint8(8)
	if err != nil {
		return err
	}

	if header.SignatureF != 'F' || header.SignatureL != 'L' || header.SignatureV != 'V' {
		err = errors.New("ERROR, 没有以FLV开始，不是标准的flv文件，解析失败")
		return err
	}

	header.Version, err = bitBuf.PeekUint8(8)
	if err != nil {
		return err
	}

	err = bitBuf.Skip(5)
	if err != nil {
		return err
	}

	header.TypeFlagsAudio, err = bitBuf.PeekUint8(1)
	if err != nil {
		return err
	}

	err = bitBuf.Skip(1)
	if err != nil {
		return err
	}

	header.TypeFlagsVideo, err = bitBuf.PeekUint8(1)
	if err != nil {
		return err
	}

	header.DataOffset, err = bitBuf.PeekUint32(32)
	if err != nil {
		return err
	}

	return nil
}

func (header *FlvHeader) Dump() {
	fmt.Printf("==========================FLV Header==========================\n")
	fmt.Printf("SignatureF:		%c\n", header.SignatureF)
	fmt.Printf("SignatureL:		%c\n", header.SignatureL)
	fmt.Printf("SignatureV:		%c\n", header.SignatureV)
	fmt.Printf("Version:			0x%x\n", header.Version)
	fmt.Printf("TypeFlagsAudio:	0x%x\n", header.TypeFlagsAudio)
	fmt.Printf("TypeFlagsVideo:	0x%x\n", header.TypeFlagsVideo)
	fmt.Printf("DataOffset:		0x%x\n", header.DataOffset)
	fmt.Printf("==========================FLV Body============================\n")
}
