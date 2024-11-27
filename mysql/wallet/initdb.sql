use wallet;
CREATE TABLE IF NOT EXISTS clients(
    id varchar(255), name varchar(255), email varchar(255), created_at datetime
);

CREATE TABLE IF NOT EXISTS accounts(
    id varchar(255), client_id varchar(255), balance int, created_at datetime
);

CREATE TABLE IF NOT EXISTS transactions(
    id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at datetime
);

INSERT INTO clients Values('e066ec93-746e-4d66-8b40-154c4502e84c','Jane Doe','jane@j.com', now());
INSERT INTO clients Values('f243266c-d196-468a-9ce8-0658b78d3e34','John Doe','john@j.com', now());

INSERT INTO accounts Values('166e5bdb-ad38-4c77-b6a4-dee97c24e36e','e066ec93-746e-4d66-8b40-154c4502e84c',1000, now());
INSERT INTO accounts Values('ff048314-f669-4044-8c21-a704ab57ffe3','f243266c-d196-468a-9ce8-0658b78d3e34',1000, now());