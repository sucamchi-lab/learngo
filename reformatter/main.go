package main

func reformat(message string, formatter func(string) string) string {
	first := formatter(message)
	second := formatter(first)
	third := formatter(second)

	final := "TEXTIO: " + third
	return final
}
