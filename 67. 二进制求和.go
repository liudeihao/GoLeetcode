package main

func addBinary(a string, b string) string {
    na, nb := len(a), len(b)
    if na < nb {
        a, b = b, a
        na, nb = nb, na
    }
    ans := make([]byte, na+1)
    carry := 0
    pc := na
    for i := 1; i <= nb; i++ {
        ia := na - i
        ib := nb - i
        va, vb := int(a[ia]-'0'), int(b[ib]-'0')
        curbit := (va + vb + carry) & 1
        carry = (va + vb + carry) >> 1
        ans[pc] = byte(curbit + '0')
        pc--
    }
    for ia := na - nb - 1; ia >= 0; ia-- {
        va := int(a[ia] - '0')
        curbit := (va + carry) & 1
        carry = (va + carry) >> 1
        ans[pc] = byte(curbit + '0')
        pc--
    }
    if carry == 1 {
        ans[0] = '1'
        return string(ans)
    } else {
        return string(ans[1:])
    }
}
