/*
 *  Copyright (c) 2022-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package badges

// Color describes the background and foreground colors of both badge parts.
type Color struct {
	TitleBG   string
	TitleFont string
	DataBG    string
	DataFont  string
}

const (
	colorTitleBackground = "#555"
	colorWhite           = "#fff"
	colorBlack           = "#000"
)

var (
	// ColorPrimary is the default primary badge color.
	ColorPrimary = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#0d6efd", DataFont: colorWhite}
	// ColorSecondary is the default secondary badge color.
	ColorSecondary = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#6c757d", DataFont: colorWhite}
	// ColorSuccess is the default success badge color.
	ColorSuccess = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#198754", DataFont: colorWhite}
	// ColorDanger is the default danger badge color.
	ColorDanger = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#dc3545", DataFont: colorWhite}
	// ColorWarning is the default warning badge color.
	ColorWarning = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#ffc107", DataFont: colorBlack}
	// ColorInfo is the default informational badge color.
	ColorInfo = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#0dcaf0", DataFont: colorBlack}
	// ColorLight is the default light badge color.
	ColorLight = Color{TitleBG: colorTitleBackground, TitleFont: colorWhite, DataBG: "#f8f9fa", DataFont: colorBlack}
)
