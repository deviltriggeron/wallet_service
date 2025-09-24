CREATE TABLE wallets ( 
    id UUID PRIMARY KEY, 
    balance BIGINT NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT now() 
);