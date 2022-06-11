package badges_test

import (
	"bytes"
	"testing"

	"github.com/deweppro/go-badges"
)

func TestBadges(t *testing.T) {
	mock := &bytes.Buffer{}
	bb, err := badges.New()
	if err != nil {
		t.Fatal(err)
	}

	err = bb.Write(mock, badges.ColorInfo, "123", "456")
	if err != nil {
		t.Fatal(err)
	}

	exp := `<svg xmlns="http://www.w3.org/2000/svg" width="67" height="20"><linearGradient id="a" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".2"/><stop offset="1" stop-opacity=".2"/></linearGradient><g shape-rendering="crispEdges"><rect width="33" height="20" rx="3" fill="#555" /><rect x="33" width="34" height="20" rx="3" fill="#0dcaf0" /></g><g shape-rendering="crispEdges"><rect x="30" width="3" height="20" fill="#555" /><rect x="33" width="3" height="20" fill="#0dcaf0" /></g><g shape-rendering="crispEdges"><rect rx="3" width="67" height="20" fill="url(#a)"/></g><g shape-rendering="crispEdges" fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" lengthAdjust="spacingAndGlyphs" font-size="11"><text x="17" y="14" fill="#010101" fill-opacity=".3" textLength="21">123</text><text x="17" y="14" fill="#fff" textLength="21">123</text><text x="49" y="14" fill="#010101" fill-opacity=".3" textLength="22">456</text><text x="49" y="14" fill="#000" textLength="22">456</text></g></svg>`
	result := mock.String()
	if result != exp {
		t.Fatal("invalid template:\n" + result)
	}
}

func BenchmarkBadges(b *testing.B) {
	mock := &mockNilWriter{}
	bb, err := badges.New()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err = bb.Write(mock, badges.ColorInfo, "123", "456")
			if err != nil {
				b.Fatal(err)
			}
		}
	})

}

type mockNilWriter struct{}

func (*mockNilWriter) Write(_ []byte) (int, error) {
	return 0, nil
}
