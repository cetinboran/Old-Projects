extern printf:proc
extern scanf:proc

.data
	help1 db "Enter two integers: ", 0
	help2 db "1-add, 2-sub, 3-multi, 4-divide (Only + in Divide)", 10, 0
	help3 db "Enter the choice: ", 0
	help4 db "Invalid choice. App Closing...", 0

	error1 db "You cannot divide by zero!", 0

	format1 db "%d %d", 0
	format2 db "Result: %d", 0
	format3 db "%d", 0
	format4 db "Division: %d, Remainder: %d", 0

	var1 QWORD 0
	var2 QWORD 0

	result QWORD 0
	remainder QWORD 0

	choice QWORD 0
.code

one:
	sub rsp, 28h
	mov rax, var1
	add rax, var2
	mov result, rax
	add rsp, 28h

	sub rsp, 28h
	lea rcx, format2
	mov rdx, result
	call printf
	add rsp, 28h

	mov rax, 1
	ret

two:
	sub rsp, 28h
	mov rax, var1
	sub rax, var2
	mov result, rax
	add rsp, 28h

	sub rsp, 28h
	lea rcx, format2
	mov rdx, result
	call printf
	add rsp, 28h

	mov rax, 2
	ret

three:
	sub rsp, 28h
	mov rax, var1
	mov rbx, var2
	mul rbx
	mov result, rax
	add rsp, 28h

	sub rsp, 28h
	lea rcx, format2
	mov rdx, result
	call printf
	add rsp, 28h

	mov rax, 0
	ret

four:
	cmp var2, 0
	je divideByZero

	sub rsp, 28h
	mov rax, var1
	xor rdx, rdx
	mov rbx, var2
	div rbx
	mov result, rax
	mov remainder, rdx
	add rsp, 28h

	sub rsp, 28h
	lea rcx, format4
	mov rdx, result
	mov r8, remainder
	call printf
	add rsp, 28h


	mov rax, result
	ret

final:
	sub rsp, 28h
	lea rcx, help4
	call printf
	add rsp, 28h

	mov rax, 5
	ret

divideByZero:
	sub rsp, 28h
	lea rcx, error1
	call printf
	add rsp, 28h

	mov rax, 6
	ret

main proc
	sub rsp, 28h
	lea rcx, help1
	call printf
	add rsp, 28h

	sub rsp, 28h
	lea rcx, format1
	lea rdx, var1
	lea r8, var2
	call scanf
	add rsp, 28h

	sub rsp, 28h
	lea rcx, help2
	call printf
	add rsp, 28h

	sub rsp, 28h
	lea rcx, help3
	call printf
	add rsp, 28h

	sub rsp, 28h
	lea rcx, format3
	lea rdx, choice
	call scanf
	add rsp, 28h

	cmp choice, 1
	je one

	cmp choice, 2
	je two
	
	cmp choice, 3
	je three

	cmp choice, 4
	je four

	jmp final

	ret
main endp

end