package main

/*
Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able
to see the 10 most recent tweets in the user's news feed.

Implement the Twitter class:

Twitter() Initializes your twitter object.
void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this
function will be made with a unique tweetId.
List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in
the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most
 recent to least recent.
void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID
followeeId.
*/

type User struct {
	Followings map[int]struct{}
}

type Feeds struct {
	UserID  int
	TweetId int
}

type Twitter struct {
	Users map[int]User
	Feeds []Feeds
}

func Constructor() Twitter {
	return Twitter{
		Users: map[int]User{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Feeds = append(this.Feeds, Feeds{
		UserID:  userId,
		TweetId: tweetId,
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	response := []int{}
	for i := len(this.Feeds) - 1; i >= 0; i-- {
		if _, ok := this.Users[userId].Followings[this.Feeds[i].UserID]; ok || this.Feeds[i].UserID == userId {
			response = append(response, this.Feeds[i].TweetId)
			if len(response) >= 10 {
				return response
			}
		}
	}

	return response
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	user, ok := this.Users[followerId]
	if !ok || user.Followings == nil {
		this.Users[followerId] = User{
			Followings: map[int]struct{}{},
		}
	}

	this.Users[followerId].Followings[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.Users[followerId].Followings, followeeId)
}
