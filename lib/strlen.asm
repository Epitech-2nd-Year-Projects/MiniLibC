section .text
global strlen

strlen:
    xor rax, rax        ; Create counter and initialize it to 0

.loop:                  ; Create loop to check each character (not null)
    cmp byte [rdi + rax], 0 ; Check if current byte is null terminator
    je .end             ; Exit loop if yes
    inc rax             ; Else increment counter
    jmp .loop           ; Repeat the loop

.end:
    ret                 ; Return result
