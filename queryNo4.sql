SELECT
    c.id AS customer_id,
    COALESCE(c.bank_account, 'No bank') AS bank_account,
    b.amount
FROM
    customer c
JOIN
    balances b ON c.id = b.customer_id
WHERE
    b.amount > 0;