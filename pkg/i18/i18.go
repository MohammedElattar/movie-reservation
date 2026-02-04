// Package i18
package i18

import "fmt"

type (
	Category int
	Locale   string
)

const (
	Messages Category = iota
	Validation
)

const (
	EnLocale Locale = "en"
	ArLocale Locale = "ar"
)

type Bundle struct {
	tables map[Locale]map[Category]map[string]string
}

func New() *Bundle {
	return &Bundle{
		tables: make(map[Locale]map[Category]map[string]string),
	}
}

func (b *Bundle) Register(locale Locale, cat Category, data map[string]string) {
	if _, ok := b.tables[locale]; !ok {
		b.tables[locale] = make(map[Category]map[string]string)
	}

	b.tables[locale][cat] = data
}

func (b *Bundle) Error(locale Locale, fieldKey, errorKey string) string {
	fieldKey = b.T(locale, Messages, fieldKey)

	return b.T(locale, Validation, errorKey, fieldKey)
}

func (b *Bundle) Success(locale Locale, subjectKey, action string) string {
	subjectKey = b.T(locale, Messages, subjectKey)

	return b.T(locale, Messages, action, subjectKey)
}

func (b *Bundle) Word(locale Locale, key string, params ...any) string {
	return b.T(locale, Messages, key, params...)
}

func (b *Bundle) T(locale Locale, cat Category, key string, params ...any) string {
	if l, ok := b.tables[locale]; ok {
		if table, ok := l[cat]; ok {
			if msg, ok := table[key]; ok {
				if len(params) > 0 {
					return fmt.Sprintf(msg, params...)
				}
				return msg
			}
		}
	}

	return fmt.Sprintf("%s.%s", catString(cat), key)
}

func catString(cat Category) string {
	switch cat {
	case Messages:
		return "messages"
	case Validation:
		return "validation"
	default:
		return "unknown"
	}
}
