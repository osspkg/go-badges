package badges

import (
	"sync"
)

var modelPool = sync.Pool{New: func() interface{} { return &model{} }}

type model struct {
	FullW, TitleW, TextW int
	TitleX, TextX        int
	TitleL, TextL        int
	D1, D2               int

	TitleBG, TitleFont, DataBG, DataFont string

	Title, Data string
}

const tmpl = `<svg xmlns="http://www.w3.org/2000/svg" width="{{.FullW}}" height="20"><linearGradient id="a" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".2"/><stop offset="1" stop-opacity=".2"/></linearGradient><g shape-rendering="crispEdges"><rect width="{{.TitleW}}" height="20" rx="3" fill="{{.TitleBG}}" /><rect x="{{.TitleW}}" width="{{.TextW}}" height="20" rx="3" fill="{{.DataBG}}" /></g><g shape-rendering="crispEdges"><rect x="{{.D1}}" width="3" height="20" fill="{{.TitleBG}}" /><rect x="{{.D2}}" width="3" height="20" fill="{{.DataBG}}" /></g><g shape-rendering="crispEdges"><rect rx="3" width="{{.FullW}}" height="20" fill="url(#a)"/></g><g shape-rendering="crispEdges" fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" lengthAdjust="spacingAndGlyphs" font-size="11"><text x="{{.TitleX}}" y="14" fill="#010101" fill-opacity=".3" textLength="{{.TitleL}}">{{.Title}}</text><text x="{{.TitleX}}" y="14" fill="{{.TitleFont}}" textLength="{{.TitleL}}">{{.Title}}</text><text x="{{.TextX}}" y="14" fill="#010101" fill-opacity=".3" textLength="{{.TextL}}">{{.Data}}</text><text x="{{.TextX}}" y="14" fill="{{.DataFont}}" textLength="{{.TextL}}">{{.Data}}</text></g></svg>`
