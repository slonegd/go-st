package source

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Source []rune

// offset - доп строки до и после для контекста
func (x Source) PositionSnippet(position, offset int) string {
	line, column := x.findLineColumn(position)
	return x.Snippet(line, column, offset)
}

// offset - доп строки до и после для контекста
// TODO смещение номера строки по всему коду, чтоб не прыгал во время дебага (по номеру последней строки)
func (x Source) Snippet(line, column, offset int) string {
	lines, messageLine, ok := x.findLines(line, offset)
	if !ok {
		return ""
	}
	var r string
	for i, line := range lines {
		if i == 0 {
			r = "\n"
		}
		r += line
		if i != messageLine {
			r += "\n"
			continue
		}
		for range column { // TODO учесть табуляцую
			r += "."
		}
		r += "^    %s\n" // форматер для сообщения
	}

	return r
}

// поиск строк с несколькими строками до и после для контекста
// а также вставка пустой строки для сообщения
// вставка номера строки
func (s Source) findLines(line int, offset int) ([]string, int, bool) {
	if len(s) == 0 {
		return nil, 0, false
	}
	line -= 1 // массивы с нуля
	lines := strings.Split(string(s), "\n")
	startLine := 0
	if line > offset {
		startLine = line - offset // захватить offset строки до для контекста
	}
	endLine := len(lines)
	if line < endLine-offset {
		endLine = line + offset // захватить offset строки после для контекста
	}
	messageLine := line - startLine + 1
	r := slices.Clone(lines[startLine : line+1])
	r = append(r, "")
	if line < endLine {
		r = append(r, lines[line+1:endLine]...)
	}
	numberLinesColumns := powerOf10(endLine)
	positionFormatter := "% " + strconv.Itoa(numberLinesColumns) + "d|"
	messageFormat := ""
	for range numberLinesColumns {
		messageFormat += " "
	}
	messageFormat += "|"
	for i := range r {
		switch {
		case i < messageLine:
			r[i] = fmt.Sprintf(positionFormatter, i+startLine+1) + r[i]
		case i == messageLine:
			r[i] = messageFormat + r[i]
		case i > messageLine:
			r[i] = fmt.Sprintf(positionFormatter, i+startLine) + r[i]
		}
	}
	return r, messageLine, true
}

func (s Source) findLineColumn(position int) (line, column int) {
	lines := strings.Split(string(s), "\n")
	count := 0
	for i, line := range lines {
		count += len(line) + 1
		if position <= count {
			return i, count - position
		}
	}
	return 0, 0
}

func powerOf10(n int) int {
	r := 2
	for n > 0 {
		n /= 10
		r++
	}
	return r
}
