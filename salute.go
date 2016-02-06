package polite

func Salute(name string) string{
	var s string

	if name == ""  {
		s = "Hello nice to meet you!"
	} else {
		s = "Good morning Mr. " + name
	}

	return s
}

func Aurevoir(name string) string{
	var s string

	if name == ""  {
		s = "Bye nice to meet you!"
	} else {
		s = "Good night Mr. " + name
	}

	return s
}