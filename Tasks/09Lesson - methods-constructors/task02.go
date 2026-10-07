package main

import "fmt"

type Song struct {
	Title    string
	Duration int
}

type Playlist struct {
	song []Song
}

func (p *Playlist) Add(s Song) {
	p.song = append(p.song, s)
}

func (p *Playlist) TotalDuration() int {
	var totDuration int
	for _, v := range p.song {
		totDuration += v.Duration
	}

	return totDuration
}

func (p *Playlist) Longest() Song {
	if len(p.song) == 0 {
		return Song{}
	}

	longest := p.song[0]
	for _, v := range p.song {
		if longest.Duration < v.Duration {
			longest = v
		}
	}

	return longest
}

func (p Playlist) String() string {
	var result string
	for _, v := range p.song {
		result += fmt.Sprintf("- %s (%d сек.)\n", v.Title, v.Duration)
	}
	return result
}

func main() {
	playlist := Playlist{}
	playlist.Add(Song{Title: "Bohemian Rhapsody", Duration: 354})
	playlist.Add(Song{Title: "Yesterday", Duration: 125})
	playlist.Add(Song{Title: "Hotel California", Duration: 391})
	playlist.Add(Song{Title: "Numb", Duration: 187})
	playlist.Add(Song{Title: "Stairway to Heaven", Duration: 482})

	totalDuration := playlist.TotalDuration()
	fmt.Printf("Общая длительность: %d секунд\n", totalDuration)

	longestSong := playlist.Longest()
	fmt.Printf("Самая длинная песня: %s %d секунд \n", longestSong.Title, longestSong.Duration)

	fmt.Println(playlist)
}
