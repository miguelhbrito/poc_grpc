CREATE TABLE IF NOT EXISTS notebook
(
    id CHARACTER varying(36) NOT NULL,
    name CHARACTER varying(200) NOT NULL,
    marca CHARACTER varying(200) NOT NULL,
    modelo CHARACTER varying(100),
    numero_serie INTEGER NOT NULL,
    PRIMARY KEY (id)
);
