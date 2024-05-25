local function word_count(s)
  local words = {}
  -- pattern match:
  -- [..] character class.
  -- %w': letter or digit (alphanumeric characters) or '
  -- +: match more than one such character
  for word in s:gmatch("[%w']+") do
      -- begins with ^, so match only at the beginning of the string.
      -- *$: until the end.
      -- *: 0 or more repetitions
      -- alphanumeric characters, ' and dash
      -- to lower character conversion.
      word = word:match("^'*([%w']-)'*$"):lower()
      -- if word exists in word table, count +1
      if words[word] then
      	words[word] = words[word] + 1
      else
        words[word] = 1
      end
  end
  return words
end

return { word_count = word_count }