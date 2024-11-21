-- https://leetcode.com/problems/product-sales-analysis-i/description/
SELECT s.sale_id, p.product_id, s.year, s.quantity, s.price
FROM Sales as s, Product as p JOIN ON s.product_id = p.product_id