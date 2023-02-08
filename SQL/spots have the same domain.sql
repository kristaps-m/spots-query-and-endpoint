-- website field, only contains the domain
SELECT dbo.Parse_For_Domain_Name([website]) as Domain, 
COUNT([coordinates]) as SpotCount
FROM [spots].[dbo].[MY_TABLE]
WHERE [website] is NOT NULL
-- how many spots have the same domain
GROUP BY dbo.Parse_For_Domain_Name([website])
HAVING COUNT([website]) > 1
-- spots which have a domain with a count greater than 1
ORDER BY SpotCount desc