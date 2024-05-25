-- Caesar Cipher module

local caesar_cipher = {}

local function ascii_base(s)
    return s:lower() == s and ('a'):byte() or ('A'):byte()
end

function caesar_cipher.caesar( str, key )
    return str:gsub('%a', function(s)
			local base = ascii_base(s)
			return string.char(((s:byte() - base + key) % 26) + base)
    end)
end

return caesar_cipher
