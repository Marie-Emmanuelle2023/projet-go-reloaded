package rules

import(
	"strconv"
	"strings"
	"unicode"
)

func Capitalize(mot string) string{
	if mot == ""{
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}


func applyCase(mots []string, mode string, n int){
	start := len(mots) - n
	if start < 0{
		start = 0
	}
	for i := start; i < len(mots); i++{
		switch mode{
		case "up":
			mots[i] = strings.ToUpper(mots[i])
		case "low":
			mots[i] = strings.ToLower(mots[i])
		case "cap":
			mots[i] = Capitalize(mots[i])
		}
	}
}

func applyCaseN(mots []string, tag string, count string) bool {
	if !strings.HasSuffix(count, ")"){
		return false
	}
	n, err := strconv.Atoi(strings.TrimSuffix(count, ")"))
	if err != nil || n < 0{
		return false
	}
	applyCase(mots, tag[1:len(tag)-1], n)
	return true
}
