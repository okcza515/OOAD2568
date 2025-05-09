package model

type PathFile struct {
	Path     string `json:"path"`      // เช่น "/uploads/pdfs/1234.pdf"
	Filename string `json:"filename"`  // เช่น "1234.pdf"
	MimeType string `json:"mime_type"` // เช่น "application/pdf"
	Size     int64  `json:"size"`      // ขนาดไฟล์ใน bytes
}
