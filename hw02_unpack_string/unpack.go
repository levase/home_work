package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack распаковывает строку с повторяющимися символами.
// Пример: "a4bc2" -> "aaaabbc"
// Поддерживает экранирование через \ (например, \4 -> цифра 4 без повтора).
func Unpack(str string) (string, error) {
	// Пустая строка на входе -> пустая строка на выходе
	if str == "" {
		return "", nil
	}

	isDigit := func(r rune) bool { return r >= '0' && r <= '9' }

	var result strings.Builder
	result.Grow(len(str))
	runes := []rune(str)
	i := 0

	for i < len(runes) {
		// Обработка escape-последовательности ('\')
		if runes[i] == '\\' {
			// '\' в конце строки - ошибка
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}
			// Экранировать можно только '\' или цифру
			if runes[i+1] != '\\' && !isDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
			char := runes[i+1]
			i += 2

			// Проверяем, есть ли цифра после экранированного символа
			if i < len(runes) && isDigit(runes[i]) {
				count := int(runes[i] - '0')
				result.WriteString(strings.Repeat(string(char), count))
				i++
			} else {
				result.WriteRune(char)
			}
			continue
		}

		// Цифра в начале или после другой цифры - ошибка
		if isDigit(runes[i]) {
			return "", ErrInvalidString
		}

		char := runes[i]
		i++

		// Если после символа идёт цифра - повторяем символ
		if i < len(runes) && isDigit(runes[i]) {
			count := int(runes[i] - '0')
			result.WriteString(strings.Repeat(string(char), count))
			i++
		} else {
			result.WriteRune(char)
		}
	}

	return result.String(), nil
}
