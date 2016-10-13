package main

// Size info
type Size struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

//Inference info
type Result struct {
	Result_Label_1 string  `json:"result_label_1"`
	Result_Score_1 float32 `json:"result_score_1"`
	Result_Label_2 string  `json:"result_label_2"`
	Result_Score_2 float32 `json:"result_score_2"`
	Result_Label_3 string  `json:"result_label_3"`
	Result_Score_3 float32 `json:"result_score_3"`
}

// Image info
type Image struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Results Result `json:"results"`
	Resize  bool   `json:"resize"`
	Size    Size   `json:"size"`
}

type Sizes []Size
type Results []Result
type Images []Image
