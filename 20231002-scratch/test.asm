section.data
msg: db "Hello, world!", 10

section.text
global _start

_start:
    mov rax, 1
    mov rdi, 1
    mov rsi, msg
    mov rdx, 14
    syscall
    mov rax, 60
    xor rdi, rdi
    syscall
    ret

"Thanks Cody for teaching me assembly!"


