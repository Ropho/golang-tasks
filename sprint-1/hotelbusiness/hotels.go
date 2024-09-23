//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	minCheckout := 9999
	maxCheckout := 0
	for _, guest := range guests {
		if guest.CheckOutDate > maxCheckout {
			maxCheckout = guest.CheckOutDate
		}
		if guest.CheckInDate < minCheckout {
			minCheckout = guest.CheckInDate
		}
	}
	if minCheckout == 9999 {
		return []Load{}
	}

	load := make([]Load, maxCheckout-minCheckout+1)

	for _, guest := range guests {
		for i := guest.CheckInDate; i < guest.CheckOutDate; i++ {
			load[i-minCheckout].GuestCount++
		}
	}

	ans := []Load{}
	for index, day := range load {
		if index == 0 || day.GuestCount != load[index-1].GuestCount {
			ans = append(ans, Load{
				GuestCount: day.GuestCount,
				StartDate:  index + minCheckout,
			})
		}
	}

	return ans
}
