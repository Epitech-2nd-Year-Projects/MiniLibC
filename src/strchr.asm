section .text
global strchr

strchr:
.loop:                     ; Create loop to check each character (not null)
    mov   al, byte [rdi]   ; al = curr char
    cmp   al, sil          ; Compare char with target (sil)
    je    .found           ; Jump to found if char match
    test  al, al           ; Check if curr char is null terminator
    jz    .not_found       ; Jump to not found if char is null terminator
    inc   rdi              ; curr -> next (with incr)
    jmp   .loop            ; Repeat the loop

.found:
    mov   rax, rdi         ; Return address of matching char
    ret

.not_found:
    xor   rax, rax         ; Return null
    ret
