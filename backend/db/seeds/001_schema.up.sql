-- users: ユーザ情報
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- items: アイテムのマスタデータ
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category INT NOT NULL,
    image_url VARCHAR(255) NOT NULL
);

-- user_items: 所持アイテム(Inventory)
CREATE TABLE user_items (
    user_id BIGINT,
    item_id INT,
    PRIMARY KEY (user_id, item_id),
    CONSTRAINT fk_users_items_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_users_items_item FOREIGN KEY (item_id) REFERENCES items(id)
);

-- avatars: 現在のアバターの着せ替え状態
CREATE TABLE avatars (
    user_id BIGINT PRIMARY KEY, -- 1人のユーザーは1つのアバター状態を持つ
    hat_id INT,
    shirt_id INT,
    jacket_id INT,
    bottoms_id INT,
    shoes_id INT,
    CONSTRAINT fk_avatars_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_avatars_hat FOREIGN KEY (hat_id) REFERENCES items(id),
    CONSTRAINT fk_avatars_shirt FOREIGN KEY (shirt_id) REFERENCES items(id),
    CONSTRAINT fk_avatars_jacket FOREIGN KEY (jacket_id) REFERENCES items(id),
    CONSTRAINT fk_avatars_bottoms FOREIGN KEY (bottoms_id) REFERENCES items(id),
    CONSTRAINT fk_avatars_shoes FOREIGN KEY (shoes_id) REFERENCES items(id)
);

