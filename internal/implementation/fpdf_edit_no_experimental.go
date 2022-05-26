//go:build !pdfium_experimental
// +build !pdfium_experimental

package implementation

import (
	pdfium_errors "github.com/klippa-app/go-pdfium/errors"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

// FPDFPage_RemoveObject removes an object from a page.
// Ownership is transferred to the caller. Call FPDFPageObj_Destroy() to free
// it.
// Experimental API.
func (p *PdfiumImplementation) FPDFPage_RemoveObject(request *requests.FPDFPage_RemoveObject) (*responses.FPDFPage_RemoveObject, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_GetMatrix returns the transform matrix of a page object.
// The matrix is composed as:
//   |a c e|
//   |b d f|
// and can be used to scale, rotate, shear and translate the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_GetMatrix(request *requests.FPDFPageObj_GetMatrix) (*responses.FPDFPageObj_GetMatrix, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_SetMatrix sets the transform matrix on a page object.
// The matrix is composed as:
//   |a c e|
//   |b d f|
// and can be used to scale, rotate, shear and translate the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_SetMatrix(request *requests.FPDFPageObj_SetMatrix) (*responses.FPDFPageObj_SetMatrix, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_CountMarks returns the count of content marks in a page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_CountMarks(request *requests.FPDFPageObj_CountMarks) (*responses.FPDFPageObj_CountMarks, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_GetMark returns the content mark of a page object at the given index.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_GetMark(request *requests.FPDFPageObj_GetMark) (*responses.FPDFPageObj_GetMark, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_AddMark adds a new content mark to the given page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_AddMark(request *requests.FPDFPageObj_AddMark) (*responses.FPDFPageObj_AddMark, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_RemoveMark removes the given content mark from the given page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_RemoveMark(request *requests.FPDFPageObj_RemoveMark) (*responses.FPDFPageObj_RemoveMark, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_GetName returns the name of a content mark.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_GetName(request *requests.FPDFPageObjMark_GetName) (*responses.FPDFPageObjMark_GetName, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_CountParams returns the number of key/value pair parameters in the given mark.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_CountParams(request *requests.FPDFPageObjMark_CountParams) (*responses.FPDFPageObjMark_CountParams, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_GetParamKey returns the key of a property in a content mark.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_GetParamKey(request *requests.FPDFPageObjMark_GetParamKey) (*responses.FPDFPageObjMark_GetParamKey, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_GetParamValueType returns the type of the value of a property in a content mark by key.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_GetParamValueType(request *requests.FPDFPageObjMark_GetParamValueType) (*responses.FPDFPageObjMark_GetParamValueType, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_GetParamIntValue returns the value of a number property in a content mark by key as int.
// FPDFPageObjMark_GetParamValueType() should have returned FPDF_OBJECT_NUMBER
// for this property.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_GetParamIntValue(request *requests.FPDFPageObjMark_GetParamIntValue) (*responses.FPDFPageObjMark_GetParamIntValue, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_GetParamStringValue returns the value of a string property in a content mark by key.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_GetParamStringValue(request *requests.FPDFPageObjMark_GetParamStringValue) (*responses.FPDFPageObjMark_GetParamStringValue, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_GetParamBlobValue returns the value of a blob property in a content mark by key.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_GetParamBlobValue(request *requests.FPDFPageObjMark_GetParamBlobValue) (*responses.FPDFPageObjMark_GetParamBlobValue, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_SetIntParam sets the value of an int property in a content mark by key. If a parameter
// with the given key exists, its value is set to the given value. Otherwise, it is added as
// a new parameter.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_SetIntParam(request *requests.FPDFPageObjMark_SetIntParam) (*responses.FPDFPageObjMark_SetIntParam, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_SetStringParam sets the value of a string property in a content mark by key. If a parameter
// with the given key exists, its value is set to the given value. Otherwise, it is added as
// a new parameter.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_SetStringParam(request *requests.FPDFPageObjMark_SetStringParam) (*responses.FPDFPageObjMark_SetStringParam, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_SetBlobParam sets the value of a blob property in a content mark by key. If a parameter
// with the given key exists, its value is set to the given value. Otherwise, it is added as
// a new parameter.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_SetBlobParam(request *requests.FPDFPageObjMark_SetBlobParam) (*responses.FPDFPageObjMark_SetBlobParam, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObjMark_RemoveParam removes a property from a content mark by key.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObjMark_RemoveParam(request *requests.FPDFPageObjMark_RemoveParam) (*responses.FPDFPageObjMark_RemoveParam, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFImageObj_GetRenderedBitmap returns a bitmap rasterization of the given image object that takes the image mask and
// image matrix into account. To render correctly, the caller must provide the
// document associated with the image object. If there is a page associated
// with the image object the caller should provide that as well.
// The returned bitmap will be owned by the caller, and FPDFBitmap_Destroy()
// must be called on the returned bitmap when it is no longer needed.
// Experimental API.
func (p *PdfiumImplementation) FPDFImageObj_GetRenderedBitmap(request *requests.FPDFImageObj_GetRenderedBitmap) (*responses.FPDFImageObj_GetRenderedBitmap, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_GetDashPhase returns the line dash phase of the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_GetDashPhase(request *requests.FPDFPageObj_GetDashPhase) (*responses.FPDFPageObj_GetDashPhase, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_SetDashPhase sets the line dash phase of the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_SetDashPhase(request *requests.FPDFPageObj_SetDashPhase) (*responses.FPDFPageObj_SetDashPhase, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_GetDashCount returns the line dash array size of the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_GetDashCount(request *requests.FPDFPageObj_GetDashCount) (*responses.FPDFPageObj_GetDashCount, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_GetDashArray returns the line dash array of the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_GetDashArray(request *requests.FPDFPageObj_GetDashArray) (*responses.FPDFPageObj_GetDashArray, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFPageObj_SetDashArray sets the line dash array of the page object.
// Experimental API.
func (p *PdfiumImplementation) FPDFPageObj_SetDashArray(request *requests.FPDFPageObj_SetDashArray) (*responses.FPDFPageObj_SetDashArray, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFText_LoadStandardFont loads one of the standard 14 fonts per PDF spec 1.7 page 416. The preferred
// way of using font style is using a dash to separate the name from the style,
// for example 'Helvetica-BoldItalic'.
// The loaded font can be closed using FPDFFont_Close.
// Experimental API.
func (p *PdfiumImplementation) FPDFText_LoadStandardFont(request *requests.FPDFText_LoadStandardFont) (*responses.FPDFText_LoadStandardFont, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFTextObj_SetTextRenderMode sets the text rendering mode of a text object.
// Experimental API.
func (p *PdfiumImplementation) FPDFTextObj_SetTextRenderMode(request *requests.FPDFTextObj_SetTextRenderMode) (*responses.FPDFTextObj_SetTextRenderMode, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFTextObj_GetFont returns the font of a text object.
// Experimental API.
func (p *PdfiumImplementation) FPDFTextObj_GetFont(request *requests.FPDFTextObj_GetFont) (*responses.FPDFTextObj_GetFont, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetFontName returns the font name of a font.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetFontName(request *requests.FPDFFont_GetFontName) (*responses.FPDFFont_GetFontName, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetFontData returns the decoded data from the given font.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetFontData(request *requests.FPDFFont_GetFontData) (*responses.FPDFFont_GetFontData, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetIsEmbedded returns whether the given font is embedded or not.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetIsEmbedded(request *requests.FPDFFont_GetIsEmbedded) (*responses.FPDFFont_GetIsEmbedded, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetFlags returns the descriptor flags of a font.
// Returns the bit flags specifying various characteristics of the font as
// defined in ISO 32000-1:2008, table 123.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetFlags(request *requests.FPDFFont_GetFlags) (*responses.FPDFFont_GetFlags, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetWeight returns the font weight of a font.
// Typical values are 400 (normal) and 700 (bold).
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetWeight(request *requests.FPDFFont_GetWeight) (*responses.FPDFFont_GetWeight, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetItalicAngle returns the italic angle of a font.
// The italic angle of a font is defined as degrees counterclockwise
// from vertical. For a font that slopes to the right, this will be negative.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetItalicAngle(request *requests.FPDFFont_GetItalicAngle) (*responses.FPDFFont_GetItalicAngle, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetAscent returns ascent distance of a font.
// Ascent is the maximum distance in points above the baseline reached by the
// glyphs of the font. One point is 1/72 inch (around 0.3528 mm).
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetAscent(request *requests.FPDFFont_GetAscent) (*responses.FPDFFont_GetAscent, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetDescent returns the descent distance of a font.
// Descent is the maximum distance in points below the baseline reached by the
// glyphs of the font. One point is 1/72 inch (around 0.3528 mm).
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetDescent(request *requests.FPDFFont_GetDescent) (*responses.FPDFFont_GetDescent, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetGlyphWidth returns the width of a glyph in a font.
// Glyph width is the distance from the end of the prior glyph to the next
// glyph. This will be the vertical distance for vertical writing.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetGlyphWidth(request *requests.FPDFFont_GetGlyphWidth) (*responses.FPDFFont_GetGlyphWidth, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFFont_GetGlyphPath returns the glyphpath describing how to draw a font glyph.
// Experimental API.
func (p *PdfiumImplementation) FPDFFont_GetGlyphPath(request *requests.FPDFFont_GetGlyphPath) (*responses.FPDFFont_GetGlyphPath, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFGlyphPath_CountGlyphSegments returns the number of segments inside the given glyphpath.
// Experimental API.
func (p *PdfiumImplementation) FPDFGlyphPath_CountGlyphSegments(request *requests.FPDFGlyphPath_CountGlyphSegments) (*responses.FPDFGlyphPath_CountGlyphSegments, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFGlyphPath_GetGlyphPathSegment returns the segment in glyphpath at the given index.
// Experimental API.
func (p *PdfiumImplementation) FPDFGlyphPath_GetGlyphPathSegment(request *requests.FPDFGlyphPath_GetGlyphPathSegment) (*responses.FPDFGlyphPath_GetGlyphPathSegment, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}
