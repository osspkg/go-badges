package badges

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Badges struct {
	faces    sync.Pool
	models   sync.Pool
	template *template.Template
	font     *opentype.Font
}

func New() (*Badges, error) {
	var err error
	vv := &Badges{}

	if err = vv.defaultFont(); err != nil {
		return nil, err
	}

	if err = vv.defaultTemplate(); err != nil {
		return nil, err
	}

	vv.poolModel()
	vv.poolFace()

	return vv, nil
}

func (v *Badges) defaultTemplate() (err error) {
	v.template, err = template.New("tmpl").Parse(tmpl)
	return
}

func (v *Badges) defaultFont() error {
	b, err := base64.StdEncoding.DecodeString(meta)
	if err != nil {
		return fmt.Errorf("badges: font corrupted: %w", err)
	}
	v.font, err = opentype.Parse(b)
	if err != nil {
		return fmt.Errorf("badges: font corrupted: %w", err)
	}
	return nil
}

func (v *Badges) poolModel() {
	v.models = sync.Pool{New: func() interface{} { return &model{} }}
}

func (v *Badges) getModel(call func(m *model) error) error {
	m, ok := v.models.Get().(*model)
	if !ok {
		m = &model{}
	}
	err := call(m)
	v.models.Put(m)
	return err
}

type face struct {
	Face font.Face
	Err  error
}

func (v *Badges) poolFace() {
	v.faces = sync.Pool{New: func() interface{} {
		f, err := opentype.NewFace(v.font, &opentype.FaceOptions{
			Size:    9,
			DPI:     96,
			Hinting: font.HintingNone,
		})
		return &face{
			Face: f,
			Err:  err,
		}
	}}
}

func (v *Badges) getFace(call func(m font.Face) error) error {
	m, ok := v.faces.Get().(*face)
	if !ok {
		fmt.Println(m)
		return fmt.Errorf("badges: cant get font.Face")
	}
	if m.Err != nil {
		return fmt.Errorf("badges: cant get font.Face: %w", m.Err)
	}
	err := call(m.Face)
	v.faces.Put(m)
	return err
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
	return v.getModel(func(m *model) error {
		if err := v.getFace(func(ff font.Face) error {
			m.TitleW, _ = bound(ff, title)
			m.TextW, _ = bound(ff, data)
			return nil
		}); err != nil {
			return err
		}

		m.Title = title
		m.Data = data
		m.TitleBG = color.TitleBG
		m.TitleFont = color.TitleFont
		m.DataBG = color.DataBG
		m.DataFont = color.DataFont

		m.TitleL, m.TitleW = m.TitleW+2, m.TitleW+14
		m.TextL, m.TextW = m.TextW+2, m.TextW+14
		m.TitleX = m.TitleW/2 + 1
		m.TextX = m.TitleW + m.TextW/2 - 1
		m.D1, m.D2 = m.TitleW-3, m.TitleW
		m.FullW = m.TitleW + m.TextW

		return v.template.ExecuteTemplate(w, "tmpl", m)
	})
}

func bound(face font.Face, data string) (int, int) {
	b, _ := font.BoundString(face, data)

	with := b.Max.X.Round() - b.Min.X.Round()
	height := b.Max.Y.Round() - b.Min.Y.Round()

	return with, height
}
