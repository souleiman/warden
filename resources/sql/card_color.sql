CREATE TABLE IF NOT EXISTS CARD_COLOR (
  card_id SERIAL REFERENCES CARD (id),
  color_id SERIAL REFERENCES COLOR (id)
)