package main

// StringSlice aliases a slice of letters to make them sortable
type StringSlice []string

// It satisfies the sorting interface
func (rs StringSlice) Len() int           { return len(rs) }
func (rs StringSlice) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs StringSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
