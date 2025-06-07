SELECT 
    a.username,
    a.email,
    (
        SELECT t.name
        FROM readings r2
        JOIN tariffs t ON r2.tariff_id = t.tariff_id
        WHERE r2.account_id = a.account_id
        ORDER BY r2.tariff_id DESC
        LIMIT 1
    ) AS highest_tariff,
    SUM(r.amount) AS consumption,
    ROUND(SUM(r.amount * t.cost), 2) AS total_cost
FROM 
    account a
JOIN 
    readings r ON a.account_id = r.account_id
JOIN 
    tariffs t ON r.tariff_id = t.tariff_id
GROUP BY 
    a.account_id, a.username, a.email
ORDER BY 
    a.username ASC;
