package main

import (
	"fmt"
	"github.com/shaalx/sstruct/service/log"
	"io/ioutil"
	"os"
)

var (
	files = [][]string{
		{
			"C34-Economy0002.txt",
			"C34-Economy0003.txt",
			"C34-Economy0008.txt",
			"C34-Economy0012.txt",
			"C34-Economy0010.txt",
			"C34-Economy0004.txt",
			// "C34-Economy0006.txt",
			"C34-Economy0014.txt",
			"C34-Economy0020.txt",
			"C34-Economy0022.txt",
			"C34-Economy0026.txt",
		},
		{
			"C32-Agriculture0002.txt",
			"C32-Agriculture0005.txt",
			"C32-Agriculture0007.txt",
			// "C32-Agriculture0011.txt",
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
			"C31-Enviornment0003.txt", //0
			"C31-Enviornment0007.txt",
			"C31-Enviornment0009.txt",
			"C31-Enviornment0011.txt",
			"C31-Enviornment0013.txt",
			"C31-Enviornment0019.txt",
			"C31-Enviornment0025.txt",
			"C31-Enviornment0029.txt",
			"C31-Enviornment0035.txt",
			"C31-Enviornment0039.txt", //0
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
			// "C38-Politics0039.txt",
			"C38-Politics0041.txt",
			// "C38-Politics0059.txt",
			// "C38-Politics0067.txt",
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
			// "C39-Sports0001.txt",
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
			// "C39-Sports0055.txt",
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
			// "C19-Computer0010.txt",
			"C19-Computer0012.txt",
			"C19-Computer0018.txt",
			"C19-Computer0020.txt",
			"C19-Computer0036.txt",
			"C19-Computer0038.txt",
			"C19-Computer0044.txt",
			"C19-Computer0048.txt",
			// "C19-Computer0058.txt",
			"C19-Computer0062.txt",
		},
	}
	items = []string{"economy", "agriculture", "art", "history", "environment", "politics", "space", "sports", "philosophy", "computer"}
)

func main() {
	v2()
}

func v2() {
	leng1 := len(files)
	var count int32
	count = int32(0)
	content := ""
	for i := 0; i < leng1; i++ {
		leng2 := len(files[i])
		// stat this item
		for ii := 0; ii < leng2; ii++ {
			c := FileWrodsCount("./origin/" + items[i] + "/" + files[i][ii])
			count += c
			content += fmt.Sprintf("%s :%v\n", files[i][ii], c)
			fmt.Printf("%s :%v\n", files[i][ii], c)
		}
	}

	content += fmt.Sprintf("total:%v\n avg:%v\n", count, count/100)
	WriteFile("avg_words_stats.txt", content)

}

func FileWrodsCount(filename string) int32 {
	// ioutil.ReadFile(filename)
	b, _ := ioutil.ReadFile(filename)
	return int32(len(b) / 3)
}

// 追加
func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE, os.ModeDevice)
	if log.IsError("{Write to file error}", err) {
		return
	}
	file.WriteString(content + "\n")
	defer file.Close()
}
