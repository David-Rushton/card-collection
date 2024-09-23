package deck

import "slices"

type Hand []Card

// Appends the first n cards from other to the hand, and returns it.
// When other contain fewer than n; we join the two hands.
func (h Hand) Append(other Hand, n int) Hand {
	if n <= 0 {
		return h
	}

	n = min(n, len(other))
	return slices.Concat(h, other[0:n])
}

// Appends the first n cards from other that match the predicate to the hand, and returns it.
// When other contain fewer than n matches; we append as many as we can.
func (h Hand) AppendWhen(other Hand, n int, predicate func(Card) bool) Hand {
	if n <= 0 {
		return h
	}

	result := h
	taken := 0
	for _, card := range other {
		if predicate(card) {
			result = append(result, card)

			taken++
			if taken >= n {
				return result
			}
		}
	}

	return result
}

// Takes the first n cards from the hand.
func (h Hand) Take(n int) Hand {
	if n > len(h) {
		return h
	}

	if n <= 0 {
		return Hand{}
	}

	return h[0:n]
}

// Takes the first n cards from the hand, that matche the predicate.
//
//   - n
//
// Returns the first n cards that match the predicate.
// If the hand does not contain n matches, we return as many as we can.
//
//   - predicate
//
// User supplied filter.
func (h Hand) TakeWhen(n int, predicate func(Card) bool) Hand {
	result := Hand{}

	if n <= 0 {
		return result
	}

	taken := 0
	for _, card := range h {
		if predicate(card) {
			result = append(result, card)

			taken++
			if taken >= n {
				return result
			}
		}
	}

	return result
}

// Returns a copy of the hand, sorted by rank.
// Suits are not considered.
// The outcome is stable.
func (h Hand) Sort() Hand {
	if len(h) <= 1 {
		return h
	}

	return mergeSort(h)
}

// A [merge sort] implementation.
//
// [merge sort]: https://en.wikipedia.org/wiki/Merge_sort
func mergeSort(h Hand) Hand {
	var left Hand
	var right Hand

	for i, card := range h {
		if i < len(h)/2 {
			left = append(left, card)
		} else {
			right = append(right, card)
		}
	}

	left = left.Sort()
	right = right.Sort()

	return merge(left, right)
}

// Merges left and right.
// Taking the lowest rank from the head of left/right on each iteration.
func merge(left, right Hand) Hand {
	var result Hand

	// Iterate until either left or right is depleted.
	// Always take lower of the two and append to result.
	for len(left) > 0 && len(right) > 0 {
		if rankOrder(left[0].Rank) <= rankOrder(right[0].Rank) {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Consume any remaining elements.
	// At most one of these for statements will be true.
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func rankOrder(r Rank) int {
	if r == Ace {
		return int(King) + 1
	}

	return int(r)
}
