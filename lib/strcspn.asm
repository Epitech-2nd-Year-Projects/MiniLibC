section .text
global strcspn

strcspn:
    xor     rax, rax        ; Init rax to null to return if needed
    mov     r8, rsi         ; Save original reject pointer

.first_loop:
    mov     dl, byte [rdi]  ; dl = curr char from s
    test    dl, dl          ; Check if curr char is null terminator
    jz      .end            ; Jump to end if true
    mov     rsi, r8         ; Reset reject pointer to beginning

.second_loop:
    mov     cl, byte [rsi]  ; cl = current character from reject
    test    cl, cl          ; Check if curr char is null terminator
    jz      .next_char      ; Jump to next char if end of reject
    cmp     dl, cl          ; Compare s character (dl) with reject character (cl)
    je      .end            ; Jump to end if char match
    inc     rsi             ; rsi++ increment reject pointer
    jmp     .second_loop    ; Repeat second loop

.next_char:
    inc     rdi             ; rdi++ (move to next char)
    inc     rax             ; rax++ increment valid char counter
    jmp     .first_loop     ; Repeat first loop

.end:
    ret

