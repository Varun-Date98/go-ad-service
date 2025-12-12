-- +goose Up

CREATE TABLE creatives (
  creative_id   TEXT PRIMARY KEY,
  campaign_id   TEXT NOT NULL REFERENCES campaigns(campaign_id) ON DELETE CASCADE,
  asset_url     TEXT NOT NULL,
  click_url     TEXT NOT NULL,
  format        TEXT NOT NULL
);

-- +goose Down

DROP TABLE creatives;
