package model

type Song struct {
	Genre string
	Title string
}

func (s Song) GetTitle() string {
	return s.Title
}

func (s Song) GetGenre() string {
	return s.Genre
}

func (s Song) ToString() string {
	var response string = "title: " + s.Title + "\ngenre: " + s.Genre
	return response
}

func New(title string, genre string) Song {
	s := Song{Title: title, Genre: genre}
	return s

}
