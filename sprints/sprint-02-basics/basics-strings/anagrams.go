package problems

import "strings"

// AreAnagrams определяет, являются ли две строки анаграммами.
//
// Вход:
//   - a string: первая строка, например "listen"
//   - b string: вторая строка, например "silent"
//
// Выход:
//   - bool: true, если строки состоят из одинакового набора символов после
//     приведения к одному регистру и удаления пробелов; иначе false.
//
// Пример:
//   - a = "listen", b = "silent" -> true
//   - a = "rat",    b = "car"    -> false
func AreAnagrams(a, b string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	a, b = strings.ToLower(strings.Join(strings.Fields(a), "")),
		strings.ToLower(strings.Join(strings.Fields(b), ""))
	m := make(map[rune]int)
	for _, val := range a {
		if _, ok := m[val]; !ok {
			m[val] = 0
		}
		m[val]++
	}
	for _, val := range b {
		if _, ok := m[val]; !ok {
			return false
		}
		m[val]--
		if m[val] < 0 {
			return false
		}
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
