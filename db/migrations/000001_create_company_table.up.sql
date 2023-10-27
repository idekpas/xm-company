CREATE TABLE company
(
    companyID        SERIAL PRIMARY KEY,
    ID               uuid               NOT NULL,
    Name             varchar(15) UNIQUE NOT NULL,
    Description      varchar(3000) NULL,
    Employees_Amount int                NOT NULL,
    Registered       bool               NOT NULL,
    Type             varchar(20)        NOT NULL
);