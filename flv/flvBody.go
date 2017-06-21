package flv

type flvBody struct {
	firstTagSize     uint32
	tag              []FlvTag
	PreviousTagSizes []uint32
}
