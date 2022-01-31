package algo4

type User struct {
	UserId    int
	Following map[int]struct{}
}
type Tweet struct {
	TweetId int
	Author  int
}
type Twitter struct {
	Users  map[int]*User
	Tweets []Tweet
}

func Constructor() Twitter {
	return Twitter{
		Users:  map[int]*User{},
		Tweets: []Tweet{},
	}
}
func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := Tweet{TweetId: tweetId, Author: userId}
	this.Tweets = append(this.Tweets, tweet)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	user := this.GetUser(userId)
	feed := []int{}
	for i := len(this.Tweets) - 1; i >= 0 && len(feed) < 10; i-- {
		tweet := this.Tweets[i]
		_, exists := user.Following[tweet.Author]
		if tweet.Author == userId || exists {
			feed = append(feed, tweet.TweetId)
		}
	}
	return feed
}
func (this *Twitter) GetUser(userID int) *User {
	if user, ok := this.Users[userID]; ok {
		return user
	}

	user := &User{
		UserId:    userID,
		Following: map[int]struct{}{},
	}
	this.Users[userID] = user

	return user
}
func (this *Twitter) Follow(followerId int, followeeId int) {
	user := this.GetUser(followerId)
	user.Following[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	user := this.GetUser(followerId)
	delete(user.Following, followeeId)
}
