package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unPackStc string) (string, error) {
	isEkr := 0
	symb := '-'
	var err error
	var b strings.Builder

	for _, r := range unPackStc {
		err = checkSymb(r)

		if err != nil {
			b.Reset()
			break
		}

		switch {
		case isEkr == 0 && r == '\\':
			// пришло экранирование, прошлый символ допечатываем если надо
			if symb != '-' {
				b.WriteRune(symb)
			}
			symb = '-'
			isEkr = 1
		case isEkr == 0 && r <= '9':
			// пришла цифра, надо повторять
			if symb == '-' { // а повторять нечего - выходим
				err = ErrInvalidString
				b.Reset()
				break
			} else {
				// раз 0 - ничего не выводим
				if r == '0' {
					symb = '-'
				} else { // значит повторим прошлы символ столько раз, сколько написано
					for i := 0; i < int(r-'0'); i++ {
						b.WriteRune(symb)
					}
					symb = '-'
				}
			}
		case isEkr == 0:
			// пришел символ, экранирования нет, просто сохраним
			if symb != '-' {
				b.WriteRune(symb)
			}
			symb = r
		case isEkr == 1:
			// пришло экраниррованное значение, сохраним его
			symb = r
			isEkr = 0
		}
	}

	if isEkr == 1 { // ошибка что последним пришло экранирование
		err = ErrInvalidString
		b.Reset()
	}

	if symb != '-' {
		b.WriteRune(symb)
	}

	return b.String(), err
}

func checkSymb(chksd rune) error {
	var err error
	switch {
	case chksd == '\\':
		err = nil
	case chksd <= '9':
		if chksd < '0' {
			err = ErrInvalidString
		}
	case chksd <= 'Z':
		if chksd < 'A' {
			err = ErrInvalidString
		}
	case chksd <= 'z':
		if chksd < 'a' {
			err = ErrInvalidString
		}
	default:
		err = nil
	}

	return err
}
