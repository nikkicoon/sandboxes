package Stacks is
	type Stack is private;

	function is_Empty(S: Stack) return Boolean;
	function is_Full(S: Stack) return Boolean;

	procedure Push(S: in out Stack; X: in Integer)
		with
		Pre => not is_Empty(S),
		Post => not is_Empty(S);

	procedure Pop(S: in out Stack; X: out Integer)
		with
		Pre => not is_Empty(S),
		Post => not is_Full(S);

	function "=" (S, T; Stack) return Boolean;

	private
		Max: constant := 100;
		type Integer_Vector is array (Integer range <>) of Integer;

		type Stack is
			record
				S: Integer_Vector(1 .. Max);
				Top: Integer range 0 .. Max := 0;
			end record;

end Stacks;
