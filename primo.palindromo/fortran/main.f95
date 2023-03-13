! @jeffotoni
program prime_palindrome
  implicit none
  integer, parameter :: precision = 100000
  real(kind=16), parameter :: pi = 3.141592653589793238_16
  character(len=precision) :: s
  integer :: i, n
  ! quero um exemplo usando a função write em fortran
  write(*,*) pi
  ! write(s, '(F', precision, '.', precision, ')') pi
  do i = 1, len(s) - 5
    read(s(i:i+4), *) n
    if (is_palindrome(n) .and. is_prime(n)) then
      write(*,'(I5)') n
      exit
    end if
  end do

contains
  function is_palindrome(n) result(ans)
    integer, intent(in) :: n
    integer :: k, l, r
    logical :: ans

    k = n
    l = 0
    do while (k > 0)
      l = l*10 + mod(k, 10)
      k = k / 10
    end do
    ans = (l == n)
  end function

  function is_prime(n) result(ans)
    integer, intent(in) :: n
    integer :: i
    logical :: ans

    ans = (n > 1)
    do i = 2, n-1
      if (mod(n, i) == 0) then
        ans = .false.
        exit
      end if
    end do
  end function
end program prime_palindrome
