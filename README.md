# Go + auth Quickstart for Jenkins X   
# Jawaban Query Test Analisa Database
1. select count(*) as jumlah from vm1dta_chdrpf where validflag = 1;
2. select case statcode WHEN 'IF' THEN 'In Force'
        WHEN 'PO' THEN 'Postpone'
        WHEN 'WD' THEN 'Withdrawn'
        WHEN 'SU' THEN 'Surrender'
        WHEN 'LA' THEN 'Lapsed'
        WHEN 'PT' THEN 'Postterm' end AS status, count(statcode) as Jumlah from vm1dta_chdrpf group by statcode;
3. SELECT 
    c.cnttype AS KODE_PRODUK,
    d.longdesc AS NAMA_PRODUK,
    COUNT(*) AS JUMLAH
    FROM vm1dta_chdrpf c
    JOIN vm1dta_descpf d 
    ON c.cnttype = d.descitem
    WHERE 
    c.VALIDFLAG = 1
    AND d.desctabl = 't5688'
    GROUP BY 
    c.cnttype, d.longdesc
    ORDER BY 
    JUMLAH DESC;     
4. SELECT DISTINCT
    a.agntnum AS NOMOR_AGEN,
    c.surname AS NAMA_AGEN,
    CONCAT_WS(', ', c.cltaddr01, c.cltaddr02, c.cltaddr03, c.cltaddr04, c.cltaddr05) AS ALAMAT_AGEN
    FROM vm1dta_chdrpf h
    JOIN vm1dta_descpf d ON h.cnttype = d.descitem
    JOIN vm1dta_agntpf a ON h.agntnum = a.agntnum
    JOIN vm1dta_clntpf c ON a.clntnum = c.clntnum
    WHERE d.desctabl = 't5688' AND h.VALIDFLAG = 1;   
5. SELECT
    p.chdrnum AS NOMOR_POLIS,
    p.cnttype AS KODE_PRODUK,
    dp.longdesc AS NAMA_PRODUK,
    ds.longdesc AS STATUS_PRODUK,
    CASE
    WHEN p.VALIDFLAG = 1 THEN 'VALID'
    ELSE 'TIDAK VALID'
    END AS STATUS_VALID,
    cp.surname AS NAMA_PEMEGANG_POLIS,
    CONCAT_WS(', ', cp.cltaddr01, cp.cltaddr02, cp.cltaddr03, cp.cltaddr04, cp.cltaddr05) AS ALAMAT_PEMEGANG_POLIS,
    a.agntnum AS NOMOR_AGEN,
    ca.surname AS NAMA_AGEN,
    CONCAT_WS(', ', ca.cltaddr01, ca.cltaddr02, ca.cltaddr03, ca.cltaddr04, ca.cltaddr05) AS ALAMAT_AGEN
    FROM vm1dta_chdrpf p
    JOIN vm1dta_descpf dp ON p.cnttype = dp.descitem AND dp.desctabl = 't5688'
    JOIN vm1dta_descpf ds ON p.statcode = ds.descitem AND ds.desctabl = 't3623'
    JOIN vm1dta_clntpf cp ON p.cownnum = cp.clntnum
    JOIN vm1dta_agntpf a ON p.agntnum = a.agntnum
    JOIN vm1dta_clntpf ca ON a.clntnum = ca.clntnum
    WHERE p.chdrnum = '90000105'; 

# Detail Query Back End
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

# Deploy Docker
1. buka cmd atau terminal docker
2. direct ke repo golang, misal cd go/src/bisnis-be
2. docker compose up --build
3. docker sudah jalan dan bisa langsung digunakan dengan link http://localhost:8080/

# Alur API
1. API dapat di akses di postman collection
2. Mengambil token dari endpoint loginagent
3. Pada saat mengakses endpoint addagent, deleteagent, updateagent, addtransaction, deletetransaction, updatetransaction, diperlukan isi body sesuai yang ada di postman dan mengisi Authorization di Header dengan mengisi token dari endpoint loginagent

# AI
AI menggunakan chatgpt, penggunaan ai digunakan pada saat mencari beberapa syntax yang lupa

# Redis
redis digunakan pada saat create token baru di endpoint loginagent (set redis), sehingga ketika agent_id yang sudah pernah dihit sebelumnya di hit kembali, tidak akan create ualng token tetapi menggunakan token lama dengan ttl / time duration sama seperti token expired.

# Alur Program
1. Hit endpoint http request
2. Server golang membaca
3. lalu masuk ke delivery->handler, disini body request dibaca dan di unmarshal
4. memanggil function service dari delivery/handler
5. di service, ada nya logika seperti if else data
6. service mengirim hasil body request ke data
7. data mengelola hasil request body dan memanggil query
8. query membaca paramater dan membaca koneksi dari http boot
9. hasil dari query dikeluarkan di data
10. data melempar hasil query ke service
11. service melakukan validasi / logika kembali pada hasil query
12. service mengirim ke delivery
13. delivery melempar ke server
14. server mengeluarkan output response hasil query