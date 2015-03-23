package bean

import (
	"fmt"
	"github.com/shaalx/sstruct/utils"
	. "github.com/shaalx/sstruct/vars"
	"math"
	"sort"
	"strings"
)

type TopicMatix []TopicSlice

var filter []string = []string{
	",", "的", "在", "和", "了", "也", "上", "还", "是", "年", "有", "都", "而", "我", "这个", "这么", "将", "一个", "家", "最", "从", "能", "就", "不", "而是", "就是", "该", "中", "", "", "", "", "", "", "", "", "",
}

func IsFilterContains(str string) bool {
	filters := strings.Join(filter, ",")
	return strings.Contains(filters, ","+str+",")
}

// 统计结果
func (t *TopicMatix) Statistics() {
	// 统计词频，映射到map
	key_freq := make(map[string]int32, 1)
	for _, key_slice := range *t {
		for _, key := range key_slice {
			if IsFilterContains(key.Const) {
				continue
			}
			key_freq[key.Const]++
		}
	}
	// 按照词频排序
	cellSlice := make(CellSlice, len(key_freq))
	i := 0
	var sum_freq int32
	for key_word, freq := range key_freq {
		cell := Cell{key_word, freq}
		cellSlice[i] = &cell
		i++
		sum_freq += freq
	}
	sort.Sort(cellSlice)
	cellSlice.String()

	// 卡方检验
	// nq 期望值为 sum_freq/len(key_freq)
	nq := float64(sum_freq) / float64(len(key_freq))
	var score float64
	// 组合关键字
	sentences := make(Sens, len(*t))
	// 后续处理
	var minFreq int32
	for i, key_slice := range *t {
		score = 0.0
		minFreq = 1000
		for _, key_word := range key_slice {
			fr, ok := key_freq[key_word.Const]
			if minFreq > fr {
				minFreq = fr
			}
			if ok {
				score += math.Pow(float64(fr)-nq, 2.0) / nq
				// score += float32(float64(math.Pow(float64(fr-int32(len(it)*2)), 2.0)) / float64(len(it)*2))
				// score += float32(float64(math.Pow(float64(fr-int32(len(it))), 2.0)) / float64(len(it)))
			}
		}
		if minFreq < 2 {
			score *= 0.3
		}
		// 降低单个词的卡方值
		score *= math.Log2(float64(len(key_slice) + 1))
		sen := Sen{Str: key_slice.WordStrings(), Sum: score, Avg: score}
		sentences[i] = &sen
	}
	// 排除重复
	sentences = *sentences.EjRepeat()
	// 按照卡方值排序
	sort.Sort(sentences)

	topN := len(key_freq)/500 + 10
	topFormatString, topNSlice := sentences.Top(topN)
	fmt.Print(topFormatString)
	statResultString := PreciseAndRecall(topNSlice)
	SaveResult(statResultString)
}

// 计算实验结果数据
func PreciseAndRecall(tops []string) string {
	resultStr := ""
	// key_words := []string{".", "扫脸", "刷脸", "刷脸支付", "支付技术", "扫脸支付", "", "."}
	// key_words := []string{".", "流通", "流通理论", "流通经济学", "经济学", "."}
	key_word, err := utils.ReadKey(KEY_DIR + CURRENT_FILENAME)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	key_words := strings.Split(key_word, "|")
	key_words_len := len(key_words) - 2
	tops = tops[:key_words_len+1]
	count := 0.0
	for _, it := range tops {
		if strings.Contains(key_word, "|"+it+"|") {
			count += 1.0
		}
	}
	pricise := count / float64(key_words_len+1)
	recall := count / float64(key_words_len)
	b := 1.0
	f_measrue := (b*b + 1) * pricise * recall / (b*b*pricise + recall)

	priciseString := fmt.Sprintf("精确率：\t%.4f\n", pricise) // len(key_words)-2
	recallString := fmt.Sprintf("召回率：\t%.4f\n", recall)   // len(tops)
	f_measrueString := fmt.Sprintf("F-measure：\t%.4f\n", f_measrue)

	resultStr += fmt.Sprintf("Top key words:\t%v\n", tops)
	resultStr += fmt.Sprintf("Key words:\t%v\n", key_words)
	resultStr += priciseString + recallString + f_measrueString
	return resultStr
}

func SaveResult(content string) {
	// 保存结果
	stringSaveChan := make(chan string, 5)
	go utils.SaveString(stringSaveChan, STAT_DIR+CURRENT_FILENAME)
	// fmt.Printf("\n%s\t %s\t %s\t %s\n", "index", "avg", "key-word", "freq")
	stringSaveChan <- content
	fmt.Println(content)
}

type Sen struct {
	Str string
	Sum float64
	Avg float64
	Fre int
}

func (s *Sen) String() string {
	// return fmt.Sprintf("%.3f\t %.3f\t %s\t", s.Avg, s.Sum, s.Str)
	return fmt.Sprintf("%.2f\t %s\t %d", s.Avg, s.Str, s.Fre)
}

type Sens []*Sen

func (c Sens) Len() int {
	return len(c)
}

func (c Sens) Less(i, j int) bool {
	// return c[i].Avg > c[j].Avg // 不可靠
	return float64(c[i].Fre)*c[i].Avg > float64(c[j].Fre)*c[j].Avg
	// return math.Log2(float64(c[i].Fre))*c[i].Avg > math.Log2(float64(c[j].Fre))*c[j].Avg
}

func (c Sens) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (s Sens) Top(n int) (string, []string) {
	resStr := ""
	topNStr := make([]string, n)
	for i, it := range s {
		if i >= n {
			return resStr, topNStr[:n]
		}
		resStr += fmt.Sprintf("%d\t %s\n", i+1, it)
		topNStr[i] = it.Str
	}
	return resStr, topNStr
}

// 排除重复值
func (t *Sens) EjRepeat() *Sens {
	length := len(*t)
	ejMap := make(map[string]*Sen, length)
	ejMapCount := make(map[string]int, length)
	for _, it := range *t {
		// _, ok := ejMap[it.Str]
		// if ok {
		// 	continue
		// } else {
		ejMap[it.Str] = it
		ejMapCount[it.Str]++
		// }
	}
	i := 0
	result := make(Sens, len(ejMap))
	for _, v := range ejMap {
		v.Fre = ejMapCount[v.Str]
		result[i] = v
		i++
	}
	return &result
}
