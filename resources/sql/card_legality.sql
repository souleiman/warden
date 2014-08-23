CREATE TABLE IF NOT EXISTS CARD_LEGALITY (
  card_id   SERIAL REFERENCES CARD (id),
  format_id SERIAL REFERENCES FORMAT (id),
  legality  TEXT,

  PRIMARY KEY (card_id, format_id)
)