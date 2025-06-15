package main

func main() {

}

type Present struct {
	lenght int
	width int
	height int
}

type PresentWrap struct {
	squareFeet int
} 

func (p *PresentWrap) GetExtraPaper() int {
	return p.squareFeet + 6
}
