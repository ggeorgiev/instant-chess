package chess

type Square struct {
	X int
	Y int
}

type Answer struct {
	BlackFrom Square
	BlackTos  []Square
}

type ToAnswer struct {
	WhiteTo Square
	Answer  Answer
}

type Move struct {
	WhiteForm Square
	ToAnswers []ToAnswer
}
