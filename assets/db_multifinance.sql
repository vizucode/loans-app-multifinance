CREATE TABLE customers (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NULL,
    date_birth DATE NOT NULL,
    born_at VARCHAR(255) NULL,
    salary DECIMAL(15, 2) NULL,
    national_identity_number VARCHAR(255) NOT NULL,
    national_identity_image_url VARCHAR(255) NULL,
    selfie_image_url VARCHAR(255) NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_customers_email (email),
    INDEX idx_customers_national_identity_number (national_identity_number)
);

CREATE TABLE tenors (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    total_month TINYINT UNSIGNED NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_tenors_total_month (total_month)
);

CREATE TABLE customer_tenors (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    customer_id VARCHAR(36) NOT NULL,
    tenor_id VARCHAR(36) NOT NULL,
    limit_loan_amount DECIMAL(15, 2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT uc_customer_tenor UNIQUE (customer_id, tenor_id),

    CONSTRAINT fk_customer_tenors_customer_id
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
        ON DELETE CASCADE ON UPDATE CASCADE,

    CONSTRAINT fk_customer_tenors_tenor_id
        FOREIGN KEY (tenor_id)
        REFERENCES tenors(id)
        ON DELETE RESTRICT ON UPDATE CASCADE,
        
    INDEX idx_customer_tenors_customer_id (customer_id),
    INDEX idx_customer_tenors_tenor_id (tenor_id)
);

CREATE TABLE customer_loans (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    customer_id VARCHAR(36) NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    otr DECIMAL(15, 2) NOT NULL,
    total_month TINYINT UNSIGNED NOT NULL,
    interest_rate DECIMAL(5, 4) NOT NULL,
    admin_fee_amount DECIMAL(15, 2) NOT NULL,
    interest_amount DECIMAL(15, 2) NOT NULL,
    installment_amount DECIMAL(15, 2) NOT NULL,
    total_installment_amount DECIMAL(15, 2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_customer_loans_customer_id
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
        
    INDEX idx_customer_loans_customer_id (customer_id)
);

CREATE TABLE tokens (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    customer_id VARCHAR(36) NOT NULL,
    refresh_token TEXT,
    access_token TEXT,
    refresh_token_revoked BOOLEAN DEFAULT FALSE,
    access_token_revoked BOOLEAN DEFAULT FALSE,
    refresh_token_expired_at DATETIME,
    access_token_expired_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_token_customer_id
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
        
    INDEX idx_tokens_refresh_token (refresh_token(255)),
    INDEX idx_tokens_access_token (access_token(255)),
    INDEX idx_tokens_customer_id (customer_id)
);

INSERT INTO tenors (id, total_month, created_at, updated_at) VALUES
    (UUID(), 1, NOW(), NOW()),
    (UUID(), 2, NOW(), NOW()),
    (UUID(), 3, NOW(), NOW()),
    (UUID(), 6, NOW(), NOW());