package gptemplate

type Template struct {
	delimiter    Delimiter
	replacements map[string]string
}

func NewTemplate(delimiter Delimiter, replacements map[string]string) *Template {

	return &Template{
		delimiter:    delimiter,
		replacements: replacements,
	}
}

func NewTemplateCurlyBrackets(replacements map[string]string) *Template {

	return &Template{
		delimiter:    DelimiterCurlyBrackets,
		replacements: replacements,
	}
}

func NewTemplateSquareBrackets(replacements map[string]string) *Template {

	return &Template{
		delimiter:    DelimiterSquareBrackets,
		replacements: replacements,
	}
}

func NewTemplateRoundBrackets(replacements map[string]string) *Template {

	return &Template{
		delimiter:    DelimiterRoundBrackets,
		replacements: replacements,
	}
}

func (that *Template) Replace(input string) string {

	return Replace(input, that.delimiter, that.replacements)
}
