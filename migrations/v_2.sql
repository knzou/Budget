-- fuzzy text matching
-- Trigrams are formed by breaking string into grups of 3 consecutive letters ex. Hello -> h, he, hel, ell, llo, lo, o
CREATE EXTENSION IF NOT EXISTS pg_trgm;