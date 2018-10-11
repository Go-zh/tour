// Go语言的作者2016年版权所有
// 此源代码的使用受BSD风格的约束
// 可以在LICENSE文件中找到的许可证

package content

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// 测试内容文件中的所有.go文件是否构建并执行（不检查输出正确性）
// 不构建包含字符串 "// +build no-build" 的文件
// 不执行包含字符串 "// +build no-run" 的文件
func TestContent(t *testing.T) {
	scratch, err := ioutil.TempDir("", "tour-content-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(scratch)

	err = filepath.Walk(".", func(path string, fi os.FileInfo, err error) error {
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if filepath.Base(path) == "content_test.go" {
			return nil
		}
		if err := testSnippet(t, path, scratch); err != nil {
			t.Errorf("%v: %v", path, err)
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func testSnippet(t *testing.T, path, scratch string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	build := string(bytes.SplitN(b, []byte{'\n'}, 2)[0])
	if !strings.HasPrefix(build, "// +build ") {
		return errors.New("first line is not a +build comment")
	}
	if !strings.Contains(build, "OMIT") {
		return errors.New(`+build comment does not contain "OMIT"`)
	}

	if strings.Contains(build, "no-build") {
		return nil
	}
	bin := filepath.Join(scratch, filepath.Base(path)+".exe")
	out, err := exec.Command("go", "build", "-o", bin, path).CombinedOutput()
	if err != nil {
		return fmt.Errorf("build error: %v\noutput:\n%s", err, out)
	}
	defer os.Remove(bin)

	if strings.Contains(build, "no-run") {
		return nil
	}
	out, err = exec.Command(bin).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v\nOutput:\n%s", err, out)
	}
	return nil
}
