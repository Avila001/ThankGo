package main

/*

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	b.WriteByte('"')
	dur := int(time.Duration(d).Minutes())
	hours := dur / 60
	if hours != 0 {
		fmt.Fprintf(b, "%dh", hours)
	}
	mins := dur % 60
	if mins != 0 {
		fmt.Fprintf(b, "%dm", mins)
	}
	b.WriteByte('"')
	return b.Bytes(), nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	b.WriteByte('"')
	for i := 0; i < int(r); i++ {
		b.WriteRune('★')
	}
	for i := 0; i < 5-int(r); i++ {
		b.WriteRune('☆')
	}
	b.WriteByte('"')
	return b.Bytes(), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	//for _, movie := range movies {
	//	movie.Duration.MarshalJSON()
	//	movie.Rating.MarshalJSON()
	//}
	var ind string
	for i := 0; i < indent; i++ {
		ind += " "
	}
	var str []byte
	if indent == 0 {
		str, _ = json.Marshal(movies)
	} else {
		str, _ = json.MarshalIndent(movies, "", ind)
	}

	return string(str), nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	}

*/
