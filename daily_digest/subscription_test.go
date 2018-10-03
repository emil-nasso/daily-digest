package daily_digest

import (
	"reflect"
	"testing"
)

func TestCreateSubscription(t *testing.T) {
	type args struct {
		user   *User
		source *Source
	}
	tests := []struct {
		name string
		args args
		want *Subscription
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSubscription(tt.args.user, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSubscription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllSubscriptions(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name string
		args args
		want []Subscription
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllSubscriptions(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllSubscriptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
