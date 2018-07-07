package daily_digest

import "strconv"

type Subscription struct {
	ID     string
	Source *Source
	user   *User
}

var subscriptions []Subscription

func init() {
	subscriptions = make([]Subscription, 0)
}

func CreateSubscription(user *User, source *Source) *Subscription {
	sub := Subscription{
		ID:     strconv.Itoa(len(subscriptions) + 1),
		Source: source,
		user:   user,
	}
	user.subscriptions = append(user.subscriptions, &sub)
	subscriptions = append(subscriptions, sub)
	return &sub
}

func ListAllSubscriptions(user *User) []Subscription {
	var subs []Subscription
	for _, s := range user.subscriptions {
		subs = append(subs, *s)
	}
	return subs
}
