package alignment

import "github.com/ggeorgiev/instant-chess/src/square"

type Vector uint8

type Vectors [square.Number][square.Number]Vector

func (v Vector) IsLeaniar() bool {
	return v == File || v == Rank
}

func (v Vector) IsDiagonalic() bool {
	return v == Diagonal || v == CounterDiagonal
}

const (
	NoVector Vector = iota
	File
	Rank
	Diagonal
	CounterDiagonal
)

const (
	n = NoVector
	f = File
	r = Rank
	d = Diagonal
	c = CounterDiagonal
)

var SquareVectors = Vectors{
	{
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
		f, n, d, n, n, n, n, n,
		f, n, n, d, n, n, n, n,
		f, n, n, n, d, n, n, n,
		f, n, n, n, n, d, n, n,
		f, n, n, n, n, n, d, n,
		f, n, n, n, n, n, n, d,
	},
	{
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
		n, f, n, d, n, n, n, n,
		n, f, n, n, d, n, n, n,
		n, f, n, n, n, d, n, n,
		n, f, n, n, n, n, d, n,
		n, f, n, n, n, n, n, d,
		n, f, n, n, n, n, n, n,
	},
	{
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
		c, n, f, n, d, n, n, n,
		n, n, f, n, n, d, n, n,
		n, n, f, n, n, n, d, n,
		n, n, f, n, n, n, n, d,
		n, n, f, n, n, n, n, n,
		n, n, f, n, n, n, n, n,
	},
	{
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
		n, c, n, f, n, d, n, n,
		c, n, n, f, n, n, d, n,
		n, n, n, f, n, n, n, d,
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, n,
	},
	{
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
		n, n, c, n, f, n, d, n,
		n, c, n, n, f, n, n, d,
		c, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
	},
	{
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
		n, n, n, c, n, f, n, d,
		n, n, c, n, n, f, n, n,
		n, c, n, n, n, f, n, n,
		c, n, n, n, n, f, n, n,
		n, n, n, n, n, f, n, n,
		n, n, n, n, n, f, n, n,
	},
	{
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
		n, n, n, n, c, n, f, n,
		n, n, n, c, n, n, f, n,
		n, n, c, n, n, n, f, n,
		n, c, n, n, n, n, f, n,
		c, n, n, n, n, n, f, n,
		n, n, n, n, n, n, f, n,
	},
	{
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
		n, n, n, n, n, c, n, f,
		n, n, n, n, c, n, n, f,
		n, n, n, c, n, n, n, f,
		n, n, c, n, n, n, n, f,
		n, c, n, n, n, n, n, f,
		c, n, n, n, n, n, n, f,
	},
	{
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
		f, n, d, n, n, n, n, n,
		f, n, n, d, n, n, n, n,
		f, n, n, n, d, n, n, n,
		f, n, n, n, n, d, n, n,
		f, n, n, n, n, n, d, n,
	},
	{
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
		n, f, n, d, n, n, n, n,
		n, f, n, n, d, n, n, n,
		n, f, n, n, n, d, n, n,
		n, f, n, n, n, n, d, n,
		n, f, n, n, n, n, n, d,
	},
	{
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
		c, n, f, n, d, n, n, n,
		n, n, f, n, n, d, n, n,
		n, n, f, n, n, n, d, n,
		n, n, f, n, n, n, n, d,
		n, n, f, n, n, n, n, n,
	},
	{
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
		n, c, n, f, n, d, n, n,
		c, n, n, f, n, n, d, n,
		n, n, n, f, n, n, n, d,
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, n,
	},
	{
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
		n, n, c, n, f, n, d, n,
		n, c, n, n, f, n, n, d,
		c, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
	},
	{
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
		n, n, n, c, n, f, n, d,
		n, n, c, n, n, f, n, n,
		n, c, n, n, n, f, n, n,
		c, n, n, n, n, f, n, n,
		n, n, n, n, n, f, n, n,
	},
	{
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
		n, n, n, n, c, n, f, n,
		n, n, n, c, n, n, f, n,
		n, n, c, n, n, n, f, n,
		n, c, n, n, n, n, f, n,
		c, n, n, n, n, n, f, n,
	},
	{
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
		n, n, n, n, n, c, n, f,
		n, n, n, n, c, n, n, f,
		n, n, n, c, n, n, n, f,
		n, n, c, n, n, n, n, f,
		n, c, n, n, n, n, n, f,
	},
	{
		f, n, c, n, n, n, n, n,
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
		f, n, d, n, n, n, n, n,
		f, n, n, d, n, n, n, n,
		f, n, n, n, d, n, n, n,
		f, n, n, n, n, d, n, n,
	},
	{
		n, f, n, c, n, n, n, n,
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
		n, f, n, d, n, n, n, n,
		n, f, n, n, d, n, n, n,
		n, f, n, n, n, d, n, n,
		n, f, n, n, n, n, d, n,
	},
	{
		d, n, f, n, c, n, n, n,
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
		c, n, f, n, d, n, n, n,
		n, n, f, n, n, d, n, n,
		n, n, f, n, n, n, d, n,
		n, n, f, n, n, n, n, d,
	},
	{
		n, d, n, f, n, c, n, n,
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
		n, c, n, f, n, d, n, n,
		c, n, n, f, n, n, d, n,
		n, n, n, f, n, n, n, d,
		n, n, n, f, n, n, n, n,
	},
	{
		n, n, d, n, f, n, c, n,
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
		n, n, c, n, f, n, d, n,
		n, c, n, n, f, n, n, d,
		c, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
	},
	{
		n, n, n, d, n, f, n, c,
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
		n, n, n, c, n, f, n, d,
		n, n, c, n, n, f, n, n,
		n, c, n, n, n, f, n, n,
		c, n, n, n, n, f, n, n,
	},
	{
		n, n, n, n, d, n, f, n,
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
		n, n, n, n, c, n, f, n,
		n, n, n, c, n, n, f, n,
		n, n, c, n, n, n, f, n,
		n, c, n, n, n, n, f, n,
	},
	{
		n, n, n, n, n, d, n, f,
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
		n, n, n, n, n, c, n, f,
		n, n, n, n, c, n, n, f,
		n, n, n, c, n, n, n, f,
		n, n, c, n, n, n, n, f,
	},
	{
		f, n, n, c, n, n, n, n,
		f, n, c, n, n, n, n, n,
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
		f, n, d, n, n, n, n, n,
		f, n, n, d, n, n, n, n,
		f, n, n, n, d, n, n, n,
	},
	{
		n, f, n, n, c, n, n, n,
		n, f, n, c, n, n, n, n,
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
		n, f, n, d, n, n, n, n,
		n, f, n, n, d, n, n, n,
		n, f, n, n, n, d, n, n,
	},
	{
		n, n, f, n, n, c, n, n,
		d, n, f, n, c, n, n, n,
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
		c, n, f, n, d, n, n, n,
		n, n, f, n, n, d, n, n,
		n, n, f, n, n, n, d, n,
	},
	{
		d, n, n, f, n, n, c, n,
		n, d, n, f, n, c, n, n,
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
		n, c, n, f, n, d, n, n,
		c, n, n, f, n, n, d, n,
		n, n, n, f, n, n, n, d,
	},
	{
		n, d, n, n, f, n, n, c,
		n, n, d, n, f, n, c, n,
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
		n, n, c, n, f, n, d, n,
		n, c, n, n, f, n, n, d,
		c, n, n, n, f, n, n, n,
	},
	{
		n, n, d, n, n, f, n, n,
		n, n, n, d, n, f, n, c,
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
		n, n, n, c, n, f, n, d,
		n, n, c, n, n, f, n, n,
		n, c, n, n, n, f, n, n,
	},
	{
		n, n, n, d, n, n, f, n,
		n, n, n, n, d, n, f, n,
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
		n, n, n, n, c, n, f, n,
		n, n, n, c, n, n, f, n,
		n, n, c, n, n, n, f, n,
	},
	{
		n, n, n, n, d, n, n, f,
		n, n, n, n, n, d, n, f,
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
		n, n, n, n, n, c, n, f,
		n, n, n, n, c, n, n, f,
		n, n, n, c, n, n, n, f,
	},
	{
		f, n, n, n, c, n, n, n,
		f, n, n, c, n, n, n, n,
		f, n, c, n, n, n, n, n,
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
		f, n, d, n, n, n, n, n,
		f, n, n, d, n, n, n, n,
	},
	{
		n, f, n, n, n, c, n, n,
		n, f, n, n, c, n, n, n,
		n, f, n, c, n, n, n, n,
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
		n, f, n, d, n, n, n, n,
		n, f, n, n, d, n, n, n,
	},
	{
		n, n, f, n, n, n, c, n,
		n, n, f, n, n, c, n, n,
		d, n, f, n, c, n, n, n,
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
		c, n, f, n, d, n, n, n,
		n, n, f, n, n, d, n, n,
	},
	{
		n, n, n, f, n, n, n, c,
		d, n, n, f, n, n, c, n,
		n, d, n, f, n, c, n, n,
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
		n, c, n, f, n, d, n, n,
		c, n, n, f, n, n, d, n,
	},
	{
		d, n, n, n, f, n, n, n,
		n, d, n, n, f, n, n, c,
		n, n, d, n, f, n, c, n,
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
		n, n, c, n, f, n, d, n,
		n, c, n, n, f, n, n, d,
	},
	{
		n, d, n, n, n, f, n, n,
		n, n, d, n, n, f, n, n,
		n, n, n, d, n, f, n, c,
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
		n, n, n, c, n, f, n, d,
		n, n, c, n, n, f, n, n,
	},
	{
		n, n, d, n, n, n, f, n,
		n, n, n, d, n, n, f, n,
		n, n, n, n, d, n, f, n,
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
		n, n, n, n, c, n, f, n,
		n, n, n, c, n, n, f, n,
	},
	{
		n, n, n, d, n, n, n, f,
		n, n, n, n, d, n, n, f,
		n, n, n, n, n, d, n, f,
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
		n, n, n, n, n, c, n, f,
		n, n, n, n, c, n, n, f,
	},
	{
		f, n, n, n, n, c, n, n,
		f, n, n, n, c, n, n, n,
		f, n, n, c, n, n, n, n,
		f, n, c, n, n, n, n, n,
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
		f, n, d, n, n, n, n, n,
	},
	{
		n, f, n, n, n, n, c, n,
		n, f, n, n, n, c, n, n,
		n, f, n, n, c, n, n, n,
		n, f, n, c, n, n, n, n,
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
		n, f, n, d, n, n, n, n,
	},
	{
		n, n, f, n, n, n, n, c,
		n, n, f, n, n, n, c, n,
		n, n, f, n, n, c, n, n,
		d, n, f, n, c, n, n, n,
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
		c, n, f, n, d, n, n, n,
	},
	{
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, c,
		d, n, n, f, n, n, c, n,
		n, d, n, f, n, c, n, n,
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
		n, c, n, f, n, d, n, n,
	},
	{
		n, n, n, n, f, n, n, n,
		d, n, n, n, f, n, n, n,
		n, d, n, n, f, n, n, c,
		n, n, d, n, f, n, c, n,
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
		n, n, c, n, f, n, d, n,
	},
	{
		d, n, n, n, n, f, n, n,
		n, d, n, n, n, f, n, n,
		n, n, d, n, n, f, n, n,
		n, n, n, d, n, f, n, c,
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
		n, n, n, c, n, f, n, d,
	},
	{
		n, d, n, n, n, n, f, n,
		n, n, d, n, n, n, f, n,
		n, n, n, d, n, n, f, n,
		n, n, n, n, d, n, f, n,
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
		n, n, n, n, c, n, f, n,
	},
	{
		n, n, d, n, n, n, n, f,
		n, n, n, d, n, n, n, f,
		n, n, n, n, d, n, n, f,
		n, n, n, n, n, d, n, f,
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
		n, n, n, n, n, c, n, f,
	},
	{
		f, n, n, n, n, n, c, n,
		f, n, n, n, n, c, n, n,
		f, n, n, n, c, n, n, n,
		f, n, n, c, n, n, n, n,
		f, n, c, n, n, n, n, n,
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
		f, d, n, n, n, n, n, n,
	},
	{
		n, f, n, n, n, n, n, c,
		n, f, n, n, n, n, c, n,
		n, f, n, n, n, c, n, n,
		n, f, n, n, c, n, n, n,
		n, f, n, c, n, n, n, n,
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
		c, f, d, n, n, n, n, n,
	},
	{
		n, n, f, n, n, n, n, n,
		n, n, f, n, n, n, n, c,
		n, n, f, n, n, n, c, n,
		n, n, f, n, n, c, n, n,
		d, n, f, n, c, n, n, n,
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
		n, c, f, d, n, n, n, n,
	},
	{
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, c,
		d, n, n, f, n, n, c, n,
		n, d, n, f, n, c, n, n,
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
		n, n, c, f, d, n, n, n,
	},
	{
		n, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
		d, n, n, n, f, n, n, n,
		n, d, n, n, f, n, n, c,
		n, n, d, n, f, n, c, n,
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
		n, n, n, c, f, d, n, n,
	},
	{
		n, n, n, n, n, f, n, n,
		d, n, n, n, n, f, n, n,
		n, d, n, n, n, f, n, n,
		n, n, d, n, n, f, n, n,
		n, n, n, d, n, f, n, c,
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
		n, n, n, n, c, f, d, n,
	},
	{
		d, n, n, n, n, n, f, n,
		n, d, n, n, n, n, f, n,
		n, n, d, n, n, n, f, n,
		n, n, n, d, n, n, f, n,
		n, n, n, n, d, n, f, n,
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
		n, n, n, n, n, c, f, d,
	},
	{
		n, d, n, n, n, n, n, f,
		n, n, d, n, n, n, n, f,
		n, n, n, d, n, n, n, f,
		n, n, n, n, d, n, n, f,
		n, n, n, n, n, d, n, f,
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
		n, n, n, n, n, n, c, f,
	},
	{
		f, n, n, n, n, n, n, c,
		f, n, n, n, n, n, c, n,
		f, n, n, n, n, c, n, n,
		f, n, n, n, c, n, n, n,
		f, n, n, c, n, n, n, n,
		f, n, c, n, n, n, n, n,
		f, c, n, n, n, n, n, n,
		n, r, r, r, r, r, r, r,
	},
	{
		n, f, n, n, n, n, n, n,
		n, f, n, n, n, n, n, c,
		n, f, n, n, n, n, c, n,
		n, f, n, n, n, c, n, n,
		n, f, n, n, c, n, n, n,
		n, f, n, c, n, n, n, n,
		d, f, c, n, n, n, n, n,
		r, n, r, r, r, r, r, r,
	},
	{
		n, n, f, n, n, n, n, n,
		n, n, f, n, n, n, n, n,
		n, n, f, n, n, n, n, c,
		n, n, f, n, n, n, c, n,
		n, n, f, n, n, c, n, n,
		d, n, f, n, c, n, n, n,
		n, d, f, c, n, n, n, n,
		r, r, n, r, r, r, r, r,
	},
	{
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, n,
		n, n, n, f, n, n, n, c,
		d, n, n, f, n, n, c, n,
		n, d, n, f, n, c, n, n,
		n, n, d, f, c, n, n, n,
		r, r, r, n, r, r, r, r,
	},
	{
		n, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
		n, n, n, n, f, n, n, n,
		d, n, n, n, f, n, n, n,
		n, d, n, n, f, n, n, c,
		n, n, d, n, f, n, c, n,
		n, n, n, d, f, c, n, n,
		r, r, r, r, n, r, r, r,
	},
	{
		n, n, n, n, n, f, n, n,
		n, n, n, n, n, f, n, n,
		d, n, n, n, n, f, n, n,
		n, d, n, n, n, f, n, n,
		n, n, d, n, n, f, n, n,
		n, n, n, d, n, f, n, c,
		n, n, n, n, d, f, c, n,
		r, r, r, r, r, n, r, r,
	},
	{
		n, n, n, n, n, n, f, n,
		d, n, n, n, n, n, f, n,
		n, d, n, n, n, n, f, n,
		n, n, d, n, n, n, f, n,
		n, n, n, d, n, n, f, n,
		n, n, n, n, d, n, f, n,
		n, n, n, n, n, d, f, c,
		r, r, r, r, r, r, n, r,
	},
	{
		d, n, n, n, n, n, n, f,
		n, d, n, n, n, n, n, f,
		n, n, d, n, n, n, n, f,
		n, n, n, d, n, n, n, f,
		n, n, n, n, d, n, n, f,
		n, n, n, n, n, d, n, f,
		n, n, n, n, n, n, d, f,
		r, r, r, r, r, r, r, n,
	},
}
