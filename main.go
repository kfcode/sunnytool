
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/kfcode/sunnytool/comm"
	"github.com/kfcode/sunnytool/generator"
)

func main() {
	g := generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("reading input err:%v", err)
	}
	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}
	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}
	g.CommandLineParameters(g.Request.GetParameter())
	err = comm.CheckParameter(g.Param)
	if err != nil {
		g.Fail(err.Error())
	}
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	switch g.Param["p"] {
	case "cli":
			if g.Param["step"] == "cli" {
				g.GenerateCliAllFiles()
			}else {
				g.GenerateCliGenAllFiles()
			}
	case "svr":
			if g.Param["step"] == "svr" {
				g.GenerateSvrAllFiles()
			}else {
				g.GenerateSvrGenGenAllFiles()
			}
	default:
		g.Fail("can not recognize -p parameter")
	}
	return
}
