# test_junior
## Тестовое задание на позицию junior golang developer

- Реализован эндпоинт, принимающий query-параметр в формате YYYY-MM-DD.
  Возвращаемое значение - количество дней до наступления указанной даты.
- Документация API доступна по URL `/docs`

### Запуск приложения

Локально:
```bash
go run ./cmd/dater/main.go
```

Запуск тестов:
```bash
go test -v
```

Запуск в Docker:
```bash
docker build -t app:0.0.1 . && docker run -p 8000:8000 app:0.0.1
```
