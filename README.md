go run main.go

Список маршрутов с параметрами:
| LEFT | CENTER | RIGHT |
|----------------|:---------:|-------------------------------------:|
| `/order` | order | `{"districtId":5,"price":[100.50,25.20]}` |
| `/pay` | receipt | `{"orderId":4,"price":[{"payment":"cash","value":100.20},{"payment":"card","value":25.50}]}`|
| `/click` | click | `{"entryid":8}` |
| `/delivered` | delivered | `{"orderId":4}` |

| LEFT | CENTER | RIGHT |
|----------------|:---------:|----------------:|
| По левому краю | По центру | По правому краю |
| текст | текст | текст |
