section .text
global strstr

strstr:
    cmp     byte [rsi], 0      ; Check if needle is empty
    je      .return_haystack   ; Jump to return_haystack if true
    mov     r8, rdi            ; Save original haystack pointer
    mov     r9, rsi            ; Save original heedle pointer
.first_loop:
    mov     al, byte [r8]      ; al = curr char from haystack
    test    al, al             ; Check if curr char is null terminator
    jz      .not_found         ; Jump to not_found if true
    mov     r9, rsi            ; Reset needle to start
    mov     r10, r8            ; Create temporary haystack pointer for comparison

.second_loop:
    mov     cl, byte [r9]      ; cl = curr char from needle
    test    cl, cl             ; Check if curr char is null terminator
    jz      .found             ; Jump to found if true
    mov     al, byte [r10]     ; al = curr char from haystack
    cmp     al, cl             ; Cmp char from needle and char from haystack 
    jne     .advance           ; Jump to advance if not found
    inc     r10                ; r10++ (move to next haystack char)
    inc     r9                 ; r9++ (move to next needle char)
    jmp     .second_loop       ; Repeat this loop

.advance:
    inc     r8                 ; r8++ (move to next haystack char)
    jmp     .first_loop        ; Jump to first_loop

.found:
    mov     rax, r8            ; Return curr haystack pos (r8 point to start of match)
    ret

.return_haystack:
    mov     rax, rdi           ; Return haystack if needle is empty
    ret

.not_found:
    xor     rax, rax           ; Return null (no match)
    ret
