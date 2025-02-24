section .text
global strcasecmp

strcasecmp:
    xor     rax, rax        ; Init rax to null (str1) 
    xor     rbx, rbx        ; Init rbx to null (str2)
.process_str1:
    mov     al, byte [rdi]  ; al = str1 char
    cmp     al, 'A'         ; Check if al > 'A'
    jb      .process_str2   ; Skip if not
    cmp     al, 'Z'         ; Check if al < 'Z'
    ja      .process_str2   ; Skip if not 
    add     al, 0x20        ; uppercase -> lowercase

.process_str2:
    mov     bl, byte [rsi]  ; bl = str2 char
    cmp     bl, 'A'         ; Check if bl > 'A'
    jb      .compare        ; Skip if not 
    cmp     bl, 'Z'         ; Check if bl < 'Z'
    ja      .compare        ; Skip if not
    add     bl, 0x20        ; uppercase -> lowercase

.compare:
    cmp     al, bl          ; Compare al and bl
    jne     .difference     ; If char al != char bl jump to difference
    test    al, al          ; Check if end of str1 is null terminator
    jz      .end            ; End program if end of str1 and str2
    inc     rdi             ; rdi++ move to next char (str1)
    inc     rsi             ; rsi++ move to next char (str2)
    jmp     .process_str1   ; Repeat the loop

.difference:
    sub     eax, ebx        ; al - bl (calculate char diff)
    movsx   rax, eax        ; Sign extend the result
    ret                     ; Return diff

.end:
    xor     eax, eax        ; Return 0 (bc string are equal)
    ret
