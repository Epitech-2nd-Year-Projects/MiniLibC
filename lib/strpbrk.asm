section .text
global strpbrk

strpbrk:
    xor     rax, rax             ; Init rax to null to return if needed
    mov     r8, rsi              ; Save original accept pointer
.first_loop:
    mov     al, byte [rdi]       ; al = curr char from s
    test    al, al               ; Check if curr char is null terminator
    jz      .not_found           ; Jump to not_found if true
    mov     rsi, r8              ; Reset rsi to the start of accept

.second_loop:
    mov     cl, byte [rsi]       ; cl = current character from accept
    test    cl, cl               ; Check if curr char is null terminator
    jz      .next_char           ; Jump to next char if end of reject
    cmp     al, cl               ; Compare s char (al) with accept char (cl)
    je      .found               ; If true jump to found
    inc     rsi                  ; rsi++ increment accept pointer
    jmp     .second_loop          ; Repeat first loop

.next_char:
    inc     rdi                  ; rdi++ (move to next char)
    jmp     .first_loop         ; Repeat second loop

.found:
    mov     rax, rdi             ; Set return pointer to current position in s (match found)
    ret

.not_found:
    ret
