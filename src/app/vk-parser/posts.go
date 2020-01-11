package parser

// import (
// 	"github.com/mrKitikat/Vk-Parser-Service/src/app/models"
// 	"fmt"
// )

// // GetLikedPosts returns
// // Returns array of (userID)user's liked posts if success else error.
// func (vk *VkParser) GetLikedPosts(userID int64) error{
// 	posts, err := vk.GetPosts(userID)
// 	if err != nil {
// 		return nil
// 	}
// 	likedPosts := []models.Post
// 	for _, post := range posts {
// 		if like, err := vk.isLiked(userID, groupdID, itemID); err != nil {
			
// 		} else if like {
// 			likedPosts = append(likedPosts, itemID)
// 		}
// 	}

// 	return nil

// }

// func (vk *VkParser) GetPosts(userID int64) error {
// 	params := "&user_id=" + fmt.Sprintf("%d", userID) + "&type=post&owned_id=-" + fmt.Sprintf("%d", groupdID) + "&item_id=" + fmt.Sprintf("%d", itemID)
// 	resp, err := vk.GET("likes.isLiked", params)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (vk *VkParser) isLiked(userID, groupdID, itemID int64) (bool, error) {
// 	params := "&user_id=" + fmt.Sprintf("%d", userID) + "&type=post&owned_id=-" + fmt.Sprintf("%d", groupdID) + "&item_id=" + fmt.Sprintf("%d", itemID)
// 	resp, err := vk.GET("likes.isLiked", params)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
