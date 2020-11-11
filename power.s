#PURPOSE: Program to illustrate how functions work
#	  This program will compute the value of
# 	  2^3 + 5^2
#

#Everything in the main program is stored in registers,
#so the data section doesn't have anything.
.section .data

.section .text

.globl _start
_start: 
movl $3, (%esp)
subl $4, (%esp)			#push second argument
movl $2, (%esp)
subl $4, (%esp) 		#push first argument
call power 		#call the function
add $8, %esp		#move the stack pointer back


movl %eax, (%esp)
subl $4, (%esp) 		#save the first answer before
			#calling the next function

movl $2, (%esp)
subl $4, (%esp)		#push second argument

movl $2, (%esp)
subl $4, (%esp)		#push first argument
call power 		#call the function
add $8, %esp		#move the stack pointer back


movl (%esp), %edx
addl $4, %esp		#The second answer is already
			#in %eax. We saved the 
			#first answer onto the stack,
			#so now we can just pop it
			#out into %ebx

add %eax, %edi		#add them together
			#the result is in %ebx

mov $60, %eax		#exit (%ebx is returned)
syscall

#PURPOSE: This function is used to compute
#	  the value of a number to a power
#
#INPUT:   First argument - the base number
# 	  Second argument - the power to 
#			  raise it to
#
#OUTPUT:  Will give the result as a return value
#
#NOTES:   The power must be 1 or greater
#
#VARIABLES: 
#	  %ebx - holds the base number
#	  %exc - holds the power
#
#	  -4(%ebp) - holds the current result
#
#	  %eax is used for temporary storage
#
.type power, @function
power: 

movl %ebp, (%esp)
addl $4, (%esp) 	#save old base pointer
mov %esp, %ebp	#make stack pointer the base pointer
sub $4, %esp	#get room for our local storage

mov 8(%ebp), %ebx #put first argument in %eax
mov 12(%ebp), %ecx #put second argument in %ecx

mov %ebx, -4(%ebp) #store current result 

power_loop_start:
cmp $1, %ecx	#if the power is 1, we are done
je end_power
mov -4(%ebp), %eax #move the current result into %eax
imul %ebx, %eax    #multiply the current result by
		    #the base number
mov %eax, -4(%ebp) #store the current result 

dec %ecx	    #decrease the power
jmp power_loop_start#run for the next power

end_power: 
mov -4(%ebp), %eax #return value goes in %eax
mov %ebp, %esp	    #restore the stack pointer

movl (%esp), %ebp
addl $4, %esp 	    #restore the base pointer
ret
