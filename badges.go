package badges

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Badges struct {
	face font.Face
	tmpl *template.Template
}

func New() (*Badges, error) {
	f, err := prepare()
	if err != nil {
		return nil, err
	}
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	return &Badges{face: f, tmpl: t}, err
}

func prepare() (font.Face, error) {
	b, err := base64.StdEncoding.DecodeString(meta)
	if err != nil {
		return nil, fmt.Errorf("badges: font corrupted: %w", err)
	}
	f, err := opentype.Parse(b)
	if err != nil {
		return nil, fmt.Errorf("badges: font corrupted: %w", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    9,
		DPI:     96,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, fmt.Errorf("badges: font corrupted: %w", err)
	}
	return face, nil
}

//Write generate badge and write it to io.Writer
func (v *Badges) Write(w io.Writer, color Color, title, data string) error {
	return v.generate(w, color, title, data)
}

//WriteResponse generate badge and write it to http.Response
func (v *Badges) WriteResponse(w http.ResponseWriter, color Color, title, data string) error {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Cache-Control", "max-age=86400, public")
	w.WriteHeader(http.StatusOK)
	return v.generate(w, color, title, data)
}

func (v *Badges) generate(w io.Writer, color Color, title, data string) error {
	m, ok := modelPool.Get().(*model)
	if !ok {
		m = &model{}
	}

	m.Title = title
	m.Data = data
	m.TitleBG = color.TitleBG
	m.TitleFont = color.TitleFont
	m.DataBG = color.DataBG
	m.DataFont = color.DataFont

	m.TitleW, _ = v.bound(title)
	m.TextW, _ = v.bound(data)

	m.TitleL, m.TitleW = m.TitleW+2, m.TitleW+14
	m.TextL, m.TextW = m.TextW+2, m.TextW+14
	m.TitleX = m.TitleW/2 + 1
	m.TextX = m.TitleW + m.TextW/2 - 1
	m.D1, m.D2 = m.TitleW-3, m.TitleW
	m.FullW = m.TitleW + m.TextW

	err := v.tmpl.ExecuteTemplate(w, "tmpl", m)

	modelPool.Put(m)

	return err
}

func (v *Badges) bound(data string) (int, int) {
	b, _ := font.BoundString(v.face, data)

	with := b.Max.X.Round() - b.Min.X.Round()
	height := b.Max.Y.Round() - b.Min.Y.Round()

	return with, height
}
