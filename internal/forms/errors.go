package forms

type errors map[string][]string

func (e errors) AddError(value, field string) {
	e[field] = append(e[field], value)
}

func (e errors) GetError(filed string) string {
	es := e[filed]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
