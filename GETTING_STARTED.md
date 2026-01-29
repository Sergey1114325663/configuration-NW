# Network Configuration Manager - Инструкции по Началу Работы

## 🚀 Быстрый Старт (5 минут)

### Шаг 1: Установка

#### Windows
```powershell
# Откройте PowerShell в папке проекта
cd "C:\Users\fosin\OneDrive\Рабочий стол\Projects\go"

# Запустите установку
.\install.ps1

# Проверьте установку
.\network-config.exe --help
```

#### Linux/macOS
```bash
# Перейти в папку проекта
cd ~/Projects/go

# Запустить установку
bash install.sh

# Проверить установку
./network-config --help
```

### Шаг 2: Первые Команды

```bash
# 1. Создать VLAN
network-config vlan create --id 10 --name "Production"

# 2. Просмотреть VLANs
network-config vlan list

# 3. Добавить firewall правило
network-config firewall rule --action add --rule "allow tcp port 80"

# 4. Просмотреть firewall статус
network-config firewall status

# 5. Начать мониторинг
network-config monitor stats
```

### Шаг 3: Загрузить Пример Конфигурации

```bash
# Копировать пример конфигурации
cp examples/network.yml ./config/network.yml

# Отредактировать под свои нужды
# На Windows: notepad config\network.yml
# На Linux: nano config/network.yml
# На macOS: open config/network.yml
```

---

## 📚 Полное Руководство (1 час)

### 1. Понимание Архитектуры (15 минут)

**Цель**: Понять как работает приложение

Читайте:
- [README.md](README.md) - Обзор проекта (5 мин)
- [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md) - Резюме проекта (10 мин)

**Что вы узнаете**:
- ✅ Основные возможности
- ✅ Поддерживаемые протоколы
- ✅ Архитектура приложения

### 2. Установка и Конфигурация (15 минут)

**Цель**: Установить и настроить приложение

Действия:
1. Запустите установку ([install.sh](install.sh) или [install.ps1](install.ps1))
2. Проверьте установку: `network-config --help`
3. Посмотрите примеры: `examples/network.yml`

**Что вы получите**:
- ✅ Рабочее приложение
- ✅ Пример конфигурации
- ✅ Справочная информация

### 3. Практические Примеры (20 минут)

**Цель**: Попробовать основные команды

Следуйте [USAGE_GUIDE.md](USAGE_GUIDE.md):

#### VLAN Management (5 мин)
```bash
# Создать несколько VLANs
./network-config vlan create --id 1 --name "Management"
./network-config vlan create --id 10 --name "Production"
./network-config vlan create --id 20 --name "Development"

# Просмотреть список
./network-config vlan list

# Удалить VLAN
./network-config vlan delete --id 20
```

#### VPN Configuration (5 мин)
```bash
# Создать L2TP подключение
./network-config vpn l2tp --action create --name "office-vpn"

# Создать OpenVPN подключение
./network-config vpn openvpn --action start --config "vpn.conf"

# Список всех команд
./network-config vpn --help
```

#### Firewall Management (5 мин)
```bash
# Добавить правила
./network-config firewall rule --action add --rule "allow tcp port 80"
./network-config firewall rule --action add --rule "allow tcp port 443"

# Просмотреть правила
./network-config firewall list

# Проверить статус
./network-config firewall status
```

#### QoS Configuration (3 мин)
```bash
# Создать политику QoS
./network-config qos policy --name "high-priority" --bandwidth "100Mbps"

# Установить приоритет
./network-config qos priority --class voip --priority 7

# Ограничить bandwidth
./network-config qos limit --interface eth0 --limit "50Mbps"
```

#### Traffic Monitoring (2 мин)
```bash
# Мониторить интерфейс
./network-config monitor traffic --interface eth0

# Получить статистику
./network-config monitor stats

# Установить алерт
./network-config monitor alerts --threshold "100Mbps"
```

### 4. Продвинутая Конфигурация (10 минут)

**Цель**: Научиться продвинутым техникам

Читайте [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md):

- VLAN оптимизация
- VPN оптимизация
- Firewall оптимизация
- QoS оптимизация
- Security best practices
- HA (High Availability) setup

---

## 🔧 Практические Сценарии

### Сценарий 1: Настройка Продакшн Сети

**Время**: 30 минут
**Сложность**: Средняя

1. Создайте VLANs для сегментации:
   ```bash
   network-config vlan create --id 1 --name "Management" --description "Управление"
   network-config vlan create --id 10 --name "Production" --description "Продакшн"
   network-config vlan create --id 100 --name "DMZ" --description "Интернет"
   ```

2. Настройте Firewall:
   ```bash
   network-config firewall rule --action add --rule "allow tcp port 80"
   network-config firewall rule --action add --rule "allow tcp port 443"
   network-config firewall rule --action add --rule "allow tcp port 22"
   ```

3. Установите QoS приоритеты:
   ```bash
   network-config qos policy --name "critical" --bandwidth "200Mbps"
   network-config qos priority --class critical --priority 7
   ```

4. Начните мониторинг:
   ```bash
   network-config monitor traffic --interface eth0
   network-config monitor alerts --threshold "150Mbps"
   ```

### Сценарий 2: Настройка Удаленного Доступа

**Время**: 20 минут
**Сложность**: Средняя

1. Создайте OpenVPN подключение:
   ```bash
   network-config vpn openvpn --action start --config "/path/to/vpn.conf"
   ```

2. Создайте резервное L2TP подключение:
   ```bash
   network-config vpn l2tp --action create --name "backup-vpn"
   ```

3. Дополните Firewall правилами:
   ```bash
   network-config firewall rule --action add --rule "allow tcp port 443"
   network-config firewall rule --action add --rule "allow udp port 1194"
   ```

4. Установите мониторинг VPN:
   ```bash
   network-config monitor traffic --interface vpn0
   ```

### Сценарий 3: Мониторинг и Оптимизация

**Время**: 25 минут
**Сложность**: Высокая

1. Начните детальный мониторинг:
   ```bash
   network-config monitor traffic --interface eth0
   network-config monitor stats
   ```

2. Проанализируйте результаты и создайте QoS политики:
   ```bash
   network-config qos policy --name "video-streaming" --bandwidth "80Mbps"
   network-config qos priority --class video --priority 6
   ```

3. Оптимизируйте Firewall:
   ```bash
   network-config firewall list  # Просмотр текущих правил
   # Удалить неиспользуемые правила
   ```

4. Установите алерты:
   ```bash
   network-config monitor alerts --threshold "100Mbps"
   ```

---

## 📖 Дополнительные Ресурсы

### Документация
| Документ | Описание |
|----------|---------|
| [README.md](README.md) | Основная информация |
| [QUICK_REFERENCE.md](QUICK_REFERENCE.md) | Краткая справка |
| [FAQ.md](FAQ.md) | Часто задаваемые вопросы |
| [API_DOCUMENTATION.md](API_DOCUMENTATION.md) | API справка |

### Примеры
| Файл | Описание |
|------|---------|
| [examples/network.yml](examples/network.yml) | Пример конфигурации |
| [examples/ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) | Продвинутые примеры |

### Установка
| Файл | Назначение |
|------|-----------|
| [install.ps1](install.ps1) | Windows установка |
| [install.sh](install.sh) | Linux/macOS установка |
| [Dockerfile](Dockerfile) | Docker образ |

---

## ✅ Контрольный Список

### День 1 - Установка и Базовая Конфигурация
- [ ] Установить приложение
- [ ] Проверить `network-config --help`
- [ ] Прочитать README.md
- [ ] Создать первый VLAN
- [ ] Добавить первое firewall правило

### День 2 - Расширенная Конфигурация
- [ ] Создать VPN подключение
- [ ] Настроить QoS политики
- [ ] Включить мониторинг
- [ ] Прочитать ADVANCED_CONFIG.md
- [ ] Подготовить конфигурацию для продакшна

### День 3 - Оптимизация и Мониторинг
- [ ] Проанализировать трафик
- [ ] Оптимизировать firewall правила
- [ ] Отрегулировать QoS приоритеты
- [ ] Установить алерты
- [ ] Создать резервную копию конфигурации

---

## 🆘 Если Что-то Не Работает

### Шаг 1: Проверяйте Логи
```bash
# Просмотреть логи
cat logs/network-config.log  # Linux/macOS
type logs\network-config.log  # Windows
```

### Шаг 2: Проверьте FAQ
Откройте [FAQ.md](FAQ.md) → Troubleshooting раздел

### Шаг 3: Проверьте USAGE_GUIDE
Откройте [USAGE_GUIDE.md](USAGE_GUIDE.md) → Troubleshooting раздел

### Шаг 4: Посмотрите Требования
Откройте [REQUIREMENTS.md](REQUIREMENTS.md) для проверки системных требований

---

## 💡 Советы для Успеха

### ✅ Лучшие Практики

1. **Версионирование Конфигурации**
   ```bash
   git init
   git add config/
   git commit -m "Initial network configuration"
   ```

2. **Резервные Копии**
   ```bash
   cp config/network.yml config/network.yml.backup
   ```

3. **Тестирование**
   - Создайте тестовую среду перед продакшном
   - Протестируйте все правила firewall
   - Проверьте VPN подключения

4. **Логирование**
   - Включите детальное логирование в начале
   - Регулярно просматривайте логи
   - Сохраняйте логи для анализа

5. **Мониторинг**
   - Начните мониторинг с первого дня
   - Установите разумные пороги алертов
   - Регулярно анализируйте метрики

### ⚠️ Частые Ошибки

1. ❌ Неправильный синтаксис команд
   ✅ Используйте `--help` для справки

2. ❌ Забыли про флаги
   ✅ Проверьте [QUICK_REFERENCE.md](QUICK_REFERENCE.md)

3. ❌ Неправильный формат конфигурации
   ✅ Используйте YAML формат из [examples/network.yml](examples/network.yml)

4. ❌ Недостаточные привилегии
   ✅ Запустите с администраторскими правами

5. ❌ Забыли сохранить конфигурацию
   ✅ Используйте версионирование (git)

---

## 🎓 Обучающие Ресурсы

### Понять VLAN
- https://en.wikipedia.org/wiki/Virtual_LAN
- Посмотрите примеры в [examples/network.yml](examples/network.yml)

### Понять VPN
- Читайте [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) → VPN Protocols

### Понять Firewall
- Читайте [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) → Firewall Rules Priority

### Понять QoS
- Читайте [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) → QoS Scheduling Algorithms

### Понять Мониторинг
- Читайте [ADVANCED_CONFIG.md](examples/ADVANCED_CONFIG.md) → Traffic Monitoring Metrics

---

## 🚀 Следующие Шаги

1. ✅ Завершить установку
2. ✅ Протестировать базовые команды
3. ✅ Прочитать документацию
4. ✅ Создать собственную конфигурацию
5. ✅ Развернуть в продакшн (с осторожностью!)

---

**Начните с [README.md](README.md) и следуйте этому руководству!**

**Удачи в управлении вашей сетью! 🎉**
