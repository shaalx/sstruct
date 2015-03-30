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
	",", "的", "在", "和", "了", "也", "上", "还", "是", "年", "有", "都", "而", "我", "这个", "这么", "将", "一个", "家", "最", "从", "能", "就", "不", "而是", "就是", "该", "中", "他", "时", "这个", "【", "】", "使", "只", "不", "不能", "没有", ",",
}

var posStop string

func init() {
	posStop = "|o|p|u|e|c|nh|r|q|m|"
}

func IsFilterContains(str string) bool {
	filters := strings.Join(filter, ",")
	return strings.Contains(filters, ","+str+",")
}

// 统计结果
func (t *TopicMatix) Statistics() {
	// 统计词频，映射到map
	key_freq := make(map[string]int32, 1)
	// key_freq_o := make(map[string]int32, 1)
	for _, key_slice := range *t {
		for _, key := range key_slice {
			if strings.Contains(posStop, "|"+key.Pos+"|") || IsFilterContains(key.Const) {
				// fmt.Println(key.Pos, key.Const)
				continue
			}
			key_freq[strings.TrimSpace(key.Const)]++
		}
		/*// 一元词
		if 1 == len(key_slice) {
			key := key_slice[0]
			if strings.Contains(posStop, "|"+key.Pos+"|") || IsFilterContains(key.Const) {
				// fmt.Println(key.Pos, key.Const)
				continue
			}
			key_freq_o[key.Const]++
		}*/
		// fmt.Println(key_slice)
	}
	// 按照词频排序
	cellSlice := make(CellSlice, len(key_freq))
	i := 0
	var sum_freq_o int32
	for key_word, freq := range key_freq {
		cell := Cell{key_word, freq}
		cellSlice[i] = &cell
		i++
		sum_freq_o += freq
	}
	sort.Sort(cellSlice)
	// cellSlice.String()

	// 卡方检验
	// nq 期望值为 sum_freq_o/len(key_freq)
	nq := float64(sum_freq_o) / float64(len(key_freq))
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
			}
			if strings.Contains(posStop, "|"+key_word.Pos+"|") || IsFilterContains(key_word.Const) {
				score *= 0.0
				break
			}
		}
		if minFreq < 2 {
			score *= 0.3
		}
		if len(key_slice) == 2 && key_slice[0].Const == key_slice[1].Const {
			score *= 0.2
		}
		// 降低单个词的卡方值
		score *= math.Log2(float64(len(key_slice) + 1))
		sen := Sen{Str: key_slice.WordStrings(), Sum: score, Avg: score}
		sentences[i] = &sen
	}
	// 排除重复
	sentences = *sentences.EjRepeat()
	// improve
	for _, sentence := range sentences {
		sentence.Sum = math.Log2(float64(sentence.Fre)) * sentence.Avg // 短文较差，长文还可以
	}
	// 按照卡方值排序
	sort.Sort(sentences)

	topN := /*len(key_freq)/500 + */ 20
	topFormatString, topNSlice := sentences.Top(topN)
	fmt.Print(topFormatString)
	statResultString := PreciseAndRecall(topNSlice)
	statResultString = "\n\n候选关键字Fi" + statResultString
	SaveResult(statResultString)
}

// 统计结果
func (t *TopicMatix) StatisticsWithOrigin(o *TopicMatix) {
	realteSet(*t)
	// 统计词频，映射到map
	key_freq := make(map[string]int32, 1)
	for _, key_slice := range *t {
		for _, key := range key_slice {
			if strings.Contains(posStop, "|"+key.Pos+"|") || IsFilterContains(key.Const) {
				// fmt.Println(key.Pos, key.Const)
				continue
			}
			key_freq[key.Const]++ //strings.Trim(, " ")
		}
		// fmt.Println(key_slice)
	}
	maxFreq := int32(0)
	sec_maxFreq := int32(0)
	// 最大词频
	for _, v := range key_freq {
		if maxFreq < v {
			sec_maxFreq = maxFreq
			maxFreq = v
		}
	}
	// 缩小最大词频群
	for k, v := range key_freq {
		if maxFreq-1 < v || sec_maxFreq-1 < v {
			key_freq[k] = (sec_maxFreq + maxFreq) / 2
			fmt.Print(k, "\t")
		}
	}
	key_freq_o := make(map[string]int32, 1)
	for _, key_slice := range *o {
		for _, key := range key_slice {
			if strings.Contains(posStop, "|"+key.Pos+"|") || IsFilterContains(key.Const) {
				// fmt.Println(key.Pos, key.Const)
				continue
			}
			key_freq_o[key.Const]++ //strings.Trim(, " ")
		}
		// fmt.Println(key_slice)
	}

	// 按照词频排序
	cellSlice := make(CellSlice, len(key_freq_o))
	i := 0
	var sum_freq_o int32
	for key_word, freq := range key_freq_o {
		cell := Cell{key_word, freq}
		cellSlice[i] = &cell
		i++
		sum_freq_o += freq
	}
	sort.Sort(cellSlice)
	// cellSlice.String()

	// 卡方检验
	// nq 期望值为 sum_freq_o/len(key_freq)
	nq := float64(sum_freq_o) / float64(len(key_freq_o))
	var score float64
	// 组合关键字
	sentences := make(Sens, len(*t))
	// 后续处理
	var minFreq int32
	relate := ""
	for i, key_slice := range *t {
		score = 0.0
		minFreq = 1000
		relate = ""
		for _, key_word := range key_slice {
			fr, ok := key_freq[key_word.Const] //strings.Trim(, " ")
			if minFreq > fr {
				minFreq = fr
			}
			if ok {
				score += math.Pow(float64(fr)-nq, 2.0) / nq
			}
			if strings.Contains(posStop, "|"+key_word.Pos+"|") || IsFilterContains(key_word.Const) {
				score *= 0.0
				break
			}
			relate += key_word.Relate + "-"
		}
		if minFreq < 2 {
			score *= 0.7
		}
		if len(key_slice) == 2 && key_slice[0].Const == key_slice[1].Const {
			score *= 0.0
		}
		// 降低单个词的卡方值
		score *= math.Log2(float64(len(key_slice) + 1))
		sen := Sen{Str: key_slice.WordStrings(), Sum: score, Avg: score, Rel: relate}
		sentences[i] = &sen
	}
	// 排除重复
	sentences = *sentences.EjRepeat()
	// improve
	for _, sentence := range sentences {
		sentence.Sum = math.Log2(float64(sentence.Fre+1)) * sentence.Avg // 短文较差，长文还可以
		// sentence.Sum = float64(sentence.Fre) * sentence.Avg
	}
	// 按照卡方值排序
	sort.Sort(sentences)
	saveCoWords(sentences)

	topN := 50 /*len(key_freq)/500 + */
	topFormatString, topNSlice := sentences.Top(topN)
	fmt.Print(topFormatString)
	statResultString := PreciseAndRecall(topNSlice)
	statResultString = "\n\n一元词Fi'" + statResultString
	SaveResult(statResultString)
}

func saveCoWords(s Sens) {
	relateMap := make(map[string]int, 1)
	for _, sen := range s {
		// if len(sen.Rel) > 6 && 0.0 < sen.Sum {
		// 	fmt.Printf("%d %s %s %.3f\n", i, sen.Str, sen.Rel, sen.Sum)
		// }
		relateMap[sen.Rel]++
	}
	fmt.Println()
	length := len(s)
	for k, v := range relateMap {
		fmt.Printf("%s\t %d\t %.5f\n", k, v, float32(v)/float32(length))
		// fmt.Printf("%s\t %d\n", k, v)
	}
}

func realteSet(m TopicMatix) {
	relateMap := make(map[string]int, 1)
	ori_relateMap := make(map[string]int, 1)
	relate := ""
	sigal_length := 0
	double_length := 0
	for _, topic_slice := range m {
		relate = ""
		if 1 == len(topic_slice) {
			sigal_length += len(topic_slice)
			ori_relateMap[topic_slice[0].Pos]++
			continue
		}
		for _, t := range topic_slice {
			// fmt.Printf("%s-", t.Relate)
			relate += t.Relate + "-"
		}
		double_length += 1
		relateMap[relate]++
	}

	for k, v := range relateMap {
		fmt.Printf("%s\t %d\t %.5f\n", k, v, float32(v)/float32(double_length))
	}
	fmt.Println()
	for k, v := range ori_relateMap {
		fmt.Printf("%s\t %d\t %.5f\n", k, v, float32(v)/float32(sigal_length))
	}
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
	key_words := strings.Split(key_word, "||")
	key_words_len := len(key_words) - 2
	// 提取的关键字比给出的关键字多bias个
	bias := 2
	key_words = key_words[1 : key_words_len+1]
	tops = tops[:key_words_len+bias] //
	count := 0.0
	containWords := strings.Join(key_words, "|")
	containWords = "|" + containWords + "|"
	fmt.Println(containWords)
	for _, it := range tops {
		if strings.Contains(containWords, "|"+it+"|") {
			count += 1.0
		}
	}
	pricise := count / float64(key_words_len+bias) //
	recall := count / float64(key_words_len)
	b := 1.0
	f_measrue := (b*b + 1) * pricise * recall / (b*b*pricise + recall)

	priciseString := fmt.Sprintf("精确率：\t%.4f\n", pricise) // len(key_words)-2
	recallString := fmt.Sprintf("召回率：\t%.4f\n", recall)   // len(tops)
	f_measrueString := fmt.Sprintf("F-measure：\t%.4f\n", f_measrue)
	resultStr += "\n Co-word Chi \n"
	resultStr += fmt.Sprintf("Top key words:\t%v\n", tops)
	resultStr += fmt.Sprintf("Key words:\t%v\n", key_words)
	resultStr += priciseString + recallString + f_measrueString
	return resultStr
}

func SaveResult(content string) {
	// 保存结果
	stringSaveChan := make(chan string, 5)
	go utils.SaveStringAppend(stringSaveChan, STAT_DIR+CURRENT_FILENAME)
	// fmt.Printf("\n%s\t %s\t %s\t %s\n", "index", "avg", "key-word", "freq")
	stringSaveChan <- content
	fmt.Println(content)
}

type Sen struct {
	Str string
	Sum float64
	Avg float64
	Fre int
	Rel string
}

func (s *Sen) String() string {
	// return fmt.Sprintf("%.3f\t %.3f\t %s\t", s.Avg, s.Sum, s.Str)
	return fmt.Sprintf("%.2f\t %.2f\t %s\t %d", s.Avg, s.Sum, s.Str, s.Fre)
}

type Sens []*Sen

func (c Sens) Len() int {
	return len(c)
}

func (c Sens) Less(i, j int) bool {
	// return float64(c[i].Fre)*c[i].Avg > float64(c[j].Fre)*c[j].Avg // 比较合理
	return c[i].Sum > c[j].Sum // 长文不可靠，短文还可以
	// return c[i].Avg > c[j].Avg // 长文不可靠，短文还可以
	// return math.Log2(float64(c[i].Fre))*c[i].Avg > math.Log2(float64(c[j].Fre))*c[j].Avg // 短文较差，长文还可以
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
