section .text
global memset

memset:
    mov     r8, rdi       ; Start pointer -> r8
    mov     al, sil       ; byte value -> al (used by stosb)
    mov     rcx, rdx      ; rcx = loop counter (number of bytes)
    rep     stosb         ; Repeat (until rcx = 0) al -> rdi; rdi++; rcx--
    mov     rax, r8       ; Original pointer -> rax
    ret                   ; Return rax
