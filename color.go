/*
 *  Copyright (c) 2022-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package badges

// Color model
type Color struct {
	TitleBG   string
	TitleFont string
	DataBG    string
	DataFont  string
}

var (
	ColorPrimary   = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#0d6efd", DataFont: "#fff"}
	ColorSecondary = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#6c757d", DataFont: "#fff"}
	ColorSuccess   = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#198754", DataFont: "#fff"}
	ColorDanger    = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#dc3545", DataFont: "#fff"}
	ColorWarning   = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#ffc107", DataFont: "#000"}
	ColorInfo      = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#0dcaf0", DataFont: "#000"}
	ColorLight     = Color{TitleBG: "#555", TitleFont: "#fff", DataBG: "#f8f9fa", DataFont: "#000"}
)
