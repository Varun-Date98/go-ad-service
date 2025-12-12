-- name: GetCandidateAd :many

SELECT 
  c.campaign_id,
  c.min_age,
  c.max_age,
  c.country,
  c.bid_cpm,
  t.placement_ids,
  t.interests_any,
  t.devices_any,
  t.languages_any,
  t.creators_any,
  cr.creative_id,
  cr.asset_url,
  cr.click_url
FROM campaigns c
JOIN campaign_targeting t ON c.campaign_id = t.campaign_id
JOIN creatives cr ON c.campaign_id = cr.campaign_id
WHERE c.status = 'active'
  AND c.country = sqlc.arg(country)
  AND sqlc.arg(age) BETWEEN c.min_age AND c.max_age
  AND (cardinality(t.placement_ids) = 0 OR sqlc.arg(placement_id)::text = ANY(t.placement_ids))
  AND (cardinality(t.devices_any)   = 0 OR sqlc.arg(device)::text = ANY(t.devices_any))
  AND (cardinality(t.languages_any) = 0 OR sqlc.arg(language)::text = ANY(t.languages_any))
  AND (cardinality(t.creators_any)  = 0 OR sqlc.arg(creator_id)::text = ANY(t.creators_any))
  AND (cardinality(t.interests_any) = 0 OR t.interests_any && sqlc.arg(interests)::text[]);
