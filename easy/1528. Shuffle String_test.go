package easy


func restoreString(s string, indices []int) string {
    newS := make([]byte, len(indices))
    for i,v:=range indices{
        newS[v] = s[i]
    }
    return string(newS)
}
// Thoi gian : n do phai duyet qua string
// Khong gian: n do phai tao slice byte