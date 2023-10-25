User - Таблица Пользователей и его персональные данные ( ник, почта, телефон и тд)
Связанные таблицы: Card, User_address, Comment
User:
{Id} -> {Password, Birthday, Icon}
{Username} ->  {Password, Birthday, Icon}
{Phone_number} ->  {Password, Birthday, Icon}
{Email} ->  {Password, Birthday, Icon}

Restaurant - Таблица Ресторанов и его данные ( название, рейтинг, и тд)
Связанные таблицы: Restaurant_category, Restaurant_Address, Comment, Menu_Type, Product
Restaurant:
{Id}-> {Name, Rating, CommentsCount, Category, Icon}
{Name}->{Rating, CommentsCount, Category, Icon}

Category - Tаблица категорий ресторанов с текстовым названием категории
Связанные таблицы: Restaurant_category
Category:
{Id}->{Text}

Restaurant_category - Таблица-связка между ресторанами(Restaurant) и категориями (Category)
Restaurant_category:
{Id}->{Restaurant_id, Category_id}

Menu_Type - Раздел меню в ресторане
Связанные таблицы: Product_Menu_Type
Menu_Type:
{Id}->{Name, Restaurant_id}

Product - Продукт в ресторане ( с его ценой, временем приготовления и тд)
Связанные таблицы: Product_Menu_Type, Order_Product
Product:
{Id}->{ Price, Cooking_time}
{Name, Restraunt_id}->{Price, Cooking_time}

Product_Menu_Type - Таблица-связка между продуктами (Product) и разделами меню (Menu_Type)
Product_Menu_Type:
{Id}->{Menu_Type_id, Product_id}

Order - Таблица заказов
Связанные таблицы: Order_Product
Order:
{Id}->{User_id, Date}

Order_Product - Таблица-связка между продуктами (Product) и заказами (Order)
Order_Product:
{Id}->{Product_id, Order_id}

Comment - Таблица комментариев пользователей к ресторанам
Comment:
{Id}->{Text, Restaurant_id, User_id, Rating}

Card - Таблица карт пользователей
Card:
{Id}->{Number, User_id}

Address - Таблица адресов 
Связанные таблицы: Restaurant_Address, User_Address
Address:
{Id}->{City, Street, House_number, Flat_number}

Restaurant_Address - Таблица адресов ресторанов (таблица-связка между адресами и ресторанами)
Restaurant_Address:
{Id}->{Restaurant_id, Address_id}

User_Address - Таблица адресов ресторанов (таблица-связка между адресами и ресторанами)
User_Address:
{Id}->{User_id, Address_id}

Cсылка на ER-диаграммы: https://www.plantuml.com/plantuml/uml/fPJFxjem3CVlUOgS9_KD46BhXeQcCSJWN6N4WceqcH97MqLvzxN1C08cCFrEk5__-IJOpiIZSXYQlG23DRPoxS79Wv3odPpf_gSeeZt8HZKGFnRA-z_M3UuFjSETLkr9tra3bGdD3t1H-DIeUjzRiuqruYWcOFI0keCkeLU2KmOPzFoCNIgipJQQfMsQFjUNDXYCokk8gbG8M-jQXdKLLMiv6as_ZCSd8ELSoHbOXh22FvXB_GecLAls5sJ_2UwOcQ8BPtNOO9KhVl6bDwVS_mCzExjJmU0Tfzu0DTV-x1jTOTsdv6ZVd7uuh3VLFog4fwMgnySdJrgLib9j7iS_nwBXVJGSkJbahaPhaGE-rzg1zmomxDtH-8uEuCJ-R8DFjdYwLqZVBsuO3ynUZ5BuCUCUgGmakjzGlbHECsMfimo1gtNrfdLTKNG7xWaIZ2QFGtWdcHYgz6zNnSKYr_sg2rBE0KyZqbqGyOjthVsD9r5ep9dZaoUDLauBwX9fIPjlVCoqfQ3NWvBRS6JgxLy0
