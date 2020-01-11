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
func (vk *VkParser) GetMembers(groups []int64, params *models.IntersecReq) ([]models.User, error) {
	var wg sync.WaitGroup
	var allMembers []models.User
	for _, group := range groups {
		offset := 0
		NMembers := 0
		var members []models.User
		wg.Add(1)
		go func(group int64, NMembers, offset int, members []models.User, wg *sync.WaitGroup) {
			for {
				// Scaned all members.
				if offset > NMembers {
					fmt.Printf("Finined for group %d: %d\n", group, len(members))
					allMembers = append(allMembers, members...)
					break
				}

				req := "&group_id=" + fmt.Sprintf("%d", group) + "&fields=sex,can_write_private_message,photo_200&count=1000&offset=" + fmt.Sprintf("%d", offset)
				resp, err := vk.GET("groups.getMembers", req)
				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				offset += 1000
				// Unmarshall into members struct
				membrs := &models.Members{}
				err = json.Unmarshal(resp, membrs)
				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}

				// Count of members in the group
				NMembers = membrs.Data.Count
				if NMembers > 100000 {
					break
				}

				// Check another constraint
				var necessaryUsers []models.User
				for _, u := range membrs.Data.Users {
					if params.Sex != 0 {
						if !u.IsClosed && u.Sex == params.Sex {
							necessaryUsers = append(necessaryUsers, u)
						}
					} else {
						if !u.IsClosed {
							necessaryUsers = append(necessaryUsers, u)
						}
					}
				}

				members = append(members, necessaryUsers...)
			}
			wg.Done()
		}(group, NMembers, offset, members, &wg)
	}

	wg.Wait()
	fmt.Println("Finish", len(allMembers))
	// Result of intersectaion.
	allMembers = intersectaion(allMembers, params.N)
	fmt.Println("Finish intersectaion", len(allMembers))
	// Check if intersectaion isn't too small.
	if len(allMembers) > 10 {
		return checkParams(allMembers, params), nil
	}

	return allMembers, nil
}

// intersectaion finds for intersection of groups with a current minimum number(N) of occurrences.
// For example: [1, 2, 3] && [1, 2, 5, 7] && [1, 8, 7]
// IF N == 2: returns [1, 2, 7]
// IF N == 3: returns [1]
// Returns responce array of user's id if success else error.
func intersectaion(members []models.User, N int) []models.User {
	uniqueUsers := make(map[models.User]int)
	for _, id := range members {
		if _, ok := uniqueUsers[id]; ok {
			uniqueUsers[id]++
		} else {
			uniqueUsers[id] = 1
		}
	}

	// Start find intersection
	intersUsers := []models.User{}
	for k, v := range uniqueUsers {
		if v >= N {
			intersUsers = append(intersUsers, k)
		}
	}

	return intersUsers
}

// Can write message check.
func checkParams(members []models.User, params *models.IntersecReq) []models.User {
	var necessaryUsers []models.User
	for _, u := range members {
		if params.Message {
			if u.CanWrite == 1 {
				necessaryUsers = append(necessaryUsers, u)
			}
		}
	}

	return necessaryUsers
}
