package handlers

type VideoListData struct {
	id       string
	name     string
	duration int
}

func GetVideoList() [3]VideoListData {
	return [3]VideoListData{
		{"d290f1ee-6c54-4b01-90e6-d701748f0851", "Black Retrospetive Woman", 15},
		{"sldjfl34-dfgj-523k-jk34-5jk3j45klj34", "Go Rally TEASER-HD", 41},
		{"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345", "Танцор", 92},
	}
}

func GetVideoListDataById(id string) *VideoListData {
	data := GetVideoList()
	for i := 0; i < len(data); i++ {
		if data[i].id == id {
			return &data[i]
		}
	}
	return nil
}
