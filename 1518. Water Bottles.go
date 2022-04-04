func numWaterBottles(numBottles  int, numExchange int) int {
    var drinkBottle int
    var emptyBottle int
    for numBottles > 0 {
        drinkBottle += numBottles
        emptyBottle += numBottles
        // Exchange (because of int so the decimal part with be passed)
        numBottles = emptyBottle / numExchange
        // Emptybottle that can exchange
        emptyBottle = emptyBottle - numBottles * numExchange
        //fmt.Println(drinkBottle, emptyBottle, numBottles)
    }
    
    return drinkBottle
    
}