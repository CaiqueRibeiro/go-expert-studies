CREATE TABLE pedidos (
    numero SERIAL PRIMARY KEY,
    cliente_id INTEGER,
    valor REAL,
    valor_entrada REAL,
    bloqueado BOOLEAN
);

CREATE INDEX idx_pedidos_cliente_id ON pedidos (cliente_id);

CREATE TABLE pagamentos (
    numero SERIAL PRIMARY KEY,
    cliente_id INTEGER,
    vencimento DATE,
    pagamento DATE,
    valor REAL,
    liquidado BOOLEAN,
    inadimplente BOOLEAN
);

CREATE INDEX idx_pagamentos_cliente_id ON pagamentos (cliente_id);

insert into pedidos(cliente_id, valor, valor_entrada, bloqueado)
    values(100, 200.00, 10.00, FALSE);
insert into pedidos(cliente_id, valor, valor_entrada, bloqueado)
    values(200, 1000.00, 10.00, FALSE);
insert into pedidos(cliente_id, valor, valor_entrada, bloqueado)
    values(300, 100.00, 10.00, TRUE);

INSERT INTO pagamentos (cliente_id, vencimento, pagamento, valor, liquidado, inadimplente)
    VALUES (100, '2024-01-10', '2024-01-10', 500.00, TRUE, FALSE);
INSERT INTO pagamentos (cliente_id, vencimento, pagamento, valor, liquidado, inadimplente)
    VALUES (200, '2024-01-10', '2024-01-10', 1000.00, TRUE, FALSE);
INSERT INTO pagamentos (cliente_id, vencimento, pagamento, valor, liquidado, inadimplente)
    VALUES (300, '2024-01-10', '2024-01-10', 1000.00, FALSE, false);
INSERT INTO pagamentos (cliente_id, vencimento, pagamento, valor, liquidado, inadimplente)
    VALUES (400, '2024-01-10', '2024-01-10', 1000.00, FALSE, true);


