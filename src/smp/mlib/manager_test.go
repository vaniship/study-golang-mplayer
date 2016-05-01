package library

import "testing"

func TestOpts(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManger falied.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManger falied. Not empty.")
	}
	m0 := &MusicEntry{
		"1",
		"My Heart Will Go On",
		"Celion Dion",
		"http://",
		"MP3",
	}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManger.Add() falied. ")
	}
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManger.Find() falied. ")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManger.Get() falied. ", err)
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManger.Remove() falied. ")
	}
}
