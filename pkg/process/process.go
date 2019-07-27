/**
 * @Author: huangw1
 * @Date: 2019/7/25 16:34
 */

package process

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var ROOT string

var TemplateDir string

func init() {
	ROOT, _ = ExecutableDir()
}

func ExecutableDir() (string, error) {
	curFilename := os.Args[0]
	binaryPath, err := exec.LookPath(curFilename)
	if err != nil {
		panic(err)
	}
	pathAbs, err := filepath.Abs(binaryPath)
	if err != nil {
		panic(err)
	}
	return path.Dir(pathAbs), nil
}
