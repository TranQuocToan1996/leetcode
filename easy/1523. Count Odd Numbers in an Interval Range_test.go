package easy

func countOdds(low int, high int) int {
    var result int
    var countNum int = high - low + 1
    // Truong hop duy nhat so le nhieu hon so chan
    if (countNum % 2 == 1) && (high % 2 == 1) {
        result = countNum / 2 + 1
    // Cac truong hop con lai so le bang so chan	
    } else {result = countNum / 2}
    
    return result
}

// Do phuc tap thoi gian: O(1) do tinh truc tiep dap an, bat chap khoang low high
// Do phuc tap khong gian: O(1) do ko phat sinh them khong gian khi chay