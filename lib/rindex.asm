section .text
global rindex

rindex:
    xor     rax, rax        ; Init rax to null to return if needed

.loop:
    mov     dl, byte [rdi]  ; dl = curr char
    cmp     dl, sil         ; Compare char with target (sil)
    jne     .not_found      ; Jump to not_found if the char didnt match
    mov     rax, rdi        ; rax -> rdi (last match)

.not_found:
    test    dl, dl          ; Check if curr char is null terminator
    jz      .end            ; Jump to end if true
    inc     rdi             ; curr -> next (rdi++)
    jmp     .loop           ; Repeat the loop

.end:
    ret                     ; Return rax in any case (null if not found)
