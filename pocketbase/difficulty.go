package pocketbase

var DifficultyList = []string{"basic", "advanced", "expert", "master", "re_master"}
var BanquetDifficultyList = []string{"banquet", "banquet2"}

type DS map[string]float64
type Level map[string]string
type Cids map[string]int

// 将数组按难度映射为 map
func mapByDifficulty[T float64 | string | int](row []T, genre string) map[string]T {
	result := make(map[string]T)

	var difficultyList []string
	if genre == "宴会場" {
		difficultyList = BanquetDifficultyList
	} else {
		difficultyList = DifficultyList
	}
	for i := range row {
		difficulty := difficultyList[i]
		result[difficulty] = row[i]
	}
	return result
}

// type DS struct {
// 	Basic    float64 `json:"basic"`
// 	Advanced float64 `json:"advanced"`
// 	Expert   float64 `json:"expert"`
// 	Master   float64 `json:"master"`
// 	ReMaster float64 `json:"re_master"`
// 	Banquet  float64 `json:"banquet"`
// 	Banquet2 float64 `json:"banquet2"`
// }

// func (d *DS) FromDivingFish(ds []float64, genre string) {
// 	if genre == "宴会場" {
// 		d.Banquet = ds[0]
// 		if len(ds) > 1 {
// 			d.Banquet2 = ds[1]
// 		}
// 	} else {
// 		d.Basic = ds[0]
// 		d.Advanced = ds[1]
// 		d.Expert = ds[2]
// 		d.Master = ds[3]
// 		if len(ds) > 4 {
// 			d.ReMaster = ds[4]
// 		}
// 	}
// }

// type Level struct {
// 	Basic    string `json:"basic"`
// 	Advanced string `json:"advanced"`
// 	Expert   string `json:"expert"`
// 	Master   string `json:"master"`
// 	ReMaster string `json:"re_master"`
// 	Banquet  string `json:"banquet"`
// 	Banquet2 string `json:"banquet2"`
// }

// func (l *Level) FromDivingFish(level []string, genre string) {
// 	if genre == "宴会場" {
// 		l.Banquet = level[0]
// 		if len(level) > 1 {
// 			l.Banquet2 = level[1]
// 		}
// 	} else {
// 		l.Basic = level[0]
// 		l.Advanced = level[1]
// 		l.Expert = level[2]
// 		l.Master = level[3]
// 		if len(level) > 4 {
// 			l.ReMaster = level[4]
// 		}
// 	}
// }

// type Cids struct {
// 	Basic    int `json:"basic"`
// 	Advanced int `json:"advanced"`
// 	Expert   int `json:"expert"`
// 	Master   int `json:"master"`
// 	ReMaster int `json:"re_master"`
// 	Banquet  int `json:"banquet"`
// 	Banquet2 int `json:"banquet2"`
// }

// func (c *Cids) FromDivingFish(cids []int, genre string) {
// 	if genre == "宴会場" {
// 		c.Banquet = cids[0]
// 		if len(cids) > 1 {
// 			c.Banquet2 = cids[1]
// 		}
// 	} else {
// 		c.Basic = cids[0]
// 		c.Advanced = cids[1]
// 		c.Expert = cids[2]
// 		c.Master = cids[3]
// 		if len(cids) > 4 {
// 			c.ReMaster = cids[4]
// 		}
// 	}
// }
