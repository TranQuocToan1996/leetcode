package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type foodHeap []*food

func (h foodHeap) Len() int { return len(h) }
func (h foodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].food < h[j].food
	}
	return h[i].rating > h[j].rating
}

func (h foodHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

func (h *foodHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	a := x.(*food)
	a.heapIndex = len(*h)
	*h = append(*h, a)
}

func (h *foodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type food struct {
	food    string
	cuisine string
	rating  int

	heapIndex int
}

type FoodRatings struct {
	foods    map[string]*food
	cuisines map[string]*foodHeap
}

// Rename when submit to leetcode
func NewFoodRating(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodRatings := FoodRatings{
		foods:    make(map[string]*food),
		cuisines: make(map[string]*foodHeap),
	}
	for i := 0; i < len(foods); i++ {
		//
	}
	return foodRatings
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	// Update rating by food
	cuisineRating, ok := this.cuisineRating[food]
	if !ok {
		return
	}
	cuisineRating.rating = newRating
	this.cuisineRating[food] = cuisineRating

	// Update highest rating
	if needUpdateFoodRating(this.foodRating[cuisineRating.cuisine].rating, newRating,
		this.foodRating[cuisineRating.cuisine].food, food) {
		this.foodRating[cuisineRating.cuisine] = FoodRating{
			food:   food,
			rating: newRating,
		}
	}
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return (*this.cuisines[cuisine])[0].food
}

func TestNewFoodRating(t *testing.T) {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}
	foodRatings := NewFoodRating(foods, cuisines, ratings)
	require.Equal(t, "kimchi", foodRatings.HighestRated("korean"))
	require.Equal(t, "ramen", foodRatings.HighestRated("japanese"))
	foodRatings.ChangeRating("sushi", 16)
	require.Equal(t, "sushi", foodRatings.HighestRated("japanese"))
	foodRatings.ChangeRating("ramen", 16)
	require.Equal(t, "ramen", foodRatings.HighestRated("japanese"))
}
