package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unPackStc string) (string, error) {
	symb := '-'
	needPrnt := false
	isEkr := false

	var err error
	var b strings.Builder

	for _, r := range unPackStc {
		switch {
		case !isEkr && r == '\\':
			// пришло экранирование, прошлый символ допечатываем если надо
			if needPrnt {
				b.WriteRune(symb)
			}
			needPrnt = false
			isEkr = true
		case !isEkr && r >= '0' && r <= '9':
			// пришла цифра, надо повторять
			if !needPrnt { // а повторять нечего - выходим
				err = ErrInvalidString
				return "", err
			}

			// значит повторим прошлый символ столько раз, сколько написано
			b.WriteString(strings.Repeat(string(symb), int(r-'0')))
			needPrnt = false

		case !isEkr:
			// пришел символ, экранирования нет, просто сохраним
			if needPrnt {
				b.WriteRune(symb)
			}
			symb = r
			needPrnt = true

		case isEkr:
			// пришло экраниррованное значение, сохраним его
			if (r > '9' || r < '0') && r != '\\' {
				return "", ErrInvalidString
			}
			symb = r
			isEkr = false
			needPrnt = true
		}
	}

	if isEkr { // ошибка что последним пришло экранирование
		return "", ErrInvalidString
	}

	if needPrnt {
		b.WriteRune(symb)
	}

	return b.String(), nil
}
