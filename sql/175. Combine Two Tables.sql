-- https://leetcode.com/problems/combine-two-tables/description/
select p.firstname, p.lastname, a.city, a.state from Person as p
left join Address as a on a.personId = p.personId