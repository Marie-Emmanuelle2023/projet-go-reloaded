package rules
import "strconv"

func ConvertBase(mot string, base int) string{
	n, err:= strconv.ParseInt(mot,base,64)
	if err!=nil {
		return mot
	}
	return strconv.	FormatInt(n,10)
}

func ConvertLast(mots []string, base int){
	if len(mots)== 0{
		return
	}
	last := len(mots) - 1
	mots[last] = ConvertBase(mots[last], base)
}
