package comm

import "errors"

func CheckParameter(paraMap map[string]string) error {
	platform, ok := paraMap["p"]
	if !ok || (platform != "cli" && platform != "svr") {
		return errors.New("can not find -p parameter")
	}
	if platform == "cli" {
		pkg, pkgOK := paraMap["pkg"]
		sub, subOK := paraMap["sub"]
		step, stepOK := paraMap["step"]
		if !pkgOK || len(pkg) == 0 {
			return errors.New("generate for client code, can not find pkg parameter")
		}
		if !subOK || len(sub) == 0 {
			return errors.New("generate for client code, can not find sub parameter")
		}
		if !stepOK || len(step) == 0 {
			return errors.New("generate for client code, can not find step parameter")
		}
	}
	return nil
}
