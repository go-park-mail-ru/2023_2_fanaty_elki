DROP TABLE restaurant CASCADE;
DROP TABLE users CASCADE;
DROP TABLE product CASCADE;
DROP TABLE orders CASCADE;
DROP TABLE orders_product CASCADE;
DROP TABLE menu_type CASCADE;
DROP TABLE product_menu_type CASCADE;

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS public.USERS
(
    ID serial NOT NULL,
    USERNAME varchar UNIQUE NOT NULL  ,
	PASSWORD varchar NOT NULL,
	BIRTHDAY date,
	PHONE_NUMBER varchar UNIQUE NOT NULL,
    EMAIL varchar UNIQUE NOT NULL,
	ICON varchar default 'deficon',
	CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID),
    CONSTRAINT VALID_USERNAME CHECK ( LENGTH(USERNAME) >= 3 and LENGTH(USERNAME) <= 20 ),
    CONSTRAINT VALID_PASSWORD CHECK ( LENGTH(PASSWORD) >= 8 and LENGTH(PASSWORD) <= 30 ),
    CONSTRAINT VALID_PHONE CHECK ( PHONE_NUMBER ~* '^[+]?[0-9]{3,25}$'),
    CONSTRAINT VALID_EMAIL CHECK ( EMAIL ~* '\S*@\S*')
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON USERS
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.RESTAURANT
(
    ID serial NOT NULL,
    NAME varchar UNIQUE NOT NULL,
	RATING numeric(2,1) default 0.0 NOT NULL,
	COMMENTS_COUNT integer default 0 NOT NULL,
	CATEGORY varchar,
	ICON varchar default 'img/defaultIcon.webp' NOT NULL,
	CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID),
    CONSTRAINT VALID_RESTAURANT CHECK ( LENGTH(NAME) > 0 ),
    CONSTRAINT VALID_RATING CHECK ( RATING >= 0.0 AND RATING <= 5.0),
    CONSTRAINT VALID_COMMENTS_COUNT CHECK (COMMENTS_COUNT >= 0)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON RESTAURANT
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


insert into restaurant(name,rating,comments_count,icon,category)
values('Burger King',3.7,60,'img/burger_king.webp','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('MacBurger',3.8,69,'img/mac_burger.webp','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('Вкусно и точка',3.2,90,'img/tasty_and..webp','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('KFC',4.0,90,'img/kfc.webp','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('Шоколадница',4.5,90,'img/chocolate.jpeg','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('Корчма Тарас Бульба',5.0,90,'img/bulba.webp','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('Subway',3.0,90,'img/subway.jpeg','Fastfood');
insert into restaurant(name,rating,comments_count,icon,category)
values('Sushiwok',4.5,90,'img/sushi_wok.webp','Fastfood');



CREATE TABLE IF NOT EXISTS public.PRODUCT
(
    ID serial NOT NULL,
    NAME varchar NOT NULL, -- У блюда не может быть дефолтного значения, иначе как нам понять что это
    PRICE numeric(10,2) default '0.0' NOT NULL,
    COOKING_TIME INT default '0' NOT NULL,
    PORTION varchar default '1 шт' NOT NULL,
    DESCRIPTION TEXT,
    ICON varchar default 'deficon' NOT NULL,
    CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID),
    CONSTRAINT VALID_PRODUCT CHECK ( LENGTH(NAME) > 0 ),
    CONSTRAINT VALID_PRICE CHECK ( PRICE >= 0.0 ),
    CONSTRAINT VALID_TIME CHECK ( COOKING_TIME >= 0 )
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON PRODUCT
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

insert into product(name, price, cooking_time, portion, description)
values('Кинг Фри станд','109.99', 5,'106 г','Горячий и свежий картофель Кинг Фри® - золотистые и хрустящие ломтики отлично дополнят любой обед');

insert into product(name, price, cooking_time, portion, description)
values('Кинг Фри большой','144.99', 5,'160 г','Горячий и свежий картофель Кинг Фри® - золотистые и хрустящие ломтики отлично дополнят любой обед');

insert into product(name, price, cooking_time, portion, description)
values('Воппер','289.99', 15,'268 г','булочка для гамбургера с кунжутом(89 гр), котлета из говядины WHOPPER, майонез для салата, салат Айсберг, томаты, огурцы маринованные, лук репчатый, кетчуп томатный');

insert into product(name, price, cooking_time, portion, description)
values('Сибирский Кинг','349.99', 25,'262 г','Ешь в БК — спасай леса! Часть выручки мы направим на восстановление лесов Сибири, пострадавших от пожаров! 100%-я говядина, ароматный бекон, горчица и сливочный хрен на ржаной булочке. А ещё внутри сыр Чеддер, хрустящие маринованные огурчики, луковый конфитюр, ломтик томата и салат Айсберг.');

insert into product(name, price, cooking_time, portion, description)
values('Сибирский Кинг с курицей','349.99', 25,'258 г','Ешь в БК — спасай леса! Часть выручки мы направим на восстановление лесов Сибири, пострадавших от пожаров! Сочная курочка и ароматный бекон под бодрящей горчицей и сливочным хреном на ржаной булочке. К ним добавили: луковый конфитюр, сыр Чеддер, маринованные огурчики, салат Айсберг и ломтик томата.');

CREATE TABLE IF NOT EXISTS public.MENU_TYPE
(
    ID serial NOT NULL,
    NAME varchar default 'FOOD' NOT NULL,
    RESTAURANT_ID int REFERENCES RESTAURANT(ID) NOT NULL,
    PRIMARY KEY (ID),
    CONSTRAINT VALID_MENU_TYPE CHECK ( LENGTH(NAME) > 0 )
);

insert into menu_type(name,restaurant_id)
values('Популярное',1);
insert into menu_type(name,restaurant_id)
values('Новинки',1);


CREATE TABLE IF NOT EXISTS public.PRODUCT_MENU_TYPE
(
    ID serial NOT NULL,
    MENU_TYPE_ID int REFERENCES public.MENU_TYPE(ID) NOT NULL,
    PRODUCT_ID int REFERENCES public.PRODUCT(ID) NOT NULL,
    PRIMARY KEY (ID)
);

insert into product_menu_type(menu_type_id,product_id)
values(1,1);
insert into product_menu_type(menu_type_id,product_id)
values(1,2);
insert into product_menu_type(menu_type_id,product_id)
values(1,3);
insert into product_menu_type(menu_type_id,product_id)
values(2,4);
insert into product_menu_type(menu_type_id,product_id)
values(2,5);


CREATE TABLE IF NOT EXISTS public.ORDERS
(
    ID serial NOT NULL,
    USER_ID int REFERENCES public.USERS(ID) NOT NULL,
    ORDER_DATE TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    STATUS int default 0 NOT NULL,
    CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID),
    CONSTRAINT VALID_STATUS CHECK (STATUS >= 0 )
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON ORDERS
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.ORDERS_PRODUCT
(
    ID serial NOT NULL,
    PRODUCT_ID int REFERENCES public.PRODUCT(ID) NOT NULL,
    ORDER_ID int REFERENCES public.ORDERS(ID) NOT NULL,
    ITEM_COUNT INT default 1 NOT NULL,
    PRIMARY KEY (ID),
    CONSTRAINT VALID_COUNT CHECK ( ITEM_COUNT > 0 )
);