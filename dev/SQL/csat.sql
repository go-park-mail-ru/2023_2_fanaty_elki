DROP TABLE IF EXISTS QUESTIONNAIRE CASCADE;
DROP TABLE IF EXISTS QUESTION CASCADE;
DROP TABLE IF EXISTS ANSWER CASCADE;
DROP TABLE IF EXISTS ADMIN CASCADE;

CREATE TABLE IF NOT EXISTS public.QUESTIONNAIRE
(
    ID serial NOT NULL,
    NAME text,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS public.QUESTION
(
    ID serial NOT NULL,
    QUESTIONNAIRE_ID int REFERENCES public.QUESTIONNAIRE(ID) NOT NULL, 
    TEXT text,
    ANSWER_TYPE INT,
    PRIMARY KEY (ID)
);
CREATE TABLE IF NOT EXISTS public.ANSWER
(
    ID serial NOT NULL,
    QUESTION_ID int REFERENCES public.QUESTION(ID) NOT NULL,
    TEXT text,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS public.ADMIN
(
    ID serial NOT NULL,
    USERNAME text,
    PASSWORD text,
    PRIMARY KEY (ID)
);

insert into QUESTIONNAIRE(name) VALUES('Общие вопросы');

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE) VALUES(1,'Насколько сильно Вы хотите порекомендовать наш сервис знакомым', 2);

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE) VALUES(1, 'Напишите что Вам понравилось, а что нет', 1);

insert into QUESTIONNAIRE(name) VALUES('Оформление заказа');

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE)
VALUES(1, 'Насколько удобно вам было оформлять заказ?', 2);

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE)
VALUES(1, 'Напишите что Вам понравилось, а что нет', 1);

insert into QUESTIONNAIRE(name) VALUES('Регистрация');

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE)
VALUES(1, 'Насколько удобно вам было регистрироваться?', 2);

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE)
VALUES(1, 'Напишите что Вам понравилось, а что нет', 1);

