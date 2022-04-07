# Микросервис для работы с балансом пользователей

Данный микросервис с HTTP API был разработан по условиям [тестового задания в Авито](https://github.com/avito-tech/autumn-2021-intern-assignment)

Используемые технологии:

- Postgres
- Docker

## Эндпоинты и примеры их работы

HTTP API принимает запросы и возвращает ответы в формате JSON

### /deposit

Метод: POST

Пополняет баланс пользователя, если account_id или balance_id отсутствует, создаёт новую строку в БД с этими id

![Deposit](https://user-images.githubusercontent.com/78679173/162110721-9cee1656-d4ba-4867-bbb4-3678a5432c41.png)

### /withdraw

Метод: POST

Снимает деньги с баланса пользователя

![Withdraw](https://user-images.githubusercontent.com/78679173/162111114-399230b7-d202-4502-8f09-93d615f87d26.png)

### /transfer

Метод: POST

Переводит деньги от одного пользователя другому

![Transfer](https://user-images.githubusercontent.com/78679173/162111945-ab426ca8-4989-4146-bfa9-00dea75608c4.png)

### /balance

Метод: GET

Возвращает баланс

![Balance](https://user-images.githubusercontent.com/78679173/162112056-cc8f73d1-3bff-403f-b702-9ae37facf059.png)

Пример ответа:

![BalanceResponse](https://user-images.githubusercontent.com/78679173/162112209-25f5f0be-83c7-49e0-bb52-a4908cecb3b7.png)

Также можно перевести баланс в другую валюту, добавив в query параметр currency=код_валюты

Для конвертации валют используется [ExchangeRate-API](https://www.exchangerate-api.com/)

Для использования своего API Key нужно поменять URL в .env в переменной EXCHANGE_RATE_API_URL

Пример:

![BalanceCurr](https://user-images.githubusercontent.com/78679173/162114317-59e9975e-e9f7-47f6-8f4d-f0448f26cb8d.png)

![BalanceCurrResponse](https://user-images.githubusercontent.com/78679173/162114320-e6fd7e50-462b-41aa-bbec-8d90fa62640f.png)


### /history

Метод: GET

Возвращает всю историю транзакций balance_id (т.е. самого кошелька)

Поддерживает сортировку по возрастанию и убыванию по дате транзакции и её сумме (указывать её необязательно):

- "date_asc", "date_desc" - сортировка транзакций по дате
- "value_asc", "value_desc" - сортировка транзакций по сумме

![History](https://user-images.githubusercontent.com/78679173/162168782-ad6f1c06-b721-4f21-b56d-e000cca2ef96.png)

Пример ответа:

![HistoryResponse](https://user-images.githubusercontent.com/78679173/162168789-a16d8aa4-bd75-4bb9-add9-01a2ded9e15c.png)

### /history/{page}

Метод: GET

Возвращает {page} страницу из истории транзакций balance_id (т.е. самого кошелька)

Сортировка работает аналогично /history

Количество записей на страницу можно поменять в .env перед деплоем, изменив MAX_HISTORY_PAGE_LEN

![HistoryPage](https://user-images.githubusercontent.com/78679173/162114034-0ec76926-b0b5-4793-9df9-2fe7e3e19ec8.png)

Пример ответа:

![HistoryResponse](https://user-images.githubusercontent.com/78679173/162112437-16e4559f-603e-4f26-aba1-d665a55d5c94.png)
