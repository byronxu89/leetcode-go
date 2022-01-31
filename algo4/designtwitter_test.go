package algo4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 5)

	ActualFeeds := twitter.GetNewsFeed(1)

	if !reflect.DeepEqual(ActualFeeds, []int{5}) {
		t.Errorf("#1 Expected feed is not same as actual feed")
	}
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)

	ActualFeeds = twitter.GetNewsFeed(1)

	if !reflect.DeepEqual(ActualFeeds, []int{6, 5}) {
		t.Errorf("#1 Expected feed is not same as actual feed")
	}
	fmt.Println(ActualFeeds)
	twitter.Unfollow(1, 2)
	ActualFeeds = twitter.GetNewsFeed(1)
	fmt.Println(ActualFeeds)
	if !reflect.DeepEqual(ActualFeeds, []int{5}) {
		t.Errorf("#1 Expected feed is not same as actual feed")
	}
}
