section .text
global ffs

ffs:
    xor     rax, rax        ; Init rax to null (return value)
    test    edi, edi        ; Check if input = 0
    jz      .end            ; Jump to end if true
    bsf     rax, edi        ; Find first set bit (0 based index) and store it to edi
    inc     rax             ; Convert to 1 based index (eax++)
.end:
    ret
