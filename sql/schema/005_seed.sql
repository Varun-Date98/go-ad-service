-- +goose Up

-- Clean slate (safe for dev)
DELETE FROM users;
DELETE FROM campaign_targeting;
DELETE FROM creatives;
DELETE FROM campaigns;

-- Campaigns (12)
INSERT INTO campaigns (campaign_id, name, status, bid_cpm, min_age, max_age, country) VALUES
('cmp_001', 'Riot - Valorant A',         'active',  8.50, 18, 35, 'US'),
('cmp_002', 'Riot - Valorant B',         'active',  7.90, 18, 35, 'US'),
('cmp_003', 'Nike - Running Shoes',      'active',  5.20, 21, 45, 'US'),
('cmp_004', 'Spotify - Premium',         'active',  4.00, 16, 40, 'US'),
('cmp_005', 'DoorDash - Delivery',       'active',  6.75, 18, 60, 'US'),
('cmp_006', 'Apple - MacBook Pro',       'paused', 12.00, 22, 50, 'US'),
('cmp_007', 'Prime Video - New Show',    'active',  6.10, 18, 55, 'US'),
('cmp_008', 'Duolingo - Learn Spanish',  'active',  3.50, 16, 50, 'US'),
('cmp_009', 'EA - FC Game',              'active',  7.10, 18, 40, 'US'),
('cmp_010', 'Samsung - Galaxy',          'active',  6.40, 18, 50, 'US'),
('cmp_011', 'Canada - Tim Hortons',      'active',  4.90, 18, 60, 'CA'),
('cmp_012', 'Germany - BMW',             'active',  9.50, 25, 60, 'DE');

-- Creatives (at least 2 per key campaigns; total 18)
INSERT INTO creatives (creative_id, campaign_id, asset_url, click_url, format) VALUES
('cr_001', 'cmp_001', 'https://cdn.ads/valorant_a_1.mp4', 'https://playvalorant.com', 'video'),
('cr_002', 'cmp_001', 'https://cdn.ads/valorant_a_2.jpg', 'https://playvalorant.com', 'image'),

('cr_003', 'cmp_002', 'https://cdn.ads/valorant_b_1.mp4', 'https://playvalorant.com', 'video'),
('cr_004', 'cmp_002', 'https://cdn.ads/valorant_b_2.jpg', 'https://playvalorant.com', 'image'),

('cr_005', 'cmp_003', 'https://cdn.ads/nike_1.jpg',       'https://nike.com/run',    'image'),
('cr_006', 'cmp_003', 'https://cdn.ads/nike_2.jpg',       'https://nike.com/run',    'image'),

('cr_007', 'cmp_004', 'https://cdn.ads/spotify_1.mp4',    'https://spotify.com/premium', 'video'),

('cr_008', 'cmp_005', 'https://cdn.ads/doordash_1.jpg',   'https://doordash.com',    'image'),

('cr_009', 'cmp_006', 'https://cdn.ads/macbook_1.jpg',    'https://apple.com/mac',   'image'),

('cr_010', 'cmp_007', 'https://cdn.ads/prime_1.mp4',      'https://primevideo.com',  'video'),
('cr_011', 'cmp_007', 'https://cdn.ads/prime_2.jpg',      'https://primevideo.com',  'image'),

('cr_012', 'cmp_008', 'https://cdn.ads/duolingo_1.jpg',   'https://duolingo.com',    'image'),

('cr_013', 'cmp_009', 'https://cdn.ads/ea_fc_1.mp4',      'https://ea.com/fc',       'video'),
('cr_014', 'cmp_009', 'https://cdn.ads/ea_fc_2.jpg',      'https://ea.com/fc',       'image'),

('cr_015', 'cmp_010', 'https://cdn.ads/galaxy_1.jpg',     'https://samsung.com',     'image'),

('cr_016', 'cmp_011', 'https://cdn.ads/tim_1.jpg',        'https://timhortons.ca',   'image'),

('cr_017', 'cmp_012', 'https://cdn.ads/bmw_1.jpg',        'https://bmw.com',         'image'),
('cr_018', 'cmp_012', 'https://cdn.ads/bmw_2.jpg',        'https://bmw.com',         'image');

-- Campaign Targeting
-- NOTE: Empty array = target ALL for that dimension
INSERT INTO campaign_targeting (
  campaign_id, placement_ids, interests_any, devices_any, languages_any, creators_any
) VALUES

-- Two Valorant campaigns both eligible for same user/placement/creator (for multi-eligible cap test)
('cmp_001', ARRAY['stream_pre_roll'], ARRAY['valorant','esports','fps'], ARRAY['desktop'], ARRAY['en'], ARRAY['shroud','tenz']),
('cmp_002', ARRAY['stream_pre_roll'], ARRAY['valorant','fps'],          ARRAY['desktop'], ARRAY['en'], ARRAY['shroud']),

-- Feed-based lifestyle
('cmp_003', ARRAY['feed'],            ARRAY['fitness','running'],       ARRAY['mobile','desktop'], ARRAY['en'], ARRAY[]::TEXT[]),

-- Spotify: audio mid-roll + feed, all devices
('cmp_004', ARRAY['feed','audio_mid_roll'], ARRAY['music'],             ARRAY[]::TEXT[], ARRAY['en','es'], ARRAY[]::TEXT[]),

-- DoorDash: ALL placements, mobile, EN+ES (so Spanish users can match)
('cmp_005', ARRAY[]::TEXT[],          ARRAY['food','delivery'],         ARRAY['mobile'], ARRAY['en','es'], ARRAY[]::TEXT[]),

-- MacBook paused anyway, keep targeting
('cmp_006', ARRAY['feed'],            ARRAY['tech'],                    ARRAY['desktop'], ARRAY['en'], ARRAY[]::TEXT[]),

-- Prime Video: video placements
('cmp_007', ARRAY['stream_pre_roll','video_mid_roll'], ARRAY['tv','movies','entertainment'], ARRAY['desktop','mobile'], ARRAY['en'], ARRAY[]::TEXT[]),

-- Duolingo: feed, language learners
('cmp_008', ARRAY['feed'],            ARRAY['learning','languages','spanish'], ARRAY['mobile','desktop'], ARRAY['en','es'], ARRAY[]::TEXT[]),

-- EA FC: stream placements, gaming
('cmp_009', ARRAY['stream_pre_roll','video_mid_roll'], ARRAY['gaming','sports'], ARRAY['desktop'], ARRAY['en'], ARRAY['shroud','nickmercs']),

-- Samsung: feed, tech broad
('cmp_010', ARRAY['feed'],            ARRAY['tech','mobile'],           ARRAY['mobile','desktop'], ARRAY['en'], ARRAY[]::TEXT[]),

-- CA only
('cmp_011', ARRAY['feed'],            ARRAY['coffee','food'],           ARRAY['mobile','desktop'], ARRAY['en','fr'], ARRAY[]::TEXT[]),

-- DE only
('cmp_012', ARRAY['feed'],            ARRAY['cars','luxury'],           ARRAY['mobile','desktop'], ARRAY['de','en'], ARRAY[]::TEXT[]);

-- Users (10)
INSERT INTO users (user_id, country, age, device, language, interests) VALUES
('user_001', 'US', 24, 'desktop', 'en', ARRAY['valorant','esports']),
('user_002', 'US', 32, 'mobile',  'en', ARRAY['fitness','running']),
('user_003', 'US', 19, 'mobile',  'en', ARRAY['music']),
('user_004', 'US', 45, 'desktop', 'en', ARRAY['tech']),
('user_005', 'US', 28, 'mobile',  'es', ARRAY['food','delivery']),
('user_006', 'US', 26, 'desktop', 'en', ARRAY['tv','movies']),
('user_007', 'US', 29, 'desktop', 'en', ARRAY['gaming','sports']),
('user_008', 'US', 20, 'mobile',  'en', ARRAY['learning','languages']),
('user_009', 'CA', 30, 'mobile',  'en', ARRAY['coffee','food']),
('user_010', 'DE', 35, 'desktop', 'de', ARRAY['cars','luxury']);

-- +goose Down
DELETE FROM users;
DELETE FROM campaign_targeting;
DELETE FROM creatives;
DELETE FROM campaigns;
