package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unPackStc string) (string, error) {
	isEkr := 0
	symb := '-'
	needPrnt := false

	var err error
	var b strings.Builder

	for _, r := range unPackStc {
		switch {
		case isEkr == 0 && r == '\\':
			// пришло экранирование, прошлый символ допечатываем если надо
			if needPrnt {
				b.WriteRune(symb)
			}
			needPrnt = false
			isEkr = 1
		case isEkr == 0 && unicode.IsDigit(r):
			// пришла цифра, надо повторять
			if !needPrnt { // а повторять нечего - выходим
				err = ErrInvalidString
				return "", err
			}

			// значит повторим прошлый символ столько раз, сколько написано
			b.WriteString(strings.Repeat(string(symb), int(r-'0')))
			needPrnt = false

		case isEkr == 0:
			// пришел символ, экранирования нет, просто сохраним
			if needPrnt {
				b.WriteRune(symb)
			}
			symb = r
			needPrnt = true

		case isEkr == 1:
			// пришло экраниррованное значение, сохраним его
			if !unicode.IsDigit(r) && r != '\\' {
				return "", ErrInvalidString
			}
			symb = r
			isEkr = 0
			needPrnt = true
		}
	}

	if isEkr == 1 { // ошибка что последним пришло экранирование
		return "", ErrInvalidString
	}

	if needPrnt {
		b.WriteRune(symb)
	}

	return b.String(), nil
}
