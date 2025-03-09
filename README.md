## Delivery Service overview

The repository of delivery service

## Endpoints

Method | Path                             | Description                                   |                                                                         
---    |----------------------------------|------------------------------------------------
GET    | `/health`                        | Health page                                   |
GET    | `/metrics`                       | Страница с метриками                          |
GET    | `/v1/couriers/list`              | Получение всех курьеров                       |
GET    | `/v1/deliveres/list`             | Получение всех доставок                       |
POST   | `/v1/couriers/book`              | Забронировать курьера на доставку             |