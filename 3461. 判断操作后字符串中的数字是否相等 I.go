package main

func hasSameDigits(s string) bool {
    bs := []byte(s)
    for len(bs) > 2 {
        nbs := make([]byte, 0)
        for i := 0; i < len(bs)-1; i++ {
            nb := (bs[i]+bs[i+1]-2*'0')%10 + '0'
            nbs = append(nbs, nb)
        }
        bs = nbs
    }
    return bs[0] == bs[1]
}
