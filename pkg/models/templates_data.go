package models

// TemplatesData estructura que contiene todos los datos que van a ser enviados de los controladores a las plantillas
type TemplatesData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
