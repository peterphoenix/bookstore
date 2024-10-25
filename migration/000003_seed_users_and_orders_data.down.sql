DELETE FROM "order_item"
WHERE order_id IN (
    'd34434fc-b380-47ee-8af6-2f87afa92968',
    'cf843510-38dc-42db-aeea-482a48b5ac27',
    'b27b75a4-3006-4046-b2c9-9ad127ff97ef'
);

DELETE FROM "order"
WHERE id IN (
   'd34434fc-b380-47ee-8af6-2f87afa92968',
   'cf843510-38dc-42db-aeea-482a48b5ac27',
   'b27b75a4-3006-4046-b2c9-9ad127ff97ef'
);

DELETE FROM "user"
WHERE id IN (
  '1ccc5f59-c962-41b4-b296-bda9dcbf0077',
  'd23699f0-96f7-4474-a6ec-5c6971c31504'
);
