procedure Show_Simple_Precondition is

   procedure DB_Entry (Name : String; Age  : Natural)
     with Pre => Name'Length > 0
   is
   begin
      --  Missing implementation
      null;
   end DB_Entry;
begin
   DB_Entry ("John", 30);
   DB_Entry ("",     21);  --  Precondition will fail!
end Show_Simple_Precondition;
