section .text
global index

; Literally strchr

index:
.loop:
    mov     al, byte [rdi]   ; al = curr char
    cmp     al, sil          ; Cmp char with target (sil)
    je      .found           ; Jump to found if char match
    test    al, al           ; Check if curr char is null terminator
    jz      .not_found       ; Jump to not found if char is null terminator
    inc     rdi              ; curr -> next (rdi++)
    jmp     .loop            ; Repeat the loop

.found:
    mov     rax, rdi         ; Return address of matching char
    ret

.not_found:
    xor     rax, rax         ; Return null
    ret
