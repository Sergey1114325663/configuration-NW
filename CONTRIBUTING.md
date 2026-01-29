# Contributing to Network Configuration Manager

Спасибо за интерес к контрибуции в Network Configuration Manager! Этот документ описывает как начать разработку.

## Начало Работы

### 1. Форкируйте репозиторий
```bash
git clone https://github.com/your-username/network-config.git
cd network-config
```

### 2. Создайте ветку для вашей фичи
```bash
git checkout -b feature/your-feature-name
# или для бага фикса
git checkout -b fix/your-bug-fix
```

### 3. Установите зависимости
```bash
go mod download
go mod tidy
```

### 4. Сделайте изменения
- Следуйте Go code style guidelines
- Добавляйте unit тесты для новых функций
- Обновляйте документацию

### 5. Запустите тесты
```bash
go test ./... -v -race -cover
```

### 6. Запустите линтер
```bash
go fmt ./...
go vet ./...
```

## Code Style

### Go Код
- Используйте `gofmt` для форматирования
- Следуйте [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Максимальная длина строки: 120 символов
- Используйте говорящие имена переменных

Пример:
```go
// Хорошо
func (vm *VLANManager) CreateVLAN(vlan *VLANInfo) error {
    vm.mu.Lock()
    defer vm.mu.Unlock()
    
    // implementation
}

// Плохо
func (v *VLANManager) create(v2 *VLANInfo) error {
    // implementation
}
```

### Тесты
- Каждая функция должна иметь тесты
- Используйте таблично-ориентированные тесты где возможно
- Тесты должны быть независимы друг от друга

```go
func TestVLANManager(t *testing.T) {
    tests := []struct {
        name    string
        id      string
        want    bool
        wantErr bool
    }{
        {"valid", "10", true, false},
        {"invalid", "", false, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Документация
- Обновляйте README.md при добавлении функций
- Добавляйте примеры в USAGE_GUIDE.md
- Документируйте API в API_DOCUMENTATION.md

## Git Commits

### Commit Message Format
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: Новая функция
- `fix`: Исправление бага
- `docs`: Изменения документации
- `style`: Форматирование, missing semicolons, etc
- `refactor`: Переструктуризация кода
- `perf`: Оптимизация производительности
- `test`: Добавление тестов
- `chore`: Build process, dependencies, etc

**Примеры:**
```
feat(vlan): add support for dynamic VLAN configuration

- Implement VLAN creation with dynamic IDs
- Add interface management
- Update documentation

Closes #123
```

```
fix(firewall): resolve rule priority ordering issue

The firewall rules were not being sorted correctly by priority.
Fixed the sorting algorithm to use stable sort.

Fixes #456
```

## Pull Request Process

### 1. Обновите ветку develop
```bash
git fetch origin
git rebase origin/develop
```

### 2. Запустите все тесты
```bash
go test ./... -v -race -cover -timeout=30s
```

### 3. Убедитесь в линтинге
```bash
go fmt ./...
go vet ./...
```

### 4. Создайте Pull Request
- Дайте понятное описание изменений
- Ссылайтесь на связанные issues
- Добавьте примеры использования если применимо
- Убедитесь что CI/CD прошел успешно

### Checklist для PR:
```markdown
- [ ] Код следует style guidelines
- [ ] Все тесты пройдены
- [ ] Документация обновлена
- [ ] Commit messages следуют формату
- [ ] Нет конфликтов с develop веткой
- [ ] Изменения не вносят breaking changes
```

## Сообщение об Ошибках

### Перед тем как сообщить
1. Проверьте [FAQ](FAQ.md)
2. Проверьте существующие issues
3. Попробуйте воспроизвести с последней версией

### Как сообщить
Используйте GitHub Issues с следующей информацией:

```markdown
## Описание
Краткое описание проблемы

## Шаги воспроизведения
1. Выполните...
2. Затем...
3. Результат...

## Ожидаемое поведение
Что должно было произойти

## Текущее поведение
Что произошло на самом деле

## Окружение
- ОС: [например Windows 11]
- Go версия: [например 1.21]
- Версия приложения: [например v1.0.0]

## Логи
```
<paste logs here>
```
```

## Request для Фич

### Заголовок
Четкое и описательное название

### Описание
- Цель фичи
- Примеры использования
- Дополнительная информация

```markdown
## Описание
Добавить возможность импорта конфигурации из JSON

## Примеры использования
```bash
network-config import --format json --file config.json
```

## Дополнительная информация
JSON формат должен быть совместим с текущей YAML схемой
```

## Development Setup

### Требуемые инструменты
```bash
# Go 1.21+
go version

# Git
git --version

# Docker (опционально)
docker --version

# Docker Compose (опционально)
docker-compose --version
```

### IDE Setup

#### VS Code
```json
{
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "go.vetOnSave": "package",
  "editor.formatOnSave": true,
  "[go]": {
    "editor.defaultFormatter": "golang.go"
  }
}
```

#### GoLand/IntelliJ IDEA
- Встроенная поддержка Go
- Включите `gofmt` при сохранении
- Настройте Linter в Settings → Go → Linter

## Тестирование

### Unit Тесты
```bash
go test ./... -v
```

### Тесты с Coverage
```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Тесты с Race Detector
```bash
go test ./... -race
```

### Бенчмарки
```bash
go test -bench=. -benchmem ./...
```

## Documentation

### Обновление Документации
1. README.md - основная информация
2. USAGE_GUIDE.md - примеры использования
3. API_DOCUMENTATION.md - API справка
4. ADVANCED_CONFIG.md - продвинутые примеры

### Markdown Style
- Используйте 2 пробела для отступа в списках
- Форматируйте code blocks с языком
- Используйте relative links для внутренних ссылок

```markdown
# Хорошо
```go
func Hello() string {
    return "world"
}
```

# Плохо
```
func Hello() string {
    return "world"
}
```
```

## License

При контрибуции вы соглашаетесь что ваш код лицензирован под MIT License.

## Вопросы?

- Откройте GitHub Discussion
- Создайте GitHub Issue с вопросом
- Свяжитесь с мейнтейнерами

---

**Спасибо за контрибуцию! 🚀**
