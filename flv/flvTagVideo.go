package flv

//video tag
type VideoTag struct {
	Header VideoTagHeader
	Body   VideoTagBody
}

type VideoTagHeader struct {
	FrameType       uint8  //4bit
	CodecID         uint8  //4bit
	AVCPacketType   uint8  //8bit
	CompositionTime uint32 //24
}

type VideoTagBody struct {
	Data []byte
}
