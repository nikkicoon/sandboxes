local function fizzbuzz(n)
	for i=1, n do
		if i % 3 == 0 and i % 5 == 0 then
			print("FizzBuzz")
		end
		if i % 3 == 0 then
			print("Fizz")
		end
		if i % 5 == 0 then
			print("Buzz")
		end
		print(i)
	end
end

fizzbuzz(100)
