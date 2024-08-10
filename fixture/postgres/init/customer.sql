CREATE TABLE customers (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  zip VARCHAR(10) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  mktsegment VARCHAR(20) NOT NULL,
  nation VARCHAR(20) NOT NULL,
  birthdate TIMESTAMP NOT NULL,
  version INT NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO customers (id, name, address, zip, phone, mktsegment, nation, birthdate, version)
VALUES
  ('XXX1', '山田 太郎', '東京都練馬区豊玉北2-13-1', '176-0013', '03-1234-5678', '個人', '日本', '1980-01-01 00:00:00', 1),
  ('YYY2', '鈴木 花子', '神奈川県横浜市中区本牧3-10-1', '231-0012', '045-222-3333', '法人', '日本', '1985-04-05 13:50:00', 2),
  ('ZZZ3', '佐藤 次郎', '大阪府大阪市北区梅田1-1-1', '530-0001', '06-6666-7777', '個人', '日本', '1990-07-20 22:10:00', 1);
