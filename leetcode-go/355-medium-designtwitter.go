package main

type Twitter struct {
	Users  map[int]map[int]bool
	Tweets []Tweet
}

type Tweet struct {
	UserId  int
	TweetId int
}

func Constructor() Twitter {
	t := Twitter{
		Users:  make(map[int]map[int]bool),
		Tweets: []Tweet{},
	}

	return t
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, userExists := this.Users[userId]; !userExists {
		this.Users[userId] = make(map[int]bool)
	}

	newTweet := Tweet{
		UserId:  userId,
		TweetId: tweetId,
	}

	this.Tweets = append(this.Tweets, newTweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := []int{}

	if _, userExists := this.Users[userId]; !userExists {
		return tweets
	}

	for i := len(this.Tweets) - 1; i >= 0; i-- {
		if len(tweets) == 10 {
			return tweets
		}

		if this.Tweets[i].UserId == userId {
			tweets = append(tweets, this.Tweets[i].TweetId)
			continue
		}

		if _, followeeExists := this.Users[userId][this.Tweets[i].UserId]; followeeExists {
			tweets = append(tweets, this.Tweets[i].TweetId)
			continue
		}
	}

	return tweets

}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.Users[followerId] == nil {
		this.Users[followerId] = make(map[int]bool)
	}

	this.Users[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.Users[followerId] == nil {
		return
	}

	if _, followeeExists := this.Users[followerId][followeeId]; followeeExists {
		delete(this.Users[followerId], followeeId)
	}

}
