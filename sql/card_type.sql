CREATE TABLE IF NOT EXISTS CARD_TYPE (
  card_id SERIAL REFERENCES CARD (id),
  type_id SERIAL REFERENCES TYPE (id),

  PRIMARY KEY (card_id, type_id)
)