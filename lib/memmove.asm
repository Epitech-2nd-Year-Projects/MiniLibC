section .text
global memmove

memmove:
    mov     r8, rdi                 ; Save dest pointer (to return value)
    test    rdx, rdx                ; Check if n = 0
    jz      .end                    ; Jump to end if true
    cmp     rdi, rsi                ; Check if dest = src
    je      .end                    ; Jump to end if same pointer
    mov     rax, rdi                ; rax = dest (overlap check)
    sub     rax, rsi                ; RAX = dest - src
    cmp     rax, rdx                ; Cmp (dest - src) and n (to detect overlap)
    jb      .backward_cpy           ; Jump to backward_cpy if (dest - src) < n (overlap)
.forward_cpy:
    mov     rcx, rdx                ; Counter = n
    cld                             ; Clear direction flag (rsi++, rdi++)
    rep     movsb                   ; Copy rcx bytes from rsi to rdi
    jmp     .end                    ; Jump to end

.backward_cpy:
    mov     rcx, rdx                ; Counter = n
    lea     rsi, [rsi + rdx - 1]    ; rsi = last src byte
    lea     rdi, [rdi + rdx - 1]    ; rdi = last dest byte
    std                             ; Set direction flag (rsi--, rdi--)
    rep     movsb                   ; Copy rcx bytes from rsi to rdi
    cld                             ; Clear direction flag (rsi++, rdi++)

.end:
    mov     rax, r8                 ; Return original dest pointer
    ret
