package main

func findKthBit(n int, k int) byte {
    if k == 1 {
        return '0'
    }
    m := 1 << (n - 1)
    if k == m {
        return '1'
    } else if k > m {
        k = 2*m - k
        return invertByte(findKthBit(n-1, k))
    }
    return findKthBit(n-1, k)
}

func invertByte(b byte) byte {
    if b == '0' {
        return '1'
    }
    return '0'
}
