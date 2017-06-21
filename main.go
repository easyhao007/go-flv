package main

import (
	"fmt"
	"flvParser/flv"
	"flag"
	"io/ioutil"
)

func main() {
	filename := flag.String("f" , "" , "被解析的文件的路径")
	flag.Parse()

	if *filename == ""{
		fmt.Println("usage: flvParse file path")
		return
	}

	buf , err := ioutil.ReadFile(*filename)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	flvHeader := new(flv.FlvHeader)
	err = flvHeader.Parse(buf[0:9])
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	flvHeader.Dump()

	//解析头
}
