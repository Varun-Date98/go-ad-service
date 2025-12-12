-- +goose Up

CREATE TABLE campaigns (
    campaign_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    bid_cpm NUMERIC NOT NULL,
    min_age INT NOT NULL,
    max_age INT NOT NULL,
    country TEXT NOT NULL,

    CONSTRAINT chk_age_range CHECK (min_age >= 0 AND max_age >= min_age),
    CONSTRAINT chk_bid_nonneg CHECK (bid_cpm >= 0)
);

-- +goose Down

DROP TABLE campaigns;
