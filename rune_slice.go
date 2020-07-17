package main

// RuneSlice aliases a slice of runes to make them sortable
type RuneSlice []rune

// It satisfies the sorting interface
func (rs RuneSlice) Len() int           { return len(rs) }
func (rs RuneSlice) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs RuneSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
