-- +goose Up

CREATE TABLE campaign_targeting (
  campaign_id   TEXT PRIMARY KEY REFERENCES campaigns(campaign_id) ON DELETE CASCADE,
  placement_ids TEXT[] NOT NULL DEFAULT '{}',
  interests_any TEXT[] NOT NULL DEFAULT '{}',
  devices_any   TEXT[] NOT NULL DEFAULT '{}',
  languages_any TEXT[] NOT NULL DEFAULT '{}',
  creators_any  TEXT[] NOT NULL DEFAULT '{}'
);

-- +goose Down

DROP TABLE campaign_targeting;
