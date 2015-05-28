package main

import (
	"bufio"
	"fmt"
	"github.com/shaalx/sstruct/service/log"
	"os"
	"strings"
)

var (
	files = [][]string{
		{
			"C34-Economy0003.txt",
			"C34-Economy0008.txt",
			"C34-Economy0012.txt",
			"C34-Economy0010.txt",
			"C34-Economy0004.txt",
			"C34-Economy0006.txt",
			"C34-Economy0014.txt",
			"C34-Economy0016.txt",
			"C34-Economy0018.txt",
			"C34-Economy0020.txt",
			"C34-Economy0022.txt",
			"C34-Economy0024.txt",
			"C34-Economy0026.txt",
			"C34-Economy0002.txt",
		},
		{
			"C19-Computer0001.txt",
			"C19-Computer0004.txt",
			"C19-Computer0012.txt",
			"C19-Computer0018.txt",
			"C19-Computer0020.txt",
			"C19-Computer0036.txt",
			"C19-Computer0038.txt",
			"C19-Computer0044.txt",
			"C19-Computer0048.txt",
			"C19-Computer0058.txt",
			"C19-Computer0062.txt",
		},
		{
			"C32-Agriculture0002.txt",
			"C32-Agriculture0005.txt",
			"C32-Agriculture0007.txt",
			"C32-Agriculture0011.txt",
			"C32-Agriculture0015.txt",
			"C32-Agriculture0017.txt",
			"C32-Agriculture0019.txt",
			"C32-Agriculture0023.txt",
			"C32-Agriculture0027.txt",
			"C32-Agriculture0031.txt",
			"C32-Agriculture0035.txt",
		},
		{
			"C3-Art0003.txt",
			"C3-Art0005.txt",
			"C3-Art0007.txt",
			"C3-Art0011.txt",
			"C3-Art0017.txt",
			"C3-Art0021.txt",
			"C3-Art0023.txt",
			"C3-Art0029.txt",
			"C3-Art0035.txt",
			"C3-Art0037.txt",
		},
		{
			"C7-History004.txt",
			"C7-History010.txt",
			"C7-History014.txt",
			"C7-History018.txt",
			"C7-History021.txt",
			"C7-History029.txt",
			"C7-History037.txt",
			"C7-History039.txt",
			"C7-History041.txt",
			"C7-History045.txt",
		},
		{
			"C31-Enviornment0003.txt",
			"C31-Enviornment0007.txt",
			"C31-Enviornment0009.txt",
			"C31-Enviornment0011.txt",
			"C31-Enviornment0013.txt",
			"C31-Enviornment0019.txt",
			"C31-Enviornment0025.txt",
			"C31-Enviornment0029.txt",
			"C31-Enviornment0035.txt",
			"C31-Enviornment0039.txt",
		},
		{
			"C38-Politics0002.txt",
			"C38-Politics0004.txt",
			"C38-Politics0008.txt",
			"C38-Politics0010.txt",
			"C38-Politics0012.txt",
			"C38-Politics0019.txt",
			"C38-Politics0021.txt",
			"C38-Politics0029.txt",
			"C38-Politics0033.txt",
			"C38-Politics0039.txt",
			"C38-Politics0041.txt",
			"C38-Politics0059.txt",
			"C38-Politics0067.txt",
		},
		{
			"C11-Space0003.txt",
			"C11-Space0007.txt",
			"C11-Space0009.txt",
			"C11-Space0015.txt",
			"C11-Space0017.txt",
			"C11-Space0025.txt",
			"C11-Space0029.txt",
			"C11-Space0037.txt",
			"C11-Space0043.txt",
			"C11-Space0051.txt",
		},
		{
			"C39-Sports0001.txt",
			"C39-Sports0003.txt",
			"C39-Sports0005.txt",
			"C39-Sports0009.txt",
			"C39-Sports0013.txt",
			"C39-Sports0019.txt",
			"C39-Sports0023.txt",
			"C39-Sports0027.txt",
			"C39-Sports0029.txt",
			"C39-Sports0041.txt",
			"C39-Sports0047.txt",
			"C39-Sports0055.txt",
		},
		{
			"C6-Philosophy22.txt",
			"C6-Philosophy32.txt",
			"C6-Philosophy39.txt",
			"C6-Philosophy47.txt",
			"C6-Philosophy08.txt",
			"C6-Philosophy14.txt",
			"C6-Philosophy43.txt",
			"C6-Philosophy71.txt",
			"C6-Philosophy80.txt",
			"C6-Philosophy86.txt",
		},
		{
			"C19-Computer0001.txt",
			"C19-Computer0004.txt",
			"C19-Computer0010.txt",
			"C19-Computer0012.txt",
			"C19-Computer0018.txt",
			"C19-Computer0020.txt",
			"C19-Computer0036.txt",
			"C19-Computer0038.txt",
			"C19-Computer0044.txt",
			"C19-Computer0048.txt",
			"C19-Computer0058.txt",
			"C19-Computer0062.txt",
		},
	}
	stats = []string{"total_", "precise_", "recall_", "f-measure_"}
)

func main() {
	v2()
}

func v2() {
	leng1 := len(files)
	for i := 0; i < leng1; i++ {
		leng2 := len(files[i])
		// stat this item
		result := make([][]string, 0, 20)
		for ii := 0; ii < leng2; ii++ {
			data := Read4Lines("./stat/" + files[i][ii])
			result = append(result, data)
		}
		length := len(result)
		for j := 0; j < 4; j++ {
			content := ""
			for jj := 0; jj < length; jj++ {
				content += result[jj][j]
				fmt.Println(result[jj][j])
			}
			content = stats[j] + "\n" + content
			WriteFile(fmt.Sprintf("%s", files[i][0]), content)
		}

	}

}

func v1() {
	d, _ := os.Open("./stat/")
	f, _ := d.Readdir(0)
	result := make([][]string, 0, 100)
	for _, it := range f {
		data := Read4Lines("./stat/" + it.Name())
		result = append(result, data)
		// fmt.Println(i, data)
	}
	fmt.Println(result)
	length := len(result)
	content := ""
	for j := 0; j < 4; j++ {
		content = ""
		for i := 0; i < length; i++ {
			content += result[i][j]
			fmt.Println(result[i][j])
		}
		WriteFile(fmt.Sprintf("stat%d.txt", j), content)
	}
}

func Read4Lines(filename string) []string {
	rfile, err := os.Open(filename)
	defer rfile.Close()
	if log.IsError("open file error", err) {
		return nil
	}
	reader := bufio.NewReader(rfile)
	ret := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		str, err := reader.ReadString('\n')
		if log.IsError("{read file error}", err) {
			return nil
		}
		ss := strings.Split(str, ":")
		ret = append(ret, ss[1])
	}
	return ret
}

// 追加
func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModeDevice)
	if log.IsError("{Write to file error}", err) {
		return
	}
	file.WriteString(content + "\n")
	defer file.Close()
}
