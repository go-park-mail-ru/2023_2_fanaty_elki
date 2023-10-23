User:
{Id} -> {Password, Birthday, Icon}
{Username} ->  {Password, Birthday, Icon}
{Phone_number} ->  {Password, Birthday, Icon}
{Email} ->  {Password, Birthday, Icon}

Restaurant:
{Id}-> {Name, Rating, CommentsCount, Category, Icon}
{Name}->{Rating, CommentsCount, Category, Icon}

Restaurant_category:
{Id}->{Restaurant_id, Category_id}

Category:
{Id}->{Text}

Menu_Type:
{Id}->{Name, Restaurant_id}

Product_Menu_Type:
{Id}->{Menu_Type_id, Product_id}

Product:
{Id}->{ Price, Cooking_time}
{Name, Restraunt_id}->{Price, Cooking_time}

Order_Product:
{Id}->{Product_id, Order_id}

Order:
{Id}->{User_id, Date}

Comment:
{Id}->{Text, Restaurant_id, User_id, Rating}

Card:
{Id}->{Number, User_id}

Address:
{Id}->{City, Street, House_number, Flat_number}

Restaurant_Address:
{Id}->{Restaurant_id, Address_id}

User_Address:
{Id}->{User_id, Address_id}

База данных соответсвует всем нижеперечисленным требованиям.

1 нормальная форма: 
Каждая ячейка таблицы может хранить только одно значение
Все данные в одной колонке могут быть только одного типа
Каждая запись в таблице должна однозначно отличаться от других записей
2 нормальная форма:
Все неключевые атрибуты таблицы должны зависеть от первичного ключа
3 нормальная форма:
В таблицах отсутствует транзитивная зависимость.
Нормальной формы Бойса-Кодда
Ключевые атрибуты составного ключа не зависят от неключевых атрибутов




