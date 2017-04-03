package po

type Inner struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TemplateData struct {
	First  Inner `json:"first"`
	Send   Inner `json:"send"`
	Text   Inner `json:"text"`
	Time   Inner `json:"time"`
	Remark Inner `json:"remark"`
}
