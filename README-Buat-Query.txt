CREATE TABLE bisnis.product (
    product_id VARCHAR(100) NOT NULL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    premium FLOAT DEFAULT 0,
    active TINYINT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bisnis.agen (
    agent_id VARCHAR(100) NOT NULL PRIMARY KEY,
    agent_name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    active TINYINT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bisnis.product_parameter (
    id int AUTO_INCREMENT PRIMARY KEY,
    product_id VARCHAR(100) NOT NULL,
    parameter_name VARCHAR(100) NOT NULL,
	parameter_value VARCHAR(100) NOT NULL,
    active TINYINT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bisnis.transaction (
    trans_id int AUTO_INCREMENT PRIMARY KEY,
    agent_id VARCHAR(100) NOT NULL,
    product_id VARCHAR(100) NOT NULL,
    nama VARCHAR(100) NOT NULL,
	usia TINYINT,
	premium FLOAT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO bisnis.product
(product_id, product_name, premium, active, created_at)
select 'T10', 'Aurora', 1000, CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END , NOW()
union
select 'L13', 'Davestera', 2000, CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END , NOW();

INSERT INTO bisnis.product_parameter
(product_id, parameter_name, parameter_value, active, created_at)
select 'T10', 'MINUSIAMASUK', '1', CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END, NOW()
union
select 'T10', 'MAXUSIAMASUK', '99', CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END, NOW()
union
select 'L13', 'MINUSIAMASUK', '18', CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END, NOW()
union
select 'L13', 'MAXUSIAMASUK', '50', CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END, NOW();

INSERT INTO bisnis.agen
(agent_id, agent_name, password, active, created_at)
select 'BFA01', 'Agent Satu', 'P@ssagent1', CASE WHEN 'Y' = 'Y' THEN 1 ELSE 0 END, NOW();