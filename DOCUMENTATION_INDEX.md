# Network Configuration Manager - Индекс Документации

## 📚 Полный Справочник Документации

### 🚀 Начало Работы

| Документ | Назначение | Для Кого |
|----------|-----------|----------|
| [README.md](README.md) | Основная информация о проекте | Все пользователи |
| [QUICK_REFERENCE.md](QUICK_REFERENCE.md) | Краткая справка команд | Опытные пользователи |
| [USAGE_GUIDE.md](USAGE_GUIDE.md) | Подробное руководство с примерами | Начинающие пользователи |
| [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md) | Проектное резюме | Менеджеры, планировщики |

### 🔧 Техническая Документация

| Документ | Назначение | Для Кого |
|----------|-----------|----------|
| [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) | Архитектура и структура кода | Разработчики |
| [API_DOCUMENTATION.md](API_DOCUMENTATION.md) | REST API (для v2.0) | API интеграторы |
| [REQUIREMENTS.md](REQUIREMENTS.md) | Системные требования | Системные администраторы |
| [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) | Продвинутая конфигурация | Опытные администраторы |

### 📖 Справочники

| Документ | Назначение | Для Кого |
|----------|-----------|----------|
| [FAQ.md](FAQ.md) | Часто задаваемые вопросы | Все пользователи |
| [CHANGELOG.md](CHANGELOG.md) | История версий | Разработчики, пользователи |
| [LICENSE](LICENSE) | Лицензионное соглашение | Все |

### 📝 Примеры Конфигурации

| Файл | Назначение |
|------|-----------|
| [examples/network.yml](examples/network.yml) | Пример конфигурации сети |
| [examples/ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) | Примеры продвинутого использования |

### 💻 Исходный Код

| Файл | Назначение |
|------|-----------|
| [main.go](main.go) | CLI точка входа |
| [cmd/commands.go](cmd/commands.go) | Реализация команд |
| [config/config.go](config/config.go) | Структуры конфигурации |
| [network/monitor.go](network/monitor.go) | Мониторинг трафика |
| [network/firewall.go](network/firewall.go) | Управление фаерволом |
| [network/vpn.go](network/vpn.go) | VPN менеджмент |
| [tests/network_test.go](tests/network_test.go) | Unit тесты |

### 🐳 Развертывание

| Файл | Назначение |
|------|-----------|
| [Dockerfile](Dockerfile) | Docker образ |
| [docker-compose.yml](docker-compose.yml) | Docker Compose конфигурация |
| [install.sh](install.sh) | Linux/macOS установка |
| [install.ps1](install.ps1) | Windows установка |
| [go.mod](go.mod) | Go модули зависимостей |

---

## 🎯 Быстрая Навигация по Задачам

### Я хочу...

#### 🔰 Начать Работу
1. Читаем [README.md](README.md)
2. Просматриваем [QUICK_REFERENCE.md](QUICK_REFERENCE.md)
3. Запускаем установку ([install.ps1](install.ps1) или [install.sh](install.sh))
4. Пробуем примеры из [USAGE_GUIDE.md](USAGE_GUIDE.md)

#### 📚 Изучить Все Возможности
1. [USAGE_GUIDE.md](USAGE_GUIDE.md) - подробные примеры
2. [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) - продвинутые техники
3. [examples/network.yml](examples/network.yml) - примеры конфигураций

#### 🔐 Настроить VPN
1. [USAGE_GUIDE.md](USAGE_GUIDE.md) → VPN Configuration Examples
2. [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) → VPN Protocols Configuration
3. [examples/network.yml](examples/network.yml) - VPN примеры

#### 🛡️ Настроить Firewall
1. [USAGE_GUIDE.md](USAGE_GUIDE.md) → Firewall Management
2. [FAQ.md](FAQ.md) → Firewall Management раздел
3. [REQUIREMENTS.md](REQUIREMENTS.md) → Firewall раздел

#### 📊 Настроить QoS
1. [USAGE_GUIDE.md](USAGE_GUIDE.md) → QoS Configuration
2. [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) → QoS Scheduling
3. [examples/network.yml](examples/network.yml) - QoS примеры

#### 👀 Мониторить Трафик
1. [USAGE_GUIDE.md](USAGE_GUIDE.md) → Traffic Monitoring
2. [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) → Traffic Monitor раздел
3. [FAQ.md](FAQ.md) → Traffic Monitoring раздел

#### 🐳 Развернуть в Docker
1. [README.md](README.md) → Docker раздел
2. [docker-compose.yml](docker-compose.yml) - конфигурация
3. [REQUIREMENTS.md](REQUIREMENTS.md) → Docker раздел

#### 🧑‍💻 Разрабатывать/Расширять
1. [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - архитектура
2. [API_DOCUMENTATION.md](API_DOCUMENTATION.md) - интерфейсы
3. Просмотр [cmd/commands.go](cmd/commands.go), [network/](network/) файлов

#### 🆘 Решить Проблему
1. [FAQ.md](FAQ.md) → Troubleshooting раздел
2. [USAGE_GUIDE.md](USAGE_GUIDE.md) → Troubleshooting раздел
3. [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - как это работает

#### 📋 Узнать Требования Системы
[REQUIREMENTS.md](REQUIREMENTS.md) - полная информация о требованиях

#### 📖 Получить API Документацию
[API_DOCUMENTATION.md](API_DOCUMENTATION.md) - REST API справочник

#### 🤔 Найти Ответ на Вопрос
[FAQ.md](FAQ.md) - часто задаваемые вопросы

#### 📜 Узнать об Истории Проекта
[CHANGELOG.md](CHANGELOG.md) - версии и изменения

#### ⚖️ Узнать Лицензию
[LICENSE](LICENSE) - MIT лицензия

---

## 📊 Структура Документации

```
Документация
├── Начало Работы
│   ├── README.md           ← Начните здесь
│   ├── QUICK_REFERENCE.md  ← Краткая справка
│   └── USAGE_GUIDE.md      ← Подробное руководство
│
├── Техническая Информация
│   ├── PROJECT_STRUCTURE.md    ← Архитектура
│   ├── API_DOCUMENTATION.md    ← API справка
│   ├── REQUIREMENTS.md         ← Системные требования
│   └── PROJECT_SUMMARY.md      ← Проектное резюме
│
├── Справочники
│   ├── FAQ.md              ← Часто задаваемые вопросы
│   ├── CHANGELOG.md        ← История версий
│   └── LICENSE             ← Лицензия
│
├── Примеры и Конфигурации
│   └── examples/
│       ├── network.yml         ← Пример конфигурации
│       └── ADVANCED_CONFIG.md  ← Продвинутые примеры
│
├── Исходный Код
│   ├── main.go
│   ├── cmd/
│   ├── config/
│   ├── network/
│   └── tests/
│
└── Развертывание
    ├── Dockerfile
    ├── docker-compose.yml
    ├── install.sh
    └── install.ps1
```

---

## 🔍 Поиск по Темам

### VLAN (Virtual Local Area Networks)
- [USAGE_GUIDE.md](USAGE_GUIDE.md#vlan-management-examples)
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md#vlan-commands)
- [examples/network.yml](examples/network.yml)
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md#vlan-support)

### VPN (Virtual Private Networks)
- [USAGE_GUIDE.md](USAGE_GUIDE.md#vpn-configuration)
- [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md#vpn-protocols-configuration)
- [FAQ.md](FAQ.md#vpn-configuration)
- [examples/network.yml](examples/network.yml)

### Firewall (Межсетевой Экран)
- [USAGE_GUIDE.md](USAGE_GUIDE.md#firewall-management)
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md#firewall-commands)
- [FAQ.md](FAQ.md#firewall-management)
- [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md#firewall-rules-priority)

### QoS (Quality of Service)
- [USAGE_GUIDE.md](USAGE_GUIDE.md#quality-of-service-qos)
- [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md#qos-scheduling-algorithms)
- [FAQ.md](FAQ.md#quality-of-service-qos)

### Мониторинг Трафика
- [USAGE_GUIDE.md](USAGE_GUIDE.md#traffic-monitoring)
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md#monitoring-commands)
- [FAQ.md](FAQ.md#traffic-monitoring)

### Конфигурация
- [USAGE_GUIDE.md](USAGE_GUIDE.md#configuration-file-usage)
- [examples/network.yml](examples/network.yml)
- [FAQ.md](FAQ.md#configuration)

### Docker & Контейнеры
- [README.md](README.md#docker)
- [Dockerfile](Dockerfile)
- [docker-compose.yml](docker-compose.yml)
- [FAQ.md](FAQ.md#docker--containers)

### Установка
- [README.md](README.md#установка)
- [install.sh](install.sh)
- [install.ps1](install.ps1)
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md#installation)

### Troubleshooting
- [FAQ.md](FAQ.md#troubleshooting)
- [USAGE_GUIDE.md](USAGE_GUIDE.md#troubleshooting)
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md#troubleshooting)

### API & Интеграция
- [API_DOCUMENTATION.md](API_DOCUMENTATION.md)
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md#api-функционал)

### Производительность & Оптимизация
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md#performance-tips)
- [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md#performance-optimization)
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md#performance-characteristics)

### Безопасность
- [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md#security-best-practices)
- [REQUIREMENTS.md](REQUIREMENTS.md#безопасность)

---

## 📞 Помощь и Поддержка

### Где Найти Информацию
1. **FAQ.md** - Ответы на частые вопросы
2. **USAGE_GUIDE.md** - Подробные примеры и объяснения
3. **QUICK_REFERENCE.md** - Быстрые команды
4. **Исходный код** - Комментарии в коде

### Сообщение об Ошибках
- GitHub Issues (будет добавлено)
- Email поддержка (будет добавлено)

### Предложение Функций
- GitHub Discussions (будет добавлено)
- Feature requests (будет добавлено)

---

## 🎓 Рекомендуемый Порядок Чтения

### Для Новичков
1. README.md (5 минут)
2. QUICK_REFERENCE.md (10 минут)
3. USAGE_GUIDE.md (30 минут)
4. examples/network.yml (5 минут)
5. Практика с примерами (30 минут)

### Для Опытных Пользователей
1. README.md (2 минуты)
2. QUICK_REFERENCE.md (5 минут)
3. ADVANCED_CONFIG.md (20 минут)
4. FAQ.md (при необходимости)

### Для Разработчиков
1. PROJECT_STRUCTURE.md (30 минут)
2. Просмотр исходного кода (1 час)
3. API_DOCUMENTATION.md (20 минут)
4. tests/network_test.go (15 минут)

### Для Администраторов
1. README.md (5 минут)
2. REQUIREMENTS.md (10 минут)
3. USAGE_GUIDE.md (30 минут)
4. ADVANCED_CONFIG.md (20 минут)
5. FAQ.md → Troubleshooting (при необходимости)

---

**Последнее обновление**: 2026-01-29
**Версия документации**: 1.0
**Статус**: Полная ✅
