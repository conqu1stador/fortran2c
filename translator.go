package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getParams(regEx, str string) (paramsMap map[string]string) {
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(str)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func buildFormatString(arguments []string, print bool) string {
	integer := regexp.MustCompile(`(?:I|J|K|L|M|N)\w{0,5}`)
	var formatString string
	for _, val := range arguments {
		if integer.MatchString(val) {
			formatString += "%d"
		} else {
			formatString += "%f"
		}
	}
	formatString += "\""
	for _, val := range arguments {
		if print {
			formatString += ", " + strings.Trim(val, " ")
		} else {
			formatString += ", &" + strings.Trim(val, " ")
		}
	}

	return formatString
}

func removeString(text []string, s int) []string {
	return append(text[:s], text[s+1:]...)
}

//ФУНКЦИИ ДЛЯ ЗАМЕН

var arrNames []string

//done
func replaceComments(text *[]string, showComments bool) {
	for i := 0; i < len(*text); i++ {
		var comment = regexp.MustCompile(`^C`)
		if comment.MatchString((*text)[i]) {
			if showComments {
				(*text)[i] = comment.ReplaceAllString((*text)[i], `// `)
			} else {
				(*text) = removeString((*text), i)
				i--
			}
		}
	}
}

//+-
func replaceRead(text *[]string) {
	read := regexp.MustCompile(`READ\s*\(\d{1,3}\,(?P<LABEL>(?:\*|\d{1,3}))\) *(?P<VARS>(?:\w*|\,| |\(|\)|\')*)`)

	for i := 0; i < len(*text); i++ {
		params := getParams(read.String(), (*text)[i])

		if read.MatchString((*text)[i]) {
			if params["LABEL"] != "*" {
				readFormatLabel(text, params["LABEL"], false)
			}

			arguments := regexp.MustCompile(`(?:\w* *\((?:\w*|\,| )+\))|\w+`).FindAllString(params["VARS"], -1)

			(*text)[i] = `scanf("` + buildFormatString(arguments, false) + ");"
		}
	}
}

// ???
func replaceWrite(text *[]string) {
	simpleString := regexp.MustCompile(`'\w+'`)

	write := regexp.MustCompile(`WRITE\s*\(\d{1,3}\,(?P<LABEL>(?:\*|\d{1,3}))\) *(?P<VARS>(?:\w*|\,| |\(|\)|\')*)`)

	for i := 0; i < len(*text); i++ {
		if write.MatchString((*text)[i]) {
			params := getParams(write.String(), (*text)[i])

			if simpleString.MatchString(params["VARS"]) {
				(*text)[i] = `printf("` + strings.Trim(params["VARS"], "'") + "\\n\");"
			} else if params["LABEL"] == "*" {
				arguments := regexp.MustCompile(`(?:\w* *\((?:\w*|\,| )+\))|\w+`).FindAllString(params["VARS"], -1)
				(*text)[i] = `printf("` + buildFormatString(arguments, true) + ");"
			} else {
				formatString := ""
				formatString = writeFormatLabel(text, params["LABEL"], false)

				formatString = regexp.MustCompile(`(?:,|)\d{1,2}H`).ReplaceAllString(formatString, "")
				formatString = regexp.MustCompile(`(?:,|)\d{1,2}X,`).ReplaceAllString(formatString, "")
				formatString = regexp.MustCompile(`\$$`).ReplaceAllString(formatString, "")
				formatString = regexp.MustCompile(`^,`).ReplaceAllString(formatString, "")

				//tmpParams := getParams(regexp.MustCompile(`\, *(?P<AMOUNT>\d*)(?P<TYPE>I|F|D)(?:\d|\.)*`).String(), formatString)
				printfParamsRegex := regexp.MustCompile(`\, *(?P<AMOUNT>\d*)(?P<TYPE>I|F|D)(?:\d|\.)*`)

				for _, val := range printfParamsRegex.FindAllString(formatString, -1) {
					tmpParams := getParams(printfParamsRegex.String(), val)
					if tmpParams["AMOUNT"] == "" {
						tmpParams["AMOUNT"] = "1"
					}

					amount, _ := strconv.ParseInt(tmpParams["AMOUNT"], 10, 64)

					var outputArgs string
					for i := 0; i < int(amount); i++ {
						if tmpParams["TYPE"] == "I" {
							outputArgs += "%d"
						} else if tmpParams["TYPE"] == "D" || tmpParams["TYPE"] == "F" {
							outputArgs += "%f"
						}
					}

					formatString = strings.Replace(formatString, val, outputArgs, 1)
				}

				formatString = strings.ReplaceAll(formatString, "'", "")

				arguments := regexp.MustCompile(`(?:\w* *\((?:\w*|\,| )+\))|\w+`).FindAllString(params["VARS"], -1)
				//fmt.Println(arguments)
				printfString := `printf("` + formatString
				printfString += "\\n\""
				for j := 0; j < len(arguments); j++ {
					printfString += ", " + strings.Trim(arguments[j], " ")
				}
				(*text)[i] = printfString + ");"
				//fmt.Println((*text)[i])
			}
		}

	}
}

func replaceDimensions(text *[]string) {

}

//++-
func readFormatLabel(text *[]string, label string, showComments bool) {
	labelFormat := regexp.MustCompile(`^.*` + label + ` `)
	for i := 0; i < len(*text); i++ {
		if labelFormat.MatchString((*text)[i]) {
			if showComments {
				(*text)[i] = "//" + (*text)[i]
			} else {
				(*text) = removeString((*text), i)
				i--
			}
		}
	}
}

//++-
func writeFormatLabel(text *[]string, label string, showComments bool) string {
	labelFormat := regexp.MustCompile(`(?mUs)^\s*` + label + `\s*FORMAT\(.*\)`)
	format := regexp.MustCompile(`FORMAT\(.*\)`)
	//labelFormat := regexp.MustCompile(`(?mUs)` + label + ` FORMAT\(.*\)`)
	var formatString string
	for i := 0; i < len(*text); i++ {
		if labelFormat.MatchString((*text)[i]) {
			formatString = strings.Trim(format.FindString((*text)[i])[len(label)+7:], "()")
			if showComments {
				(*text)[i] = "//" + (*text)[i]
			} else {
				(*text) = removeString((*text), i)
				i--
			}
		}
	}
	return formatString
}

//++-
func changeLabel(text *[]string, label string) {
	labelFormat := regexp.MustCompile(`^\s+` + label + ` `)
	for i := 0; i < len(*text); i++ {
		if labelFormat.MatchString((*text)[i]) {
			(*text) = append((*text), "")
			copy((*text)[i+1:], (*text)[i:])
			(*text)[i+1] = strings.Trim((*text)[i], " ")[len(label)+1:]
			(*text)[i] = "l_" + label + ":"
		}
	}
}

//done
func changeDoLabel(text *[]string, label string) {
	labelFormat := regexp.MustCompile(`^\s+` + label + ` `)
	for i := 0; i < len(*text); i++ {
		if labelFormat.MatchString((*text)[i]) {
			(*text)[i] = "}"
		}
	}
}

//done
func replaceIf(text *[]string) {
	ifStatement := regexp.MustCompile(`^\s*IF.*\(.*\)`)
	logicalOperators := []string{".LT.", ".LE.", ".EQ.", ".NE.", ".GE.", ".GT.", ".NOT.", ".AND.", ".OR."}
	logicalOperatorsC := []string{" < ", " <= ", " == ", " != ", " >= ", " > ", "!", "&&", "||"}
	for i := 0; i < len(*text); i++ {
		if ifStatement.MatchString((*text)[i]) {
			for index, operator := range logicalOperators {
				tmp := regexp.MustCompile(operator)
				if tmp.MatchString((*text)[i]) {
					(*text)[i] = tmp.ReplaceAllString((*text)[i], logicalOperatorsC[index])
				}
			}
			//1. if () goto
			//2. if () lbl1,lbl2,lbl3
			//3. if () do smth
			ifGoTo := regexp.MustCompile(`GOTO.*`)
			ifManyGoTo := regexp.MustCompile(`\).*,.*,.*`)
			//ifDoSmth := regexp.MustCompile(`.*=`)

			if ifGoTo.MatchString((*text)[i]) {
				label := strings.Trim(regexp.MustCompile(`GOTO`).ReplaceAllString(ifGoTo.FindString((*text)[i]), ""), " ")
				changeLabel(text, label)
				(*text)[i] = strings.TrimRight((*text)[i], " ")[:len((*text)[i])-len(label)] + "l_" + label + ";"
			} else if ifManyGoTo.MatchString((*text)[i]) {
				labels := strings.Split(ifManyGoTo.FindString((*text)[i])[2:], ",")
				for _, label := range labels {
					changeLabel(text, label)
				}

				expression := regexp.MustCompile(`IF.*\(`).ReplaceAllString(ifStatement.FindString((*text)[i]), "")
				expression = strings.Trim(expression[:len(expression)-1], " ")

				(*text)[i] = "if (" + expression + " < 0) goto l_" + labels[0] + ";\n" +
					"if (" + expression + " == 0) goto l_" + labels[1] + ";\n" +
					"if (" + expression + " > 0) goto l_" + labels[2] + ";"
			}

			(*text)[i] = regexp.MustCompile(`GOTO`).ReplaceAllString((*text)[i], "goto")
			(*text)[i] = regexp.MustCompile(`IF`).ReplaceAllString((*text)[i], "if")
		}
	}
}

//done
func replaceDo(text *[]string) {
	doStatement := regexp.MustCompile(`DO [^=].*`)
	doLabel := regexp.MustCompile(`DO \d+`)
	initNameRegex := regexp.MustCompile(`(?mU)\w{0,5}\s*=`)
	integer := regexp.MustCompile(`(?mU)(?:I|J|K|L|M|N)\w{0,5}\s*=`)

	for i := 0; i < len(*text); i++ {
		if doStatement.MatchString((*text)[i]) {
			label := regexp.MustCompile(`\d+`).FindString(doLabel.FindString((*text)[i]))
			cycleValues := strings.Split(strings.Trim(doLabel.ReplaceAllString((*text)[i], " "), " "), ",")
			changeDoLabel(text, label)

			var initValue, optionalIncrement string
			//fmt.Println(integer.MatchString((*text)[i]), (*text)[i])
			if integer.MatchString((*text)[i]) {

				initValue = "int "
			} else {
				initValue = "double "
			}

			initValue += cycleValues[0][:len(cycleValues[0])]

			initName := strings.TrimRight(initNameRegex.FindString((*text)[i])[:len(initNameRegex.FindString((*text)[i]))-1], " ")
			finalValue := initName + " <= " + cycleValues[1]

			if len(cycleValues) == 2 {
				optionalIncrement = initName + "++"
			} else {
				optionalIncrement = initName + " += " + cycleValues[2]
			}

			(*text)[i] = "for (" + initValue + "; " + finalValue + "; " + optionalIncrement + ") {"
		}
	}
}

//+-
func replaceStop(text *[]string) {
	stop := regexp.MustCompile(`STOP(\s|)(?P<ARG>\d+|)`)
	end := regexp.MustCompile(`^\s*END`)
	goTo := regexp.MustCompile(`^\s+GOTO\s*`)
	for i := 0; i < len(*text); i++ {
		if stop.MatchString((*text)[i]) {
			params := getParams(stop.String(), (*text)[i])
			(*text)[i] = stop.ReplaceAllString((*text)[i], "return ")
			if params["ARG"] == "" {
				(*text)[i] += "0"
			} else {
				(*text)[i] += params["ARG"]
			}

		} else if end.MatchString((*text)[i]) {
			(*text)[i] = strings.TrimLeft(end.ReplaceAllString((*text)[i], "}"), " ")
		} else if goTo.MatchString((*text)[i]) {
			label := strings.Trim(goTo.ReplaceAllString((*text)[i], " "), " ")
			(*text)[i] = "goto l_" + label + ";"
			changeLabel(text, label)
		}
	}
}

//+--
func parseVariables(text *[]string) {
	integer := regexp.MustCompile(`^\s*INTEGER`)
	double := regexp.MustCompile(`^\s*DOUBLE PRECISION`)
	float := regexp.MustCompile(`^\s*REAL`)
	arr := regexp.MustCompile(`^\s*DIMENSION *(?P<VARS>(?:\w*|\,| |\(|\)|\')*)`)
	arrNamesRegex := regexp.MustCompile(`(?P<NAME>\w*)\((?:\w|\,)*\)`)
	//complex := regexp.MustCompile(`^\s*DOUBLE PRECISION`)
	for i := 0; i < len(*text); i++ {
		if integer.MatchString((*text)[i]) {
			(*text)[i] = integer.ReplaceAllString((*text)[i], "int")
		} else if double.MatchString((*text)[i]) {
			(*text)[i] = double.ReplaceAllString((*text)[i], "double")
		} else if float.MatchString((*text)[i]) {
			(*text)[i] = double.ReplaceAllString((*text)[i], "float")
		} else if arr.MatchString((*text)[i]) {
			params := getParams(arr.String(), (*text)[i])

			for _, val := range arrNamesRegex.FindAllString((*text)[i], -1) {
				arrNames = append(arrNames, getParams(arrNamesRegex.String(), val)["NAME"])
			}

			arguments := regexp.MustCompile(`(?:\w* *\((?:\w*|\,| )+\))|\w+`).FindAllString(params["VARS"], -1)
			(*text)[i] = "float "
			for index, val := range arguments {
				initName := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(val, ",", "]["), ")", "]"), "(", "[")
				(*text)[i] += initName + ", "

				replaceDimensionByName(text, arrNames[index])
			}
			(*text)[i] = (*text)[i][:len((*text)[i])-2] + ";"
		}
	}
}

func replaceDimensionByName(text *[]string, arrName string) {
	dimensionRegex := regexp.MustCompile(arrName + `\((?:\w|\,)*\)`)
	for i := 0; i < len(*text); i++ {
		if dimensionRegex.MatchString((*text)[i]) {
			//(*text)[i] = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(dimensionRegex.FindString((*text)[i]), ",", "]["), ")", "]"), "(", "[")
			newName := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(dimensionRegex.FindString((*text)[i]), ",", " - 1]["), ")", " - 1]"), "(", "[")
			(*text)[i] = dimensionRegex.ReplaceAllString((*text)[i], newName)
			//fmt.Println(newName, "----", dimensionRegex.ReplaceAllString((*text)[i], newName))
		}
	}
}

//+-
func checkSemicolons(text *[]string) {
	for i := 0; i < len(*text); i++ {
		if len(strings.Trim((*text)[i], " ")) > 2 &&
			(*text)[i][(len((*text)[i]))-1:] != ";" &&
			(*text)[i][(len((*text)[i]))-1:] != "{" &&
			(*text)[i][(len((*text)[i]))-1:] != ":" {
			(*text)[i] += ";"
		}
		(*text)[i] = strings.TrimLeft((*text)[i], " ")
	}
}

//-
func replaceMathFunctions(text *[]string) {
	sqrt := regexp.MustCompile(`(?:D|C|)SQRT`)
	abs := regexp.MustCompile(`(?:D|C|)ABS`)
	sqr := regexp.MustCompile(`\*\*`)
	for i := 0; i < len(*text); i++ {
		if sqrt.MatchString((*text)[i]) {
			(*text)[i] = sqrt.ReplaceAllString((*text)[i], "sqrt")
		}
		if abs.MatchString((*text)[i]) {
			(*text)[i] = abs.ReplaceAllString((*text)[i], "abs")
		}
		if sqr.MatchString((*text)[i]) {
			(*text)[i] = sqr.ReplaceAllString((*text)[i], "^")
		}
	}
}

func beautify(text *[]string) {
	var offset string
	for i := 0; i < len(*text); i++ {
		if len((*text)[i]) != 0 {
			if (*text)[i][len((*text)[i])-1] == '}' {
				offset = offset[:len(offset)-1]
			}
			(*text)[i] = offset + (*text)[i]
			if (*text)[i][len((*text)[i])-1] == '{' {
				offset += "\t"
			}
		}
	}
}

func main() {
	fileRead, err := os.Open("f2.f")

	check(err)
	scanner := bufio.NewScanner(fileRead)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	fileRead.Close()

	fileWrite, err := os.Create("test.c")
	datawriter := bufio.NewWriter(fileWrite)

	var textC []string
	textC = append(textC, "#include <stdio.h>\n#include <math.h>\n ", "int main() {")

	replaceComments(&text, false)
	replaceStop(&text)

	replaceIf(&text)
	replaceDo(&text)

	replaceRead(&text)
	replaceWrite(&text)

	parseVariables(&text)
	replaceMathFunctions(&text)

	checkSemicolons(&text)

	textC = append(textC, text...)

	//beautify(&textC)
	for _, data := range textC {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	fileWrite.Close()
}
