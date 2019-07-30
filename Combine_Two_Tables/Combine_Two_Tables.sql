# Write your MySQL query statement below
select Person.FirstNme, Person.LastName, Address.City, Address.State
from Person, Address
where Person.PersonId = Address.PersonId
;
