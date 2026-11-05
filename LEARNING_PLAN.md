# 🚀 Go Backend Learning Roadmap: Smart Finance & Subscription Tracker

Этот документ фиксирует контекст, формат обучения, архитектуру и текущий статус разработки Telegram-бота для учета расходов и подписок.

---

## 🧑‍💻 Формат обучения и роли
* **Студент (User)**: пишет весь код самостоятельно, изучает конструкции и идиомы языка Go, задает вопросы по сложным моментам.
* **Ментор / Tech Lead (AI Assistant)**: 
  * Ведет студента по этапам.
  * Объясняет теорию (concurrency, interfaces, memory management, project layout).
  * Проводит подробное **код-ревью** написанного студентом кода (выявляет race conditions, неидиоматичный Go, утечки ресурсов, ошибки архитектуры).
  * **Не пишет готовые решения за студента**, а ставит задачи и спецификации.

---

## 🛠️ Стек и ключевые механики Go в проекте
* **Язык**: Go 1.21+
* **Архитектура**: Clean Architecture (Domain -> Repository -> Service -> Delivery).
* **Многопоточность (Concurrency)**: Goroutines, Channels, `sync.RWMutex`, `sync.WaitGroup`, `time.Ticker`, проверка через `go test/run -race`.
* **Управление жизненным циклом**: `context.Context` (таймауты, отмена, Graceful Shutdown).
* **Сетевое взаимодействие**: `net/http` клиент, сериализация `encoding/json`.
* **Базы данных**: `database/sql` (сначала in-memory, затем SQLite / PostgreSQL).
* **Интерфейс (UI)**: Telegram Bot API (`github.com/go-telegram-bot-api/telegram-bot-api/v5`).

---

## 📂 Структура проекта
```text
go-finance-bot/
├── cmd/
│   └── bot/
│       └── main.go              # Точка входа приложения
├── internal/                    # Приватный код (недоступен внешним модулям)
│   ├── config/                  # Конфигурация (.env, переменные окружения)
│   ├── domain/                  # Модели данных, кастомные типы, доменные ошибки
│   ├── repository/              # Интерфейсы и реализации хранилищ (in-memory, sql)
│   ├── service/                 # Бизнес-логика приложения
│   └── delivery/                # Транспортный слой (Telegram-хэндлеры и роутинг)
├── .env                         # Локальные переменные окружения (в .gitignore!)
├── .gitignore
├── go.mod
├── go.sum
└── LEARNING_PLAN.md             # Этот файл контекста
```

---

## 📈 Прогресс и этапы разработки

- [x] **Этап 0: Подготовка окружения и Git**
  - Инициализация Git-репозитория, подключение удаленного репозитория на GitHub.
  - Настройка правильного `.gitignore` (секреты `.env`, бинарники, сессии).

- [x] **Этап 1: Конфигурация и точка входа (`internal/config`)**
  - Реализован пакет `internal/config`.
  - Чтение `.env` через `godotenv`, валидация обязательного `TELEGRAM_BOT_TOKEN`.
  - Избавление от stuttering (`config.Load()`), использование guard clauses.

- [x] **Этап 2: Доменный слой (`internal/domain`)**
  - Кастомный тип `Category` со строгой типизацией констант (`category.go`).
  - Доменные сущности `Expense` и `Subscription` со структурированными полями.
  - Доменные ошибки в `errors.go` (`ErrNotFound`, `ErrInvalidAmount`, etc.).
  - Методы валидации `Validate() error` с pointer receiver `*Expense` и `*Subscription`.

- [x] **Этап 3.1: Контракты хранилища (`internal/repository/repository.go`)**
  - Объявлены интерфейсы `ExpenseRepository` и `SubscriptionRepository` с поддержкой `context.Context`.

- [ ] **Этап 3.2 (ТЕКУЩИЙ ШАГ): In-Memory хранилище и многопоточность (`internal/repository/memory.go`)**
  - **Задача**: написать потокобезопасную реализацию интерфейса `ExpenseRepository` в оперативной памяти.
  - **Ключевые концепции Go**:
    - Структура хранилища с `sync.RWMutex`, мапой `map[int64]*domain.Expense` и счетчиком `nextID int64`.
    - Разница между полным локом `mu.Lock()` (для `Save`, `Delete`) и локом на чтение `mu.RLock()` (для `GetByID`, `ListByUserID`).
    - Идиома `defer mu.Unlock()`.
    - Неявная реализация интерфейса (Duck Typing).

- [ ] **Этап 4: Сервисный слой и бизнес-логика (`internal/service`)**
  - Подсчет статистики расходов по категориям за период (месяц/день).
  - Конвертация и валидация перед сохранением.
  - Оборачивание ошибок через `fmt.Errorf("...: %w", err)`.

- [ ] **Этап 5: Слой доставки Telegram (`internal/delivery/telegram`)**
  - Вынос логики бота из `main.go` в модульные хэндлеры.
  - Роутинг команд: `/start`, `/add`, `/expenses`, `/subscribe`, `/stats`.
  - Клавиатуры и интерактивные кнопки.

- [ ] **Этап 6: Фоновые процессы и планировщик (Concurrency)**
  - Фоновый воркер на горутинах и каналах.
  - Использование `time.NewTicker` для проверки приближающихся подписок и отправки напоминаний в Telegram.
  - Корректный **Graceful Shutdown** через сигналы ОС и `context.WithCancel`.

- [ ] **Этап 7: Сетевое взаимодействие (HTTP-клиент и JSON)**
  - Интеграция с внешним REST API курсов валют через стандартный `net/http` клиент.
  - Парсинг JSON-ответов (`encoding/json`), дедлайны через контекст.

- [ ] **Этап 8: Персистентность (База данных SQLite / PostgreSQL)**
  - Замена `memory.go` на реализацию через `database/sql` без изменения остального кода.

- [ ] **Этап 9: Тестирование (Unit-тесты и Table-Driven Testing)**
  - Написание каноничных табличных тестов для бизнес-логики и валидации.

---

## 🔄 Как продолжить разработку на новом компьютере (Windows):
1. Установить Go и Git.
2. Клонировать репозиторий:
   ```bash
   git clone https://github.com/rruslek/go-finance-bot.git
   cd go-finance-bot
   ```
3. Создать `.env` файл в корне проекта:
   ```env
   TELEGRAM_BOT_TOKEN="твой_токен_бота"
   ```
4. Открыть проект в Antigravity.
5. Написать ассистенту:
   > *"Привет! Я развернул проект на Windows. Продолжаем работу по LEARNING_PLAN.md. Текущий шаг — Этап 3.2 (реализация MemoryStorage в `internal/repository/memory.go`)."*

