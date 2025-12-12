-- +goose Up

-- Campaigns
INSERT INTO campaigns (campaign_id, name, status, bid_cpm, min_age, max_age, country) VALUES
('cmp_001', 'Riot Games - Valorant Launch', 'active', 8.50, 18, 35, 'US'),
('cmp_002', 'Nike Running Shoes',           'active', 5.20, 21, 45, 'US'),
('cmp_003', 'Spotify Premium',              'active', 4.00, 16, 40, 'US'),
('cmp_004', 'Apple MacBook Pro',             'paused', 12.00, 22, 50, 'US'),
('cmp_005', 'DoorDash Food Delivery',        'active', 6.75, 18, 60, 'US');

-- Creatives
INSERT INTO creatives (creative_id, campaign_id, asset_url, click_url, format) VALUES
('cr_001', 'cmp_001', 'https://cdn.ads/valorant1.mp4', 'https://playvalorant.com', 'video'),
('cr_002', 'cmp_001', 'https://cdn.ads/valorant2.jpg', 'https://playvalorant.com', 'image'),
('cr_003', 'cmp_002', 'https://cdn.ads/nike1.jpg',     'https://nike.com/run',     'image'),
('cr_004', 'cmp_003', 'https://cdn.ads/spotify1.mp4',  'https://spotify.com/premium', 'video'),
('cr_005', 'cmp_005', 'https://cdn.ads/doordash1.jpg', 'https://doordash.com',     'image');

-- Campaign Targeting
-- Empty array = target ALL
INSERT INTO campaign_targeting (
  campaign_id,
  placement_ids,
  interests_any,
  devices_any,
  languages_any,
  creators_any
) VALUES
(
  'cmp_001',
  ARRAY['stream_pre_roll'],
  ARRAY['valorant', 'esports', 'fps'],
  ARRAY['desktop'],
  ARRAY['en'],
  ARRAY['shroud', 'tenz']
),
(
  'cmp_002',
  ARRAY['feed'],
  ARRAY['fitness', 'running'],
  ARRAY['mobile', 'desktop'],
  ARRAY['en'],
  ARRAY[]::TEXT[]
),
(
  'cmp_003',
  ARRAY['feed', 'audio_mid_roll'],
  ARRAY['music'],
  ARRAY[]::TEXT[],
  ARRAY['en', 'es'],
  ARRAY[]::TEXT[]
),
(
  'cmp_004',
  ARRAY['feed'],
  ARRAY['tech'],
  ARRAY['desktop'],
  ARRAY['en'],
  ARRAY[]::TEXT[]
),
(
  'cmp_005',
  ARRAY[]::TEXT[],
  ARRAY['food', 'delivery'],
  ARRAY['mobile'],
  ARRAY['en'],
  ARRAY[]::TEXT[]
);

-- Users
INSERT INTO users (user_id, country, age, device, language, interests) VALUES
('user_001', 'US', 24, 'desktop', 'en', ARRAY['valorant', 'esports']),
('user_002', 'US', 32, 'mobile',  'en', ARRAY['fitness', 'running']),
('user_003', 'US', 19, 'mobile',  'en', ARRAY['music']),
('user_004', 'US', 45, 'desktop', 'en', ARRAY['tech']),
('user_005', 'US', 28, 'mobile',  'es', ARRAY['food', 'delivery']);

-- +goose Down

DELETE FROM users;
DELETE FROM campaign_targeting;
DELETE FROM creatives;
DELETE FROM campaigns;
