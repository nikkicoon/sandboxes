function gcd(a, b)
	-- gcd(0, b) == b; gcd(a, 0) == a
	-- gcd(0, 0) == 0
	if a == 0 then
		return b
	end
	if b == 0 then
		return b
	end
end
