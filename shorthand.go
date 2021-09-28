package spidol

func Red(text string) Spidol {
	red := New(text).Red()
	return Spidol(red)
}

func Green(text string) Spidol {
	green := New(text).Green()
	return Spidol(green)
}

func Yellow(text string) Spidol {
	yellow := New(text).Yellow()
	return Spidol(yellow)
}

func Blue(text string) Spidol {
	blue := New(text).Blue()
	return Spidol(blue)
}

func Purple(text string) Spidol {
	purple := New(text).Purple()
	return Spidol(purple)
}

func White(text string) Spidol {
	white := New(text).White()
	return Spidol(white)
}
