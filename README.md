go run main.go

Список маршрутов с параметрами:

| POST-запрос | параметр запроса |
|----------------|------------------------------|
| `/order` | `{"districtId":5,"price":[100.50,25.20]}` |
| `/pay` | `{"orderId":4,"price":[{"payment":"cash","value":100.20},{"payment":"card","value":25.50}]}`|
| `/click` | `{"entryid":8}` |
| `/delivered` | `{"orderId":4}` |
