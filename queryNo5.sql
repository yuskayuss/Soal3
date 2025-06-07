SELECT 
    a.account_id,
    a.username,
    a.email,
    MAX(r.tariff_id) AS highest_tariff,
    SUM(r.amount) AS consumption,
    SUM(r.amount) AS total_cost
FROM 
    account a
JOIN 
    readings r ON a.account_id = r.account_id
GROUP BY 
    a.account_id, a.username, a.email;
