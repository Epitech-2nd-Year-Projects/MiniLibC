section .text
global memcpy

memcpy:
    mov     rax, rdi      ; Save dest pointer (to return value)
    mov     rcx, rdx      ; Set counter to n bytes
    cld                   ; Clear direction flag 
    rep     movsb         ; Copy rcx bytes from rsi to rdi
    ret                   ; Return original dest pointer
