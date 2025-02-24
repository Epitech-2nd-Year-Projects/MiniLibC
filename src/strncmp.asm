section .text
global strncmp

strncmp:
    xor     rax, rax        ; Init rax to null to return if needed 
    test    rdx, rdx        ; Check if n = 0
    jz      .end            ; If true return 0
.loop:
    mov     al, byte [rdi]  ; al = str1 char
    mov     bl, byte [rsi]  ; bl = str2 char
    cmp     al, bl          ; Compare al and bl
    jne     .difference     ; if char al != char bl jump to difference   
    test    al, al          ; Check if end of str1 is null terminator
    jz      .end            ; End program if end of str1 and str2
    dec     rdx             ; n--
    jz      .end            ; End program if n = 0
    inc     rdi             ; rdi++ move to next char (str1)
    inc     rsi             ; rsi++ move to next char (str2)
    jmp     .loop           ; Repeat the loop

.difference:
    sub     al, bl          ; al - bl (calculate char diff)
    movsx   rax, al         ; Sign extend the result
    ret                     ; Return diff

.end:
    xor     eax, eax        ; Return 0 (bc string are equal)
    ret
