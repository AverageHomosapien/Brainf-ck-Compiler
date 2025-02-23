package translators

type Translator interface {
	Translate(code string)
}
