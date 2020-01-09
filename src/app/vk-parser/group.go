package parser

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/mrKitikat/Vk-Parser-Service/src/app/models"
)

// GetUserSubscriptions gets vk-api method "/users.getSubscriptions" use vk.GET method.
// Returns responce current user's group's ids if success else error.
func (vk *VkParser) GetUserSubscriptions(userId int64) ([]int64, int, error) {
	params := "&user_id=" + fmt.Sprintf("%d", userId)
	resp, err := vk.GET("users.getSubscriptions", params)
	if err != nil {
		return nil, -1, err
	}

	groups := &models.Groups{}
	err = json.Unmarshal(resp, groups)
	if err != nil {
		return nil, -1, err
	}

	return groups.Data.G.Items, groups.Data.G.Count, nil
}

// GetUserGroups gets vk-api method "/users.getSubscriptions" use vk.GET method.
// Returns response current user's group ids if success else error.
func (vk *VkParser) GetMembers(groups []int64, params *models.IntersecReq) ([]int64, error) {
	var allMembers []int64
	var wg sync.WaitGroup
	for _, group := range groups {
		offset := 0
		NMembers := 0
		n := 0
		wg.Add(1)
		go func(group int64, n int) {
			for {
				// Scaned all members.
				if offset*n > NMembers {
					//fmt.Printf("All finded members in group %d : %d\n", g, len(allMembers))
					break
				}

				params := "&group_id=" + fmt.Sprintf("%d", group) + "&fields=sex,can_write_private_message&count=1000&offset=" + fmt.Sprintf("%d", offset)
				resp, err := vk.GET("groups.getMembers", params)
				if err != nil {
					fmt.Println(err, string(resp))
					time.Sleep(2 * time.Second)
					continue
				}
				offset += 1000
				// Unmarshall into members struct
				membrs := &models.Members{}
				err = json.Unmarshal(resp, membrs)
				if err != nil {
					fmt.Println(err, string(resp))
					time.Sleep(2 * time.Second)
					continue
				}
				// Count of members in the group
				NMembers = membrs.Data.Count
				if NMembers > 100000 {
					break
				}

				allMembers = append(allMembers, membrs.Data.Users...)
				n++
			}
			wg.Done()
		}(group, n)

	}

	wg.Wait()
	fmt.Println("Finish", len(allMembers))

	return intersectaion(allMembers, params.N), nil
}

// intersectaion finds for intersection of groups with a current minimum number(N) of occurrences.
// For example: [1, 2, 3] && [1, 2, 5, 7] && [1, 8, 7]
// IF N == 2: returns [1, 2, 7]
// IF N == 3: returns [1]
// Returns responce array of user's id if success else error.
func intersectaion(members []int64, N int) []int64 {

	uniqueUsers := make(map[int64]int)
	for _, id := range members {
		if _, ok := uniqueUsers[id]; ok {
			uniqueUsers[id]++
		} else {
			uniqueUsers[id] = 1
		}
	}

	// Start find intersection
	intersUsers := []int64{}
	for k, v := range uniqueUsers {
		if v >= N {
			intersUsers = append(intersUsers, k)
		}
	}

	return intersUsers
}
