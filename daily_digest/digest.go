package daily_digest

// Daily

type Digest struct {
	Subscription Subscription
	Entries      []*Entry
}

func GetDigest(user *User, date string) []Digest {
	digests := make([]Digest, 0)
	subscriptions := ListAllSubscriptions(user)

	for i := range subscriptions {
		d := Digest{
			Subscription: subscriptions[i],
			Entries:      subscriptions[i].Source.EntriesForDate(date),
		}
		digests = append(digests, d)
	}

	return digests
}
