CREATE TABLE IF NOT EXISTS CARD (
  id        SERIAL PRIMARY KEY,
  set_id    SERIAL REFERENCES SET (id),
  name      TEXT NOT NULL,
  ascii     TEXT NOT NULL,
  image     TEXT,
  layout    TEXT,
  type      TEXT NOT NULL,
  mana      TEXT,
  cost      INT,
  text      TEXT,
  flavor    TEXT,
  power     TEXT,
  toughness TEXT,
  loyalty   INT,
  hand      INT,
  life      INT,

  CONSTRAINT card_uniqueness UNIQUE (id, set_id, name),
  CHECK (text IS NOT NULL AND flavor IS NOT NULL)
)