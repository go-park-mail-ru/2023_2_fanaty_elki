CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    birthday DATE,
    phone_number TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    icon TEXT DEFAULT 'deficon',
    created_at TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (id),
    CONSTRAINT valid_username CHECK ( LENGTH(username) > 3 and LENGTH(username) < 20 ),
    CONSTRAINT valid_password CHECK ( LENGTH(password) > 8 and LENGTH(password) < 30 ),
    CONSTRAINT valid_phone CHECK ( phone_number ~* '/\+7[0-9]{10}/'),
    CONSTRAINT valid_email CHECK ( email ~* '\S*@\S*' and LENGTH(email) < 40)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.restaurant
(
    id SERIAL NOT NULL,
    name TEXT UNIQUE NOT NULL,
    rating NUMERIC(2,1) DEFAULT 0.0 NOT NULL,
    comments_count INT DEFAULT 0 NOT NULL,
    icon TEXT DEFAULT 'deficon' NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (id),
    CONSTRAINT valid_restaurant CHECK ( LENGTH(NAME) > 0 and LENGTH(NAME) < 40 ),
    CONSTRAINT valid_rating CHECK ( rating >= 0.0 AND rating <= 5.0),
    CONSTRAINT valid_comments_count CHECK (comments_count >= 0)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON restaurant
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


CREATE TABLE IF NOT EXISTS public.category
(
    id SERIAL NOT NULL,
    name TEXT UNIQUE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT valid_text CHECK ( LENGTH(name) > 0 and LENGTH(name) < 40 )
);

CREATE TABLE IF NOT EXISTS public.restaurant_category
(
    id SERIAL NOT NULL,
    restaurant_id INT REFERENCES public.restaurant(id) NOT NULL,
    category_id INT REFERENCES public.category(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.menu_type
(
    id SERIAL NOT NULL,
    name TEXT DEFAULT 'FOOD' NOT NULL,
    restaurant_id INT REFERENCES restaurant(id) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT valid_menu_type CHECK ( LENGTH(name) > 0 and LENGTH(name) < 40 )
);

CREATE TABLE IF NOT EXISTS public.product
(
    id SERIAL NOT NULL,
    name TEXT NOT NULL, -- У блюда не может быть дефолтного значения, иначе как нам понять что это
    restaurant_id INT REFERENCES RESTAURANT(ID) NOT NULL,
    price NUMERIC(10,2) DEFAULT '0.0' NOT NULL,
    cooking_time INT DEFAULT '0' NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (id),
    CONSTRAINT valid_product CHECK ( LENGTH(name) > 0 and LENGTH(name) < 40 ),
    CONSTRAINT valid_price CHECK ( price >= 0.0 ),
    CONSTRAINT valid_time CHECK ( cooking_time >= 0 )
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON product
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.product_menu_type
(
    id SERIAL NOT NULL,
    menu_type_id INT REFERENCES public.menu_type(id) NOT NULL,
    product_id INT REFERENCES public.product(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.orders
(
    id SERIAL NOT NULL,
    user_id INT REFERENCES public.users(id) NOT NULL,
    order_date TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    status TEXT DEFAULT 'CREATED' NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (id),
    CONSTRAINT valid_status CHECK (LENGTH(status) >= 0 and LENGTH(status) < 40 )
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON orders
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


CREATE TABLE IF NOT EXISTS public.orders_product
(
    id SERIAL NOT NULL,
    product_id INT REFERENCES public.product(id) NOT NULL,
    order_id INT REFERENCES public.orders(id) NOT NULL,
    item_count INT DEFAULT 1 NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT valid_count CHECK ( item_count > 0 )
);

CREATE TABLE IF NOT EXISTS public.comment
(
    id SERIAL NOT NULL,
    comment_text TEXT,  -- бывают комментарии без текста
    restaurant_id INT REFERENCES restaurant(id) NOT NULL,
    user_id INT REFERENCES public.users(id) NOT NULL,
    rating NUMERIC(2,1) DEFAULT 0.0 NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT valid_count CHECK ( rating >= 0.0 AND rating <= 5.0)
);

CREATE TABLE IF NOT EXISTS public.address
(
    id SERIAL NOT NULL,
    city TEXT NOT NULL,
    street TEXT NOT NULL,
    house_number INT NOT NULL,
    flat_number INT, -- бывают ДОМА БЕЗ КВАРТИР
    PRIMARY KEY (id),
    CONSTRAINT valid_city CHECK (LENGTH(city) > 0 and LENGTH(city) < 40 ),
    CONSTRAINT valid_street CHECK (LENGTH(street) > 0 and LENGTH(street) < 40 ),
    CONSTRAINT valid_house_number CHECK (house_number > 0 ),
    CONSTRAINT valid_flat_number CHECK (flat_number > 0 )
);

CREATE TABLE IF NOT EXISTS public.restaurant_address
(
    id SERIAL NOT NULL,
    restaurant_id INT REFERENCES restaurant(id) NOT NULL,
    address_id INT REFERENCES address(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.user_address
(
    id SERIAL NOT NULL,
    user_id INT REFERENCES public.USERS(id) NOT NULL,
    address_id INT REFERENCES address(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.card
(
    id SERIAL NOT NULL,
    card_number TEXT UNIQUE NOT NULL, --ДЕФОЛТНЫЙ НОМЕР КАРТЫ ЭТО СТРАННО
    user_id INT REFERENCES public.users(id) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT valid_card CHECK (LENGTH(card_number) > 0 and LENGTH(card_number) < 40)
);

CREATE TABLE IF NOT EXISTS public.user_card
(
    id SERIAL NOT NULL,
    user_id INT REFERENCES public.users(id) NOT NULL,
    card_id INT REFERENCES card(id) NOT NULL,
    PRIMARY KEY (id)
);
