CREATE TABLE company
(
    companyid        SERIAL PRIMARY KEY,
    id               uuid               NOT NULL,
    name             varchar(15) UNIQUE NOT NULL,
    description      varchar(3000) NULL,
    employees_amount int                NOT NULL,
    registered       bool               NOT NULL,
    type             varchar(20)        NOT NULL
);