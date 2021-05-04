package robotname

import (
	"math/rand"
	"time"
)

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var nums = "0123456789"

type Robot string

var localSeen = map[string]bool{}

func (r *Robot) Name() (string, error) {
	if len(*r) == 5 {
		return string(*r), nil
	}

	name := make([]byte, 5)
	for {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 5; i++ {
			if i < 2 {
				name[i] = letters[rand.Intn(len(letters))]
			} else {
				name[i] = nums[rand.Intn(len(nums))]
			}
		}


		if _, found := localSeen[string(name)]; !found {
			break
		}
	}

	localSeen[string(name)] = true
	*r = Robot(name)

	return string(name), nil
}

func (r *Robot) Reset() {
	//delete(localSeen, string(*r))
	*r = ""
}
