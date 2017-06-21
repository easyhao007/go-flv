package flv

//audio tags
type AudioTag struct {
	Header AudioTagHeader
	Body   AudioTagBody
}

type AudioTagHeader struct {
	SoundFormat   uint8 //4bit
	SoundRate     uint8 //2bit
	SoundSize     uint8 //1bit
	SoundType     uint8 //1bit
	AACPacketType uint8 //8bit if SoundFormat== 10  defined
}

type AudioTagBody struct {
	Data []byte
}
