DROP TABLE IF EXISTS restaurant CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS product CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS orders_product CASCADE;
DROP TABLE IF EXISTS menu_type CASCADE;
DROP TABLE IF EXISTS product_menu_type CASCADE;
DROP TABLE IF EXISTS cart CASCADE;
DROP TABLE IF EXISTS cart_product CASCADE;
DROP TABLE IF EXISTS address CASCADE;
DROP TABLE IF EXISTS orders_address CASCADE;
DROP TABLE IF EXISTS category CASCADE;
DROP TABLE IF EXISTS restaurant_category CASCADE;
DROP TABLE IF EXISTS comment CASCADE;
DROP TABLE IF EXISTS promo CASCADE;
DROP TABLE IF EXISTS user_promo CASCADE;
DROP TABLE IF EXISTS users_address CASCADE;
DROP TABLE IF EXISTS cart_promo CASCADE;


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
    CONSTRAINT VALID_PHONE CHECK ( PHONE_NUMBER ~* '^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{6,10}$'),
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
	ICON varchar default 'deficon' NOT NULL,
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


insert into restaurant(name, icon)
values('Burger King', 'img/burger_king.jpg');
insert into restaurant(name, icon)
values('Якитория', 'img/yakitoria.jpg');
insert into restaurant(name, icon)
values('Вкусно и точка', 'img/tasty_and..jpg');
insert into restaurant(name, icon)
values('KFC', 'img/kfc.jpg');
insert into restaurant(name, icon)
values('Шоколадница', 'img/chocolate.jpeg');
insert into restaurant(name, icon)
values('Корчма Тарас Бульба', 'img/bulba.jpg');
insert into restaurant(name, icon)
values('Subway', 'img/subway.jpeg');
insert into restaurant(name, icon)
values('Sushiwok', 'img/sushi_wok.png');

CREATE TABLE IF NOT EXISTS public.category
(
    id SERIAL NOT NULL,
    name TEXT UNIQUE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT valid_text CHECK ( LENGTH(name) > 0 and LENGTH(name) < 40 )
);

insert into category(name)
values('Бургеры');
insert into category(name)
values('Суши');
insert into category(name)
values('Завтраки');
insert into category(name)
values('Обед');
insert into category(name)
values('Русская');
insert into category(name)
values('Кофе');

CREATE TABLE IF NOT EXISTS public.restaurant_category
(
    id SERIAL NOT NULL,
    restaurant_id INT REFERENCES public.restaurant(id) NOT NULL,
    category_id INT REFERENCES public.category(id) NOT NULL,
    PRIMARY KEY (id)
);

insert into restaurant_category(restaurant_id, category_id)
values(1,1);
insert into restaurant_category(restaurant_id, category_id)
values(1,4);
insert into restaurant_category(restaurant_id, category_id)
values(2,2);
insert into restaurant_category(restaurant_id, category_id)
values(3,1);
insert into restaurant_category(restaurant_id, category_id)
values(3,4);
insert into restaurant_category(restaurant_id, category_id)
values(4,1);
insert into restaurant_category(restaurant_id, category_id)
values(4,3);
insert into restaurant_category(restaurant_id, category_id)
values(4,6);
insert into restaurant_category(restaurant_id, category_id)
values(5,1);
insert into restaurant_category(restaurant_id, category_id)
values(5,3);
insert into restaurant_category(restaurant_id, category_id)
values(5,6);
insert into restaurant_category(restaurant_id, category_id)
values(6,4);
insert into restaurant_category(restaurant_id, category_id)
values(6,5);
insert into restaurant_category(restaurant_id, category_id)
values(7,4);
insert into restaurant_category(restaurant_id, category_id)
values(8,2);



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

insert into product(name, price, cooking_time, portion, icon, description)
values('Кинг Фри станд','110', 5,'106 г','img/king_stand.png', 'Горячий и свежий картофель Кинг Фри® - золотистые и хрустящие ломтики отлично дополнят любой обед');

insert into product(name, price, cooking_time, portion, icon, description)
values('Кинг Фри большой','145', 5,'160 г', 'img/king_big.png', 'Горячий и свежий картофель Кинг Фри® - золотистые и хрустящие ломтики отлично дополнят любой обед');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Воппер','290', 15,'268 г', 'img/vopper.png', 'булочка для гамбургера с кунжутом(89 гр), котлета из говядины WHOPPER, майонез для салата, салат Айсберг, томаты, огурцы маринованные, лук репчатый, кетчуп томатный');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Сибирский Кинг','350', 25,'262 г', 'img/sib.png','Ешь в БК — спасай леса! Часть выручки мы направим на восстановление лесов Сибири, пострадавших от пожаров! 100%-я говядина, ароматный бекон, горчица и сливочный хрен на ржаной булочке. А ещё внутри сыр Чеддер, хрустящие маринованные огурчики, луковый конфитюр, ломтик томата и салат Айсберг.');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Сибирский Кинг с курицей','350', 25,'258 г', 'img/sib_chick.png','Ешь в БК — спасай леса! Часть выручки мы направим на восстановление лесов Сибири, пострадавших от пожаров! Сочная курочка и ароматный бекон под бодрящей горчицей и сливочным хреном на ржаной булочке. К ним добавили: луковый конфитюр, сыр Чеддер, маринованные огурчики, салат Айсберг и ломтик томата.');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Пирожок Абрикосовый','77', 5,'82 г', 'img/abr_pir.png','Пирожок с абрикосом – это настоящее лакомство для любителей фруктов. Ароматный, сочный и сладкий, он обязательно понравится всем, кто его попробует. Горячий пирожок с абрикосом уже ждет тебя в KFC. Состав: Масло растительное; Пирожок с начинкой "Абрикосовый"');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Кофе Капучино со вкусом груши с корицей большой','195', 15,'311 мл', 'img/cofe_pear.png','Вкус сезона! Кофейный напиток с ароматом эспрессо, нежной густой молочной пенкой и нотками груши с корицей. Состав: Кипяченая вода; Молоко питьевое; Кофе жаренный в зернах; Сироп со вкусом и ароматом груши и корицы "Груша-Корица"');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Пирожное Макарон Фисташка и Манго-Маракуйя','130', 5,'24 г', 'img/pir_mar.png','Нежные, чуть хрустящие миндальные печенья, соединенные кремовой начинкой и покрытые тонкой, хрупкой глазурью, никого не оставят равнодушными.');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Комбо на компанию','1100', 35,'1258 г', 'img/combo_company.png','3 Шефбургера ориг. / остр. + 2 Твистера ориг + Картофель Фри Малый 5 шт. + 5 Соусов на выбор');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Комбо c Шефбургер Де Люкс','300', 25,'558 г', 'img/combo_chef.png','Шефбургер Де Люкс ориг. / остр. + 9 Наггетсов + Соус на выбор');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Осенний салат с печёной тыквой и свёклой','411', 25,'210 г', 'img/autome_salad.png','Запеченная тыква, свекла отварная, сыр брынза, айсберг, рукола, грецкие орехи, горчичный соус, соус песто, крем бальзамик');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Салат из печёных баклажанов с соусом гамадари','491', 25,'200 г', 'img/salad_garry.png','Запечённые баклажаны, свежие помидоры, салат айсберг, рукола, свежий шпинат, кинза, грецкие орехи, соус гамадари');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Чизбургер','491', 15,'230 г', 'img/cheeseburger.png','Котлета из мраморной говядины, чеддер, кетчуп, горчица, огурцы маринованные, помидоры, пшеничная булочка');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Боул с индейкой','451', 25,'200 г', 'img/boul_ind.png','Индейка, овощная сальса (перец болгарский, огурец свежий, авокадо, манго, кинза, мята, соль, сок лайма, помидоры), смесь отварных круп, айсберг, романо, соус медово-горчичный (горчица зернистая, мед, масло оливковое, лимонный сок), помидоры черри, авокадо, лайм, шпинат, рукола, мята');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Боул с креветками','501', 25,'200 г', 'img/boul_crev.png','Креветки, овощная сальса (перец болгарский, огурец свежий, авокадо, манго, кинза, мята, соль, сок лайма, помидоры), смесь отварных круп, айсберг, романо, соус медово-горчичный (горчица зернистая, мед, масло оливковое, лимонный сок), помидоры черри, авокадо, лайм, шпинат, рукола, мята');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Ролл Якитория','638', 25,'245 г', 'img/roll_yakitoria.png','Ролл в слайсах тунца и лосося, спайси угорь, сыр, огурец, авокадо, снежный краб, рис для суси, нори, масаго, васаби (8 шт). ');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Ролл Калифорния','638', 35,'215 г', 'img/california.png','Мясо краба, огурец, авокадо, рис для суси, нори, тобико, васаби (8 шт).');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Сет с лососем','2178', 45,'730 г', 'img/set_salmon.png','О рицу маки (6 шт), Киото рору 8 шт., Филадельфия 8 шт., Сякэ рору (6 шт) - 4 порции роллов 28 шт');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Гаспачо','458', 25,'296 г', 'img/gaspacho.png','Холодный суп из свежих овощей с острым соусом из трав на оливковом масле, хрустящие крутоны.');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Куксу','438', 35,'675 г', 'img/cucsu.png','Яичная лапша, говядина, омлет, капуста кимчи, редис, огурцы, кинза и кунжут в охлажденном кисло-остром бульоне. ');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Чизбургер','201', 5,'108 г', 'img/cheesebur.png','Домашняя булочка, котлета из натуральной говядины, сыр чеддер, маринованные огурцы, кетчуп, горчичный соус');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Бургер Классика','451', 5,'264 г', 'img/classic_burger.png','Классика, которая придется по вкусу каждому. Фирменный бургер с домашней булочкой, сочной котлетой с ломтиками сыра чеддер, которые идеально дополняют свежие листья салата, помидоры, маринованные огурчики и красный лук');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Наггетсы куриные','131', 5,'70 г', 'img/naggets.png','');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Круас Бургер','431', 15,'200 г', 'img/cruas.png','Круассан, котлета из говядины, маринованные огурцы, сыр чеддер, сыр моцарелла, листовой салат, свежий лук, соус техасский');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Spicy бургер','491', 15,'270 г', 'img/spicy_burger.png','Пикантная новинка! Домашняя булочка, сочная котлета из говядины, обжаренная с сыром чеддер, свежий салат, гавайский соус, ломтики помидора, маринованные огурчики и незабываемое сочетание острого соуса и джема из черной смородины - настоящее гастрономическое удовольствие');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Фирменный борщ Корчма','471', 25,'300 г', 'img/borsh.png','Говядина, свинина, капуста, картофель, лук репчатый, морковь, перец болгарский, помидоры, свекла, фасоль, чеснок. Подается с домашней сметаной и пампушкой');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Борщ щавелевый на домашнем бульоне','431', 25,'300 г', 'img/shi.png','На домашнем курином бульоне, подается с курицей, яйцом и сметаной');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Котлета по-киевски','550', 15,'200 г', 'img/cotletakiev.png','Легендарная котлета из курицы со сливочным маслом, перцем. Подается с картофельным пюре, помидором и луком фри');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Лаваш','51', 1,'50 г', 'img/lavash.png','Лаваш');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Шашлык из телятины','631', 25,'170 г', 'img/shash_tel.png','Шашлык из телятины');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Джуниор саб Двойной Сыр','140', 5,'230 г', 'img/jun_sub_double_cheese.png','Джуниор саб Двойной Сыр готовится на маленьком белом итальянском хлебе с сыром Чеддер и Моцарелла. Из овощей: листья салата, томаты, свежие огурцы и 2 соуса на выбор (рекомендуем кетчуп и майонез)');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Джуниор саб Салями-пепперони','156', 5,'246 г', 'img/jun_sub_sal_pep.png','Джуниор саб Салями-пепперони готовится на маленьком белом итальянском хлебе, салями (3 шт.), пепперони (3 шт.), сыр Чеддер (1 ломтик). Из овощей: листья салата, томаты, свежие огурцы и 2 соуса на выбор (рекомендуем кетчуп и майонез)');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Джуниор саб Индейка-ветчина','156', 5,'242 г', 'img/jun_sab_ind.png','Джуниор саб Индейка-ветчина готовится на маленьком белом итальянском хлебе. Ветчина (2 шт.), индейка (1 шт.), сыр чеддер (1 шт.), из овощей: листья салата, томаты, свежие огурцы и 2 соуса на выбор (рекомендуем кетчуп и майонез)');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Ролл Овощной','198', 5,'182 г', 'img/roll_veg.png','Ролл, завернутый в пшеничную лепешку с соусом и овощами на выбор');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Ролл Индейка','276', 5,'227 г', 'img/roll_ind.png','Мясо индейки и свежие овощи в пшеничной лепешке в комбинации с соусом на выбор');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Сет Жаркий сезон','2370', 25,'1,05 кг', 'img/set_hot_season.png','Запеченный ролл Аяши, запеченный ролл Румяный, запеченный ролл Сырный, запеченный ролл Хот Фиш, запеченный ролл Яки шиитаке, запеченный ролл Крабик Хот');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Соус соевый','101', 1,'100 г', 'img/soe_souce.png','');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Сет Филомания','2580', 35,'692 г', 'img/set_filomania.png','Ролл Филадельфия в масаго, ролл Филадельфия, ролл Калифорния в кунжуте, ролл с огурцом');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Мидии запеченные','696', 15,'150 г', 'img/midii_zap.png','Яки соус, мидии, лимон');

insert into product(name, price, cooking_time, portion, icon,  description)
values('Мидии спайси','696', 15,'150 г', 'img/midii_spicy.png','Спайси соус, мидии, лимон');

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

insert into menu_type(name,restaurant_id)
values('Новинки',4);
insert into menu_type(name,restaurant_id)
values('Ланчи и Комбо',4);

insert into menu_type(name,restaurant_id)
values('Осеннее предложение',5);
insert into menu_type(name,restaurant_id)
values('Завтраки',5);

insert into menu_type(name,restaurant_id)
values('Популярные блюда',2);
insert into menu_type(name,restaurant_id)
values('Fall In Love',2);

insert into menu_type(name,restaurant_id)
values('Популярные блюда',3);
insert into menu_type(name,restaurant_id)
values('Бургеры классические',3);

insert into menu_type(name,restaurant_id)
values('Популярные блюда',6);
insert into menu_type(name,restaurant_id)
values('Шашлыки',6);

insert into menu_type(name,restaurant_id)
values('Популярные блюда',7);
insert into menu_type(name,restaurant_id)
values('Роллы',7);

insert into menu_type(name,restaurant_id)
values('Популярные блюда',8);
insert into menu_type(name,restaurant_id)
values('Салаты и закуски',8);




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

insert into product_menu_type(menu_type_id,product_id)
values(3,6);
insert into product_menu_type(menu_type_id,product_id)
values(3,7);
insert into product_menu_type(menu_type_id,product_id)
values(3,8);
insert into product_menu_type(menu_type_id,product_id)
values(4,9);
insert into product_menu_type(menu_type_id,product_id)
values(4,10);

insert into product_menu_type(menu_type_id,product_id)
values(5,11);
insert into product_menu_type(menu_type_id,product_id)
values(5,12);
insert into product_menu_type(menu_type_id,product_id)
values(5,13);
insert into product_menu_type(menu_type_id,product_id)
values(6,14);
insert into product_menu_type(menu_type_id,product_id)
values(6,15);

insert into product_menu_type(menu_type_id,product_id)
values(7,16);
insert into product_menu_type(menu_type_id,product_id)
values(7,17);
insert into product_menu_type(menu_type_id,product_id)
values(7,18);
insert into product_menu_type(menu_type_id,product_id)
values(8,19);
insert into product_menu_type(menu_type_id,product_id)
values(8,20);

insert into product_menu_type(menu_type_id,product_id)
values(9,21);
insert into product_menu_type(menu_type_id,product_id)
values(9,22);
insert into product_menu_type(menu_type_id,product_id)
values(9,23);
insert into product_menu_type(menu_type_id,product_id)
values(10,24);
insert into product_menu_type(menu_type_id,product_id)
values(10,25);

insert into product_menu_type(menu_type_id,product_id)
values(11,26);
insert into product_menu_type(menu_type_id,product_id)
values(11,27);
insert into product_menu_type(menu_type_id,product_id)
values(11,28);
insert into product_menu_type(menu_type_id,product_id)
values(12,29);
insert into product_menu_type(menu_type_id,product_id)
values(12,30);

insert into product_menu_type(menu_type_id,product_id)
values(13,31);
insert into product_menu_type(menu_type_id,product_id)
values(13,32);
insert into product_menu_type(menu_type_id,product_id)
values(13,33);
insert into product_menu_type(menu_type_id,product_id)
values(14,34);
insert into product_menu_type(menu_type_id,product_id)
values(14,35);

insert into product_menu_type(menu_type_id,product_id)
values(15,36);
insert into product_menu_type(menu_type_id,product_id)
values(15,37);
insert into product_menu_type(menu_type_id,product_id)
values(15,38);
insert into product_menu_type(menu_type_id,product_id)
values(16,39);
insert into product_menu_type(menu_type_id,product_id)
values(16,40);




CREATE TABLE IF NOT EXISTS public.ORDERS
(
    id            SERIAL NOT NULL,
    user_id       INT REFERENCES public.USERS(ID) NOT NULL,
    order_date    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    status        INT default 0 NOT NULL,
    price         INT DEFAULT 0 NOT NULL, 
    delivery_time INT default 0 NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	updated_at    TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (id),
    CONSTRAINT valid_status CHECK (status >= 0 AND status <= 2)
    
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

CREATE TABLE IF NOT EXISTS public.CART
(
    ID serial NOT NULL,
    USER_ID int REFERENCES public.USERS(ID) NOT NULL,
    CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID)
);
 
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON CART
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.CART_PRODUCT
(
    ID serial NOT NULL,
    PRODUCT_ID int REFERENCES public.PRODUCT(ID) NOT NULL,
    CART_ID int REFERENCES public.CART(ID) NOT NULL,
    ITEM_COUNT INT default 1 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS public.ADDRESS
(
    ID serial NOT NULL,
    CITY text NOT NULL,
    STREET text NOT NULL,
    HOUSE_NUMBER text NOT NULL,
    FLAT_NUMBER INT default 0, -- бывают ДОМА БЕЗ КВАРТИР
    CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID),
    CONSTRAINT VALID_CITY CHECK (LENGTH(CITY) > 0 ),
    CONSTRAINT VALID_STREET CHECK (LENGTH(STREET) > 0 ),
    CONSTRAINT VALID_HOUSE_VALID CHECK (LENGTH(HOUSE_NUMBER) > 0 ),
    CONSTRAINT VALID_FLAT_NUMBER CHECK (FLAT_NUMBER >= 0 )
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON address
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.ORDERS_ADDRESS
(
    ID serial NOT NULL,
    ORDERS_ID INT REFERENCES ORDERS(ID) NOT NULL,
    ADDRESS_ID INT REFERENCES ADDRESS(ID) NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS public.users_address
(
    ID         SERIAL NOT NULL,
    user_id    INT REFERENCES users(id) NOT NULL,
    address_id INT REFERENCES address(id) NOT NULL,
    current    BOOLEAN NOT NULL DEFAULT true,
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS public.comment
(
    id            SERIAL                                 NOT NULL,
    content       TEXT,
    rating        INT                                    NOT NULL,
    restaurant_id INT REFERENCES restaurant(id)          NOT NULL,
    user_id       INT REFERENCES users(id)               NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	updated_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    PRIMARY KEY (ID),
    CONSTRAINT valid_rating CHECK (rating >= 1 AND rating <= 5)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON comment
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.promo
(
    id            SERIAL                                 NOT NULL,
    code          TEXT,
    promo_type    INT                                    NOT NULL,
    sale          INT,
    restaurant_id INT REFERENCES restaurant(id),
    active_from   TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    active_to     TIMESTAMP WITH TIME ZONE               NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	updated_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON promo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

insert into promo(code, promo_type, sale, restaurant_id, active_from, active_to)
values('KORCHMA15', 0, 15, 6, '2023-12-5', '2023-12-25');
insert into promo(code,  promo_type, sale, restaurant_id, active_from, active_to)
values('SUBWAY35',0, 35, 7, '2023-12-5', '2023-12-25');
insert into promo(code,  promo_type, sale, restaurant_id, active_from, active_to)
values('YAKITORIA50',0, 50, 2, '2023-12-5', '2023-12-25');
insert into promo(code,  promo_type,  restaurant_id, active_from, active_to)
values('BURGERKINGFREE', 1 , 1, '2023-12-5', '2023-12-25');
insert into promo(code,  promo_type, sale, restaurant_id, active_from, active_to)
values('VKUSNO20', 0 , 20, 3 , '2023-12-5', '2023-12-10');
insert into promo(code,  promo_type, active_from, active_to)
values('PRINESYFREE', 1 , '2023-12-5', '2023-12-25');
insert into promo(code,  promo_type, sale, active_from, active_to)
values('SHYSH30', 0, 30, '2023-12-5', '2023-12-25');


CREATE TABLE IF NOT EXISTS public.user_promo
(
    id            SERIAL                                 NOT NULL,
    user_id       INT REFERENCES users(id)               NOT NULL,
    promo_id      INT REFERENCES promo(id)               NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	updated_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON user_promo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.cart_promo
(
    id            SERIAL                                 NOT NULL,
    cart_id       INT REFERENCES cart(id)                NOT NULL,
    promo_id      INT REFERENCES promo(id)               NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	updated_at    TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON cart_promo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
