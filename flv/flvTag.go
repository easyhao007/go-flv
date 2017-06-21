package flv

type FlvTag struct {
	Header FlvTagHeader
	Body   FlvTagBody
}

//flv tags
type FlvTagHeader struct {
	//Reserved					//2bit - Reserved for FMS, should be 0
	Filter            uint8  //1bit
	TagType           uint8  //5bit	- 8 = audio , 9 = video , 18 = script data
	DataSize          uint32 //24bit
	Timestamp         uint32 //24bit
	TimestampExtended uint8  //8bit
	StreamID          uint32 //24bit
}

type FlvTagBody struct {
	Audio AudioTag
	Video VideoTag
	//matedate
}
