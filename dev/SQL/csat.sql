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

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE) VALUES(1, 'Оцените общее впечатление от сервиса', 2);

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE) VALUES(1, 'Насколько сильно Вы хотите порекомендовать наш сервис знакомым', 3);

insert into QUESTION(QUESTIONNAIRE_ID, text, ANSWER_TYPE) VALUES(1, 'Напишите что Вам понравилось, а что нет', 3);